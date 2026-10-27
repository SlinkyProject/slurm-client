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

	// syncPollInterval is how often waitForSyncGet/waitForSyncList poll the in-memory
	// dirty flag while waiting for a pending refresh to land, clamped to the wait's own
	// timeout so a very low CacheSyncPeriod still gets at least one poll inside its window.
	// The check itself is a cheap map lookup under a read lock, so this can be much finer
	// than syncPeriod without added cost; a coarse interval only adds pure latency to every
	// RefreshCache/WaitRefreshCache caller, since writes wait on this internally unless
	// their readback is skipped.
	syncPollInterval = 50 * time.Millisecond
)

// refreshResult is immutable after done closes. A batch covers requests through sequence.
type refreshResult struct {
	sequence uint64
	done     chan struct{}
	object   object.Object
	list     object.ObjectList
	err      error
}
type refreshSequence struct {
	requested, completed uint64
	pending              *refreshResult
}

func (s *refreshSequence) request() *refreshResult {
	s.requested++
	if s.pending == nil {
		s.pending = &refreshResult{done: make(chan struct{})}
	}
	s.pending.sequence = s.requested
	return s.pending
}
func (s *refreshSequence) take() *refreshResult {
	result := s.pending
	s.pending = nil
	return result
}

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

	// mu guards the cache, sync state, and handler registration.
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

	// fetching holds keys with a Get sync in flight; true means another was requested.
	fetching map[object.ObjectKey]bool

	// listing is true while a List sync is in flight; listAgain means another was requested.
	listing, listAgain bool
	getSequences       map[object.ObjectKey]*refreshSequence
	listSequence       refreshSequence

	// handler runs for each read event from eventCh.
	handler cache.ResourceEventHandler

	// syncPeriod is the frequency to run the informer.
	syncPeriod time.Duration
}

// SetEventHandler implements InformerCache.
func (i *informerCache) SetEventHandler(handler cache.ResourceEventHandler) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.handler = handler
}

// UnsetEventHandler implements InformerCache.
func (i *informerCache) UnsetEventHandler() {
	i.SetEventHandler(nil)
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

	for {
		select {
		case _, ok := <-i.syncCh:
			if !ok {
				// syncCh was closed!
				return
			}
			// List right away, but one at a time so an older response cannot
			// overwrite a newer one.
			i.mu.Lock()
			i.dirty = true
			running := i.listing
			i.listing, i.listAgain = true, running
			i.mu.Unlock()
			if !running {
				go i.runListFetch(stopCh)
			}
		case <-ticker.C:
			i.syncCh <- struct{}{}
		case <-stopCh:
			return
		}
	}
}

// runListFetch lists until no refresh arrived during the last list or stopCh closes.
func (i *informerCache) runListFetch(stopCh <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go cancelOnStop(ctx, stopCh, cancel)
	for {
		i.mu.Lock()
		batch := i.listSequence.take()
		i.mu.Unlock()
		list, err := i.doListInformer(ctx)
		i.mu.Lock()
		i.syncErrorList = err
		if err == nil {
			i.processObjects(list)
			i.dirty = false
		}
		if batch != nil {
			batch.list, batch.err = list, err
			i.listSequence.completed = batch.sequence
			close(batch.done)
		}
		if ctx.Err() != nil && i.listSequence.pending != nil {
			pending := i.listSequence.take()
			pending.err = ctx.Err()
			close(pending.done)
		}
		if !i.listAgain || ctx.Err() != nil {
			i.listing, i.listAgain = false, false
			i.mu.Unlock()
			return
		}
		// Keep waiters blocked until the list that started after their request lands.
		i.listAgain, i.dirty = false, true
		i.mu.Unlock()
	}
}

// cancelOnStop cancels ctx when stopCh closes, so stopping aborts in-flight requests.
func cancelOnStop(ctx context.Context, stopCh <-chan struct{}, cancel context.CancelFunc) {
	select {
	case <-stopCh:
		cancel()
	case <-ctx.Done():
	}
}

// doListInformer marks the cache dirty and lists every object from remote.
func (i *informerCache) doListInformer(ctx context.Context) (object.ObjectList, error) {
	i.mu.Lock()
	i.dirty = true
	i.mu.Unlock()

	r, ok := resourceCatalog[i.objectType]
	if !ok || !r.cacheable || r.newList == nil {
		panic(fmt.Sprintf("doListInformer: attempting to cache '%s', this should never have happened.", i.objectType))
	}
	list := r.newList()

	opts := &ListOptions{SkipCache: true}
	return list, i.reader.List(ctx, list, opts)
}

func (i *informerCache) runGetInformer(stopCh <-chan struct{}) {
	for {
		select {
		case key, ok := <-i.syncObjCh:
			if !ok {
				// syncObjCh was closed!
				return
			}
			// Fetch right away, but one at a time per key so an older response
			// cannot overwrite a newer one.
			i.mu.Lock()
			_, running := i.fetching[key]
			i.fetching[key] = running
			i.mu.Unlock()
			if !running {
				go i.runGetFetch(stopCh, key)
			}
		case <-stopCh:
			return
		}
	}
}

