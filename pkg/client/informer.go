// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/SlinkyProject/slurm-client/internal/wait"
	"github.com/SlinkyProject/slurm-client/pkg/cache"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	defaultSyncPeriod = 30 * time.Second
	batchPeriod       = 1 * time.Second

	// syncPollInterval is how often waitForSyncGet/waitForSyncList poll the in-memory
	// dirty flag while waiting for a pending refresh to land, clamped to the wait's own
	// timeout so a very low CacheSyncPeriod still gets at least one poll inside its window.
	// The check itself is a cheap map lookup under a read lock, so this can be much finer
	// than syncPeriod without added cost; a coarse interval only adds pure latency to every
	// RefreshCache/WaitRefreshCache caller, since Create and Update always wait on this
	// internally after every write.
	syncPollInterval = 50 * time.Millisecond
)

type cacheEntry struct {
	lastUpdate time.Time
	object     object.Object
	dirty      bool
}

var _ InformerCache = &informerCache{}

type informerCache struct {
	// reader knows how to read from remote.
	reader Reader

	// objectType tracks the object type this informer backs.
	objectType object.ObjectType

	// mu guards access to the map.
	mu sync.RWMutex

	// cache holds the actual object cache.
	cache   map[object.ObjectKey]*cacheEntry
	indexes map[string]*secondaryIndex

	// started is true if the informers have been started.
	started bool

	// dirty indicates if the informer cache as a whole should be considered dirty.
	dirty bool

	// eventCh holds events for handler.
	eventCh chan event.Event

	// syncCh holds sync requests.
	syncCh chan struct{}

	// syncObjCh holds sync requests by ObjectKey.
	syncObjCh chan object.ObjectKey

	// syncError is the last List sync error
	syncErrorList error

	// syncErrorGet is the last Get sync error per object
	syncErrorGet map[object.ObjectKey]error

	// handler runs for each read event from eventCh.
	handler cache.ResourceEventHandler

	// syncPeriod is the frequency to run the informer.
	syncPeriod time.Duration
}

// SetEventHandler implements InformerCache.
func (i *informerCache) SetEventHandler(handler cache.ResourceEventHandler) {
	i.handler = handler
}

// UnsetEventHandler implements InformerCache.
func (i *informerCache) UnsetEventHandler() {
	i.handler = nil
}

// Run implements InformerCache.
func (i *informerCache) Run(stopCh <-chan struct{}) {
	i.mu.RLock()
	if i.started {
		defer i.mu.RUnlock()
		return
	}
	i.mu.RUnlock()

	i.mu.Lock()
	i.started = true
	i.mu.Unlock()

	go i.runListInformer(stopCh)
	go i.runGetInformer(stopCh)
	go i.runHandler(stopCh)

	<-stopCh
	i.mu.Lock()
	i.started = false
	i.mu.Unlock()
}

func (i *informerCache) runListInformer(stopCh <-chan struct{}) {
	ticker := time.NewTicker(i.syncPeriod)
	defer ticker.Stop()

	batchTimer := time.NewTimer(batchPeriod)
	defer batchTimer.Stop()

	requestSync := false
	for {
		select {
		case _, ok := <-i.syncCh:
			if !ok {
				// syncCh was closed!
				return
			}
			i.mu.Lock()
			i.dirty = true
			i.mu.Unlock()
			requestSync = true
		case <-batchTimer.C:
			if requestSync {
				go i.doListInformer()
				requestSync = false
			}
			batchTimer.Reset(batchPeriod)
		case <-ticker.C:
			i.syncCh <- struct{}{}
		case <-stopCh:
			return
		}
	}
}

func (i *informerCache) doListInformer() {
	i.mu.Lock()
	i.dirty = true
	i.mu.Unlock()

	r, ok := resourceCatalog[i.objectType]
	if !ok || !r.cacheable || r.newList == nil {
		panic(fmt.Sprintf("doListInformer: attempting to cache '%s', this should never have happened.", i.objectType))
	}
	list := r.newList()

	opts := &ListOptions{SkipCache: true}
	err := i.reader.List(context.TODO(), list, opts)
	i.mu.Lock()
	i.syncErrorList = err
	if err == nil {
		i.processObjects(list)
		i.dirty = false
	}
	i.mu.Unlock()
}

func (i *informerCache) runGetInformer(stopCh <-chan struct{}) {
	batchTimer := time.NewTimer(batchPeriod)
	defer batchTimer.Stop()

	requestSync := make(map[object.ObjectKey]struct{})
	for {
		select {
		case key, ok := <-i.syncObjCh:
			if !ok {
				// syncObjCh was closed!
				return
			}
			i.mu.Lock()
			if obj := i.cache[key]; obj != nil {
				obj.dirty = true
			} else {
				i.cache[key] = &cacheEntry{dirty: true}
			}
			i.mu.Unlock()
			requestSync[key] = struct{}{}
		case <-batchTimer.C:
			if len(requestSync) > 0 {
				for key := range requestSync {
					go i.doGetInformer(key)
				}
				requestSync = make(map[object.ObjectKey]struct{})
			}
			batchTimer.Reset(batchPeriod)
		case <-stopCh:
			return
		}
	}
}

