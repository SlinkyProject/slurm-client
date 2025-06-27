// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0042DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobsWithResponse: func(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042DeleteJobsWithResponse(ctx, api.V0042KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteJobsWithResponse: func(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042DeleteJobsWithResponse(ctx, api.V0042KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobAllocateWithResponse(ctx, api.V0042JobAllocReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobAllocateWithResponse(ctx, api.V0042JobAllocReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobSubmitWithResponse(ctx, api.V0042JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobSubmitWithResponse(ctx, api.V0042JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostJobWithResponse(ctx, "", api.V0042JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostJobWithResponse(ctx, "", api.V0042JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0042PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0042PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0042PostNodeWithResponse(ctx, "", api.V0042UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0042PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0042PostNodeWithResponse(ctx, "", api.V0042UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAccountsAssociationWithResponse(ctx, api.V0042OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAccountsAssociationWithResponse(ctx, api.V0042OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsWithResponse: func(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAccountsWithResponse(ctx, api.V0042OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAccountsWithResponse: func(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAccountsWithResponse(ctx, api.V0042OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAssociationsWithResponse: func(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostAssociationsWithResponse(ctx, api.V0042OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostAssociationsWithResponse: func(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostAssociationsWithResponse(ctx, api.V0042OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostClustersWithResponse(ctx, nil, api.V0042OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostClustersWithResponse(ctx, nil, api.V0042OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostConfigWithResponse: func(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostConfigWithResponse(ctx, api.V0042OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostConfigWithResponse: func(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostConfigWithResponse(ctx, api.V0042OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostQosWithResponse(ctx, nil, api.V0042OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostQosWithResponse(ctx, nil, api.V0042OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostTresWithResponse: func(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostTresWithResponse(ctx, api.V0042OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostTresWithResponse: func(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostTresWithResponse(ctx, api.V0042OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostUsersAssociationWithResponse(ctx, nil, api.V0042OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostUsersAssociationWithResponse(ctx, nil, api.V0042OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersWithResponse: func(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostUsersWithResponse(ctx, api.V0042OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostUsersWithResponse: func(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostUsersWithResponse(ctx, api.V0042OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0042PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0042PostWckeysWithResponse(ctx, nil, api.V0042OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0042PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0042PostWckeysWithResponse(ctx, nil, api.V0042OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0042PostNodesWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostNodesWithResponse implements v0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostNodesWithResponse(ctx context.Context, body api.SlurmV0042PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetPingWithResponse implements v0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042DeleteJobWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042DeleteJobsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042DeleteJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042DeleteJobsWithResponse(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042DeleteNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetDiagWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetJobWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetJobsStateWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetLicensesWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetNodesWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetNodesWithResponse(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetPartitionWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetPartitionsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetPingWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetReconfigureWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetReservationWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetReservationsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042GetSharesWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042GetSharesWithResponse(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobAllocateWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobAllocateWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobAllocateWithResponse(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobSubmitWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobSubmitWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobSubmitWithResponse(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostJobWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostJobWithResponse(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostNodeWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0042PostNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0042PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteAccountWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteClusterWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteSingleQosWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteUserWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042DeleteWckeyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetAccountWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetAccountsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetClusterWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetClustersWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetConfigWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetDiagWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetInstanceWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetInstancesWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetJobWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetQosWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetSingleQosWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetTresWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetUserWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetUsersWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetWckeyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042GetWckeysWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAccountsAssociationWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAccountsAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAccountsAssociationWithResponse(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAccountsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAccountsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAccountsWithResponse(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAssociationsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostAssociationsWithResponse(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostClustersWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostClustersWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostConfigWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostConfigWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostConfigWithResponse(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostQosWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostQosWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostTresWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostTresWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostTresWithResponse(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostUsersAssociationWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostUsersAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostUsersWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostUsersWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostUsersWithResponse(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostWckeysWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0042PostWckeysWithResponse implements V0042.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0042PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
