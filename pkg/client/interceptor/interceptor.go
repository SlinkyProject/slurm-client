// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"

	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// Funcs contains functions that are called instead of the underlying client's methods.
type Funcs struct {
	Get         func(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error
	List        func(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error
	Create      func(ctx context.Context, obj object.Object, opts ...client.CreateOption) error
	Delete      func(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error
	DeleteAllOf func(ctx context.Context, obj object.Object, opts ...client.DeleteAllOfOption) error
	Update      func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error
	GetInformer func(obj object.ObjectType) client.InformerCache
	GetServer   func() string
	GetToken    func() string
	Start       func(ctx context.Context)
	Stop        func()
}

// NewClient returns a new interceptor client that calls the functions in funcs instead of the underlying client's methods, if they are not nil.
func NewClient(interceptedClient client.Client, funcs Funcs) client.Client {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}

type interceptor struct {
	client client.Client
	funcs  Funcs
}

func (c *interceptor) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
	if c.funcs.Get != nil {
		return c.funcs.Get(ctx, key, obj, opts...)
	}
	return c.client.Get(ctx, key, obj, opts...)
}

func (c *interceptor) List(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
	if c.funcs.List != nil {
		return c.funcs.List(ctx, list, opts...)
	}
	return c.client.List(ctx, list, opts...)
}

func (c *interceptor) Create(ctx context.Context, obj object.Object, opts ...client.CreateOption) error {
	if c.funcs.Create != nil {
		return c.funcs.Create(ctx, obj, opts...)
	}
	return c.client.Create(ctx, obj, opts...)
}

func (c *interceptor) Delete(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
	if c.funcs.Delete != nil {
		return c.funcs.Delete(ctx, obj, opts...)
	}
	return c.client.Delete(ctx, obj, opts...)
}

func (c *interceptor) DeleteAllOf(ctx context.Context, obj object.Object, opts ...client.DeleteAllOfOption) error {
	if c.funcs.DeleteAllOf != nil {
		return c.funcs.DeleteAllOf(ctx, obj, opts...)
	}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *interceptor) Update(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
	if c.funcs.Update != nil {
		return c.funcs.Update(ctx, obj, req, opts...)
	}
	return c.client.Update(ctx, obj, req, opts...)
}

func (c *interceptor) GetInformer(objectType object.ObjectType) client.InformerCache {
	if c.funcs.GetInformer != nil {
		return c.funcs.GetInformer(objectType)
	}
	return c.client.GetInformer(objectType)
}

func (c *interceptor) GetServer() string {
	if c.funcs.GetServer != nil {
		return c.funcs.GetServer()
	}
	return c.client.GetServer()
}

func (c *interceptor) GetToken() string {
	if c.funcs.GetToken != nil {
		return c.funcs.GetToken()
	}
	return c.client.GetToken()
}

func (c *interceptor) Start(ctx context.Context) {
	if c.funcs.Start != nil {
		c.funcs.Start(ctx)
	}
	c.client.Start(ctx)
}

func (c *interceptor) Stop() {
	if c.funcs.Stop != nil {
		c.funcs.Stop()
	}
	c.client.Stop()
}

var _ client.Client = &interceptor{}