func (i *informerCache) doGetInformer(key object.ObjectKey) {
	i.mu.Lock()
	if obj := i.cache[key]; obj != nil {
		obj.dirty = true
	} else {
		i.cache[key] = &cacheEntry{dirty: true}
	}
	i.mu.Unlock()

	r, ok := resourceCatalog[i.objectType]
	if !ok || !r.cacheable || r.newObject == nil {
		panic(fmt.Sprintf("doGetInformer: attempting to cache '%s', this should never have happened.", i.objectType))
	}
	obj := r.newObject()

	opts := &GetOptions{SkipCache: true}
	err := i.reader.Get(context.TODO(), key, obj, opts)

	i.mu.Lock()
	if err != nil && !errors.Is(err, apierrors.ErrNotFound) {
		i.syncErrorGet[key] = err
	} else {
		i.syncErrorGet[key] = nil
	}
	if errors.Is(err, apierrors.ErrNotFound) {
		i.deleteObject(key)
	} else if i.syncErrorGet[key] == nil {
		i.processObject(obj)
	}
	i.mu.Unlock()
}

func (i *informerCache) runHandler(stopCh <-chan struct{}) {
	for {
		select {
		case e := <-i.eventCh:
			if i.handler == nil {
				continue
			}
			go i.doHandler(e)
		case <-stopCh:
			return
		}
	}
}

func (i *informerCache) doHandler(evt event.Event) {
	switch evt.Type {
	case event.Added:
		i.handler.OnAdd(evt.Object, !i.dirty)
	case event.Modified:
		i.handler.OnUpdate(evt.ObjectOld, evt.Object)
	case event.Deleted:
		i.handler.OnDelete(evt.Object)
	}
}

func (i *informerCache) pushEvent(e event.Event) {
	if i.eventCh == nil || i.handler == nil {
		return
	}
	i.eventCh <- e
}

func (i *informerCache) processObjects(list object.ObjectList) {
	now := time.Now()
	for _, item := range list.GetItems() {
		i.processObject(item)
	}

	for key, entry := range i.cache {
		if entry == nil {
			continue
		}
		if entry.object == nil {
			i.deleteObject(key)
		} else if now.After(entry.lastUpdate) {
			i.deleteObject(key)
		}
	}
}

// deleteObject removes an object and its secondary keys under mu.
func (i *informerCache) deleteObject(key object.ObjectKey) {
	entry := i.cache[key]
	for _, index := range i.indexes {
		index.remove(key)
	}
	delete(i.cache, key)
	delete(i.syncErrorGet, key)
	if entry != nil && entry.object != nil {
		i.pushEvent(event.Event{Type: event.Deleted, Object: entry.object.DeepCopyObject().(object.Object)})
	}
}

func (i *informerCache) processObject(obj object.Object) {
	now := time.Now()
	key := obj.GetKey()

	entry, ok := i.cache[key]
	if !ok || entry == nil || entry.object == nil {
		i.cache[key] = &cacheEntry{
			lastUpdate: now,
			object:     obj.DeepCopyObject().(object.Object),
			dirty:      false,
		}
		delete(i.syncErrorGet, key)
		for _, index := range i.indexes {
			index.update(i.cache[key].object)
		}
		e := event.Event{
			Type:   event.Added,
			Object: obj.DeepCopyObject().(object.Object),
		}
		i.pushEvent(e)
	} else if ok && entry.object != nil && !now.Before(entry.lastUpdate) {
		entry.lastUpdate = now
		entry.dirty = false
		delete(i.syncErrorGet, key)
		if !reflect.DeepEqual(entry.object, obj) {
			old := entry.object
			entry.object = obj.DeepCopyObject().(object.Object)
			for _, index := range i.indexes {
				index.update(entry.object)
			}
			e := event.Event{
				Type:      event.Modified,
				Object:    obj.DeepCopyObject().(object.Object),
				ObjectOld: old.DeepCopyObject().(object.Object),
			}
			i.pushEvent(e)
		}
	}
}

// HasSynced implements InformerCache.
func (i *informerCache) HasSynced() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if !i.started {
		return false, fmt.Errorf("informer cache %s has not started, cannot sync", i.objectType)
	}
	return !i.dirty, i.syncErrorList
}

func (i *informerCache) hasSyncedList() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if !i.started {
		return true, fmt.Errorf("informer cache %s has not started, cannot sync", i.objectType)
	}
	return !i.dirty, i.syncErrorList
}

