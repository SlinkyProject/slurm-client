// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"errors"
	"net/http"

	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/client/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

type fakeClient struct {
	cache map[object.ObjectType]map[object.ObjectKey]object.Object

	updateFn updateFunc

	server    string
	authToken string
}

type updateFunc func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error

var _ client.Client = &fakeClient{}
var FakeSecret = "slurm-token"
var FakeServer = "fakeserver"

// NewFakeClient creates a new fake client for testing.
func NewFakeClient() client.Client {
	return NewClientBuilder().Build()
}

// NewClientBuilder returns a new builder to create a fake client.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

// ClientBuilder builds a fake client.
type ClientBuilder struct {
	updateFn         updateFunc
	initLists        []object.ObjectList
	initObject       []object.Object
	interceptorFuncs *interceptor.Funcs
}

// WithObjects can be optionally used to initialize this fake client with object.Object(s).
func (f *ClientBuilder) WithObjects(initObjs ...object.Object) *ClientBuilder {
	f.initObject = append(f.initObject, initObjs...)
	return f
}

// WithLists can be optionally used to initialize this fake client with object.ObjectList(s).
func (f *ClientBuilder) WithLists(initLists ...object.ObjectList) *ClientBuilder {
	f.initLists = append(f.initLists, initLists...)
	return f
}

// WithUpdateFn configures the client with the server side update function.
// Mutations to function parameter object.Object are preserved in cache.
func (f *ClientBuilder) WithUpdateFn(updateFn updateFunc) *ClientBuilder {
	f.updateFn = updateFn
	return f
}

// WithInterceptorFuncs configures the client methods to be intercepted using the provided interceptor.Funcs.
func (f *ClientBuilder) WithInterceptorFuncs(interceptorFuncs interceptor.Funcs) *ClientBuilder {
	f.interceptorFuncs = &interceptorFuncs
	return f
}

// Build builds and returns a new fake client.
func (f *ClientBuilder) Build() client.Client {
	cache := make(map[object.ObjectType]map[object.ObjectKey]object.Object)

	for _, list := range f.initLists {
		objType := list.GetType()
		if cache[objType] == nil {
			cache[objType] = make(map[object.ObjectKey]object.Object)
		}
		for _, obj := range list.GetItems() {
			cache[objType][obj.GetKey()] = obj.DeepCopyObject()
		}
	}

	for _, obj := range f.initObject {
		objType := obj.GetType()
		if cache[objType] == nil {
			cache[objType] = make(map[object.ObjectKey]object.Object)
		}
		cache[objType][obj.GetKey()] = obj.DeepCopyObject()
	}

	var result client.Client = &fakeClient{
		updateFn:  f.updateFn,
		server:    FakeServer,
		authToken: FakeSecret,
		cache:     cache,
	}

	if f.interceptorFuncs != nil {
		result = interceptor.NewClient(result, *f.interceptorFuncs)
	}

	return result
}

func (c *fakeClient) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
	entry, exists := c.cache[obj.GetType()][key]
	if !exists {
		return errors.New(http.StatusText(http.StatusNotFound))
	}
	switch o := obj.(type) {
	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0041ControllerPing:
		cache := entry.(*types.V0041ControllerPing)
		*o = *cache
	case *types.V0041JobInfo:
		cache := entry.(*types.V0041JobInfo)
		*o = *cache
	case *types.V0041Node:
		cache := entry.(*types.V0041Node)
		*o = *cache
	case *types.V0041PartitionInfo:
		cache := entry.(*types.V0041PartitionInfo)
		*o = *cache
	case *types.V0041Stats:
		cache := entry.(*types.V0041Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0042ControllerPing:
		cache := entry.(*types.V0042ControllerPing)
		*o = *cache
	case *types.V0042JobInfo:
		cache := entry.(*types.V0042JobInfo)
		*o = *cache
	case *types.V0042Node:
		cache := entry.(*types.V0042Node)
		*o = *cache
	case *types.V0042PartitionInfo:
		cache := entry.(*types.V0042PartitionInfo)
		*o = *cache
	case *types.V0042Stats:
		cache := entry.(*types.V0042Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	case *types.V0043ControllerPing:
		cache := entry.(*types.V0043ControllerPing)
		*o = *cache
	case *types.V0043JobInfo:
		cache := entry.(*types.V0043JobInfo)
		*o = *cache
	case *types.V0043Node:
		cache := entry.(*types.V0043Node)
		*o = *cache
	case *types.V0043PartitionInfo:
		cache := entry.(*types.V0043PartitionInfo)
		*o = *cache
	case *types.V0043Stats:
		cache := entry.(*types.V0043Stats)
		*o = *cache

	/////////////////////////////////////////////////////////////////////////////////

	default:
		return errors.New(http.StatusText(http.StatusNotImplemented))
	}
	return nil
}

func (c *fakeClient) List(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
	for _, entry := range c.cache[list.GetType()] {
		list.AppendItem(entry)
	}
	return nil
}

func (c *fakeClient) Create(ctx context.Context, obj object.Object, req any, opts ...client.CreateOption) error {
	t := obj.GetType()
	k := obj.GetKey()
	_, exists := c.cache[t][k]
	if exists {
		return errors.New(http.StatusText(http.StatusConflict))
	}
	if _, ok := c.cache[t]; !ok {
		c.cache[t] = make(map[object.ObjectKey]object.Object)
	}
	c.cache[t][k] = obj.DeepCopyObject()
	return nil
}

func (c *fakeClient) Delete(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
	t := obj.GetType()
	k := obj.GetKey()
	if _, ok := c.cache[t][k]; !ok {
		return errors.New(http.StatusText(http.StatusNotFound))
	}
	delete(c.cache[t], k)
	return nil
}

func (c *fakeClient) Update(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
	t := obj.GetType()
	k := obj.GetKey()
	if _, ok := c.cache[t][k]; !ok {
		return errors.New(http.StatusText(http.StatusNotFound))
	}
	if _, ok := c.cache[t]; !ok {
		c.cache[t] = make(map[object.ObjectKey]object.Object)
	}
	if c.updateFn != nil {
		if err := c.updateFn(ctx, obj, req, opts...); err != nil {
			return err
		}
	}
	c.cache[t][k] = obj.DeepCopyObject()
	return nil
}

func (c *fakeClient) GetInformer(obj object.ObjectType) client.InformerCache {
	return nil
}

func (c *fakeClient) GetServer() string {
	return c.server
}

func (c *fakeClient) SetServer(server string) {
	c.server = server
}

func (c *fakeClient) GetToken() string {
	return c.authToken
}

func (c *fakeClient) SetToken(token string) {
	c.authToken = token
}

func (c *fakeClient) Start(ctx context.Context) {
}

func (c *fakeClient) Stop() {
}
