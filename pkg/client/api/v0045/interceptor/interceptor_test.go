// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0045DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobsWithResponse: func(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045DeleteJobsWithResponse(ctx, api.V0045KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteJobsWithResponse: func(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045DeleteJobsWithResponse(ctx, api.V0045KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostReservationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostReservationWithResponse(ctx, api.SlurmV0045PostReservationJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostReservationWithResponse(ctx, api.SlurmV0045PostReservationJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045DeleteReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045DeleteReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045DeleteReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostReservationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationsWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostReservationsWithResponse(ctx, api.SlurmV0045PostReservationsJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostReservationsWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostReservationsWithResponse(ctx, api.SlurmV0045PostReservationsJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetResourcesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetResourcesWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetResourcesWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobAllocateWithResponse: func(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobAllocateWithResponse(ctx, api.V0045JobAllocReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobAllocateWithResponse: func(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobAllocateWithResponse(ctx, api.V0045JobAllocReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobSubmitWithResponse: func(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobSubmitWithResponse(ctx, api.V0045JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobSubmitWithResponse: func(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobSubmitWithResponse(ctx, api.V0045JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostJobWithResponse(ctx, "", api.V0045JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostJobWithResponse(ctx, "", api.V0045JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNodeWithResponse(ctx, "", api.V0045UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNodeWithResponse(ctx, "", api.V0045UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNodesWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodesWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNodesWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodesWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNodesWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodesWithResponse: func(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNodesWithResponse(ctx, api.V0045UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNodesWithResponse: func(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNodesWithResponse(ctx, api.V0045UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNewNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNewNodeWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNewNodeWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNewNodeWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNewNodeWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0045PostNewNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNewNodeWithResponse: func(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0045PostNewNodeWithResponse(ctx, api.V0045OpenapiCreateNodeReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0045PostNewNodeWithResponse: func(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0045PostNewNodeWithResponse(ctx, api.V0045OpenapiCreateNodeReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAccountsAssociationWithResponse(ctx, api.V0045OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAccountsAssociationWithResponse(ctx, api.V0045OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsWithResponse: func(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAccountsWithResponse(ctx, api.V0045OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAccountsWithResponse: func(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAccountsWithResponse(ctx, api.V0045OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAssociationsWithResponse: func(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostAssociationsWithResponse(ctx, api.V0045OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostAssociationsWithResponse: func(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostAssociationsWithResponse(ctx, api.V0045OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostClustersWithResponse(ctx, nil, api.V0045OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostClustersWithResponse(ctx, nil, api.V0045OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostConfigWithResponse: func(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostConfigWithResponse(ctx, api.V0045OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostConfigWithResponse: func(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostConfigWithResponse(ctx, api.V0045OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostQosWithResponse(ctx, nil, api.V0045OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostQosWithResponse(ctx, nil, api.V0045OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostTresWithResponse: func(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostTresWithResponse(ctx, api.V0045OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostTresWithResponse: func(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostTresWithResponse(ctx, api.V0045OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostUsersAssociationWithResponse(ctx, nil, api.V0045OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostUsersAssociationWithResponse(ctx, nil, api.V0045OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersWithResponse: func(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostUsersWithResponse(ctx, api.V0045OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostUsersWithResponse: func(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostUsersWithResponse(ctx, api.V0045OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0045PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0045PostWckeysWithResponse(ctx, nil, api.V0045OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0045PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0045PostWckeysWithResponse(ctx, nil, api.V0045OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0045PostNodesWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostNodesWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNodesWithResponse(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetPingWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeleteJobWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeleteJobsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeleteJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeleteJobsWithResponse(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeleteNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeletePartitionWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeletePartitionWithResponse(ctx context.Context, partitionName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeletePartitionResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetConfWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetConfWithResponse(ctx context.Context, params *api.SlurmV0045GetConfParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetConfResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetDiagWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetJobRequeueWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetJobRequeueWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobRequeueParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobRequeueResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetJobWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetJobsStateWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetLicensesWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetNodesWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetNodesWithResponse(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetPartitionWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetPartitionsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetPingWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetReconfigureWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostReservationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostReservationWithResponse(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045DeleteReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetReservationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetReservationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetReservationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostReservationsWithResponse(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetResourcesWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetResourcesWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045GetSharesWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045GetSharesWithResponse(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobAllocateWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobAllocateWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobAllocateWithResponse(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobSubmitWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobSubmitWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobSubmitWithResponse(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobWithResponse(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostNodeWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobsRequeueWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobsRequeueWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostJobsRequeueWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostJobsRequeueWithResponse(ctx context.Context, body api.SlurmV0045PostJobsRequeueJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostPartitionsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostPartitionsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostPartitionsWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostPartitionsWithResponse(ctx context.Context, body api.SlurmV0045PostPartitionsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostNewNodeWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0045PostNewNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0045PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteAccountWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteClusterWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteSingleQosWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteUserWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045DeleteWckeyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetAccountWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetAccountsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetClusterWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetClustersWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetConfWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetConfWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetConfigWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetDiagWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetInstanceWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetInstancesWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetJobWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetQosWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetSingleQosWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetTresWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetUserWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetUsersWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetWckeyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045GetWckeysWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAccountsAssociationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAccountsAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAccountsAssociationWithResponse(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAccountsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAccountsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAccountsWithResponse(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAssociationsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostAssociationsWithResponse(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostClustersWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostClustersWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostConfigWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostConfigWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostConfigWithResponse(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostQosWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostQosWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostTresWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostTresWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostTresWithResponse(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostUsersAssociationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostUsersAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostUsersWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostUsersWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostUsersWithResponse(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostWckeysWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostWckeysWithResponse implements V0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostJobWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostJobWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0045PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostJobsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0045PostJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0045PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0045PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
