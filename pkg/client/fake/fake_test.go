// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	v0041 "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/client/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()

	Context("Get", func() {
		It("should return Not Found", func() {
			client := NewFakeClient()
			obj := &types.V0041Node{}
			key := object.ObjectKey("node-0")
			err := client.Get(ctx, key, obj)
			Expect(err).To(HaveOccurred())
		})
		It("should return Found object", func() {
			client := NewClientBuilder().
				WithObjects(&types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}).
				Build()
			obj := &types.V0041Node{}
			key := object.ObjectKey("node-0")
			err := client.Get(ctx, key, obj)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return error", func() {
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Get: func(ctx context.Context, key object.ObjectKey, obj object.Object, opts ...client.GetOption) error {
						return errors.New(http.StatusText(http.StatusInternalServerError))
					},
				}).
				Build()
			obj := &types.V0041Node{}
			key := object.ObjectKey("node-0")
			err := client.Get(ctx, key, obj)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("List", func() {
		It("should return empty list", func() {
			client := NewFakeClient()
			list := &types.V0041NodeList{}
			err := client.List(ctx, list)
			Expect(err).NotTo(HaveOccurred())
			Expect(list.Items).To(BeEmpty())
		})
		It("should return non-empty list", func() {
			client := NewClientBuilder().
				WithLists(&types.V0041NodeList{Items: []types.V0041Node{
					{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}},
					{V0041Node: v0041.V0041Node{Name: ptr.To("node-1")}},
					{V0041Node: v0041.V0041Node{Name: ptr.To("node-2")}},
				}}).
				Build()
			list := &types.V0041NodeList{}
			err := client.List(ctx, list)
			Expect(err).To(BeNil())
			Expect(list.Items).To(HaveLen(3))
		})
		It("should return error", func() {
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					List: func(ctx context.Context, list object.ObjectList, opts ...client.ListOption) error {
						return errors.New(http.StatusText(http.StatusInternalServerError))
					},
				}).
				Build()
			list := &types.V0041NodeList{}
			err := client.List(ctx, list)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Create", func() {
		It("should succeed", func() {
			client := NewFakeClient()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			err := client.Create(ctx, obj, nil)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return conflict error", func() {
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			client := NewClientBuilder().WithObjects(obj).Build()
			err := client.Create(ctx, obj, nil)
			Expect(err).To(HaveOccurred())
		})
		It("should return error", func() {
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Create: func(ctx context.Context, obj object.Object, req any, opts ...client.CreateOption) error {
						return errors.New(http.StatusText(http.StatusInternalServerError))
					},
				}).
				Build()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			err := client.Create(ctx, obj, nil)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Delete", func() {
		It("should succeed", func() {
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			client := NewClientBuilder().WithObjects(obj).Build()
			err := client.Delete(ctx, obj)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return Not Found", func() {
			client := NewFakeClient()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			err := client.Delete(ctx, obj)
			Expect(err).To(HaveOccurred())
		})
		It("should return error", func() {
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Delete: func(ctx context.Context, obj object.Object, opts ...client.DeleteOption) error {
						return errors.New(http.StatusText(http.StatusInternalServerError))
					},
				}).
				Build()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			err := client.Delete(ctx, obj)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Update", func() {
		It("should update existing object", func() {
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			client := NewClientBuilder().WithObjects(obj).Build()
			req := v0041.V0041UpdateNodeMsg{}
			err := client.Update(ctx, obj, req)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should mutate object after update", func() {
			comment := "test update with mutation"
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			updateFn := func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
				switch o := obj.(type) {
				case *types.V0041Node:
					r, ok := req.(v0041.V0041UpdateNodeMsg)
					if !ok {
						return errors.New("failed to cast request object")
					}
					o.Comment = r.Comment
				default:
					return errors.New("failed to cast slurm object")
				}
				return nil
			}
			client := NewClientBuilder().WithObjects(obj).WithUpdateFn(updateFn).Build()
			req := v0041.V0041UpdateNodeMsg{Comment: &comment}
			err := client.Update(ctx, obj, req)
			Expect(err).NotTo(HaveOccurred())
			Expect(obj.Comment).To(BeEquivalentTo(&comment))
		})
		It("should return Not Found", func() {
			client := NewFakeClient()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			req := v0041.V0041UpdateNodeMsg{}
			err := client.Update(ctx, obj, req)
			Expect(err).To(HaveOccurred())
		})
		It("should return error", func() {
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Update: func(ctx context.Context, obj object.Object, req any, opts ...client.UpdateOption) error {
						return errors.New(http.StatusText(http.StatusInternalServerError))
					},
				}).
				Build()
			obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("node-0")}}
			req := v0041.V0041UpdateNodeMsg{}
			err := client.Update(ctx, obj, req)
			Expect(err).To(HaveOccurred())
		})
	})

	// Context("GetInformer", func() {
	// 	It("should return informer", func() {
	// 		client := NewFakeClient()
	// 		informer := client.GetInformer(types.ObjectTypeV0041Node)
	// 		Expect(informer).NotTo(BeNil())
	// 	})
	// })

	Context("SetServer", func() {
		It("should return", func() {
			const server = "server-foo"
			client := NewFakeClient()
			client.SetServer(server)
			Expect(client.GetServer()).To(Not(BeNil()))
		})
	})

	Context("SetToken", func() {
		It("should return", func() {
			const token = "token-foo"
			client := NewFakeClient()
			client.SetToken(token)
			Expect(client.GetToken()).NotTo(BeNil())
		})
	})

	Context("Start", func() {
		It("should return", func() {
			var called bool
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Start: func(ctx context.Context) {
						called = true
					},
				}).
				Build()
			client.Start(context.Background())
			Expect(called).To(BeTrue())
		})
	})

	Context("Stop", func() {
		It("should return", func() {
			var called bool
			client := NewClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					Stop: func() {
						called = true
					},
				}).
				Build()
			client.Stop()
			Expect(called).To(BeTrue())
		})
	})
})
