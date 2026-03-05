// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
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
	node0 := &types.V0044Node{
		V0044Node: api.V0044Node{
			Name: ptr.To("node-0"),
		},
	}
	now := time.Now()
	type testCase struct {
		name          string
		objectType    object.ObjectType
		cache         map[object.ObjectKey]*cacheEntry
		list          object.ObjectList
		wantCacheLen  int
		wantAddCnt    int64
		wantModifyCnt int64
		wantDeleteCnt int64
	}
	tests := []testCase{
		func() testCase {
			return testCase{
				name:       "add",
				objectType: node0.GetType(),
				cache:      make(map[object.ObjectKey]*cacheEntry),
				list: &types.V0044NodeList{
					Items: []types.V0044Node{
						*node0.DeepCopy(),
					},
				},
				wantCacheLen: 1,
				wantAddCnt:   1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "modify",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						lastUpdate: time.Now(),
						object: func() object.Object {
							node := node0.DeepCopy()
							node.Comment = ptr.To("foo")
							return node
						}(),
						dirty: false,
					},
				},
				list: &types.V0044NodeList{
					Items: []types.V0044Node{
						*node0.DeepCopy(),
					},
				},
				wantCacheLen:  1,
				wantModifyCnt: 1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "no modify",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						lastUpdate: time.Now(),
						object:     node0.DeepCopy(),
						dirty:      false,
					},
				},
				list: &types.V0044NodeList{
					Items: []types.V0044Node{
						*node0.DeepCopy(),
					},
				},
				wantCacheLen: 1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "delete",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						lastUpdate: time.Now(),
						object:     node0.DeepCopy(),
						dirty:      false,
					},
				},
				list:          &types.V0044NodeList{},
				wantCacheLen:  0,
				wantDeleteCnt: 1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "delete dirty, no object",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						dirty: true,
					},
				},
				list: &types.V0044NodeList{
					Items: []types.V0044Node{},
				},
				wantCacheLen: 0,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stopCh := make(chan struct{})
			var addCnt int64
			var modifyCnt int64
			var deleteCnt int64
			i := newInformerWithData(tt.objectType, tt.cache)
			i.SetEventHandler(cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj any) {
					atomic.AddInt64(&addCnt, 1)
				},
				UpdateFunc: func(oldObj, newObj any) {
					atomic.AddInt64(&modifyCnt, 1)
				},
				DeleteFunc: func(obj any) {
					atomic.AddInt64(&deleteCnt, 1)
				},
			})
			go i.runHandler(stopCh)
			i.processObjects(tt.list)
			for {
				time.Sleep(time.Second)
				if len(i.eventCh) == 0 {
					close(stopCh)
					break
				}
			}
			if len(i.cache) != len(tt.list.GetItems()) {
				t.Errorf("len(cache) = %v, want %v", len(i.cache), tt.wantCacheLen)
			}
			if addCnt != tt.wantAddCnt {
				t.Errorf("processObjects() add events = %v, want %v", addCnt, tt.wantAddCnt)
			}
			if modifyCnt != tt.wantModifyCnt {
				t.Errorf("processObjects() modify events = %v, want %v", modifyCnt, tt.wantModifyCnt)
			}
			if deleteCnt != tt.wantDeleteCnt {
				t.Errorf("processObjects() delete events = %v, want %v", deleteCnt, tt.wantDeleteCnt)
			}
			for _, entry := range i.cache {
				if now.After(entry.lastUpdate) {
					t.Errorf("entry %v : now %v not after lastUpdate %v", entry.object.GetKey(), now, entry.lastUpdate)
				}
			}
		})
	}
}

func Test_informerCache_processObject(t *testing.T) {
	node0 := &types.V0044Node{
		V0044Node: api.V0044Node{
			Name: ptr.To("node-0"),
		},
	}
	type testCase struct {
		name          string
		objectType    object.ObjectType
		cache         map[object.ObjectKey]*cacheEntry
		obj           object.Object
		wantCacheLen  int
		wantAddCnt    int64
		wantModifyCnt int64
		wantDeleteCnt int64
	}
	tests := []testCase{
		func() testCase {
			return testCase{
				name:         "add",
				objectType:   node0.GetType(),
				cache:        make(map[object.ObjectKey]*cacheEntry),
				obj:          node0.DeepCopy(),
				wantCacheLen: 1,
				wantAddCnt:   1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "no modify",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						lastUpdate: time.Now(),
						object:     node0.DeepCopy(),
						dirty:      false,
					},
				},
				obj:          node0.DeepCopy(),
				wantCacheLen: 1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "modify",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						lastUpdate: time.Now(),
						object: func() object.Object {
							node := node0.DeepCopy()
							node.Comment = ptr.To("foo")
							return node
						}(),
						dirty: false,
					},
				},
				obj:           node0.DeepCopy(),
				wantCacheLen:  1,
				wantModifyCnt: 1,
			}
		}(),
		func() testCase {
			return testCase{
				name:       "add, dirty, no object",
				objectType: node0.GetType(),
				cache: map[object.ObjectKey]*cacheEntry{
					node0.GetKey(): {
						dirty: true,
					},
				},
				obj:          node0.DeepCopy(),
				wantCacheLen: 1,
				wantAddCnt:   1,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stopCh := make(chan struct{})
			var addCnt int64
			var modifyCnt int64
			var deleteCnt int64
			i := newInformerWithData(tt.objectType, tt.cache)
			i.SetEventHandler(cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj any) {
					atomic.AddInt64(&addCnt, 1)
				},
				UpdateFunc: func(oldObj, newObj any) {
					atomic.AddInt64(&modifyCnt, 1)
				},
				DeleteFunc: func(obj any) {
					atomic.AddInt64(&deleteCnt, 1)
				},
			})
			go i.runHandler(stopCh)
			i.processObject(tt.obj)
			for {
				time.Sleep(time.Second)
				if len(i.eventCh) == 0 {
					close(stopCh)
					break
				}
			}
			if len(i.cache) != tt.wantCacheLen {
				t.Errorf("len(cache) = %v, want %v", len(i.cache), tt.wantCacheLen)
			}
			if addCnt != tt.wantAddCnt {
				t.Errorf("processObject() add events = %v, want %v", addCnt, tt.wantAddCnt)
			}
			if modifyCnt != tt.wantModifyCnt {
				t.Errorf("processObject() modify events = %v, want %v", modifyCnt, tt.wantModifyCnt)
			}
			if deleteCnt != tt.wantDeleteCnt {
				t.Errorf("processObject() delete events = %v, want %v", deleteCnt, tt.wantDeleteCnt)
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
