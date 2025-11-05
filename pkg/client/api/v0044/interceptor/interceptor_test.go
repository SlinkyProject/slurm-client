// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
)

var _ = Describe("NewClient", func() {
	wrappedClient := &emptyClient{}
	ctx := context.Background()
	Context("SlurmV0044DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044DeleteJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobsWithResponse: func(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044DeleteJobsWithResponse(ctx, api.V0044KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteJobsWithResponse: func(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044DeleteJobsWithResponse(ctx, api.V0044KillJobsMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044DeleteNodeWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetJobWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobsStateWithResponse: func(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetJobsStateWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetLicensesWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetLicensesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetNodeWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetNodesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetPartitionWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetPartitionsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetPingWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetPingWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetReconfigureWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostReservationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostReservationWithResponse(ctx, api.SlurmV0044PostReservationJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostReservationWithResponse(ctx, api.SlurmV0044PostReservationJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044DeleteReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044DeleteReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044DeleteReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetReservationWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetReservationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostReservationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostReservationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationsWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostReservationsWithResponse(ctx, api.SlurmV0044PostReservationsJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostReservationsWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostReservationsWithResponse(ctx, api.SlurmV0044PostReservationsJSONRequestBody{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetResourcesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetResourcesWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetResourcesWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044GetSharesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044GetSharesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobAllocateWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobAllocateWithResponse: func(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobAllocateWithResponse(ctx, api.V0044JobAllocReq{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobAllocateWithResponse: func(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobAllocateWithResponse(ctx, api.V0044JobAllocReq{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobSubmitWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobSubmitWithResponse: func(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobSubmitWithResponse(ctx, api.V0044JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobSubmitWithResponse: func(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobSubmitWithResponse(ctx, api.V0044JobSubmitReq{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobWithBodyWithResponse: func(ctx context.Context, jobId, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostJobWithResponse(ctx, "", api.V0044JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostJobWithResponse(ctx, "", api.V0044JobDescMsg{}, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostNodeWithBodyWithResponse: func(ctx context.Context, nodeName, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmV0044PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmV0044PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmV0044PostNodeWithResponse(ctx, "", api.V0044UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmV0044PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmV0044PostNodeWithResponse(ctx, "", api.V0044UpdateNodeMsg{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteSingleQosWithResponse: func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteUserWithResponse: func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044DeleteWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044DeleteWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetAccountWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetAccountsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetAssociationWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetAssociationsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetClusterWithResponse: func(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetClusterWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetClustersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetConfigWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetConfigWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetDiagWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetInstanceWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetInstanceWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetInstancesWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetInstancesWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetJobWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetJobWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetJobWithResponse(ctx, "")
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetJobsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetJobsWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetQosWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetSingleQosWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetTresWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetTresWithResponse(ctx)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetUserWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetUserWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetUsersWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetWckeyWithResponse: func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetWckeyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044GetWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044GetWckeysWithResponse(ctx, nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsAssociationWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAccountsAssociationWithResponse(ctx, api.V0044OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsAssociationWithResponse: func(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAccountsAssociationWithResponse(ctx, api.V0044OpenapiAccountsAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsWithResponse: func(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAccountsWithResponse(ctx, api.V0044OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAccountsWithResponse: func(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAccountsWithResponse(ctx, api.V0044OpenapiAccountsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAssociationsWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAssociationsWithResponse: func(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostAssociationsWithResponse(ctx, api.V0044OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostAssociationsWithResponse: func(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostAssociationsWithResponse(ctx, api.V0044OpenapiAssocsResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostClustersWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostClustersWithResponse(ctx, nil, api.V0044OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostClustersWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostClustersWithResponse(ctx, nil, api.V0044OpenapiClustersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostConfigWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostConfigWithResponse: func(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostConfigWithResponse(ctx, api.V0044OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostConfigWithResponse: func(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostConfigWithResponse(ctx, api.V0044OpenapiSlurmdbdConfigResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostQosWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostQosWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostQosWithResponse(ctx, nil, api.V0044OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostQosWithResponse(ctx, nil, api.V0044OpenapiSlurmdbdQosResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostTresWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostTresWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostTresWithResponse: func(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostTresWithResponse(ctx, api.V0044OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostTresWithResponse: func(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostTresWithResponse(ctx, api.V0044OpenapiTresResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersAssociationWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostUsersAssociationWithResponse(ctx, nil, api.V0044OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostUsersAssociationWithResponse(ctx, nil, api.V0044OpenapiUsersAddCondResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersWithBodyWithResponse: func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersWithResponse: func(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostUsersWithResponse(ctx, api.V0044OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostUsersWithResponse: func(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostUsersWithResponse(ctx, api.V0044OpenapiUsersResp{})
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostWckeysWithBodyWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(called).To(BeTrue())
		})
	})
	Context("SlurmdbV0044PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			var called bool
			client := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			_, _ = client.SlurmdbV0044PostWckeysWithResponse(ctx, nil, api.V0044OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
		It("should call the underlying client if the provided function is nil", func() {
			var called bool
			client1 := NewClient(wrappedClient, Funcs{
				SlurmdbV0044PostWckeysWithResponse: func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
					called = true
					return nil, nil //nolint:nilnil
				},
			})
			client2 := NewClient(client1, Funcs{})
			_, _ = client2.SlurmdbV0044PostWckeysWithResponse(ctx, nil, api.V0044OpenapiWckeyResp{})
			Expect(called).To(BeTrue())
		})
	})
})

////////////////////////////////////////////////////////////////////////////////

type emptyClient struct{}

// SlurmV0044PostNodesWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostNodesWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNodesWithResponse(ctx context.Context, body api.SlurmV0044PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetPingWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044DeleteJobWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044DeleteJobsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044DeleteJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044DeleteJobsWithResponse(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044DeleteNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetDiagWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetJobWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetJobsStateWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetLicensesWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetNodesWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetNodesWithResponse(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetPartitionWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetPartitionsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetPingWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetReconfigureWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostReservationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostReservationWithResponse(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044DeleteReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetReservationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetReservationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetReservationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostReservationsWithResponse(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetResourcesWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetResourcesWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044GetSharesWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044GetSharesWithResponse(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobAllocateWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobAllocateWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobAllocateWithResponse(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobSubmitWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobSubmitWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobSubmitWithResponse(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostJobWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostJobWithResponse(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostNodeWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostNewNodeWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmV0044PostNewNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmV0044PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0044PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteAccountWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteClusterWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteSingleQosWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteUserWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044DeleteWckeyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetAccountWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetAccountsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetClusterWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetClustersWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetConfigWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetDiagWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetInstanceWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetInstancesWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetJobWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetQosWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetSingleQosWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetTresWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetUserWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetUsersWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetWckeyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044GetWckeysWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAccountsAssociationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAccountsAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAccountsAssociationWithResponse(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAccountsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAccountsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAccountsWithResponse(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAssociationsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostAssociationsWithResponse(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostClustersWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostClustersWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostConfigWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostConfigWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostConfigWithResponse(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostQosWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostQosWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostTresWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostTresWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostTresWithResponse(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostUsersAssociationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostUsersAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostUsersWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostUsersWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostUsersWithResponse(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostWckeysWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostWckeysWithResponse implements V0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostJobWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostJobWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0044PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostJobsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

// SlurmdbV0044PostJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (e *emptyClient) SlurmdbV0044PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0044PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	return nil, nil //nolint:nilnil
}

var _ api.ClientWithResponsesInterface = &emptyClient{}
