// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()
	Context("SlurmV0045DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045DeleteJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045DeleteJobsWithResponse(ctx, api.V0045KillJobsMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045DeleteNodeWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetJobsStateWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetLicensesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetNodeWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetNodesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetPartitionWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetPartitionsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetPingWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetPingWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetReconfigureWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostReservationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostReservationWithResponse(ctx, api.SlurmV0045PostReservationJSONRequestBody{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045DeleteReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045DeleteReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetReservationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostReservationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostReservationsWithResponse(ctx, api.SlurmV0045PostReservationsJSONRequestBody{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetResourcesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetResourcesWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045GetSharesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobAllocateWithResponse(ctx, api.V0045JobAllocReq{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobSubmitWithResponse(ctx, api.V0045JobSubmitReq{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostJobWithResponse(ctx, "", api.V0045JobDescMsg{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNodeWithResponse(ctx, "", api.V0045UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNodesWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNodesWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNodesWithResponse(ctx, api.V0045UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNewNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNewNodeWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0045PostNewNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0045PostNewNodeWithResponse(ctx, api.V0045OpenapiCreateNodeReq{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045DeleteWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetAccountsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetClustersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetConfigWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetInstanceWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetInstancesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetJobWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetQosWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetTresWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetUsersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045GetWckeysWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAccountsAssociationWithResponse(ctx, api.V0045OpenapiAccountsAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAccountsWithResponse(ctx, api.V0045OpenapiAccountsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostAssociationsWithResponse(ctx, api.V0045OpenapiAssocsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostClustersWithResponse(ctx, nil, api.V0045OpenapiClustersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostConfigWithResponse(ctx, api.V0045OpenapiSlurmdbdConfigResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostQosWithResponse(ctx, nil, api.V0045OpenapiSlurmdbdQosResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostTresWithResponse(ctx, api.V0045OpenapiTresResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostUsersAssociationWithResponse(ctx, nil, api.V0045OpenapiUsersAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostUsersWithResponse(ctx, api.V0045OpenapiUsersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0045PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0045PostWckeysWithResponse(ctx, nil, api.V0045OpenapiWckeyResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})
