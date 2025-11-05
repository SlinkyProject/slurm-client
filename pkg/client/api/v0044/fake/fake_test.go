// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()
	Context("SlurmV0044DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044DeleteJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044DeleteJobsWithResponse(ctx, api.V0044KillJobsMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044DeleteNodeWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetJobsStateWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetLicensesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetNodeWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetNodesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetPartitionWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetPartitionsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetPingWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetPingWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetReconfigureWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostReservationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostReservationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostReservationWithResponse(ctx, api.SlurmV0044PostReservationJSONRequestBody{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044DeleteReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044DeleteReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetReservationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostReservationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostReservationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostReservationsWithResponse(ctx, api.SlurmV0044PostReservationsJSONRequestBody{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetResourcesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetResourcesWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044GetSharesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobAllocateWithResponse(ctx, api.V0044JobAllocReq{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobSubmitWithResponse(ctx, api.V0044JobSubmitReq{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostJobWithResponse(ctx, "", api.V0044JobDescMsg{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0044PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0044PostNodeWithResponse(ctx, "", api.V0044UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044DeleteWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetAccountsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetClustersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetConfigWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetInstanceWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetInstancesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetJobWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetQosWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetTresWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetUsersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044GetWckeysWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAccountsAssociationWithResponse(ctx, api.V0044OpenapiAccountsAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAccountsWithResponse(ctx, api.V0044OpenapiAccountsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostAssociationsWithResponse(ctx, api.V0044OpenapiAssocsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostClustersWithResponse(ctx, nil, api.V0044OpenapiClustersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostConfigWithResponse(ctx, api.V0044OpenapiSlurmdbdConfigResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostQosWithResponse(ctx, nil, api.V0044OpenapiSlurmdbdQosResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostTresWithResponse(ctx, api.V0044OpenapiTresResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostUsersAssociationWithResponse(ctx, nil, api.V0044OpenapiUsersAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostUsersWithResponse(ctx, api.V0044OpenapiUsersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0044PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0044PostWckeysWithResponse(ctx, nil, api.V0044OpenapiWckeyResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})
