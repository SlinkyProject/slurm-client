// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0040DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobsWithResponse: func(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040DeleteJobsWithResponse(ctx, api.V0040KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteJobsWithResponse: func(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040DeleteJobsWithResponse(ctx, api.V0040KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostJobSubmitWithResponse(ctx, api.V0040JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostJobSubmitWithResponse(ctx, api.V0040JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostJobWithResponse(ctx, "", api.V0040JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostJobWithResponse(ctx, "", api.V0040JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0040PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0040PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmV0040PostNodeWithResponse(ctx, "", api.V0040UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0040PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0040PostNodeWithResponse(ctx, "", api.V0040UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAccountsAssociationWithResponse(ctx, api.V0040OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAccountsAssociationWithResponse(ctx, api.V0040OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsWithResponse: func(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAccountsWithResponse(ctx, api.V0040OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAccountsWithResponse: func(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAccountsWithResponse(ctx, api.V0040OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAssociationsWithResponse: func(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostAssociationsWithResponse(ctx, api.V0040OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostAssociationsWithResponse: func(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostAssociationsWithResponse(ctx, api.V0040OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostClustersWithResponse(ctx, nil, api.V0040OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostClustersWithResponse(ctx, nil, api.V0040OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostConfigWithResponse: func(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostConfigWithResponse(ctx, api.V0040OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostConfigWithResponse: func(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostConfigWithResponse(ctx, api.V0040OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostQosWithResponse(ctx, nil, api.V0040OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostQosWithResponse(ctx, nil, api.V0040OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostTresWithResponse: func(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostTresWithResponse(ctx, api.V0040OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostTresWithResponse: func(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostTresWithResponse(ctx, api.V0040OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostUsersAssociationWithResponse(ctx, nil, api.V0040OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostUsersAssociationWithResponse(ctx, nil, api.V0040OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersWithResponse: func(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostUsersWithResponse(ctx, api.V0040OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostUsersWithResponse: func(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostUsersWithResponse(ctx, api.V0040OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0040PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			_, _ = client.SlurmdbV0040PostWckeysWithResponse(ctx, nil, api.V0040OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0040PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
					called = true
					return nil, nil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0040PostWckeysWithResponse(ctx, nil, api.V0040OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0040DeleteJobWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
	return nil, nil
}

// SlurmV0040DeleteJobsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0040DeleteJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040DeleteJobsWithResponse(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	return nil, nil
}

// SlurmV0040DeleteNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
	return nil, nil
}

// SlurmV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error) {
	return nil, nil
}

// SlurmV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
	return nil, nil
}

// SlurmV0040GetJobsStateWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error) {
	return nil, nil
}

// SlurmV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
	return nil, nil
}

// SlurmV0040GetLicensesWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error) {
	return nil, nil
}

// SlurmV0040GetNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
	return nil, nil
}

// SlurmV0040GetNodesWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetNodesWithResponse(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
	return nil, nil
}

// SlurmV0040GetPartitionWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
	return nil, nil
}

// SlurmV0040GetPartitionsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
	return nil, nil
}

// SlurmV0040GetPingWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error) {
	return nil, nil
}

// SlurmV0040GetReconfigureWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error) {
	return nil, nil
}

// SlurmV0040GetReservationWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error) {
	return nil, nil
}

// SlurmV0040GetReservationsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error) {
	return nil, nil
}

// SlurmV0040GetSharesWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040GetSharesWithResponse(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error) {
	return nil, nil
}

// SlurmV0040PostJobSubmitWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0040PostJobSubmitWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostJobSubmitWithResponse(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	return nil, nil
}

// SlurmV0040PostJobWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	return nil, nil
}

// SlurmV0040PostJobWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostJobWithResponse(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	return nil, nil
}

// SlurmV0040PostNodeWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	return nil, nil
}

// SlurmV0040PostNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0040PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteUserWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error) {
	return nil, nil
}

// SlurmdbV0040DeleteWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetInstanceWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetInstancesWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetQosWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetTresWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetUserWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error) {
	return nil, nil
}

// SlurmdbV0040GetWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAccountsAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAccountsAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAccountsAssociationWithResponse(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAccountsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAccountsWithResponse(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAssociationsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostAssociationsWithResponse(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostClustersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostConfigWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostConfigWithResponse(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostQosWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostQosWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostTresWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostTresWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostTresWithResponse(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostUsersAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostUsersAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostUsersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostUsersWithResponse(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostWckeysWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	return nil, nil
}

// SlurmdbV0040PostWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0040PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	return nil, nil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
