// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/set"

	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

const (
	defaultSyncPeriod = 30 * time.Second
	waitSyncPeriod    = 1 * time.Second
)

type cacheEntry struct {
	lastUpdate time.Time
	object     object.Object
}

type informerCache struct {
	// reader knows how to read from remote.
	reader Reader

	// writer knows how to write to the remote.
	writer Writer

	// objectType tracks the object type this informer backs.
	objectType object.ObjectType

	// mu guards access to the map.
	mu sync.RWMutex

	// cache holds the actual object cache.
	cache map[object.ObjectKey]*cacheEntry

	// started is true if the informers have been started.
	started bool

	// hasSynced is true if the informers have run at least once and the cache is not dirty.
	hasSynced bool

	// eventCh holds events for handler.
	eventCh chan event.Event

	// syncCh holds sync requests.
	syncCh chan bool

	// syncObjCh holds sync requests by ObjectKey.
	syncObjCh chan object.ObjectKey

	// syncError is the last List sync error
	syncErrorList error

	// syncErrorGet is the last Get sync error
	syncErrorGet error

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

	go i.runInformer(stopCh)
	go i.runGetInformer(stopCh)
	go i.runHandler(stopCh)
}

func (i *informerCache) runInformer(stopCh <-chan struct{}) {
	i.mu.RLock()
	ticker := time.NewTicker(i.syncPeriod)
	defer ticker.Stop()
	i.mu.RUnlock()

	for {
		var list object.ObjectList
		switch i.objectType {
		case slurmtypes.ObjectTypeNode:
			list = &slurmtypes.NodeList{}
		case slurmtypes.ObjectTypePing:
			list = &slurmtypes.PingList{}
		case slurmtypes.ObjectTypeJobInfo:
			list = &slurmtypes.JobInfoList{}
		case slurmtypes.ObjectTypePartitionInfo:
			list = &slurmtypes.PartitionInfoList{}
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
			i.hasSynced = true
		}
		i.mu.Unlock()

		select {
		case <-i.syncCh:
			// wait for sync request
			i.mu.Lock()
			i.hasSynced = false
			i.mu.Unlock()
		case <-ticker.C:
			// wait for tick
		case <-stopCh:
			i.mu.Lock()
			defer i.mu.Unlock()
			i.started = false
			return
		}
	}
}

func (i *informerCache) runGetInformer(stopCh <-chan struct{}) {
	for {
		var key object.ObjectKey
		select {
		case key = <-i.syncObjCh:
			// wait for sync request
			i.mu.Lock()
			i.hasSynced = false
			i.mu.Unlock()
		case <-stopCh:
			return
		}

		var obj object.Object
		switch i.objectType {
		case slurmtypes.ObjectTypeNode:
			obj = &slurmtypes.Node{}
		case slurmtypes.ObjectTypePing:
			obj = &slurmtypes.Ping{}
		case slurmtypes.ObjectTypeJobInfo:
			obj = &slurmtypes.JobInfo{}
		case slurmtypes.ObjectTypePartitionInfo:
			obj = &slurmtypes.PartitionInfo{}
		default:
			// NOTE: We must handle every Slurm type otherwise panic.
			// We cannot recover from here because the informer has started a
			// number of go-routines that must all start and stop together.
			panic("unhandled object type")
		}

		opts := &GetOptions{SkipCache: true}
		err := i.reader.Get(context.TODO(), key, obj, opts)

		i.mu.Lock()
		i.syncErrorGet = err
		if err == nil {
			i.processObject(obj)
			i.hasSynced = true
		}
		i.mu.Unlock()
	}
}

