// SPDX-FileCopyrightText: Copyright (C) NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	v42 "github.com/SlinkyProject/slurm-client/api/v0042"
	v43 "github.com/SlinkyProject/slurm-client/api/v0043"
	v44 "github.com/SlinkyProject/slurm-client/api/v0044"
	v45 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

var _ = DescribeTable("Job reads omit resource details",
	func(ctx SpecContext,
		newJob func() object.Object,
		newList func() object.ObjectList,
		readDefaultJob func(context.Context, Gomega, Client, string) object.Object,
		listDefaultJobs func(context.Context, Client) object.ObjectList,
	) {
		cl, err := NewClient(&Config{
			Server: restapiServer, TokenProvider: token.StaticProvider(slurmJwt),
		})
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(cl.Stop)

		By("submitting a job that retains an allocation throughout the assertions")
		submitted := &v45.V0045JobInfo{}
		Expect(cl.Create(ctx, submitted, v45.V0045JobSubmitReq{
			Job: &v45.V0045JobDescMsg{
				Environment:             &v45.V0045StringArray{"PATH=/usr/bin:/bin"},
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
			job := readDefaultJob(ctx, g, cl, string(submitted.GetKey()))
			g.Expect(job).To(HaveField("JobResources", Not(BeNil())))
			g.Expect(job).To(HaveField("JobResources.Nodes.Allocation", HaveValue(Not(BeEmpty()))))
		}, 30*time.Second, time.Second).Should(Succeed())

		By("confirming the default list response also contains the allocation")
		list := listDefaultJobs(ctx, cl)
		Expect(list.GetItems()).To(ContainElement(And(
			HaveField("JobId", Equal(submitted.JobId)),
			HaveField("JobResources.Nodes.Allocation", HaveValue(Not(BeEmpty()))),
		)))

		By("reading the same job through the client without resource details")
		job := newJob()
		Expect(cl.Get(ctx, submitted.GetKey(), job, &GetOptions{SkipCache: true})).To(Succeed())
		Expect(job.GetKey()).To(Equal(submitted.GetKey()))
		Expect(job).To(HaveField("JobResources", BeNil()))

		By("listing through the client without resource details")
		list = newList()
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
	},
	Entry("v0042", SpecTimeout(time.Minute),
		func() object.Object { return &v42.V0042JobInfo{} },
		func() object.ObjectList { return &v42.V0042JobInfoObjectList{} },
		func(ctx context.Context, g Gomega, cl Client, jobID string) object.Object {
			GinkgoHelper()
			response, err := cl.Versioned().V0042().SlurmV0042GetJobWithResponse(ctx, jobID, nil)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode()).To(Equal(http.StatusOK))
			g.Expect(response.JSON200).NotTo(BeNil())
			g.Expect(response.JSON200.Jobs).To(HaveLen(1))
			return &response.JSON200.Jobs[0]
		},
		func(ctx context.Context, cl Client) object.ObjectList {
			GinkgoHelper()
			response, err := cl.Versioned().V0042().SlurmV0042GetJobsWithResponse(ctx, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(response.JSON200).NotTo(BeNil())
			return &v42.V0042JobInfoObjectList{Items: response.JSON200.Jobs}
		},
	),
	Entry("v0043", SpecTimeout(time.Minute),
		func() object.Object { return &v43.V0043JobInfo{} },
		func() object.ObjectList { return &v43.V0043JobInfoObjectList{} },
		func(ctx context.Context, g Gomega, cl Client, jobID string) object.Object {
			GinkgoHelper()
			response, err := cl.Versioned().V0043().SlurmV0043GetJobWithResponse(ctx, jobID, nil)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode()).To(Equal(http.StatusOK))
			g.Expect(response.JSON200).NotTo(BeNil())
			g.Expect(response.JSON200.Jobs).To(HaveLen(1))
			return &response.JSON200.Jobs[0]
		},
		func(ctx context.Context, cl Client) object.ObjectList {
			GinkgoHelper()
			response, err := cl.Versioned().V0043().SlurmV0043GetJobsWithResponse(ctx, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(response.JSON200).NotTo(BeNil())
			return &v43.V0043JobInfoObjectList{Items: response.JSON200.Jobs}
		},
	),
	Entry("v0044", SpecTimeout(time.Minute),
		func() object.Object { return &v44.V0044JobInfo{} },
		func() object.ObjectList { return &v44.V0044JobInfoObjectList{} },
		func(ctx context.Context, g Gomega, cl Client, jobID string) object.Object {
			GinkgoHelper()
			response, err := cl.Versioned().V0044().SlurmV0044GetJobWithResponse(ctx, jobID, nil)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode()).To(Equal(http.StatusOK))
			g.Expect(response.JSON200).NotTo(BeNil())
			g.Expect(response.JSON200.Jobs).To(HaveLen(1))
			return &response.JSON200.Jobs[0]
		},
		func(ctx context.Context, cl Client) object.ObjectList {
			GinkgoHelper()
			response, err := cl.Versioned().V0044().SlurmV0044GetJobsWithResponse(ctx, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(response.JSON200).NotTo(BeNil())
			return &v44.V0044JobInfoObjectList{Items: response.JSON200.Jobs}
		},
	),
	Entry("v0045", SpecTimeout(time.Minute),
		func() object.Object { return &v45.V0045JobInfo{} },
		func() object.ObjectList { return &v45.V0045JobInfoObjectList{} },
		func(ctx context.Context, g Gomega, cl Client, jobID string) object.Object {
			GinkgoHelper()
			response, err := cl.Versioned().V0045().SlurmV0045GetJobWithResponse(ctx, jobID, nil)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.StatusCode()).To(Equal(http.StatusOK))
			g.Expect(response.JSON200).NotTo(BeNil())
			g.Expect(response.JSON200.Jobs).To(HaveLen(1))
			return &response.JSON200.Jobs[0]
		},
		func(ctx context.Context, cl Client) object.ObjectList {
			GinkgoHelper()
			response, err := cl.Versioned().V0045().SlurmV0045GetJobsWithResponse(ctx, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
			Expect(response.JSON200).NotTo(BeNil())
			return &v45.V0045JobInfoObjectList{Items: response.JSON200.Jobs}
		},
	),
)
