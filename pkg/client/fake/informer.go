// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"time"

	"k8s.io/client-go/tools/cache"

	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

var _ client.InformerCache = &fakeInformer{}

type fakeInformer struct {
	started    bool
	reader     client.Reader
	objectType object.ObjectType
	syncPeriod time.Duration
	handler    cache.ResourceEventHandler
	hasSynced  bool
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
	for {
		_, ok := <-stopCh
		if !ok {
			break
		}
	}
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
		objectType: objectType,
		reader:     reader,
		syncPeriod: syncPeriod,
	}
}
