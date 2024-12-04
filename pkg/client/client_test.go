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

	v0040 "github.com/SlinkyProject/slurm-client/api/v0040"
	v0041 "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

const testTimeout = 30 * time.Second

var _ = Describe("Client", func() {
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			Server:    restapiServer,
			AuthToken: slurmJwt,
		}
	})

	Describe("Client", func() {
		var cl Client

		BeforeEach(func() {
			var err error
			cl, err = NewClient(cfg)
			Expect(err).NotTo(HaveOccurred())
			Expect(cl).NotTo(BeNil())
		})

		Context("New", func() {
			It("should return a new Client", func() {
				cl, err := NewClient(cfg)
				Expect(err).NotTo(HaveOccurred())
				Expect(cl).NotTo(BeNil())
			})
			It("should fail if config is nil", func() {
				cl, err := NewClient(nil)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
			It("should fail if Server is empty", func() {
				cfg2 := &Config{
					AuthToken: cfg.AuthToken,
				}
				cl, err := NewClient(cfg2)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
			It("should fail if AuthToken is empty", func() {
				cfg2 := &Config{
					Server: cfg.Server,
				}
				cl, err := NewClient(cfg2)
				Expect(err).To(HaveOccurred())
				Expect(cl).To(BeNil())
			})
		})

		Context("Stop", func() {
			var cl Client

			BeforeEach(func() {
				var err error
				cl, err = NewClient(cfg)
				Expect(err).NotTo(HaveOccurred())
				Expect(cl).NotTo(BeNil())
			})

			It("should stop by Stop function", func() {
				By("waiting for the client cache to Start then invoking Stop")
				ctx := context.TODO()
				startCh := make(chan bool)
				stopCh := make(chan bool)
				go func() {
					startCh <- true
					cl.Start(ctx)
					stopCh <- true
				}()
				<-startCh
				time.Sleep(2 * time.Second)
				cl.Stop()
				Eventually(stopCh).Within(testTimeout).Should(Receive())
			})
			It("should stop by context cancel", func() {
				By("waiting for the client cache to Start then invoking Stop")
				ctx, cancel := context.WithCancel(context.TODO())
				startCh := make(chan bool)
				stopCh := make(chan bool)
				go func() {
					startCh <- true
					cl.Start(ctx)
					stopCh <- true
				}()
				<-startCh
				time.Sleep(2 * time.Second)
				cancel()
				Eventually(stopCh).Within(testTimeout).Should(Receive())
			})
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0040ControllerPing", func() {
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
				obj := &types.V0040ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0040ControllerPingList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0040JobInfo", func() {
		var cl Client
		var req v0040.V0040JobSubmitReq

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

		BeforeEach(func() {
			req = v0040.V0040JobSubmitReq{
				Job: &v0040.V0040JobDescMsg{
					Environment: &v0040.V0040StringArray{
						"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
					},
					CurrentWorkingDirectory: ptr.To("/tmp"),
					Script:                  ptr.To("#!/usr/bin/bash\nexit 0"),
				},
			}
		})

		Context("Create", func() {
			It("should create a new object", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0040JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0040JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0040JobInfo{}
				req := v0040.V0040JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0040JobInfo{V0040JobInfo: v0040.V0040JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0040JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0040JobInfo{}
				key := object.ObjectKey("0")
				err := cl.Get(ctx, key, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0040JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &types.V0040JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0040JobInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0040JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0040JobInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0040Node", func() {
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
				obj := &types.V0040Node{V0040Node: v0040.V0040Node{Name: ptr.To("")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := v0040.V0040UpdateNodeMsg{
				Comment: ptr.To("v0040"),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0040Node{V0040Node: v0040.V0040Node{Name: ptr.To("")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0040Node{V0040Node: v0040.V0040Node{Name: ptr.To("slurmd")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(equality.Semantic.DeepEqual(obj.Comment, req.Comment)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0040Node{V0040Node: v0040.V0040Node{Name: ptr.To("")}}
				actual := &types.V0040Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0040Node{V0040Node: v0040.V0040Node{Name: ptr.To("slurmd")}}
				actual := &types.V0040Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0040NodeList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0040PartitionInfo", func() {
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
				obj := &types.V0040PartitionInfo{V0040PartitionInfo: v0040.V0040PartitionInfo{Name: ptr.To("")}}
				actual := &types.V0040PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0040PartitionInfo{V0040PartitionInfo: v0040.V0040PartitionInfo{Name: ptr.To("all")}}
				actual := &types.V0040PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0040PartitionInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	////////////////////////////////////////////////////////////////////////////

	Describe("V0041ControllerPing", func() {
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
				obj := &types.V0041ControllerPing{}
				key := object.ObjectKey("slurmctld")
				err := cl.Get(ctx, key, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a non-empty list ", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0041ControllerPingList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0041JobInfo", func() {
		var cl Client
		req := v0041.V0041JobSubmitReq{
			Job: &v0041.V0041JobDescMsg{
				Environment: &v0041.V0041StringArray{
					"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
				},
				CurrentWorkingDirectory: ptr.To("/tmp"),
				Script:                  ptr.To("#!/usr/bin/bash\nexit 0"),
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
				obj := &types.V0041JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				actual := &types.V0041JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("writing the result back to the go struct")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
			It("should fail if the object request is invalid", func(ctx SpecContext) {
				By("creating the object")
				obj := &types.V0041JobInfo{}
				req := v0041.V0041JobSubmitReq{}
				err := cl.Create(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		// Context("Create With Allocation", func() {
		// 	opts := &CreateOptions{Allocation: true}
		// 	req := v0041.V0041JobAllocReq{
		// 		Job: &v0041.V0041JobDescMsg{
		// 			Environment: &v0041.V0041StringArray{
		// 				"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin",
		// 			},
		// 			CurrentWorkingDirectory: ptr.To("/tmp"),
		// 			Script:                  ptr.To("#!/usr/bin/bash\nexit 0"),
		// 		},
		// 	}

		// 	It("should create a new object", func(ctx SpecContext) {
		// 		By("creating the object")
		// 		obj := &types.V0041JobInfo{}
		// 		err := cl.Create(ctx, obj, req, opts)
		// 		Expect(err).NotTo(HaveOccurred())

		// 		actual := &types.V0041JobInfo{}
		// 		err = cl.Get(ctx, obj.GetKey(), actual)
		// 		Expect(err).NotTo(HaveOccurred())

		// 		By("writing the result back to the go struct")
		// 		Expect(obj).To(Equal(actual))
		// 	}, SpecTimeout(testTimeout))
		// 	It("should fail if the object request is invalid", func(ctx SpecContext) {
		// 		By("creating the object")
		// 		obj := &types.V0041JobInfo{}
		// 		req := v0041.V0041JobSubmitReq{}
		// 		err := cl.Create(ctx, obj, req, opts)
		// 		Expect(err).To(HaveOccurred())
		// 	}, SpecTimeout(testTimeout))
		// })

		Context("Delete", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("deleting the object")
				obj := &types.V0041JobInfo{V0041JobInfo: v0041.V0041JobInfo{JobId: ptr.To[int32](0)}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should delete a new object", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0041JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("deleting the object")
				err = cl.Delete(ctx, obj)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0041JobInfo{}
				key := object.ObjectKey("0")
				err := cl.Get(ctx, key, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should fetch an existing object for a go struct", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0041JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("fetching the created the object")
				actual := &types.V0041JobInfo{}
				err = cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())

				By("validating the fetched object equals the created one")
				Expect(equality.Semantic.DeepEqual(obj, actual)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0041JobInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return a non-empty", func(ctx SpecContext) {
				By("initially creating an object")
				obj := &types.V0041JobInfo{}
				err := cl.Create(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("listing all objects")
				list := &types.V0041JobInfoList{}
				err = cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())

				By("validating no objects are returned")
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0041Node", func() {
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
				obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("")}}
				err := cl.Delete(ctx, obj)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("Update", func() {
			req := v0041.V0041UpdateNodeMsg{
				Comment: ptr.To("v0041"),
			}

			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should update the existing object", func(ctx SpecContext) {
				By("update the object")
				obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("slurmd")}}
				err := cl.Update(ctx, obj, req)
				Expect(err).NotTo(HaveOccurred())

				By("validating the object field was updated")
				Expect(equality.Semantic.DeepEqual(obj.Comment, req.Comment)).To(BeTrue())
			}, SpecTimeout(testTimeout))
		})

		Context("Get", func() {
			It("should fail if the object does not exist", func(ctx SpecContext) {
				By("fetching non-existent object")
				obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("")}}
				actual := &types.V0041Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0041Node{V0041Node: v0041.V0041Node{Name: ptr.To("slurmd")}}
				actual := &types.V0041Node{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0041NodeList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})

	Describe("V0041PartitionInfo", func() {
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
				obj := &types.V0041PartitionInfo{V0041PartitionInfo: v0041.V0041PartitionInfo{Name: ptr.To("")}}
				actual := &types.V0041PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).To(HaveOccurred())
			}, SpecTimeout(testTimeout))
			It("should return existing object", func(ctx SpecContext) {
				By("fetching existent object")
				obj := &types.V0041PartitionInfo{V0041PartitionInfo: v0041.V0041PartitionInfo{Name: ptr.To("all")}}
				actual := &types.V0041PartitionInfo{}
				err := cl.Get(ctx, obj.GetKey(), actual)
				Expect(err).NotTo(HaveOccurred())
			}, SpecTimeout(testTimeout))
		})

		Context("List", func() {
			It("should return a list", func(ctx SpecContext) {
				By("listing all objects")
				list := &types.V0041PartitionInfoList{}
				err := cl.List(ctx, list)
				Expect(err).NotTo(HaveOccurred())
				Expect(list.Items).NotTo(BeEmpty())
			}, SpecTimeout(testTimeout))
		})
	})
})
