// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	v0041 "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("Create", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Create: func(ctx context.Context, obj object.Object, req any, opts ...client.CreateOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			_ = client.Create(ctx, obj, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Create: func(ctx context.Context, obj object.Object, req any, opts ...client.CreateOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.Create(ctx, obj, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("Delete", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Delete: func(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			_ = client.Delete(ctx, obj)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Delete: func(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.Delete(ctx, obj)
			Expect(called).To(BeTrue())
		})
	})
	Context("Update", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Update: func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			req := v0041.V0041UpdateNodeMsg{}
			_ = client.Update(ctx, obj, req)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Update: func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			req := v0041.V0041UpdateNodeMsg{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.Update(ctx, obj, req)
			Expect(called).To(BeTrue())
		})
	})
	Context("Get", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Get: func(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			_ = client.Get(ctx, obj.GetKey(), obj)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Get: func(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.Get(ctx, obj.GetKey(), obj)
			Expect(called).To(BeTrue())
		})
	})
	Context("List", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				List: func(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
					called = true
					return nil
				},
			})
			list := &types.V0041NodeList{}
			_ = client.List(ctx, list)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				List: func(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
					called = true
					return nil
				},
			})
			list := &types.V0041NodeList{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.List(ctx, list)
			Expect(called).To(BeTrue())
		})
	})
	Context("GetInformer", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				GetInformer: func(obj object.ObjectType) client.InformerCache {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			_ = client.GetInformer(obj.GetType())
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				GetInformer: func(obj object.ObjectType) client.InformerCache {
					called = true
					return nil
				},
			})
			obj := &types.V0041Node{}
			client2 := NewClient(client1, Funcs{})
			_ = client2.GetInformer(obj.GetType())
			Expect(called).To(BeTrue())
		})
	})
	Context("SetServer", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SetServer: func(string) {
					called = true
				},
			})
			client.SetServer("")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SetServer: func(string) {
					called = true
				},
			})
			client2 := NewClient(client1, Funcs{})
			client2.SetServer("")
			Expect(called).To(BeTrue())
		})
	})
	Context("SetToken", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SetToken: func(string) {
					called = true
				},
			})
			client.SetToken("")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SetToken: func(string) {
					called = true
				},
			})
			client2 := NewClient(client1, Funcs{})
			client2.SetToken("")
			Expect(called).To(BeTrue())
		})
	})
	Context("Start", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Start: func(ctx context.Context) {
					called = true
				},
			})
			client.Start(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Start: func(ctx context.Context) {
					called = true
				},
			})
			client2 := NewClient(client1, Funcs{})
			client2.Start(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("Stop", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				Stop: func() {
					called = true
				},
			})
			client.Stop()
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				Stop: func() {
					called = true
				},
			})
			client2 := NewClient(client1, Funcs{})
			client2.Stop()
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// Create implements client.Client.
func (e *emptyClient) Create(ctx context.Context, obj object.Object, req any, opts ...client.CreateOption) error {
	return nil
}

// Delete implements client.Client.
func (e *emptyClient) Delete(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
	return nil
}

// Update implements client.Client.
func (e *emptyClient) Update(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
	return nil
}

// Get implements client.Client.
func (e *emptyClient) Get(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
	return nil
}

// List implements client.Client.
func (e *emptyClient) List(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
	return nil
}

// GetInformer implements client.Client.
func (e *emptyClient) GetInformer(objectType object.ObjectType) client.InformerCache {
	return nil
}

// GetServer implements client.Client.
func (e *emptyClient) GetServer() string {
	return ""
}

// SetServer implements client.Client.
func (e *emptyClient) SetServer(server string) {
}

// GetToken implements client.Client.
func (e *emptyClient) GetToken() string {
	return ""
}

// SetToken implements client.Client.
func (e *emptyClient) SetToken(token string) {
}

// Start implements client.Client.
func (e *emptyClient) Start(ctx context.Context) {
}

// Stop implements client.Client.
func (e *emptyClient) Stop() {
}

var _ client.Client = &emptyClient{}
