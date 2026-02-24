// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/set"

	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

const (
	defaultSyncPeriod = 30 * time.Second
	waitSyncPeriod    = 1 * time.Second
	batchPeriod       = 1 * time.Second
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
	cache map[object.ObjectKey]*cacheEntry

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

	for {
		_, ok := <-stopCh
		if !ok {
			i.mu.Lock()
			i.started = false
			i.mu.Unlock()
			break
		}
	}
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

	var list object.ObjectList
	switch i.objectType {
	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0041ControllerPing:
		list = &types.V0041ControllerPingList{}
	case types.ObjectTypeV0041JobInfo:
		list = &types.V0041JobInfoList{}
	case types.ObjectTypeV0041Node:
		list = &types.V0041NodeList{}
	case types.ObjectTypeV0041PartitionInfo:
		list = &types.V0041PartitionInfoList{}
	case types.ObjectTypeV0041Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0041Stats:
		list = &types.V0041StatsList{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0042ControllerPing:
		list = &types.V0042ControllerPingList{}
	case types.ObjectTypeV0042JobInfo:
		list = &types.V0042JobInfoList{}
	case types.ObjectTypeV0042Node:
		list = &types.V0042NodeList{}
	case types.ObjectTypeV0042PartitionInfo:
		list = &types.V0042PartitionInfoList{}
	case types.ObjectTypeV0042Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0042Stats:
		list = &types.V0042StatsList{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0043ControllerPing:
		list = &types.V0043ControllerPingList{}
	case types.ObjectTypeV0043JobInfo:
		list = &types.V0043JobInfoList{}
	case types.ObjectTypeV0043Node:
		list = &types.V0043NodeList{}
	case types.ObjectTypeV0043PartitionInfo:
		list = &types.V0043PartitionInfoList{}
	case types.ObjectTypeV0043Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0043Stats:
		list = &types.V0043StatsList{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0044ControllerPing:
		list = &types.V0044ControllerPingList{}
	case types.ObjectTypeV0044JobInfo:
		list = &types.V0044JobInfoList{}
	case types.ObjectTypeV0044Node:
		list = &types.V0044NodeList{}
	case types.ObjectTypeV0044PartitionInfo:
		list = &types.V0044PartitionInfoList{}
	case types.ObjectTypeV0044Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0044ReservationInfo:
		list = &types.V0044ReservationInfoList{}
	case types.ObjectTypeV0044NodeResourceLayout:
		panic("NodeResouceLayout is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0044Stats:
		list = &types.V0044StatsList{}

	/////////////////////////////////////////////////////////////////////////////////

	default:
		// NOTE: We must handle every Slurm type otherwise panic.
		// We cannot recover from here because the informer has started a
		// number of go-routines that must all start and stop together.
		panic(errors.New(http.StatusText(http.StatusNotImplemented)))
	}

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

	var obj object.Object
	switch i.objectType {
	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0041ControllerPing:
		obj = &types.V0041ControllerPing{}
	case types.ObjectTypeV0041JobInfo:
		obj = &types.V0041JobInfo{}
	case types.ObjectTypeV0041Node:
		obj = &types.V0041Node{}
	case types.ObjectTypeV0041PartitionInfo:
		obj = &types.V0041PartitionInfo{}
	case types.ObjectTypeV0041Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0041Stats:
		obj = &types.V0041Stats{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0042ControllerPing:
		obj = &types.V0042ControllerPing{}
	case types.ObjectTypeV0042JobInfo:
		obj = &types.V0042JobInfo{}
	case types.ObjectTypeV0042Node:
		obj = &types.V0042Node{}
	case types.ObjectTypeV0042PartitionInfo:
		obj = &types.V0042PartitionInfo{}
	case types.ObjectTypeV0042Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0042Stats:
		obj = &types.V0042Stats{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0043ControllerPing:
		obj = &types.V0043ControllerPing{}
	case types.ObjectTypeV0043JobInfo:
		obj = &types.V0043JobInfo{}
	case types.ObjectTypeV0043Node:
		obj = &types.V0043Node{}
	case types.ObjectTypeV0043PartitionInfo:
		obj = &types.V0043PartitionInfo{}
	case types.ObjectTypeV0043Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0043Stats:
		obj = &types.V0043Stats{}

	/////////////////////////////////////////////////////////////////////////////////

	case types.ObjectTypeV0044ControllerPing:
		obj = &types.V0044ControllerPing{}
	case types.ObjectTypeV0044JobInfo:
		obj = &types.V0044JobInfo{}
	case types.ObjectTypeV0044Node:
		obj = &types.V0044Node{}
	case types.ObjectTypeV0044PartitionInfo:
		obj = &types.V0044PartitionInfo{}
	case types.ObjectTypeV0044Reconfigure:
		panic("Reconfigure is not supported, this scenario should have been avoided.")
	case types.ObjectTypeV0044ReservationInfo:
		obj = &types.V0044ReservationInfo{}
	case types.ObjectTypeV0044Stats:
		obj = &types.V0044Stats{}

	/////////////////////////////////////////////////////////////////////////////////

	default:
		// NOTE: We must handle every Slurm type otherwise panic.
		// We cannot recover from here because the informer has started a
		// number of go-routines that must all start and stop together.
		panic("unhandled object type")
	}

	opts := &GetOptions{SkipCache: true}
	err := i.reader.Get(context.TODO(), key, obj, opts)

	i.mu.Lock()
	if err != nil && err.Error() != http.StatusText(http.StatusNotFound) {
		i.syncErrorGet[key] = err
	} else {
		i.syncErrorGet[key] = nil
	}
	if i.syncErrorGet[key] == nil {
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
		i.handler.OnUpdate(evt.Object, evt.ObjectOld)
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
	fresh := make(set.Set[object.ObjectKey])
	for _, item := range list.GetItems() {
		key := item.GetKey()
		fresh.Insert(key)
		insert := false

		e := event.Event{}
		entry, ok := i.cache[key]
		if !ok || entry.object == nil {
			insert = true
			e.Type = event.Added
		} else if ok && entry.object != nil && !now.Before(entry.lastUpdate) && !reflect.DeepEqual(entry.object, item) {
			insert = true
			e.Type = event.Modified
			e.ObjectOld = entry.object.DeepCopyObject()
		}

		if insert {
			i.cache[key] = &cacheEntry{
				lastUpdate: now,
				object:     item,
				dirty:      false,
			}
			e.Object = item.DeepCopyObject()
			i.pushEvent(e)
		}
	}

	for _, entry := range i.cache {
		key := entry.object.GetKey()
		if !fresh.Has(key) {
			e := event.Event{
				Type:   event.Deleted,
				Object: entry.object.DeepCopyObject(),
			}
			delete(i.cache, key)
			i.pushEvent(e)
		}
	}
}

func (i *informerCache) processObject(obj object.Object) {
	now := time.Now()
	key := obj.GetKey()
	insert := false

	e := event.Event{}
	entry, ok := i.cache[key]
	if !ok || entry.object == nil {
		insert = true
		e.Type = event.Added
	} else if ok && entry.object != nil && !now.Before(entry.lastUpdate) && !reflect.DeepEqual(entry.object, obj) {
		insert = true
		e.Type = event.Modified
		e.ObjectOld = entry.object.DeepCopyObject()
	}

	if insert {
		i.cache[key] = &cacheEntry{
			lastUpdate: now,
			object:     obj,
			dirty:      false,
		}
		e.Object = obj.DeepCopyObject()
		i.pushEvent(e)
	}
}

// HasSynced implements InformerCache.
func (i *informerCache) HasSynced() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return !i.dirty, i.syncErrorList
}

func (i *informerCache) hasSyncedList() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return !i.dirty, i.syncErrorList
}

func (i *informerCache) hasSyncedGet(key object.ObjectKey) (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
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

// WaitForSyncList implements InformerCache.
func (i *informerCache) WaitForSyncList(ctx context.Context, interval time.Duration) error {
	err := wait.PollUntilContextTimeout(ctx, interval, i.syncPeriod, true,
		func(_ context.Context) (bool, error) {
			return i.hasSyncedList()
		})
	return err
}

// WaitForSyncGet implements InformerCache.
func (i *informerCache) WaitForSyncGet(ctx context.Context, key object.ObjectKey, interval time.Duration) error {
	err := wait.PollUntilContextTimeout(ctx, interval, i.syncPeriod, true,
		func(_ context.Context) (bool, error) {
			return i.hasSyncedGet(key)
		})
	return err
}

// Get implements InformerCache.
func (i *informerCache) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...GetOption) error {
	options := &GetOptions{}
	options.ApplyOptions(opts)

	if options.RefreshCache {
		// Mark dirty before sync request to avoid lock race between channel
		// receiver and WaitForSyncGet().
		i.mu.Lock()
		if obj := i.cache[key]; obj != nil {
			obj.dirty = true
		} else {
			i.cache[key] = &cacheEntry{dirty: true}
		}
		i.syncObjCh <- key
		i.mu.Unlock()
	} else if options.WaitRefreshCache {
		i.mu.Lock()
		if obj := i.cache[key]; obj != nil {
			obj.dirty = true
		} else {
			i.cache[key] = &cacheEntry{dirty: true}
		}
		i.mu.Unlock()
	}

	if err := i.WaitForSyncGet(ctx, key, waitSyncPeriod); err != nil {
		return fmt.Errorf("failed to wait on type %s object %s cache sync: %w", obj.GetType(), key, err)
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	entry, ok := i.cache[key]
	if !ok || entry.object == nil {
		return errors.New(http.StatusText(http.StatusNotFound))
	}

	switch o := obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0041ControllerPing:
		cache := entry.object.(*types.V0041ControllerPing)
		*o = *cache
	case *types.V0041JobInfo:
		cache := entry.object.(*types.V0041JobInfo)
		*o = *cache
	case *types.V0041Node:
		cache := entry.object.(*types.V0041Node)
		*o = *cache
	case *types.V0041PartitionInfo:
		cache := entry.object.(*types.V0041PartitionInfo)
		*o = *cache
	case *types.V0041Stats:
		cache := entry.object.(*types.V0041Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042ControllerPing:
		cache := entry.object.(*types.V0042ControllerPing)
		*o = *cache
	case *types.V0042JobInfo:
		cache := entry.object.(*types.V0042JobInfo)
		*o = *cache
	case *types.V0042Node:
		cache := entry.object.(*types.V0042Node)
		*o = *cache
	case *types.V0042PartitionInfo:
		cache := entry.object.(*types.V0042PartitionInfo)
		*o = *cache
	case *types.V0042Stats:
		cache := entry.object.(*types.V0042Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043ControllerPing:
		cache := entry.object.(*types.V0043ControllerPing)
		*o = *cache
	case *types.V0043JobInfo:
		cache := entry.object.(*types.V0043JobInfo)
		*o = *cache
	case *types.V0043Node:
		cache := entry.object.(*types.V0043Node)
		*o = *cache
	case *types.V0043PartitionInfo:
		cache := entry.object.(*types.V0043PartitionInfo)
		*o = *cache
	case *types.V0043Stats:
		cache := entry.object.(*types.V0043Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0044ControllerPing:
		cache := entry.object.(*types.V0044ControllerPing)
		*o = *cache
	case *types.V0044JobInfo:
		cache := entry.object.(*types.V0044JobInfo)
		*o = *cache
	case *types.V0044Node:
		cache := entry.object.(*types.V0044Node)
		*o = *cache
	case *types.V0044PartitionInfo:
		cache := entry.object.(*types.V0044PartitionInfo)
		*o = *cache
	case *types.V0044ReservationInfo:
		cache := entry.object.(*types.V0044ReservationInfo)
		*o = *cache
	case *types.V0044Stats:
		cache := entry.object.(*types.V0044Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}

	return nil
}

// List implements InformerCache.
func (i *informerCache) List(ctx context.Context, list object.ObjectList, opts ...ListOption) error {
	options := &ListOptions{}
	options.ApplyOptions(opts)

	if options.RefreshCache {
		// Mark dirty before sync request to avoid lock race between channel
		// receiver and WaitForSyncList().
		i.mu.Lock()
		i.dirty = true
		i.syncCh <- struct{}{}
		i.mu.Unlock()
	} else if options.WaitRefreshCache {
		i.mu.Lock()
		i.dirty = true
		i.mu.Unlock()
	}

	if err := i.WaitForSyncList(ctx, waitSyncPeriod); err != nil {
		return fmt.Errorf("failed to wait on type %s cache sync: %w", list.GetType(), err)
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	for _, entry := range i.cache {
		if entry.object == nil {
			continue
		}
		list.AppendItem(entry.object)
	}

	return nil
}