func (i *informerCache) runHandler(stopCh <-chan struct{}) {
	for {
		select {
		case e := <-i.eventCh:
			if i.handler == nil {
				continue
			}
			switch e.Type {
			case event.Added:
				i.handler.OnAdd(e.Object, i.hasSynced)
			case event.Modified:
				i.handler.OnUpdate(e.Object, e.ObjectOld)
			case event.Deleted:
				i.handler.OnDelete(e.Object)
			}
		case <-stopCh:
			i.mu.Lock()
			defer i.mu.Unlock()
			i.started = false
			return
		}
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
		if !ok {
			insert = true
			e.Type = event.Added
		} else if ok && !now.Before(entry.lastUpdate) && !entry.object.DeepEqualObject(item) {
			insert = true
			e.Type = event.Modified
			e.ObjectOld = entry.object.DeepCopyObject()
		}

		if insert {
			i.cache[key] = &cacheEntry{
				lastUpdate: now,
				object:     item,
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
	if !ok {
		insert = true
		e.Type = event.Added
	} else if ok && !now.Before(entry.lastUpdate) && !entry.object.DeepEqualObject(obj) {
		insert = true
		e.Type = event.Modified
		e.ObjectOld = entry.object.DeepCopyObject()
	}

	if insert {
		i.cache[key] = &cacheEntry{
			lastUpdate: now,
			object:     obj,
		}
		e.Object = obj.DeepCopyObject()
		i.pushEvent(e)
	}
}

// HasSynced implements InformerCache.
func (i *informerCache) HasSynced() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.hasSynced, i.syncErrorList
}

func (i *informerCache) hasSyncedList() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.hasSynced, i.syncErrorList
}

func (i *informerCache) hasSyncedGet() (bool, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.hasSynced, i.syncErrorGet
}

// HasStarted implements InformerCache.
func (i *informerCache) HasStarted() bool {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.started
}

// WaitForSyncList implements InformerCache.
func (i *informerCache) WaitForSyncList(ctx context.Context, interval time.Duration) error {
	err := wait.PollUntilContextCancel(ctx, interval, true,
		func(_ context.Context) (bool, error) {
			return i.hasSyncedList()
		})
	return err
}

// WaitForSyncGet implements InformerCache.
func (i *informerCache) WaitForSyncGet(ctx context.Context, interval time.Duration) error {
	err := wait.PollUntilContextCancel(ctx, interval, true,
		func(_ context.Context) (bool, error) {
			return i.hasSyncedGet()
		})
	return err
}

// Create implements Client.
func (i *informerCache) Create(
	ctx context.Context,
	obj object.Object,
	opts ...CreateOption,
) error {
	i.mu.Lock()
	i.hasSynced = false
	i.mu.Unlock()
	i.syncCh <- true

	err := i.writer.Create(ctx, obj, opts...)
	i.syncCh <- true
	return err
}

// Delete implements Client.
func (i *informerCache) Delete(
	ctx context.Context,
	obj object.Object,
	opts ...DeleteOption,
) error {
	i.mu.Lock()
	i.hasSynced = false
	i.mu.Unlock()
	i.syncCh <- true

	err := i.writer.Delete(ctx, obj, opts...)
	i.syncCh <- true
	return err
}

// DeleteAllOf implements Client.
func (i *informerCache) DeleteAllOf(
	ctx context.Context,
	obj object.Object,
	opts ...DeleteAllOfOption,
) error {
	i.mu.Lock()
	i.hasSynced = false
	i.mu.Unlock()
	i.syncCh <- true

	err := i.writer.DeleteAllOf(ctx, obj, opts...)
	i.syncCh <- true
	return err
}

// Get implements InformerCache.
func (i *informerCache) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...GetOption) error {
	options := &GetOptions{}
	options.ApplyOptions(opts)

	if options.RefreshCache {
		i.mu.Lock()
		i.hasSynced = false
		i.mu.Unlock()
		i.syncObjCh <- key
	}

	if err := i.WaitForSyncGet(ctx, waitSyncPeriod); err != nil {
		return err
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	entry, ok := i.cache[key]
	if !ok {
		return errors.New(http.StatusText(http.StatusNotFound))
	}

	switch o := obj.(type) {
	case *slurmtypes.Node:
		cache := entry.object.(*slurmtypes.Node)
		*o = *cache
	case *slurmtypes.Ping:
		cache := entry.object.(*slurmtypes.Ping)
		*o = *cache
	case *slurmtypes.JobInfo:
		cache := entry.object.(*slurmtypes.JobInfo)
		*o = *cache
	case *slurmtypes.PartitionInfo:
		cache := entry.object.(*slurmtypes.PartitionInfo)
		*o = *cache
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
		i.mu.Lock()
		i.hasSynced = false
		i.mu.Unlock()
		i.syncCh <- true
	}

	if err := i.WaitForSyncList(ctx, waitSyncPeriod); err != nil {
		return err
	}

	i.mu.RLock()
	defer i.mu.RUnlock()

	for _, entry := range i.cache {
		list.AppendItem(entry.object)
	}

	return nil
}

// Update implements InformerCache.
func (i *informerCache) Update(
	ctx context.Context,
	obj object.Object,
	opts ...UpdateOption,
) error {
	i.mu.Lock()
	i.hasSynced = false
	i.mu.Unlock()
	i.syncCh <- true

	err := i.writer.Update(ctx, obj, opts...)
	i.syncObjCh <- obj.GetKey()
	return err
}
