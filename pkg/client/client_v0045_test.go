// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

var _ = Describe("Client v0045", func() {
	const testTimeout = 30 * time.Second
	const cacheSyncPeriod = 5 * time.Second
	const comment = "v0045"
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:    restapiServer,
			AuthToken: slurmJwt,
		}
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0045ControllerPing", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&types.V0045ControllerPing{},
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
				obj := &types.V0045ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0045ControllerPingList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045JobInfo", func() {
		var cl Client
		req := api.V0045JobSubmitReq{
			Job: &api.V0045JobDescMsg{
				Environment: &api.V0045StringArray{
					"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
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
					&types.V0045JobInfo{},
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
				obj := &types.V0045JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0045JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0045JobInfo{}
				req := api.V0045JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0045JobInfo{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0045JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			updateReq := api.V0045JobDescMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0045JobInfo{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0045JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("update the object")
				err = cl.Update(ctx, obj, updateReq)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(equality.Semantic.DeepEqual(obj.Comment, updateReq.Comment)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0045JobInfo{}
				key := object.ObjectKey("0")
				err := cl.Get(ctx, key, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0045JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &types.V0045JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0045JobInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0045JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0045JobInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045NodeResourceLayout", func() {
		var cl Client
		req := api.V0045JobSubmitReq{
			Job: &api.V0045JobDescMsg{
				Environment: &api.V0045StringArray{
					"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
				},
				CurrentWorkingDirectory: ptr.To("/tmp"),
				Script:                  ptr.To("#!/usr/bin/bash\nsleep infinity"),
			},
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
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
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("creating the job object")
				job := &types.V0045JobInfo{}
				err := cl.Create(ctx, job, req)
				Expect(err).NotTo(HaveOccurred())
				defer func() {
					By("cleanup the job object")
					_ = cl.Delete(ctx, job)
				}()

				By("fetching the job resource layout")
				layout := &types.V0045NodeResourceLayout{}
				Eventually(func(g Gomega) {
					// Job must be running otherwise Get() will always return an error.
					g.Expect(cl.Get(ctx, job.GetKey(), layout)).To(Succeed())
				}, testTimeout, 2*time.Second).Should(Succeed())

				By("checking the layout is not empty")
				Expect(equality.Semantic.DeepEqual(layout, &types.V0045NodeResourceLayout{})).NotTo(BeTrue())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045Node", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&types.V0045Node{},
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
				obj := &types.V0045Node{V0045Node: api.V0045Node{Name: ptr.To("does-not-exist")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := api.V0045UpdateNodeMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0045Node{V0045Node: api.V0045Node{Name: ptr.To("does-not-exist")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0045Node{V0045Node: api.V0045Node{Name: ptr.To("slurmd")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(equality.Semantic.DeepEqual(obj.Comment, req.Comment)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0045Node{V0045Node: api.V0045Node{Name: ptr.To("does-not-exist")}}
				actual := &types.V0045Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0045Node{V0045Node: api.V0045Node{Name: ptr.To("slurmd")}}
				actual := &types.V0045Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0045NodeList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045PartitionInfo", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&types.V0045PartitionInfo{},
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
				obj := &types.V0045PartitionInfo{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("does-not-exist")}}
				actual := &types.V0045PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0045PartitionInfo{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("all")}}
				actual := &types.V0045PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0045PartitionInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045Reconfigure", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
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
			It("should reconfigure", func(ctx SpecContext) {
				By("making request")
				obj := &types.V0045Reconfigure{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should reconfigure", func(ctx SpecContext) {
				By("making requests")
				list := &types.V0045ReconfigureList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045ReservationInfo", func() {
		var cl Client

		nodeList := api.V0045HostlistString{}
		nodeList = append(nodeList, "slurmd")
		users := api.V0045CsvString{"slurm"}
		startTime := api.V0045Uint64NoValStruct{Number: ptr.To(time.Now().Unix()), Set: ptr.To(true)}
		duration := api.V0045Uint32NoValStruct{Number: ptr.To(int32(10000)), Set: ptr.To(true)}

		req := api.V0045ReservationDescMsg{
			NodeList:  &nodeList,
			Users:     &users,
			StartTime: &startTime,
			Duration:  &duration,
			Flags: &[]api.V0045ReservationDescMsgFlags{
				api.V0045ReservationDescMsgFlagsOVERLAP, // for testing, allow overlap
			},
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&types.V0045ReservationInfo{},
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
				const reservationName = "create-v45"
				By("creating the object")
				obj := &types.V0045ReservationInfo{}
				req := req // shallow copy and scope variable
				req.Name = ptr.To(reservationName)
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0045ReservationInfo{}
				req := api.V0045ReservationDescMsg{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0045ReservationInfo{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("does-not-exist")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete an existing object", func(ctx SpecContext) {
				const reservationName = "delete-v45"
				By("creating the object")
				obj := &types.V0045ReservationInfo{}
				req := req // shallow copy and scope variable
				req.Name = ptr.To(reservationName)
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			users := api.V0045CsvString{"slurm", "root"}
			updateReq := api.V0045ReservationDescMsg{
				NodeList: &nodeList,
				Users:    &users,
			}
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0045ReservationInfo{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("does-not-exist")}}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				const reservationName = "update-v45"
				By("creating the object")
				obj := &types.V0045ReservationInfo{}
				req := req // shallow copy and scope variable
				req.Name = ptr.To(reservationName)
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("update the object")
				err = cl.Update(ctx, obj, updateReq)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0045ReservationInfo{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("does-not-exist")}}
				actual := &types.V0045ReservationInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				const reservationName = "get-v45"
				By("creating the object")
				obj := &types.V0045ReservationInfo{}
				req := req // shallow copy and scope variable
				req.Name = ptr.To(reservationName)
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching existing object")
				actual := &types.V0045ReservationInfo{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To(reservationName)}}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				const reservationName = "list-v45"
				By("creating the object")
				obj := &types.V0045ReservationInfo{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To(reservationName)}}
				req := req // shallow copy and scope variable
				req.Name = ptr.To(reservationName)
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0045ReservationInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0045Stats", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg, &ClientOptions{
				EnableFor: []object.Object{
					&types.V0045Stats{},
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
				obj := &types.V0045Stats{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0045StatsList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})
})
