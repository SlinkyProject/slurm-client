// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
)

var _ = Describe("NewFakeClient", func() {
	ctx := context.Background()
	Context("SlurmV0043DeleteJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043DeleteJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043DeleteJobsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043DeleteJobsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043DeleteJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043DeleteJobsWithResponse(ctx, api.V0043KillJobsMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043DeleteNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043DeleteNodeWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetJobWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetJobsStateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetJobsStateWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetLicensesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetLicensesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetNodeWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetNodesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetNodesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetPartitionWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetPartitionWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetPartitionsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetPartitionsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetPingWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetPingWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetReconfigureWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetReconfigureWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetReservationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetReservationWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetReservationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetReservationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043GetSharesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043GetSharesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobAllocateWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobAllocateWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobAllocateWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobAllocateWithResponse(ctx, api.V0043JobAllocReq{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobSubmitWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobSubmitWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobSubmitWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobSubmitWithResponse(ctx, api.V0043JobSubmitReq{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostJobWithResponse(ctx, "", api.V0043JobDescMsg{}, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostNodeWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostNodeWithBodyWithResponse(ctx, "", "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmV0043PostNodeWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmV0043PostNodeWithResponse(ctx, "", api.V0043UpdateNodeMsg{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043DeleteWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043DeleteWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetAccountWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetAccountWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetAccountsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetAssociationWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetAssociationsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetClusterWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetClusterWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetClustersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetConfigWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetDiagWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetDiagWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetInstanceWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetInstanceWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetInstancesWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetInstancesWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetJobWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetJobWithResponse(ctx, "")
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetJobsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetJobsWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetQosWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetSingleQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetSingleQosWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetTresWithResponse(ctx)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetUserWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetUserWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetUsersWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetWckeyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetWckeyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043GetWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043GetWckeysWithResponse(ctx, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAccountsAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAccountsAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAccountsAssociationWithResponse(ctx, api.V0043OpenapiAccountsAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAccountsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAccountsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAccountsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAccountsWithResponse(ctx, api.V0043OpenapiAccountsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAssociationsWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostAssociationsWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostAssociationsWithResponse(ctx, api.V0043OpenapiAssocsResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostClustersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostClustersWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostClustersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostClustersWithResponse(ctx, nil, api.V0043OpenapiClustersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostConfigWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostConfigWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostConfigWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostConfigWithResponse(ctx, api.V0043OpenapiSlurmdbdConfigResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostQosWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostQosWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostQosWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostQosWithResponse(ctx, nil, api.V0043OpenapiSlurmdbdQosResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostTresWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostTresWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostTresWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostTresWithResponse(ctx, api.V0043OpenapiTresResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostUsersAssociationWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostUsersAssociationWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostUsersAssociationWithResponse(ctx, nil, api.V0043OpenapiUsersAddCondResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostUsersWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostUsersWithBodyWithResponse(ctx, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostUsersWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostUsersWithResponse(ctx, api.V0043OpenapiUsersResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostWckeysWithBodyWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostWckeysWithBodyWithResponse(ctx, nil, "", nil)
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
	Context("SlurmdbV0043PostWckeysWithResponse", func() {
		It("should call the provided function", func() {
			client := NewFakeClient()
			res, err := client.SlurmdbV0043PostWckeysWithResponse(ctx, nil, api.V0043OpenapiWckeyResp{})
			Expect(err).To(BeNil())
			Expect(res).To(Not(BeNil()))
		})
	})
})
