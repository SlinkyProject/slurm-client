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

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

var _ = Describe("Client v0044", func() {
	const testTimeout = 30 * time.Second
	const comment = "v0044"
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:    restapiServer,
			AuthToken: slurmJwt,
		}
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0044ControllerPing", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0044ControllerPingList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044JobInfo", func() {
		var cl Client
		req := api.V0044JobSubmitReq{
			Job: &api.V0044JobDescMsg{
				Environment: &api.V0044StringArray{
					"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
				},
				CurrentWorkingDirectory: ptr.To("/tmp"),
				Script:                  ptr.To("#!/usr/bin/sh\nexit 0"),
				Hold:                    ptr.To(true),
			},
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0044JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044JobInfo{}
				req := api.V0044JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		// Context("Create With Allocation", func() {
		// 	opts := &CreateOptions{Allocation: true}
		// 	req := api.V0044JobAllocReq{
		// 		Job: &api.V0044JobDescMsg{
		// 			Environment: &api.V0044StringArray{
		// 				"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
		// 			},
		// 			CurrentWorkingDirectory: ptr.To("/tmp"),
		// 			Script:                  ptr.To("#!/usr/bin/sh\nexit 0"),
		// 			Hold:                    ptr.To(true),
		// 		},
		// 	}

		// 	It("should create a new object", func(ctx SpecContext) {
		// 		By("creating the object")
		// 		obj := &types.V0044JobInfo{}
		// 		err := cl.Create(ctx, obj, req, opts)
		// 		Expect(err).NotTo(HaveOccurred())

		// 		actual := &types.V0044JobInfo{}
		// 		err = cl.Get(ctx, obj.GetKey(), actual)
		// 		Expect(err).NotTo(HaveOccurred())

		// 		By("writing the result back to the go struct")
		// 		Expect(obj).To(Equal(actual))
		// 	}, SpecTimeout(testTimeout))
		// 	It("should fail if the object request is invalid", func(ctx SpecContext) {
		// 		By("creating the object")
		// 		obj := &types.V0044JobInfo{}
		// 		req := api.V0044JobSubmitReq{}
		// 		err := cl.Create(ctx, obj, req, opts)
		// 		Expect(err).To(HaveOccurred())
		// 	}, SpecTimeout(testTimeout))
		// })

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0044JobInfo{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0044JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			updateReq := api.V0044JobDescMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0044JobInfo{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044JobInfo{}
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
				obj := &types.V0044JobInfo{}
				key := object.ObjectKey("0")
				err := cl.Get(ctx, key, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0044JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &types.V0044JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0044JobInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0044JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0044JobInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044NodeResourceLayout", func() {
		var cl Client
		req := api.V0044JobSubmitReq{
			Job: &api.V0044JobDescMsg{
				Environment: &api.V0044StringArray{
					"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
				},
				CurrentWorkingDirectory: ptr.To("/tmp"),
				Script:                  ptr.To("#!/usr/bin/bash\nsleep infinity"),
			},
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				job := &types.V0044JobInfo{}
				err := cl.Create(ctx, job, req)
				Expect(err).NotTo(HaveOccurred())
				defer func() {
					By("cleanup the job object")
					_ = cl.Delete(ctx, job)
				}()

				By("fetching the job resource layout")
				layout := &types.V0044NodeResourceLayout{}
				Eventually(func(g Gomega) {
					// Job must be running otherwise Get() will always return an error.
					g.Expect(cl.Get(ctx, job.GetKey(), layout)).To(Succeed())
				}, testTimeout, 2*time.Second).Should(Succeed())

				By("checking the layout is not empty")
				Expect(equality.Semantic.DeepEqual(layout, &types.V0044NodeResourceLayout{})).NotTo(BeTrue())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044Node", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044Node{V0044Node: api.V0044Node{Name: ptr.To("")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := api.V0044UpdateNodeMsg{
				Comment: ptr.To(comment),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0044Node{V0044Node: api.V0044Node{Name: ptr.To("")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0044Node{V0044Node: api.V0044Node{Name: ptr.To("slurmd")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(equality.Semantic.DeepEqual(obj.Comment, req.Comment)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0044Node{V0044Node: api.V0044Node{Name: ptr.To("")}}
				actual := &types.V0044Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0044Node{V0044Node: api.V0044Node{Name: ptr.To("slurmd")}}
				actual := &types.V0044Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0044NodeList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044PartitionInfo", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044PartitionInfo{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("")}}
				actual := &types.V0044PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0044PartitionInfo{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("all")}}
				actual := &types.V0044PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0044PartitionInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044Reconfigure", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044Reconfigure{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should reconfigure", func(ctx SpecContext) {
				By("making requests")
				list := &types.V0044ReconfigureList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044ReservationInfo", func() {
		var cl Client

		nodeList := api.V0044HostlistString{}
		nodeList = append(nodeList, "slurmd")
		users := api.V0044CsvString{"slurm"}
		startTime := api.V0044Uint64NoValStruct{Number: ptr.To(time.Now().Unix()), Set: ptr.To(true)}
		duration := api.V0044Uint32NoValStruct{Number: ptr.To(int32(10000)), Set: ptr.To(true)}

		req := api.V0044ReservationDescMsg{
			Name:      ptr.To("test"),
			NodeList:  &nodeList,
			Users:     &users,
			StartTime: &startTime,
			Duration:  &duration,
		}

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044ReservationInfo{}
				req := api.V0044ReservationDescMsg{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test2")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete an existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			users := api.V0044CsvString{"slurm,root"}
			updateReq := api.V0044ReservationDescMsg{
				Name:     ptr.To("test"),
				NodeList: &nodeList,
				Users:    &users,
			}
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Update(ctx, obj, updateReq)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("update the object")
				err = cl.Update(ctx, obj, updateReq)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should create a new object", func(ctx SpecContext) {

			}, SpecTimeout(testTimeout))
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("")}}
				actual := &types.V0044ReservationInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("fetching existent object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				actual = &types.V0044ReservationInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0044ReservationInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())

				By("deleting the object")
				obj = &types.V0044ReservationInfo{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test")}}
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0044Stats", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
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
				obj := &types.V0044Stats{}
				err := cl.Get(ctx, obj.GetKey(), obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0044StatsList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})
})
