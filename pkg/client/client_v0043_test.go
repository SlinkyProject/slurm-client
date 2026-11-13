// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

var _ = Describe("Client v0043", func() {
	const testTimeout = 30 * time.Second
	const cacheSyncPeriod = 5 * time.Second
	const comment = "v0043"
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:        restapiServer,
			TokenProvider: token.StaticProvider(slurmJwt),
		}
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0043ControllerPing", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0043ControllerPing{},
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
				obj := &api.V0043ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0043ControllerPingObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0043JobInfo", func() {
		var cl Client
		req := api.V0043JobSubmitReq{
			Job: &api.V0043JobDescMsg{
				Environment: &api.V0043StringArray{
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
					&api.V0043JobInfo{},
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

		It("should omit resource details from job reads", func(ctx SpecContext) {
			By("submitting a job that retains an allocation throughout the assertions")
			submitted := &api.V0043JobInfo{}
			Expect(cl.Create(ctx, submitted, api.V0043JobSubmitReq{
				Job: &api.V0043JobDescMsg{
					Environment:             &api.V0043StringArray{"PATH=/usr/bin:/bin"},
					CurrentWorkingDirectory: ptr.To("/tmp"),
					Script:                  ptr.To("#!/bin/sh\nsleep 300"),
					Tasks:                   ptr.To[int32](1),
					MinimumCpus:             ptr.To[int32](1),
					MinimumNodes:            ptr.To[int32](1),
				},
			})).To(Succeed())
			DeferCleanup(func(ctx SpecContext) {
				Expect(cl.Delete(ctx, submitted, &DeleteOptions{SkipReadAfterDelete: true})).To(Succeed())
			}, NodeTimeout(10*time.Second))

			By("waiting for the default job response to contain an allocated node")
			Eventually(func(g Gomega) {
				response, err := cl.Versioned().V0043().SlurmV0043GetJobWithResponse(ctx, string(submitted.GetKey()), nil)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(response.StatusCode()).To(Equal(http.StatusOK))
				g.Expect(response.JSON200).NotTo(BeNil())
				g.Expect(response.JSON200.Jobs).To(HaveLen(1))
				job := &response.JSON200.Jobs[0]
				g.Expect(job).To(HaveField("JobResources", Not(BeNil())))
				g.Expect(job).To(HaveField("JobResources.Nodes.Allocation", HaveValue(Not(BeEmpty()))))
			}, 30*time.Second, time.Second).Should(Succeed())

			By("confirming the default list response also contains the allocation")
			response, err := cl.Versioned().V0043().SlurmV0043GetJobsWithResponse(ctx, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(response.JSON200).NotTo(BeNil())
			list := &api.V0043JobInfoObjectList{Items: response.JSON200.Jobs}
			Expect(list.GetItems()).To(ContainElement(And(
				HaveField("JobId", Equal(submitted.JobId)),
				HaveField("JobResources.Nodes.Allocation", HaveValue(Not(BeEmpty()))),
			)))

			By("reading the same job through the client without resource details")
			job := &api.V0043JobInfo{}
			Expect(cl.Get(ctx, submitted.GetKey(), job, &GetOptions{SkipCache: true})).To(Succeed())
			Expect(job.GetKey()).To(Equal(submitted.GetKey()))
			Expect(job).To(HaveField("JobResources", BeNil()))

			By("listing through the client without resource details")
			list = &api.V0043JobInfoObjectList{}
			Expect(cl.List(ctx, list, &ListOptions{SkipCache: true})).To(Succeed())
			var found object.Object
			for _, item := range list.GetItems() {
				if item.GetKey() == submitted.GetKey() {
					found = item
					break
				}
			}
			Expect(found).NotTo(BeNil(), "the allocated job must remain in the list")
			Expect(found).To(HaveField("JobResources", BeNil()))
		}, SpecTimeout(time.Minute))

		Context("Create", func() {
			It("should create a new object", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0043JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &api.V0043JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(actual).To(BeEquivalentTo(obj))
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0043JobInfo{}
				req := api.V0043JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &api.V0043JobInfo{JobId: ptr.To[int32](0)}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0043JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			updateReq := api.V0043JobDescMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0043JobInfo{JobId: ptr.To[int32](0)}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &api.V0043JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("update the object")
				err = cl.Update(ctx, obj, updateReq)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(obj.Comment).To(BeEquivalentTo(updateReq.Comment))
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should return ErrNotFound from cache if the object does not exist", func(ctx SpecContext) {
				By("waiting for the informer cache to start")
				Eventually(cl.GetInformer(api.ObjectTypeV0043JobInfo).HasStarted).Should(BeTrue())

				By("fetching a non-existent object from cache")
				obj := &api.V0043JobInfo{}
				key := object.ObjectKey("2147483647")
				err := cl.Get(ctx, key, obj)
				Expect(errors.Is(err, apierrors.ErrNotFound)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should return ErrNotFound from the live API if the object does not exist", func(ctx SpecContext) {
				By("fetching a non-existent object from the live API")
				obj := &api.V0043JobInfo{}
				key := object.ObjectKey("2147483647")
				err := cl.Get(ctx, key, obj, &GetOptions{SkipCache: true})
				Expect(errors.Is(err, apierrors.ErrNotFound)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0043JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &api.V0043JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(actual).To(BeEquivalentTo(obj))
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0043JobInfoObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &api.V0043JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &api.V0043JobInfoObjectList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0043Node", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0043Node{},
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
				obj := &api.V0043Node{Name: ptr.To("does-not-exist")}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := api.V0043UpdateNodeMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0043Node{Name: ptr.To("does-not-exist")}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &api.V0043Node{Name: ptr.To("slurmd")}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(obj.Comment).To(BeEquivalentTo(req.Comment))
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &api.V0043Node{Name: ptr.To("does-not-exist")}
				actual := &api.V0043Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &api.V0043Node{Name: ptr.To("slurmd")}
				actual := &api.V0043Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0043NodeObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0043PartitionInfo", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0043PartitionInfo{},
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
				obj := &api.V0043PartitionInfo{Name: ptr.To("does-not-exist")}
				actual := &api.V0043PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &api.V0043PartitionInfo{Name: ptr.To("all")}
				actual := &api.V0043PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0043PartitionInfoObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0043Stats", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&api.V0043StatsMsg{},
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
				obj := &api.V0043StatsMsg{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &api.V0043StatsMsgObjectList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.GetItems()).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})
})
