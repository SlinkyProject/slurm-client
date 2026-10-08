// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

var _ = Describe("Client v0042", func() {
	const testTimeout = 30 * time.Second
	const cacheSyncPeriod = 5 * time.Second
	const comment = "v0042"
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:        restapiServer,
			TokenProvider: token.StaticProvider(slurmJwt),
		}
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0042ControllerPing", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0042ControllerPing{},
				},
				CacheSyncPeriod: cacheSyncPeriod,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())

			go cl.Start(context.TODO())

			DeferCleanup(func() {
				cl.Stop()
			})
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &api.V0042ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0042ControllerPingObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0042JobInfo", func() {
		var cl Client
		req := api.V0042JobSubmitReq{
			Job: &api.V0042JobDescMsg{
				Environment: &api.V0042StringArray{
					"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
				},
				CurrentWorkingDirectory: ptr.To("/tmp"),
				Script:                  ptr.To("#!/usr/bin/sh\nexit 0"),
				Hold:                    ptr.To(true),
			},
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0042JobInfo{},
				},
				CacheSyncPeriod: cacheSyncPeriod,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())

			go cl.Start(context.TODO())

			DeferCleanup(func() {
				cl.Stop()
			})
		})

		Context("Create", func() {
			It("should create a new object", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0042JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &api.V0042JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(actual).To(BeEquivalentTo(obj))
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0042JobInfo{}
				req := api.V0042JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &api.V0042JobInfo{JobId: ptr.To[int32](0)}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0042JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			updateReq := api.V0042JobDescMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0042JobInfo{JobId: ptr.To[int32](0)}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0042JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("update the object")
				err = cl.Update(ctx, obj, updateReq)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(updateReq.Comment).To(BeEquivalentTo(obj.Comment))
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should return ErrNotFound from cache if the object does not exist", func(ctx SpecContext) {
				By("waiting for the informer cache to start")
				Eventually(cl.GetInformer(api.ObjectTypeV0042JobInfo).HasStarted).Should(BeTrue())

				By("fetching a non-existent object from cache")
				obj := &api.V0042JobInfo{}
				key := object.ObjectKey("2147483647")
				err := cl.Get(ctx, key, obj)
				Expect(errors.Is(err, apierrors.ErrNotFound)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should return ErrNotFound from the live API if the object does not exist", func(ctx SpecContext) {
				By("fetching a non-existent object from the live API")
				obj := &api.V0042JobInfo{}
				key := object.ObjectKey("2147483647")
				err := cl.Get(ctx, key, obj, &GetOptions{SkipCache: true})
				Expect(errors.Is(err, apierrors.ErrNotFound)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0042JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &api.V0042JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(actual).To(BeEquivalentTo(obj))
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0042JobInfoObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0042JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &api.V0042JobInfoObjectList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0042Node", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0042Node{},
				},
				CacheSyncPeriod: cacheSyncPeriod,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())

			go cl.Start(context.TODO())

			DeferCleanup(func() {
				cl.Stop()
			})
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &api.V0042Node{Name: ptr.To("does-not-exist")}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := api.V0042UpdateNodeMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0042Node{Name: ptr.To("does-not-exist")}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0042Node{Name: ptr.To("slurmd")}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(obj.Comment).To(BeEquivalentTo(req.Comment))
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &api.V0042Node{Name: ptr.To("does-not-exist")}
				actual := &api.V0042Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &api.V0042Node{Name: ptr.To("slurmd")}
				actual := &api.V0042Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0042NodeObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0042PartitionInfo", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0042PartitionInfo{},
				},
				CacheSyncPeriod: cacheSyncPeriod,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())

			go cl.Start(context.TODO())

			DeferCleanup(func() {
				cl.Stop()
			})
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &api.V0042PartitionInfo{Name: ptr.To("does-not-exist")}
				actual := &api.V0042PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &api.V0042PartitionInfo{Name: ptr.To("all")}
				actual := &api.V0042PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0042PartitionInfoObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0042Stats", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0042StatsMsg{},
				},
				CacheSyncPeriod: cacheSyncPeriod,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())

			go cl.Start(context.TODO())

			DeferCleanup(func() {
				cl.Stop()
			})
		})

		Context("Get", func() {
			It("should fetch stats", func(ctx SpecContext) {
				By("fetching data")
				obj := &api.V0042StatsMsg{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0042StatsMsgObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})
})
