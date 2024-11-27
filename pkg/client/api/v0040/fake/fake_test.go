// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()
	Context("SlurmV0040DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040DeleteJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040DeleteJobsWithResponse(ctx, api.V0040KillJobsMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040DeleteNodeWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetJobsStateWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetLicensesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetNodeWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetNodesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetPartitionWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetPartitionsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetPingWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetPingWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetReconfigureWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetReservationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040GetSharesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostJobSubmitWithResponse(ctx, api.V0040JobSubmitReq{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostJobWithResponse(ctx, "", api.V0040JobDescMsg{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0040PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0040PostNodeWithResponse(ctx, "", api.V0040UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040DeleteWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetAccountsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetClustersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetConfigWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetInstanceWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetInstancesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetJobWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetQosWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetTresWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetUsersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040GetWckeysWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAccountsAssociationWithResponse(ctx, api.V0040OpenapiAccountsAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAccountsWithResponse(ctx, api.V0040OpenapiAccountsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostAssociationsWithResponse(ctx, api.V0040OpenapiAssocsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostClustersWithResponse(ctx, nil, api.V0040OpenapiClustersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostConfigWithResponse(ctx, api.V0040OpenapiSlurmdbdConfigResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostQosWithResponse(ctx, nil, api.V0040OpenapiSlurmdbdQosResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostTresWithResponse(ctx, api.V0040OpenapiTresResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostUsersAssociationWithResponse(ctx, nil, api.V0040OpenapiUsersAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostUsersWithResponse(ctx, api.V0040OpenapiUsersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0040PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0040PostWckeysWithResponse(ctx, nil, api.V0040OpenapiWckeyResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})
