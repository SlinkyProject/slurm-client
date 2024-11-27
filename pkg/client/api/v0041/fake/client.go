// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0041/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0041DeleteJobWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error) {
	res := &api.SlurmV0041DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmV0041DeleteJobsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	res := &api.SlurmV0041DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0041DeleteJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041DeleteJobsWithResponse(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	res := &api.SlurmV0041DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0041DeleteNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error) {
	res := &api.SlurmV0041DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
	res := &api.SlurmV0041GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error) {
	res := &api.SlurmV0041GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0041GetJobsStateWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error) {
	res := &api.SlurmV0041GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error) {
	res := &api.SlurmV0041GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0041GetLicensesWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error) {
	res := &api.SlurmV0041GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0041GetNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error) {
	res := &api.SlurmV0041GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0041GetNodesWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetNodesWithResponse(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error) {
	res := &api.SlurmV0041GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0041GetPartitionWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error) {
	res := &api.SlurmV0041GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0041GetPartitionsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error) {
	res := &api.SlurmV0041GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0041GetPingWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error) {
	res := &api.SlurmV0041GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0041GetReconfigureWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
	res := &api.SlurmV0041GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmV0041GetReservationWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
	res := &api.SlurmV0041GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0041GetReservationsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
	res := &api.SlurmV0041GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0041GetSharesWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041GetSharesWithResponse(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error) {
	res := &api.SlurmV0041GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0041PostJobAllocateWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	res := &api.SlurmV0041PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0041PostJobAllocateWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobAllocateWithResponse(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	res := &api.SlurmV0041PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0041PostJobSubmitWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	res := &api.SlurmV0041PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0041PostJobSubmitWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobSubmitWithResponse(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	res := &api.SlurmV0041PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0041PostJobWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	res := &api.SlurmV0041PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0041PostJobWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostJobWithResponse(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	res := &api.SlurmV0041PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0041PostNodeWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	res := &api.SlurmV0041PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmV0041PostNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0041PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	res := &api.SlurmV0041PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error) {
	res := &api.SlurmdbV0041DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0041DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0041DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error) {
	res := &api.SlurmdbV0041DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0041DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteUserWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error) {
	res := &api.SlurmdbV0041DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041DeleteWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0041DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0041GetAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error) {
	res := &api.SlurmdbV0041GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error) {
	res := &api.SlurmdbV0041GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error) {
	res := &api.SlurmdbV0041GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error) {
	res := &api.SlurmdbV0041GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error) {
	res := &api.SlurmdbV0041GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0041GetClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error) {
	res := &api.SlurmdbV0041GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0041GetConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error) {
	res := &api.SlurmdbV0041GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error) {
	res := &api.SlurmdbV0041GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetInstanceWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error) {
	res := &api.SlurmdbV0041GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0041GetInstancesWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error) {
	res := &api.SlurmdbV0041GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error) {
	res := &api.SlurmdbV0041GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error) {
	res := &api.SlurmdbV0041GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0041GetQosWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error) {
	res := &api.SlurmdbV0041GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0041GetSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error) {
	res := &api.SlurmdbV0041GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0041GetTresWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error) {
	res := &api.SlurmdbV0041GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0041GetUserWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error) {
	res := &api.SlurmdbV0041GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0041GetUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error) {
	res := &api.SlurmdbV0041GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0041GetWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error) {
	res := &api.SlurmdbV0041GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0041GetWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error) {
	res := &api.SlurmdbV0041GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0041PostAccountsAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0041PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0041PostAccountsAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAccountsAssociationWithResponse(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0041PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0041PostAccountsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	res := &api.SlurmdbV0041PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAccountsWithResponse(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	res := &api.SlurmdbV0041PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostAssociationsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	res := &api.SlurmdbV0041PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostAssociationsWithResponse(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	res := &api.SlurmdbV0041PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostClustersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	res := &api.SlurmdbV0041PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	res := &api.SlurmdbV0041PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostConfigWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	res := &api.SlurmdbV0041PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostConfigWithResponse(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	res := &api.SlurmdbV0041PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostQosWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	res := &api.SlurmdbV0041PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostQosWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	res := &api.SlurmdbV0041PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostTresWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	res := &api.SlurmdbV0041PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostTresWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostTresWithResponse(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	res := &api.SlurmdbV0041PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostUsersAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0041PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0041PostUsersAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0041PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0041PostUsersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	res := &api.SlurmdbV0041PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostUsersWithResponse(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	res := &api.SlurmdbV0041PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostWckeysWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	res := &api.SlurmdbV0041PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0041PostWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0041PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	res := &api.SlurmdbV0041PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0041OpenapiResp{},
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