// runGetFetch fetches key until no refresh for it arrived during the last fetch or stopCh closes.
func (i *informerCache) runGetFetch(stopCh <-chan struct{}, key object.ObjectKey) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go cancelOnStop(ctx, stopCh, cancel)
	for {
		i.mu.Lock()
		sequence := i.getSequences[key]
		var batch *refreshResult
		if sequence != nil {
			batch = sequence.take()
		}
		i.mu.Unlock()
		obj, err := i.doGetInformer(ctx, key)
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
		if batch != nil {
			batch.object, batch.err = obj, err
			sequence.completed = batch.sequence
			close(batch.done)
		}
		if ctx.Err() != nil && sequence != nil && sequence.pending != nil {
			pending := sequence.take()
			pending.err = ctx.Err()
			close(pending.done)
		}
		if !i.fetching[key] || ctx.Err() != nil {
			delete(i.fetching, key)
			i.mu.Unlock()
			return
		}
		// Keep waiters blocked until the fetch that started after their request lands.
		i.fetching[key] = false
		i.markDirty(key)
		i.mu.Unlock()
	}
}

// doGetInformer marks key dirty and fetches it from remote.
func (i *informerCache) doGetInformer(ctx context.Context, key object.ObjectKey) (object.Object, error) {
	i.mu.Lock()
	i.markDirty(key)
	i.mu.Unlock()

	r, ok := resourceCatalog[i.objectType]
	if !ok || !r.cacheable || r.newObject == nil {
		panic(fmt.Sprintf("doGetInformer: attempting to cache '%s', this should never have happened.", i.objectType))
	}
	obj := r.newObject()

	opts := &GetOptions{SkipCache: true}
	return obj, i.reader.Get(ctx, key, obj, opts)
}

// refresh marks key dirty and queues a fetch without waiting for it to land.
// It does nothing on a stopped informer, which has no fetcher to drain the queue.
func (i *informerCache) refresh(ctx context.Context, key object.ObjectKey) error {
	i.mu.Lock()
	started := i.started
	if started {
		i.markDirty(key)
	}
	i.mu.Unlock()
	if !started {
		return nil
	}
	select {
	case i.syncObjCh <- key:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// markDirty requires i.mu to be held.
func (i *informerCache) markDirty(key object.ObjectKey) {
	if obj := i.cache[key]; obj != nil {
		obj.dirty = true
	} else {
		i.cache[key] = &cacheEntry{dirty: true}
	}
}

func (i *informerCache) runHandler(stopCh <-chan struct{}) {
	for {
		select {
		case e := <-i.eventCh:
			go i.doHandler(e)
		case <-stopCh:
			return
		}
	}
}

func (i *informerCache) doHandler(evt event.Event) {
	i.mu.RLock()
	handler := i.handler
	initial := !i.dirty
	i.mu.RUnlock()
	if handler == nil {
		return
	}

	switch evt.Type {
	case event.Added:
		handler.OnAdd(evt.Object, initial)
	case event.Modified:
		handler.OnUpdate(evt.ObjectOld, evt.Object)
	case event.Deleted:
		handler.OnDelete(evt.Object)
	}
}

// pushEvent requires mu to be held by the cache mutation that emits the event.
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
		if err := ctx.Err(); err != nil {
			return err
		}
		i.mu.Lock()
		if !i.started {
			i.mu.Unlock()
			return fmt.Errorf("informer has not started")
		}
		sequence := i.getSequences[key]
		if sequence == nil {
			sequence = &refreshSequence{}
			i.getSequences[key] = sequence
		}
		batch := sequence.request()
		i.markDirty(key)
		i.mu.Unlock()
		select {
		case i.syncObjCh <- key:
		case <-ctx.Done():
			return ctx.Err()
		}
		waitCtx, cancel := context.WithTimeout(ctx, 2*i.syncPeriod)
		defer cancel()
		select {
		case <-batch.done:
			if batch.err != nil {
				return batch.err
			}
			return copyObject(obj, batch.object.DeepCopyObject().(object.Object))
		case <-waitCtx.Done():
			return waitCtx.Err()
		}
	} else if options.WaitRefreshCache {
		i.mu.Lock()
		i.markDirty(key)
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
		if err := ctx.Err(); err != nil {
			return err
		}
		i.mu.Lock()
		if !i.started {
			i.mu.Unlock()
			return fmt.Errorf("informer has not started")
		}
		batch := i.listSequence.request()
		i.dirty = true
		i.mu.Unlock()
		select {
		case i.syncCh <- struct{}{}:
		case <-ctx.Done():
			return ctx.Err()
		}
		waitCtx, cancel := context.WithTimeout(ctx, 2*i.syncPeriod)
		defer cancel()
		select {
		case <-batch.done:
			if batch.err != nil {
				return batch.err
			}
			for _, item := range batch.list.GetItems() {
				appendCachedObject(list, item)
			}
			return nil
		case <-waitCtx.Done():
			return waitCtx.Err()
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
		fetching:     make(map[object.ObjectKey]bool),
		getSequences: make(map[object.ObjectKey]*refreshSequence),
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
		return fmt.Errorf("cannot copy object of type %T into %T", src, dst)
	}
	dv.Elem().Set(sv.Elem())
	return nil
}
