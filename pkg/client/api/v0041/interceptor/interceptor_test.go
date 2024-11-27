// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0041DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobsWithResponse: func(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041DeleteJobsWithResponse(ctx, api.V0041KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteJobsWithResponse: func(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041DeleteJobsWithResponse(ctx, api.V0041KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobAllocateWithResponse: func(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobAllocateWithResponse(ctx, api.V0041JobAllocReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobAllocateWithResponse: func(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobAllocateWithResponse(ctx, api.V0041JobAllocReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobSubmitWithResponse: func(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobSubmitWithResponse(ctx, api.V0041JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobSubmitWithResponse: func(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobSubmitWithResponse(ctx, api.V0041JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostJobWithResponse(ctx, "", api.V0041JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostJobWithResponse(ctx, "", api.V0041JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0041PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0041PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0041PostNodeWithResponse(ctx, "", api.V0041UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0041PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0041PostNodeWithResponse(ctx, "", api.V0041UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAccountsAssociationWithResponse(ctx, api.V0041OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAccountsAssociationWithResponse(ctx, api.V0041OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsWithResponse: func(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAccountsWithResponse(ctx, api.V0041OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAccountsWithResponse: func(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAccountsWithResponse(ctx, api.V0041OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAssociationsWithResponse: func(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostAssociationsWithResponse(ctx, api.V0041OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostAssociationsWithResponse: func(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostAssociationsWithResponse(ctx, api.V0041OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostClustersWithResponse(ctx, nil, api.V0041OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostClustersWithResponse(ctx, nil, api.V0041OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostConfigWithResponse: func(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostConfigWithResponse(ctx, api.V0041OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostConfigWithResponse: func(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostConfigWithResponse(ctx, api.V0041OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostQosWithResponse(ctx, nil, api.V0041OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostQosWithResponse(ctx, nil, api.V0041OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostTresWithResponse: func(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostTresWithResponse(ctx, api.V0041OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostTresWithResponse: func(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostTresWithResponse(ctx, api.V0041OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostUsersAssociationWithResponse(ctx, nil, api.V0041OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostUsersAssociationWithResponse(ctx, nil, api.V0041OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersWithResponse: func(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostUsersWithResponse(ctx, api.V0041OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostUsersWithResponse: func(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostUsersWithResponse(ctx, api.V0041OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0041PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0041PostWckeysWithResponse(ctx, nil, api.V0041OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0041PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0041PostWckeysWithResponse(ctx, nil, api.V0041OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0041DeleteJobWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error) {
	return nil, nil
}

// SlurmV0041DeleteJobsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0041DeleteJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041DeleteJobsWithResponse(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0041DeleteNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error) {
	return nil, nil
}

// SlurmV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
	return nil, nil
}

// SlurmV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error) {
	return nil, nil
}

// SlurmV0041GetJobsStateWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error) {
	return nil, nil
}

// SlurmV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error) {
	return nil, nil
}

// SlurmV0041GetLicensesWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error) {
	return nil, nil
}

// SlurmV0041GetNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error) {
	return nil, nil
}

// SlurmV0041GetNodesWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetNodesWithResponse(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error) {
	return nil, nil
}

// SlurmV0041GetPartitionWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error) {
	return nil, nil
}

// SlurmV0041GetPartitionsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error) {
	return nil, nil
}

// SlurmV0041GetPingWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error) {
	return nil, nil
}

// SlurmV0041GetReconfigureWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
	return nil, nil
}

// SlurmV0041GetReservationWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
	return nil, nil
}

// SlurmV0041GetReservationsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
	return nil, nil
}

// SlurmV0041GetSharesWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041GetSharesWithResponse(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobAllocateWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobAllocateWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobAllocateWithResponse(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobSubmitWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobSubmitWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobSubmitWithResponse(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	return nil, nil
}

// SlurmV0041PostJobWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostJobWithResponse(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	return nil, nil
}

// SlurmV0041PostNodeWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	return nil, nil
}

// SlurmV0041PostNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0041PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteUserWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error) {
	return nil, nil
}

// SlurmdbV0041DeleteWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetInstanceWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetInstancesWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetQosWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetTresWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetUserWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0041GetWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAccountsAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAccountsAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAccountsAssociationWithResponse(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAccountsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAccountsWithResponse(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAssociationsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostAssociationsWithResponse(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostClustersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostConfigWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostConfigWithResponse(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostQosWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostQosWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostTresWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostTresWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostTresWithResponse(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostUsersAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostUsersAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostUsersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostUsersWithResponse(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostWckeysWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0041PostWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0041PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	return nil, nil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
