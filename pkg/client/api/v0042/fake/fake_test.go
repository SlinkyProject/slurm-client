// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()
	Context("SlurmV0042DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042DeleteJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042DeleteJobsWithResponse(ctx, api.V0042KillJobsMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042DeleteNodeWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetJobsStateWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetLicensesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetNodeWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetNodesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetPartitionWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetPartitionsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetPingWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetPingWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetReconfigureWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetReservationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042GetSharesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobAllocateWithResponse(ctx, api.V0042JobAllocReq{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobSubmitWithResponse(ctx, api.V0042JobSubmitReq{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostJobWithResponse(ctx, "", api.V0042JobDescMsg{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0042PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0042PostNodeWithResponse(ctx, "", api.V0042UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042DeleteWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetAccountsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetClustersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetConfigWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetInstanceWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetInstancesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetJobWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetQosWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetTresWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetUsersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042GetWckeysWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAccountsAssociationWithResponse(ctx, api.V0042OpenapiAccountsAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAccountsWithResponse(ctx, api.V0042OpenapiAccountsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostAssociationsWithResponse(ctx, api.V0042OpenapiAssocsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostClustersWithResponse(ctx, nil, api.V0042OpenapiClustersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostConfigWithResponse(ctx, api.V0042OpenapiSlurmdbdConfigResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostQosWithResponse(ctx, nil, api.V0042OpenapiSlurmdbdQosResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostTresWithResponse(ctx, api.V0042OpenapiTresResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostUsersAssociationWithResponse(ctx, nil, api.V0042OpenapiUsersAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostUsersWithResponse(ctx, api.V0042OpenapiUsersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0042PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0042PostWckeysWithResponse(ctx, nil, api.V0042OpenapiWckeyResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})
