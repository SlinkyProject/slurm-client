// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0043DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobsWithResponse: func(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043DeleteJobsWithResponse(ctx, api.V0043KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteJobsWithResponse: func(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043DeleteJobsWithResponse(ctx, api.V0043KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobAllocateWithResponse: func(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobAllocateWithResponse(ctx, api.V0043JobAllocReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobAllocateWithResponse: func(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobAllocateWithResponse(ctx, api.V0043JobAllocReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobSubmitWithResponse(ctx, api.V0043JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobSubmitWithResponse(ctx, api.V0043JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostJobWithResponse(ctx, "", api.V0043JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostJobWithResponse(ctx, "", api.V0043JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0043PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0043PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0043PostNodeWithResponse(ctx, "", api.V0043UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0043PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0043PostNodeWithResponse(ctx, "", api.V0043UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAccountsAssociationWithResponse(ctx, api.V0043OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAccountsAssociationWithResponse(ctx, api.V0043OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsWithResponse: func(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAccountsWithResponse(ctx, api.V0043OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAccountsWithResponse: func(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAccountsWithResponse(ctx, api.V0043OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAssociationsWithResponse: func(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostAssociationsWithResponse(ctx, api.V0043OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostAssociationsWithResponse: func(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostAssociationsWithResponse(ctx, api.V0043OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostClustersWithResponse(ctx, nil, api.V0043OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostClustersWithResponse(ctx, nil, api.V0043OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostConfigWithResponse: func(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostConfigWithResponse(ctx, api.V0043OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostConfigWithResponse: func(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostConfigWithResponse(ctx, api.V0043OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostQosWithResponse(ctx, nil, api.V0043OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostQosWithResponse(ctx, nil, api.V0043OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostTresWithResponse: func(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostTresWithResponse(ctx, api.V0043OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostTresWithResponse: func(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostTresWithResponse(ctx, api.V0043OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostUsersAssociationWithResponse(ctx, nil, api.V0043OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostUsersAssociationWithResponse(ctx, nil, api.V0043OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersWithResponse: func(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostUsersWithResponse(ctx, api.V0043OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostUsersWithResponse: func(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostUsersWithResponse(ctx, api.V0043OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0043PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0043PostWckeysWithResponse(ctx, nil, api.V0043OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0043PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0043PostWckeysWithResponse(ctx, nil, api.V0043OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0043PostNodesWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	return nil, nil
}

// SlurmV0043PostNodesWithResponse implements v0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostNodesWithResponse(ctx context.Context, body api.SlurmV0043PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetPingWithResponse implements v0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetPingResponse, error) {
	return nil, nil
}

// SlurmV0043DeleteJobWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
	return nil, nil
}

// SlurmV0043DeleteJobsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0043DeleteJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043DeleteJobsWithResponse(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0043DeleteNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error) {
	return nil, nil
}

// SlurmV0043GetDiagWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error) {
	return nil, nil
}

// SlurmV0043GetJobWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
	return nil, nil
}

// SlurmV0043GetJobsStateWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error) {
	return nil, nil
}

// SlurmV0043GetJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
	return nil, nil
}

// SlurmV0043GetLicensesWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error) {
	return nil, nil
}

// SlurmV0043GetNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error) {
	return nil, nil
}

// SlurmV0043GetNodesWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetNodesWithResponse(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error) {
	return nil, nil
}

// SlurmV0043GetPartitionWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
	return nil, nil
}

// SlurmV0043GetPartitionsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
	return nil, nil
}

// SlurmV0043GetPingWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
	return nil, nil
}

// SlurmV0043GetReconfigureWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
	return nil, nil
}

// SlurmV0043PostReservationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	return nil, nil
}

// SlurmV0043PostReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostReservationWithResponse(ctx context.Context, body api.SlurmV0043PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	return nil, nil
}

// SlurmV0043DeleteReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteReservationResponse, error) {
	return nil, nil
}

// SlurmV0043GetReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error) {
	return nil, nil
}

// SlurmV0043GetReservationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error) {
	return nil, nil
}

// SlurmV0043GetReservationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	return nil, nil
}

// SlurmV0043GetReservationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostReservationsWithResponse(ctx context.Context, body api.SlurmV0043PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	return nil, nil
}

// SlurmV0043GetSharesWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043GetSharesWithResponse(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobAllocateWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobAllocateWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobAllocateWithResponse(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobSubmitWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobSubmitWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobSubmitWithResponse(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	return nil, nil
}

// SlurmV0043PostJobWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostJobWithResponse(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	return nil, nil
}

// SlurmV0043PostNodeWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	return nil, nil
}

// SlurmV0043PostNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0043PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteAccountWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteClusterWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteSingleQosWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteUserWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error) {
	return nil, nil
}

// SlurmdbV0043DeleteWckeyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetAccountWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetAccountsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetClusterWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetClustersWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetConfigWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetDiagWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetInstanceWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetInstancesWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetJobWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetQosWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetSingleQosWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetTresWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetUserWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetUsersWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetWckeyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0043GetWckeysWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAccountsAssociationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAccountsAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAccountsAssociationWithResponse(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAccountsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAccountsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAccountsWithResponse(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAssociationsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostAssociationsWithResponse(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostClustersWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostClustersWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostConfigWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostConfigWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostConfigWithResponse(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostQosWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostQosWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostTresWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostTresWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostTresWithResponse(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostUsersAssociationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostUsersAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostUsersWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostUsersWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostUsersWithResponse(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostWckeysWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0043PostWckeysWithResponse implements V0043.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0043PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	return nil, nil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
