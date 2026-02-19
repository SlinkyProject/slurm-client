// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"sync"
	"testing"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/ptr"

	v0041 "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func Test_informerCache_processObjects(t *testing.T) {
	f := &emptyClient{}
	type fields struct {
		reader        Reader
		objectType    object.ObjectType
		cache         map[object.ObjectKey]*cacheEntry
		started       bool
		hasSynced     bool
		eventCh       chan event.Event
		syncCh        chan struct{}
		syncObjCh     chan object.ObjectKey
		syncErrorList error
		syncErrorGet  map[object.ObjectKey]error
		handler       cache.ResourceEventHandler
		syncPeriod    time.Duration
	}
	type args struct {
		list object.ObjectList
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "V0041Node",
			fields: fields{
				reader:        f,
				objectType:    types.ObjectTypeV0041Node,
				cache:         make(map[object.ObjectKey]*cacheEntry),
				started:       false,
				hasSynced:     false,
				eventCh:       nil,
				syncCh:        nil,
				syncObjCh:     nil,
				syncErrorList: nil,
				syncErrorGet:  nil,
				handler:       nil,
				syncPeriod:    30 * time.Second,
			},
			args: args{
				list: &types.V0041NodeList{
					Items: []types.V0041Node{
						{
							V0041Node: v0041.V0041Node{
								Name: ptr.To("node-0"),
							},
						},
						{
							V0041Node: v0041.V0041Node{
								Name: ptr.To("node-1"),
							},
						},
					},
				},
			},
		},
		{
			name: "V0041JobInfo",
			fields: fields{
				reader:        f,
				objectType:    types.ObjectTypeV0041JobInfo,
				cache:         make(map[object.ObjectKey]*cacheEntry),
				started:       false,
				hasSynced:     false,
				eventCh:       nil,
				syncCh:        nil,
				syncObjCh:     nil,
				syncErrorList: nil,
				syncErrorGet:  nil,
				handler:       nil,
				syncPeriod:    30 * time.Second,
			},
			args: args{
				list: &types.V0041JobInfoList{
					Items: []types.V0041JobInfo{
						{
							V0041JobInfo: v0041.V0041JobInfo{
								JobId: ptr.To[int32](1),
							},
						},
						{
							V0041JobInfo: v0041.V0041JobInfo{
								JobId: ptr.To[int32](2),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &informerCache{
				reader:        tt.fields.reader,
				objectType:    tt.fields.objectType,
				mu:            sync.RWMutex{},
				cache:         tt.fields.cache,
				started:       tt.fields.started,
				hasSynced:     tt.fields.hasSynced,
				eventCh:       tt.fields.eventCh,
				syncCh:        tt.fields.syncCh,
				syncObjCh:     tt.fields.syncObjCh,
				syncErrorList: tt.fields.syncErrorList,
				syncErrorGet:  tt.fields.syncErrorGet,
				handler:       tt.fields.handler,
				syncPeriod:    tt.fields.syncPeriod,
			}
			i.processObjects(tt.args.list)
			if len(i.cache) != len(tt.args.list.GetItems()) {
				t.Errorf("len(cache) = %v, expected %v", len(i.cache), len(tt.args.list.GetItems()))
			}
		})
	}
}

func Test_informerCache_processObject(t *testing.T) {
	f := &emptyClient{}
	type fields struct {
		reader        Reader
		objectType    object.ObjectType
		cache         map[object.ObjectKey]*cacheEntry
		started       bool
		hasSynced     bool
		eventCh       chan event.Event
		syncCh        chan struct{}
		syncObjCh     chan object.ObjectKey
		syncErrorList error
		syncErrorGet  map[object.ObjectKey]error
		handler       cache.ResourceEventHandler
		syncPeriod    time.Duration
	}
	type args struct {
		obj object.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "V0041Node",
			fields: fields{
				reader:        f,
				objectType:    types.ObjectTypeV0041Node,
				cache:         make(map[object.ObjectKey]*cacheEntry),
				started:       false,
				hasSynced:     false,
				eventCh:       nil,
				syncCh:        nil,
				syncObjCh:     nil,
				syncErrorList: nil,
				syncErrorGet:  nil,
				handler:       nil,
				syncPeriod:    30 * time.Second,
			},
			args: args{
				obj: &types.V0041Node{
					V0041Node: v0041.V0041Node{
						Name: ptr.To("node-0"),
					},
				},
			},
		},
		{
			name: "V0041JobInfo",
			fields: fields{
				reader:        f,
				objectType:    types.ObjectTypeV0041JobInfo,
				cache:         make(map[object.ObjectKey]*cacheEntry),
				started:       false,
				hasSynced:     false,
				eventCh:       nil,
				syncCh:        nil,
				syncObjCh:     nil,
				syncErrorList: nil,
				syncErrorGet:  nil,
				handler:       nil,
				syncPeriod:    30 * time.Second,
			},
			args: args{
				obj: &types.V0041JobInfo{
					V0041JobInfo: v0041.V0041JobInfo{
						JobId: ptr.To[int32](1),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &informerCache{
				reader:        tt.fields.reader,
				objectType:    tt.fields.objectType,
				mu:            sync.RWMutex{},
				cache:         tt.fields.cache,
				started:       tt.fields.started,
				hasSynced:     tt.fields.hasSynced,
				eventCh:       tt.fields.eventCh,
				syncCh:        tt.fields.syncCh,
				syncObjCh:     tt.fields.syncObjCh,
				syncErrorList: tt.fields.syncErrorList,
				syncErrorGet:  tt.fields.syncErrorGet,
				handler:       tt.fields.handler,
				syncPeriod:    tt.fields.syncPeriod,
			}
			i.processObject(tt.args.obj)
			if len(i.cache) != 1 {
				t.Errorf("len(cache) = %v, expected 1", len(i.cache))
			}
		})
	}
}

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// Create implements Client.
func (f *emptyClient) Create(ctx context.Context, obj object.Object, req any, opts ...CreateOption) error {
	return nil
}

// Delete implements Client.
func (f *emptyClient) Delete(ctx context.Context, obj object.Object, opts ...DeleteOption) error {
	return nil
}

// Update implements Client.
func (f *emptyClient) Update(ctx context.Context, obj object.Object, req any, opts ...UpdateOption) error {
	return nil
}

// Get implements Client.
func (f *emptyClient) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...GetOption) error {
	return nil
}

// List implements Client.
func (f *emptyClient) List(ctx context.Context, list object.ObjectList, opts ...ListOption) error {
	return nil
}

// GetInformer implements Client.
func (f *emptyClient) GetInformer(objectType object.ObjectType) InformerCache {
	return nil
}

// GetServer implements Client.
func (f *emptyClient) GetServer() string {
	return ""
}

// GetServer implements Client.
func (f *emptyClient) SetServer(server string) {
}

// GetToken implements Client.
func (f *emptyClient) GetToken() string {
	return ""
}

// GetToken implements Client.
func (f *emptyClient) SetToken(token string) {
}

// Start implements Client.
func (f *emptyClient) Start(ctx context.Context) {
}

// Stop implements Client.
func (f *emptyClient) Stop() {
}

var _ Client = &emptyClient{}
