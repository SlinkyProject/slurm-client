// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/client/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type fakeClient struct {
	cache map[object.ObjectType]map[object.ObjectKey]object.Object

	updateFn updateFunc

	server        string
	authToken     string
	tokenProvider token.Provider

	ctx    context.Context
	cancel context.CancelFunc
}

type updateFunc func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error

var _ client.Client = &fakeClient{}
var FakeSecret = "slurm-token"
var FakeServer = "fakeserver"

// NewFakeClient creates a new fake client for testing.
func NewFakeClient(initObjs ...object.RuntimeObject) client.Client {
	return NewClientBuilder().WithRuntimeObjects(initObjs...).Build()
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

// WithRuntimeObjects initializes this fake client with object.RuntimeObject(s).
func (f *ClientBuilder) WithRuntimeObjects(initObjs ...object.RuntimeObject) *ClientBuilder {
	for _, initObj := range initObjs {
		switch obj := initObj.(type) {
		case object.Object:
			f.WithObjects(obj)
		case object.ObjectList:
			f.WithLists(obj)
		default:
			panic(fmt.Sprintf("ClientBuilder: unsupported type %T", initObj))
		}
	}
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
		for _, obj := range list.GetItems() {
			store(cache, obj)
		}
	}
	for _, obj := range f.initObject {
		store(cache, obj)
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
		return apierrors.ErrNotFound
	}
	copyInto(obj, entry)
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
	store(c.cache, obj)
	return nil
}

func (c *fakeClient) Delete(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
	t := obj.GetType()
	k := obj.GetKey()
	if _, ok := c.cache[t][k]; !ok {
		return apierrors.ErrNotFound
	}
	delete(c.cache[t], k)
	return nil
}

func (c *fakeClient) Update(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
	t := obj.GetType()
	k := obj.GetKey()
	if _, ok := c.cache[t][k]; !ok {
		return apierrors.ErrNotFound
	}
	if c.updateFn != nil {
		if err := c.updateFn(ctx, obj, req, opts...); err != nil {
			return err
		}
	}
	store(c.cache, obj)
	return nil
}

func (c *fakeClient) GetInformer(obj object.ObjectType) client.InformerCache {
	return newInformer(obj, c, client.DefaultWatchInterval)
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

func (c *fakeClient) SetTokenProvider(tokenProvider token.Provider) {
	c.tokenProvider = tokenProvider
	c.authToken = ""
}

func (c *fakeClient) Start(ctx context.Context) {
	if c.ctx != nil {
		return
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	<-c.ctx.Done()
	c.ctx = nil
}

func (c *fakeClient) Stop() {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
}

func store(cache map[object.ObjectType]map[object.ObjectKey]object.Object, obj object.Object) {
	t := obj.GetType()
	if cache[t] == nil {
		cache[t] = make(map[object.ObjectKey]object.Object)
	}
	cache[t][obj.GetKey()] = obj.DeepCopyObject().(object.Object)
}

func copyInto(dst, src object.Object) {
	reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(src.DeepCopyObject()).Elem())
}

func (c *fakeClient) Versioned() client.VersionedClients {
	return client.VersionedClients{}
}
