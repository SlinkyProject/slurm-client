// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/SlinkyProject/slurm-client/pkg/cache"
	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

var _ client.InformerCache = &fakeInformer{}

type fakeInformer struct {
	mu         sync.RWMutex
	indexes    map[string]client.IndexFunc
	started    bool
	reader     client.Reader
	objectType object.ObjectType
	syncPeriod time.Duration
	handler    cache.ResourceEventHandler
	hasSynced  bool
}

func (f *fakeInformer) AddIndex(name string, extract client.IndexFunc) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if name == "" || extract == nil {
		return fmt.Errorf("index name and extractor must be set")
	}
	if _, exists := f.indexes[name]; exists {
		return fmt.Errorf("index %q already exists", name)
	}
	f.indexes[name] = extract
	return nil
}

// ByIndex evaluates the index against the fake store, including writes made
// since registration. Production informers maintain the index incrementally.
func (f *fakeInformer) ByIndex(ctx context.Context, name, value string, list object.ObjectList) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if list.GetType() != f.objectType {
		return fmt.Errorf("index on %s cannot populate %s", f.objectType, list.GetType())
	}
	f.mu.RLock()
	extract, exists := f.indexes[name]
	f.mu.RUnlock()
	if !exists {
		return fmt.Errorf("index %q is not registered", name)
	}
	all := &indexObjects{objectType: f.objectType}
	if err := f.reader.List(ctx, all); err != nil {
		return err
	}
	for _, obj := range all.items {
		if slices.Contains(extract(obj), value) {
			list.AppendItem(obj.DeepCopyObject().(object.Object))
		}
	}
	return nil
}

type indexObjects struct {
	objectType object.ObjectType
	items      []object.Object
}

func (l *indexObjects) GetType() object.ObjectType   { return l.objectType }
func (l *indexObjects) GetItems() []object.Object    { return l.items }
func (l *indexObjects) AppendItem(obj object.Object) { l.items = append(l.items, obj) }
func (l *indexObjects) DeepCopyObject() object.RuntimeObject {
	out := &indexObjects{objectType: l.objectType}
	for _, obj := range l.items {
		out.AppendItem(obj.DeepCopyObject().(object.Object))
	}
	return out
}

// Get implements [client.InformerCache].
func (f *fakeInformer) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
	return f.reader.Get(ctx, key, obj, opts...)
}

// List implements [client.InformerCache].
func (f *fakeInformer) List(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
	return f.reader.List(ctx, list, opts...)
}

// Run implements [client.InformerCache].
func (f *fakeInformer) Run(stopCh <-chan struct{}) {
	f.started = true
	<-stopCh
	f.started = false
}

// HasStarted implements [client.InformerCache].
func (f *fakeInformer) HasStarted() bool {
	return f.started
}

// HasSynced implements [client.InformerCache].
func (f *fakeInformer) HasSynced() (bool, error) {
	return f.hasSynced, nil
}

// SetEventHandler implements [client.InformerCache].
func (f *fakeInformer) SetEventHandler(handler cache.ResourceEventHandler) {
	f.handler = handler
}

// UnsetEventHandler implements [client.InformerCache].
func (f *fakeInformer) UnsetEventHandler() {
	f.handler = nil
}

func newInformer(objectType object.ObjectType, reader client.Reader, syncPeriod time.Duration) client.InformerCache {
	return &fakeInformer{
		indexes:    make(map[string]client.IndexFunc),
		objectType: objectType,
		reader:     reader,
		syncPeriod: syncPeriod,
	}
}
