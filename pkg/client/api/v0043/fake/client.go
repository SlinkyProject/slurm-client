// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0043/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0043DeleteJobWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
	res := &api.SlurmV0043DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiKillJobResp{},
	}
	return res, nil
}

// SlurmV0043DeleteJobsWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	res := &api.SlurmV0043DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0043DeleteJobsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043DeleteJobsWithResponse(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	res := &api.SlurmV0043DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0043DeleteNodeWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error) {
	res := &api.SlurmV0043DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043GetDiagWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error) {
	res := &api.SlurmV0043GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0043GetJobWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
	res := &api.SlurmV0043GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0043GetJobsStateWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error) {
	res := &api.SlurmV0043GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0043GetJobsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
	res := &api.SlurmV0043GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0043GetLicensesWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error) {
	res := &api.SlurmV0043GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0043GetNodeWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error) {
	res := &api.SlurmV0043GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0043GetNodesWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetNodesWithResponse(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error) {
	res := &api.SlurmV0043GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0043GetPartitionWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
	res := &api.SlurmV0043GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0043GetPartitionsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
	res := &api.SlurmV0043GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0043GetPingWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
	res := &api.SlurmV0043GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0043GetReconfigureWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
	res := &api.SlurmV0043GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043PostReservationWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	res := &api.SlurmV0043PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0043PostReservationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostReservationWithResponse(ctx context.Context, body api.SlurmV0043PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	res := &api.SlurmV0043PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0043DeleteReservationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteReservationResponse, error) {
	res := &api.SlurmV0043DeleteReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043GetReservationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error) {
	res := &api.SlurmV0043GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0043GetReservationsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error) {
	res := &api.SlurmV0043GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0043PostReservationsWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	res := &api.SlurmV0043PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0043PostReservationsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostReservationsWithResponse(ctx context.Context, body api.SlurmV0043PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	res := &api.SlurmV0043PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0043GetSharesWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043GetSharesWithResponse(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error) {
	res := &api.SlurmV0043GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0043PostJobAllocateWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	res := &api.SlurmV0043PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0043PostJobAllocateWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobAllocateWithResponse(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	res := &api.SlurmV0043PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0043PostJobSubmitWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	res := &api.SlurmV0043PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0043PostJobSubmitWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobSubmitWithResponse(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	res := &api.SlurmV0043PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0043PostJobWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	res := &api.SlurmV0043PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0043PostJobWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostJobWithResponse(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	res := &api.SlurmV0043PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0043PostNodeWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	res := &api.SlurmV0043PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043PostNodeWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	res := &api.SlurmV0043PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043PostNodesWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	res := &api.SlurmV0043PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmV0043PostNodesWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0043PostNodesWithResponse(ctx context.Context, body api.SlurmV0043PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	res := &api.SlurmV0043PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteAccountWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error) {
	res := &api.SlurmdbV0043DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteAssociationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0043DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteAssociationsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0043DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteClusterWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error) {
	res := &api.SlurmdbV0043DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteSingleQosWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0043DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteUserWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error) {
	res := &api.SlurmdbV0043DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043DeleteWckeyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0043DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0043GetAccountWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error) {
	res := &api.SlurmdbV0043GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetAccountsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error) {
	res := &api.SlurmdbV0043GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetAssociationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error) {
	res := &api.SlurmdbV0043GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetAssociationsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error) {
	res := &api.SlurmdbV0043GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetClusterWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error) {
	res := &api.SlurmdbV0043GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0043GetClustersWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error) {
	res := &api.SlurmdbV0043GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0043GetConfigWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error) {
	res := &api.SlurmdbV0043GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0043GetDiagWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error) {
	res := &api.SlurmdbV0043GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetInstanceWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error) {
	res := &api.SlurmdbV0043GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0043GetInstancesWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error) {
	res := &api.SlurmdbV0043GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0043GetJobWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error) {
	res := &api.SlurmdbV0043GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetJobsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error) {
	res := &api.SlurmdbV0043GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0043GetPingWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetPingResponse, error) {
	res := &api.SlurmdbV0043GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdPingResp{},
	}
	return res, nil
}

// SlurmdbV0043GetQosWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error) {
	res := &api.SlurmdbV0043GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0043GetSingleQosWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error) {
	res := &api.SlurmdbV0043GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0043GetTresWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error) {
	res := &api.SlurmdbV0043GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0043GetUserWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error) {
	res := &api.SlurmdbV0043GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0043GetUsersWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error) {
	res := &api.SlurmdbV0043GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0043GetWckeyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error) {
	res := &api.SlurmdbV0043GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0043GetWckeysWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error) {
	res := &api.SlurmdbV0043GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0043PostAccountsAssociationWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0043PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0043PostAccountsAssociationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAccountsAssociationWithResponse(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0043PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0043PostAccountsWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	res := &api.SlurmdbV0043PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostAccountsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAccountsWithResponse(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	res := &api.SlurmdbV0043PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostAssociationsWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	res := &api.SlurmdbV0043PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostAssociationsWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostAssociationsWithResponse(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	res := &api.SlurmdbV0043PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostClustersWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	res := &api.SlurmdbV0043PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostClustersWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	res := &api.SlurmdbV0043PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostConfigWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	res := &api.SlurmdbV0043PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostConfigWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostConfigWithResponse(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	res := &api.SlurmdbV0043PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostQosWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	res := &api.SlurmdbV0043PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostQosWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	res := &api.SlurmdbV0043PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostTresWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	res := &api.SlurmdbV0043PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostTresWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostTresWithResponse(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	res := &api.SlurmdbV0043PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostUsersAssociationWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0043PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0043PostUsersAssociationWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0043PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0043PostUsersWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	res := &api.SlurmdbV0043PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostUsersWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostUsersWithResponse(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	res := &api.SlurmdbV0043PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostWckeysWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	res := &api.SlurmdbV0043PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0043PostWckeysWithResponse implements v0043.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0043PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	res := &api.SlurmdbV0043PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0043OpenapiResp{},
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
