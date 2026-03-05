// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"testing"

	"k8s.io/utils/ptr"

	v0041 "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func newInformerWithData(objectType object.ObjectType, cache map[object.ObjectKey]*cacheEntry) *informerCache {
	fakeReader := &emptyClient{}
	ic := newInformer(objectType, fakeReader, defaultSyncPeriod)
	i, ok := ic.(*informerCache)
	if !ok {
		panic("failed to convert to *informerCache")
	}
	i.cache = cache
	return i
}

func Test_informerCache_processObjects(t *testing.T) {
	tests := []struct {
		name       string
		objectType object.ObjectType
		cache      map[object.ObjectKey]*cacheEntry
		list       object.ObjectList
	}{
		{
			name:       "V0041Node",
			objectType: types.ObjectTypeV0041Node,
			cache:      make(map[object.ObjectKey]*cacheEntry),
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
		{
			name:       "V0041JobInfo",
			objectType: types.ObjectTypeV0041JobInfo,
			cache:      make(map[object.ObjectKey]*cacheEntry),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := newInformerWithData(tt.objectType, tt.cache)
			i.processObjects(tt.list)
			if len(i.cache) != len(tt.list.GetItems()) {
				t.Errorf("len(cache) = %v, expected %v", len(i.cache), len(tt.list.GetItems()))
			}
		})
	}
}

func Test_informerCache_processObject(t *testing.T) {
	tests := []struct {
		name       string
		objectType object.ObjectType
		cache      map[object.ObjectKey]*cacheEntry
		obj        object.Object
	}{
		{
			name:       "V0041Node",
			objectType: types.ObjectTypeV0041Node,
			cache:      make(map[object.ObjectKey]*cacheEntry),
			obj: &types.V0041Node{
				V0041Node: v0041.V0041Node{
					Name: ptr.To("node-0"),
				},
			},
		},
		{
			name:       "V0041JobInfo",
			objectType: types.ObjectTypeV0041JobInfo,
			cache:      make(map[object.ObjectKey]*cacheEntry),
			obj: &types.V0041JobInfo{
				V0041JobInfo: v0041.V0041JobInfo{
					JobId: ptr.To[int32](1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := newInformerWithData(tt.objectType, tt.cache)
			i.processObject(tt.obj)
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