func (i *informerCache) hasSyncedGet(key object.ObjectKey) (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	if !i.started {
		return true, fmt.Errorf("informer cache %s has not started, cannot sync", i.objectType)
	}
	if obj := i.cache[key]; obj != nil {
		return !obj.dirty, i.syncErrorGet[key]
	}
	return !i.dirty, i.syncErrorList
}

// HasStarted implements InformerCache.
func (i *informerCache) HasStarted() bool {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.started
}

func (i *informerCache) waitForSyncList(ctx context.Context) error {
	timeout := 2 * i.syncPeriod
	conditionFn := func(ctx context.Context) (bool, error) {
		return i.hasSyncedList()
	}
	return wait.PollUntilContextTimeout(ctx, min(syncPollInterval, timeout), timeout, true, conditionFn)
}

func (i *informerCache) waitForSyncGet(ctx context.Context, key object.ObjectKey) error {
	timeout := 2 * i.syncPeriod
	conditionFn := func(ctx context.Context) (bool, error) {
		return i.hasSyncedGet(key)
	}
	return wait.PollUntilContextTimeout(ctx, min(syncPollInterval, timeout), timeout, true, conditionFn)
}

// Get implements InformerCache.
func (i *informerCache) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...GetOption) error {
	options := &GetOptions{}
	options.ApplyOptions(opts)

	if options.RefreshCache {
		// Mark dirty before sync request to avoid lock race between channel
		// receiver and waitForSyncGet().
		i.mu.Lock()
		if obj := i.cache[key]; obj != nil {
			obj.dirty = true
		} else {
			i.cache[key] = &cacheEntry{dirty: true}
		}
		i.mu.Unlock()

		select {
		case i.syncObjCh <- key:
		case <-ctx.Done():
			return ctx.Err()
		}
	} else if options.WaitRefreshCache {
		i.mu.Lock()
		if obj := i.cache[key]; obj != nil {
			obj.dirty = true
		} else {
			i.cache[key] = &cacheEntry{dirty: true}
		}
		i.mu.Unlock()
	}

	if err := i.waitForSyncGet(ctx, key); err != nil {
		return fmt.Errorf("failed to wait on type %s object %s cache sync: %w", obj.GetType(), key, err)
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	entry, ok := i.cache[key]
	if !ok || entry.object == nil {
		return apierrors.ErrNotFound
	}

	if err := copyObject(obj, entry.object.DeepCopyObject().(object.Object)); err != nil {
		return err
	}

	return nil
}

// List implements InformerCache.
func (i *informerCache) List(ctx context.Context, list object.ObjectList, opts ...ListOption) error {
	options := &ListOptions{}
	options.ApplyOptions(opts)

	if options.RefreshCache {
		// Mark dirty before sync request to avoid lock race between channel
		// receiver and waitForSyncList().
		i.mu.Lock()
		i.dirty = true
		i.mu.Unlock()

		select {
		case i.syncCh <- struct{}{}:
		case <-ctx.Done():
			return ctx.Err()
		}
	} else if options.WaitRefreshCache {
		i.mu.Lock()
		i.dirty = true
		i.mu.Unlock()
	}

	if err := i.waitForSyncList(ctx); err != nil {
		return fmt.Errorf("failed to wait on type %s cache sync: %w", list.GetType(), err)
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	for _, entry := range i.cache {
		if entry.object == nil {
			continue
		}
		appendCachedObject(list, entry.object)
	}

	return nil
}

// appendCachedObject lets registered lists make their own deep copy. Comparing
// concrete types keeps custom wrappers that override AppendItem protected.
func appendCachedObject(list object.ObjectList, obj object.Object) {
	if reflect.TypeOf(list) != resourceCatalog[list.GetType()].listType {
		obj = obj.DeepCopyObject().(object.Object)
	}
	list.AppendItem(obj)
}

func newInformer(objectType object.ObjectType, reader Reader, syncPeriod time.Duration) InformerCache {
	return &informerCache{
		reader:       reader,
		objectType:   objectType,
		cache:        make(map[object.ObjectKey]*cacheEntry),
		dirty:        true,
		syncErrorGet: make(map[object.ObjectKey]error),
		syncPeriod:   syncPeriod,
		eventCh:      make(chan event.Event, 8),
		syncCh:       make(chan struct{}, 8),
		syncObjCh:    make(chan object.ObjectKey, 8),
	}
}

func copyObject(dst, src object.Object) error {
	dv := reflect.ValueOf(dst)
	sv := reflect.ValueOf(src)
	if !dv.IsValid() || !sv.IsValid() || dv.Kind() != reflect.Pointer || dv.Type() != sv.Type() {
		return apierrors.ErrNotImplemented
	}
	dv.Elem().Set(sv.Elem())
	return nil
}
