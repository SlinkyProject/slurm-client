// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0044DeleteJobWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error) {
	res := &api.SlurmV0044DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiKillJobResp{},
	}
	return res, nil
}

// SlurmV0044DeleteJobsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	res := &api.SlurmV0044DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0044DeleteJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044DeleteJobsWithResponse(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	res := &api.SlurmV0044DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0044DeleteNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
	res := &api.SlurmV0044DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044GetDiagWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error) {
	res := &api.SlurmV0044GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0044GetJobWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error) {
	res := &api.SlurmV0044GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0044GetJobsStateWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error) {
	res := &api.SlurmV0044GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0044GetJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error) {
	res := &api.SlurmV0044GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0044GetLicensesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error) {
	res := &api.SlurmV0044GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0044GetNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
	res := &api.SlurmV0044GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0044GetNodesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetNodesWithResponse(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
	res := &api.SlurmV0044GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0044GetPartitionWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error) {
	res := &api.SlurmV0044GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0044GetPartitionsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error) {
	res := &api.SlurmV0044GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0044GetPingWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error) {
	res := &api.SlurmV0044GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0044GetReconfigureWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error) {
	res := &api.SlurmV0044GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostReservationWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	res := &api.SlurmV0044PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0044PostReservationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostReservationWithResponse(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	res := &api.SlurmV0044PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0044DeleteReservationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
	res := &api.SlurmV0044DeleteReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044GetReservationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
	res := &api.SlurmV0044GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0044GetReservationsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
	res := &api.SlurmV0044GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0044PostReservationsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	res := &api.SlurmV0044PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0044PostReservationsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostReservationsWithResponse(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	res := &api.SlurmV0044PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0044GetResourcesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetResourcesWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
	res := &api.SlurmV0044GetResourcesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResourceLayoutResp{},
	}
	return res, nil
}

// SlurmV0044GetSharesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044GetSharesWithResponse(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error) {
	res := &api.SlurmV0044GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0044PostJobAllocateWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	res := &api.SlurmV0044PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0044PostJobAllocateWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobAllocateWithResponse(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	res := &api.SlurmV0044PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0044PostJobSubmitWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	res := &api.SlurmV0044PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0044PostJobSubmitWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobSubmitWithResponse(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	res := &api.SlurmV0044PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0044PostJobWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	res := &api.SlurmV0044PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0044PostJobWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostJobWithResponse(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	res := &api.SlurmV0044PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0044PostNodeWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	res := &api.SlurmV0044PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	res := &api.SlurmV0044PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostNodesWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	res := &api.SlurmV0044PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostNodesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNodesWithResponse(ctx context.Context, body api.SlurmV0044PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	res := &api.SlurmV0044PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostNewNodeWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	res := &api.SlurmV0044PostNewNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmV0044PostNewNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0044PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0044PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	res := &api.SlurmV0044PostNewNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteAccountWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
	res := &api.SlurmdbV0044DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteAssociationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0044DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteAssociationsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0044DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteClusterWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error) {
	res := &api.SlurmdbV0044DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteSingleQosWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0044DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteUserWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error) {
	res := &api.SlurmdbV0044DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044DeleteWckeyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0044DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0044GetAccountWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
	res := &api.SlurmdbV0044GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetAccountsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
	res := &api.SlurmdbV0044GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetAssociationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error) {
	res := &api.SlurmdbV0044GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetAssociationsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
	res := &api.SlurmdbV0044GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetClusterWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error) {
	res := &api.SlurmdbV0044GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0044GetClustersWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error) {
	res := &api.SlurmdbV0044GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0044GetConfigWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error) {
	res := &api.SlurmdbV0044GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0044GetDiagWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error) {
	res := &api.SlurmdbV0044GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetInstanceWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error) {
	res := &api.SlurmdbV0044GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0044GetInstancesWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error) {
	res := &api.SlurmdbV0044GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0044GetJobWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error) {
	res := &api.SlurmdbV0044GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error) {
	res := &api.SlurmdbV0044GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0044GetPingWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetPingResponse, error) {
	res := &api.SlurmdbV0044GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdPingResp{},
	}
	return res, nil
}

// SlurmdbV0044GetQosWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
	res := &api.SlurmdbV0044GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0044GetSingleQosWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
	res := &api.SlurmdbV0044GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0044GetTresWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error) {
	res := &api.SlurmdbV0044GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0044GetUserWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error) {
	res := &api.SlurmdbV0044GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0044GetUsersWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error) {
	res := &api.SlurmdbV0044GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0044GetWckeyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error) {
	res := &api.SlurmdbV0044GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0044GetWckeysWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error) {
	res := &api.SlurmdbV0044GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0044PostAccountsAssociationWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0044PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0044PostAccountsAssociationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAccountsAssociationWithResponse(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0044PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0044PostAccountsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	res := &api.SlurmdbV0044PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostAccountsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAccountsWithResponse(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	res := &api.SlurmdbV0044PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostAssociationsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	res := &api.SlurmdbV0044PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostAssociationsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostAssociationsWithResponse(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	res := &api.SlurmdbV0044PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostClustersWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	res := &api.SlurmdbV0044PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostClustersWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	res := &api.SlurmdbV0044PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostConfigWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	res := &api.SlurmdbV0044PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostConfigWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostConfigWithResponse(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	res := &api.SlurmdbV0044PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostQosWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	res := &api.SlurmdbV0044PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostQosWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	res := &api.SlurmdbV0044PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostTresWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	res := &api.SlurmdbV0044PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostTresWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostTresWithResponse(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	res := &api.SlurmdbV0044PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostUsersAssociationWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0044PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0044PostUsersAssociationWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0044PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0044PostUsersWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	res := &api.SlurmdbV0044PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostUsersWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostUsersWithResponse(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	res := &api.SlurmdbV0044PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostWckeysWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	res := &api.SlurmdbV0044PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostWckeysWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	res := &api.SlurmdbV0044PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0044PostJobWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	res := &api.SlurmdbV0044PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0044PostJobWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0044PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	res := &api.SlurmdbV0044PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0044PostJobsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	res := &api.SlurmdbV0044PostJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0044PostJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0044PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0044PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	res := &api.SlurmdbV0044PostJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0044OpenapiJobModifyResp{},
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
