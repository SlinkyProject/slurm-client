// SPDX-FileCopyrightText: Copyright (C) NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"encoding/json"
	"fmt"
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

// Read the default server response independently of the client's job-read flags.
type jobResourceResponse struct {
	Jobs []struct {
		JobID        int32 `json:"job_id"`
		JobResources *struct {
			Nodes struct {
				Allocation []json.RawMessage `json:"allocation"`
			} `json:"nodes"`
		} `json:"job_resources"`
	} `json:"jobs"`
}

func readDefaultJobResources(ctx context.Context, path string) (jobResourceResponse, error) {
	var result jobResourceResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, restapiServer+path, nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("X-SLURM-USER-TOKEN", slurmJwt)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("default job read: HTTP %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

var _ = DescribeTable("Job reads omit resource details",
	func(ctx SpecContext, version string, newJob func() object.Object, newList func() object.ObjectList) {
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

		prefix := "/slurm/v0.0." + version
		By("waiting for the default job response to contain an allocated node")
		Eventually(func(g Gomega) {
			response, err := readDefaultJobResources(ctx, prefix+"/job/"+string(submitted.GetKey()))
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(response.Jobs).To(HaveLen(1))
			g.Expect(response.Jobs[0].JobResources).NotTo(BeNil())
			g.Expect(response.Jobs[0].JobResources.Nodes.Allocation).NotTo(BeEmpty())
		}, 30*time.Second, time.Second).Should(Succeed())

		By("confirming the default list response also contains the allocation")
		response, err := readDefaultJobResources(ctx, prefix+"/jobs/")
		Expect(err).NotTo(HaveOccurred())
		Expect(response.Jobs).To(ContainElement(And(
			HaveField("JobID", Equal(*submitted.JobId)),
			HaveField("JobResources.Nodes.Allocation", Not(BeEmpty())),
		)))

		By("reading the same job through the client without resource details")
		job := newJob()
		Expect(cl.Get(ctx, submitted.GetKey(), job, &GetOptions{SkipCache: true})).To(Succeed())
		Expect(job.GetKey()).To(Equal(submitted.GetKey()))
		Expect(job).To(HaveField("JobResources", BeNil()))

		By("listing through the client without resource details")
		list := newList()
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
	Entry("v0042", SpecTimeout(time.Minute), "42", func() object.Object { return &v42.V0042JobInfo{} }, func() object.ObjectList { return &v42.V0042JobInfoObjectList{} }),
	Entry("v0043", SpecTimeout(time.Minute), "43", func() object.Object { return &v43.V0043JobInfo{} }, func() object.ObjectList { return &v43.V0043JobInfoObjectList{} }),
	Entry("v0044", SpecTimeout(time.Minute), "44", func() object.Object { return &v44.V0044JobInfo{} }, func() object.ObjectList { return &v44.V0044JobInfoObjectList{} }),
	Entry("v0045", SpecTimeout(time.Minute), "45", func() object.Object { return &v45.V0045JobInfo{} }, func() object.ObjectList { return &v45.V0045JobInfoObjectList{} }),
)
