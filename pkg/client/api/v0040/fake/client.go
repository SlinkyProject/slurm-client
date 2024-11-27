// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0040/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0040DeleteJobWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
	res := &api.SlurmV0040DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmV0040DeleteJobsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	res := &api.SlurmV0040DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0040DeleteJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040DeleteJobsWithResponse(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	res := &api.SlurmV0040DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0040DeleteNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
	res := &api.SlurmV0040DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error) {
	res := &api.SlurmV0040GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
	res := &api.SlurmV0040GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0040GetJobsStateWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error) {
	res := &api.SlurmV0040GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
	res := &api.SlurmV0040GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0040GetLicensesWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error) {
	res := &api.SlurmV0040GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0040GetNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
	res := &api.SlurmV0040GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0040GetNodesWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetNodesWithResponse(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
	res := &api.SlurmV0040GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0040GetPartitionWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
	res := &api.SlurmV0040GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0040GetPartitionsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
	res := &api.SlurmV0040GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0040GetPingWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error) {
	res := &api.SlurmV0040GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0040GetReconfigureWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error) {
	res := &api.SlurmV0040GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmV0040GetReservationWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error) {
	res := &api.SlurmV0040GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0040GetReservationsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error) {
	res := &api.SlurmV0040GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0040GetSharesWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040GetSharesWithResponse(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error) {
	res := &api.SlurmV0040GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0040PostJobSubmitWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	res := &api.SlurmV0040PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0040PostJobSubmitWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostJobSubmitWithResponse(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	res := &api.SlurmV0040PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0040PostJobWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	res := &api.SlurmV0040PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0040PostJobWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostJobWithResponse(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	res := &api.SlurmV0040PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0040PostNodeWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	res := &api.SlurmV0040PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmV0040PostNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0040PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	res := &api.SlurmV0040PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error) {
	res := &api.SlurmdbV0040DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0040DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0040DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error) {
	res := &api.SlurmdbV0040DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0040DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteUserWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error) {
	res := &api.SlurmdbV0040DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040DeleteWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0040DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0040GetAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error) {
	res := &api.SlurmdbV0040GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error) {
	res := &api.SlurmdbV0040GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error) {
	res := &api.SlurmdbV0040GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error) {
	res := &api.SlurmdbV0040GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error) {
	res := &api.SlurmdbV0040GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0040GetClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error) {
	res := &api.SlurmdbV0040GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0040GetConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error) {
	res := &api.SlurmdbV0040GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error) {
	res := &api.SlurmdbV0040GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetInstanceWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error) {
	res := &api.SlurmdbV0040GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0040GetInstancesWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error) {
	res := &api.SlurmdbV0040GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error) {
	res := &api.SlurmdbV0040GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error) {
	res := &api.SlurmdbV0040GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0040GetQosWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error) {
	res := &api.SlurmdbV0040GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0040GetSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error) {
	res := &api.SlurmdbV0040GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0040GetTresWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error) {
	res := &api.SlurmdbV0040GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0040GetUserWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error) {
	res := &api.SlurmdbV0040GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0040GetUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error) {
	res := &api.SlurmdbV0040GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0040GetWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error) {
	res := &api.SlurmdbV0040GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0040GetWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error) {
	res := &api.SlurmdbV0040GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0040PostAccountsAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0040PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0040PostAccountsAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAccountsAssociationWithResponse(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0040PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0040PostAccountsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	res := &api.SlurmdbV0040PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAccountsWithResponse(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	res := &api.SlurmdbV0040PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostAssociationsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	res := &api.SlurmdbV0040PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostAssociationsWithResponse(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	res := &api.SlurmdbV0040PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostClustersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	res := &api.SlurmdbV0040PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	res := &api.SlurmdbV0040PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostConfigWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	res := &api.SlurmdbV0040PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostConfigWithResponse(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	res := &api.SlurmdbV0040PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostQosWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	res := &api.SlurmdbV0040PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostQosWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	res := &api.SlurmdbV0040PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostTresWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	res := &api.SlurmdbV0040PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostTresWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostTresWithResponse(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	res := &api.SlurmdbV0040PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostUsersAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0040PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0040PostUsersAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0040PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0040PostUsersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	res := &api.SlurmdbV0040PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostUsersWithResponse(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	res := &api.SlurmdbV0040PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostWckeysWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	res := &api.SlurmdbV0040PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0040PostWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0040PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	res := &api.SlurmdbV0040PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0040OpenapiResp{},
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
