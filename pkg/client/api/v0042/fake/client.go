// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0042DeleteJobWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
	res := &api.SlurmV0042DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiKillJobResp{},
	}
	return res, nil
}

// SlurmV0042DeleteJobsWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	res := &api.SlurmV0042DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0042DeleteJobsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042DeleteJobsWithResponse(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	res := &api.SlurmV0042DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0042DeleteNodeWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error) {
	res := &api.SlurmV0042DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmV0042GetDiagWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
	res := &api.SlurmV0042GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0042GetJobWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
	res := &api.SlurmV0042GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0042GetJobsStateWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error) {
	res := &api.SlurmV0042GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0042GetJobsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
	res := &api.SlurmV0042GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0042GetLicensesWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error) {
	res := &api.SlurmV0042GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0042GetNodeWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error) {
	res := &api.SlurmV0042GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0042GetNodesWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetNodesWithResponse(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error) {
	res := &api.SlurmV0042GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0042GetPartitionWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error) {
	res := &api.SlurmV0042GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0042GetPartitionsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error) {
	res := &api.SlurmV0042GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0042GetPingWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error) {
	res := &api.SlurmV0042GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0042GetReconfigureWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error) {
	res := &api.SlurmV0042GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmV0042GetReservationWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error) {
	res := &api.SlurmV0042GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0042GetReservationsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error) {
	res := &api.SlurmV0042GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0042GetSharesWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042GetSharesWithResponse(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error) {
	res := &api.SlurmV0042GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0042PostJobAllocateWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	res := &api.SlurmV0042PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0042PostJobAllocateWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobAllocateWithResponse(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	res := &api.SlurmV0042PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0042PostJobSubmitWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	res := &api.SlurmV0042PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0042PostJobSubmitWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobSubmitWithResponse(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	res := &api.SlurmV0042PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0042PostJobWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	res := &api.SlurmV0042PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0042PostJobWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostJobWithResponse(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	res := &api.SlurmV0042PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0042PostNodeWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	res := &api.SlurmV0042PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmV0042PostNodeWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	res := &api.SlurmV0042PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmV0042PostNodesWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	res := &api.SlurmV0042PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmV0042PostNodesWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0042PostNodesWithResponse(ctx context.Context, body api.SlurmV0042PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	res := &api.SlurmV0042PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteAccountWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error) {
	res := &api.SlurmdbV0042DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteAssociationWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0042DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteAssociationsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0042DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteClusterWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error) {
	res := &api.SlurmdbV0042DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteSingleQosWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0042DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteUserWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error) {
	res := &api.SlurmdbV0042DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042DeleteWckeyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0042DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0042GetAccountWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
	res := &api.SlurmdbV0042GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetAccountsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
	res := &api.SlurmdbV0042GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetAssociationWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error) {
	res := &api.SlurmdbV0042GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetAssociationsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
	res := &api.SlurmdbV0042GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetClusterWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error) {
	res := &api.SlurmdbV0042GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0042GetClustersWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error) {
	res := &api.SlurmdbV0042GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0042GetConfigWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error) {
	res := &api.SlurmdbV0042GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0042GetDiagWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error) {
	res := &api.SlurmdbV0042GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetInstanceWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error) {
	res := &api.SlurmdbV0042GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0042GetInstancesWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error) {
	res := &api.SlurmdbV0042GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0042GetJobWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error) {
	res := &api.SlurmdbV0042GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetJobsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error) {
	res := &api.SlurmdbV0042GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0042GetPingWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetPingResponse, error) {
	res := &api.SlurmdbV0042GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdPingResp{},
	}
	return res, nil
}

// SlurmdbV0042GetQosWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error) {
	res := &api.SlurmdbV0042GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0042GetSingleQosWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error) {
	res := &api.SlurmdbV0042GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0042GetTresWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error) {
	res := &api.SlurmdbV0042GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0042GetUserWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error) {
	res := &api.SlurmdbV0042GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0042GetUsersWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error) {
	res := &api.SlurmdbV0042GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0042GetWckeyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error) {
	res := &api.SlurmdbV0042GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0042GetWckeysWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error) {
	res := &api.SlurmdbV0042GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0042PostAccountsAssociationWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0042PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0042PostAccountsAssociationWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAccountsAssociationWithResponse(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0042PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0042PostAccountsWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	res := &api.SlurmdbV0042PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostAccountsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAccountsWithResponse(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	res := &api.SlurmdbV0042PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostAssociationsWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	res := &api.SlurmdbV0042PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostAssociationsWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostAssociationsWithResponse(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	res := &api.SlurmdbV0042PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostClustersWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	res := &api.SlurmdbV0042PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostClustersWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	res := &api.SlurmdbV0042PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostConfigWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	res := &api.SlurmdbV0042PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostConfigWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostConfigWithResponse(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	res := &api.SlurmdbV0042PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostQosWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	res := &api.SlurmdbV0042PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostQosWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	res := &api.SlurmdbV0042PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostTresWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	res := &api.SlurmdbV0042PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostTresWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostTresWithResponse(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	res := &api.SlurmdbV0042PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostUsersAssociationWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0042PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0042PostUsersAssociationWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0042PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0042PostUsersWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	res := &api.SlurmdbV0042PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostUsersWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostUsersWithResponse(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	res := &api.SlurmdbV0042PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostWckeysWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	res := &api.SlurmdbV0042PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0042PostWckeysWithResponse implements v0042.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0042PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	res := &api.SlurmdbV0042PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0042OpenapiResp{},
	}
	return res, nil
}

var _ api.ClientWithResponsesInterface = &fakeClient{}

// NewFakeClient creates a new fake client for testing.
func NewFakeClient() api.ClientWithResponsesInterface {
	return NewFakeClientBuilder().Build()
}

// ClientBuilder builds a fake client.
type ClientBuilder struct {
	interceptorFuncs *interceptor.Funcs
}

// WithInterceptorFuncs configures the client methods to be intercepted using the provided interceptor.Funcs.
func (c *ClientBuilder) WithInterceptorFuncs(interceptorFuncs interceptor.Funcs) *ClientBuilder {
	c.interceptorFuncs = &interceptorFuncs
	return c
}

// Build returns a fake client.
func (c *ClientBuilder) Build() api.ClientWithResponsesInterface {
	client := &fakeClient{}
	if c.interceptorFuncs != nil {
		return interceptor.NewClient(client, *c.interceptorFuncs)
	}
	return client
}

// NewFakeClientBuilder returns a new builder to create a fake client.
func NewFakeClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}
