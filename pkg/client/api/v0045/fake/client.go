// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"io"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/interceptor"
)

var (
	HttpSuccess = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	}
)

type fakeClient struct{}

// SlurmV0045DeleteJobWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error) {
	res := &api.SlurmV0045DeleteJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiKillJobResp{},
	}
	return res, nil
}

// SlurmV0045DeleteJobsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	res := &api.SlurmV0045DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0045DeleteJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeleteJobsWithResponse(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	res := &api.SlurmV0045DeleteJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiKillJobsResp{},
	}
	return res, nil
}

// SlurmV0045DeleteNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
	res := &api.SlurmV0045DeleteNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045DeletePartitionWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeletePartitionWithResponse(ctx context.Context, partitionName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeletePartitionResponse, error) {
	res := &api.SlurmV0045DeletePartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045GetConfWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetConfWithResponse(ctx context.Context, params *api.SlurmV0045GetConfParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetConfResponse, error) {
	res := &api.SlurmV0045GetConfResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiConfResp{},
	}
	return res, nil
}

// SlurmV0045GetDiagWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error) {
	res := &api.SlurmV0045GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiDiagResp{},
	}
	return res, nil
}

// SlurmV0045GetJobWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error) {
	res := &api.SlurmV0045GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0045GetJobRequeueWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetJobRequeueWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobRequeueParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobRequeueResponse, error) {
	res := &api.SlurmV0045GetJobRequeueResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobRequeueResp{},
	}
	return res, nil
}

// SlurmV0045GetJobsStateWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error) {
	res := &api.SlurmV0045GetJobsStateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0045GetJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error) {
	res := &api.SlurmV0045GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobInfoResp{},
	}
	return res, nil
}

// SlurmV0045GetLicensesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error) {
	res := &api.SlurmV0045GetLicensesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiLicensesResp{},
	}
	return res, nil
}

// SlurmV0045GetNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
	res := &api.SlurmV0045GetNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0045GetNodesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetNodesWithResponse(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
	res := &api.SlurmV0045GetNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiNodesResp{},
	}
	return res, nil
}

// SlurmV0045GetPartitionWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error) {
	res := &api.SlurmV0045GetPartitionResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0045GetPartitionsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error) {
	res := &api.SlurmV0045GetPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiPartitionResp{},
	}
	return res, nil
}

// SlurmV0045GetPingWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error) {
	res := &api.SlurmV0045GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiPingArrayResp{},
	}
	return res, nil
}

// SlurmV0045GetReconfigureWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error) {
	res := &api.SlurmV0045GetReconfigureResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostReservationWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	res := &api.SlurmV0045PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0045PostReservationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostReservationWithResponse(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	res := &api.SlurmV0045PostReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0045DeleteReservationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
	res := &api.SlurmV0045DeleteReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045GetReservationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
	res := &api.SlurmV0045GetReservationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0045GetReservationsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
	res := &api.SlurmV0045GetReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationResp{},
	}
	return res, nil
}

// SlurmV0045PostReservationsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	res := &api.SlurmV0045PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0045PostReservationsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostReservationsWithResponse(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	res := &api.SlurmV0045PostReservationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiReservationModResp{},
	}
	return res, nil
}

// SlurmV0045GetResourcesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetResourcesWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error) {
	res := &api.SlurmV0045GetResourcesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResourceLayoutResp{},
	}
	return res, nil
}

// SlurmV0045GetSharesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045GetSharesWithResponse(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error) {
	res := &api.SlurmV0045GetSharesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSharesResp{},
	}
	return res, nil
}

// SlurmV0045PostJobAllocateWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	res := &api.SlurmV0045PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0045PostJobAllocateWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobAllocateWithResponse(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	res := &api.SlurmV0045PostJobAllocateResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobAllocResp{},
	}
	return res, nil
}

// SlurmV0045PostJobSubmitWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	res := &api.SlurmV0045PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0045PostJobSubmitWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobSubmitWithResponse(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	res := &api.SlurmV0045PostJobSubmitResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobSubmitResponse{},
	}
	return res, nil
}

// SlurmV0045PostJobWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	res := &api.SlurmV0045PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0045PostJobWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobWithResponse(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	res := &api.SlurmV0045PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobPostResponse{},
	}
	return res, nil
}

// SlurmV0045PostNodeWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	res := &api.SlurmV0045PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	res := &api.SlurmV0045PostNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostNodesWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	res := &api.SlurmV0045PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostNodesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNodesWithResponse(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	res := &api.SlurmV0045PostNodesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostJobsRequeueWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobsRequeueWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	res := &api.SlurmV0045PostJobsRequeueResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobsRequeueResp{},
	}
	return res, nil
}

// SlurmV0045PostJobsRequeueWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostJobsRequeueWithResponse(ctx context.Context, body api.SlurmV0045PostJobsRequeueJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	res := &api.SlurmV0045PostJobsRequeueResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobsRequeueResp{},
	}
	return res, nil
}

// SlurmV0045PostPartitionsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostPartitionsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	res := &api.SlurmV0045PostPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostPartitionsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostPartitionsWithResponse(ctx context.Context, body api.SlurmV0045PostPartitionsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	res := &api.SlurmV0045PostPartitionsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostNewNodeWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	res := &api.SlurmV0045PostNewNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmV0045PostNewNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmV0045PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	res := &api.SlurmV0045PostNewNodeResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteAccountWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error) {
	res := &api.SlurmdbV0045DeleteAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAccountsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteAssociationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error) {
	res := &api.SlurmdbV0045DeleteAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteAssociationsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error) {
	res := &api.SlurmdbV0045DeleteAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAssocsRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteClusterWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error) {
	res := &api.SlurmdbV0045DeleteClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiClustersRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteSingleQosWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error) {
	res := &api.SlurmdbV0045DeleteSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdQosRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteUserWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error) {
	res := &api.SlurmdbV0045DeleteUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045DeleteWckeyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error) {
	res := &api.SlurmdbV0045DeleteWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiWckeyRemovedResp{},
	}
	return res, nil
}

// SlurmdbV0045GetAccountWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error) {
	res := &api.SlurmdbV0045GetAccountResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetAccountsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error) {
	res := &api.SlurmdbV0045GetAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAccountsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetAssociationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error) {
	res := &api.SlurmdbV0045GetAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetAssociationsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error) {
	res := &api.SlurmdbV0045GetAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAssocsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetClusterWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error) {
	res := &api.SlurmdbV0045GetClusterResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0045GetClustersWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error) {
	res := &api.SlurmdbV0045GetClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiClustersResp{},
	}
	return res, nil
}

// SlurmdbV0045GetConfWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetConfWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfResponse, error) {
	res := &api.SlurmdbV0045GetConfResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdConfResp{},
	}
	return res, nil
}

// SlurmdbV0045GetConfigWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error) {
	res := &api.SlurmdbV0045GetConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdConfigResp{},
	}
	return res, nil
}

// SlurmdbV0045GetDiagWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error) {
	res := &api.SlurmdbV0045GetDiagResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdStatsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetInstanceWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error) {
	res := &api.SlurmdbV0045GetInstanceResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0045GetInstancesWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error) {
	res := &api.SlurmdbV0045GetInstancesResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiInstancesResp{},
	}
	return res, nil
}

// SlurmdbV0045GetJobWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error) {
	res := &api.SlurmdbV0045GetJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error) {
	res := &api.SlurmdbV0045GetJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdJobsResp{},
	}
	return res, nil
}

// SlurmdbV0045GetPingWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetPingResponse, error) {
	res := &api.SlurmdbV0045GetPingResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdPingResp{},
	}
	return res, nil
}

// SlurmdbV0045GetQosWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error) {
	res := &api.SlurmdbV0045GetQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0045GetSingleQosWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error) {
	res := &api.SlurmdbV0045GetSingleQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiSlurmdbdQosResp{},
	}
	return res, nil
}

// SlurmdbV0045GetTresWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error) {
	res := &api.SlurmdbV0045GetTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiTresResp{},
	}
	return res, nil
}

// SlurmdbV0045GetUserWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
	res := &api.SlurmdbV0045GetUserResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0045GetUsersWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
	res := &api.SlurmdbV0045GetUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiUsersResp{},
	}
	return res, nil
}

// SlurmdbV0045GetWckeyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error) {
	res := &api.SlurmdbV0045GetWckeyResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0045GetWckeysWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error) {
	res := &api.SlurmdbV0045GetWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiWckeyResp{},
	}
	return res, nil
}

// SlurmdbV0045PostAccountsAssociationWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0045PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0045PostAccountsAssociationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAccountsAssociationWithResponse(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	res := &api.SlurmdbV0045PostAccountsAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiAccountsAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0045PostAccountsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	res := &api.SlurmdbV0045PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostAccountsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAccountsWithResponse(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	res := &api.SlurmdbV0045PostAccountsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostAssociationsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	res := &api.SlurmdbV0045PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostAssociationsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostAssociationsWithResponse(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	res := &api.SlurmdbV0045PostAssociationsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostClustersWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	res := &api.SlurmdbV0045PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostClustersWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	res := &api.SlurmdbV0045PostClustersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostConfigWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	res := &api.SlurmdbV0045PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostConfigWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostConfigWithResponse(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	res := &api.SlurmdbV0045PostConfigResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostQosWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	res := &api.SlurmdbV0045PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostQosWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	res := &api.SlurmdbV0045PostQosResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostTresWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	res := &api.SlurmdbV0045PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostTresWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostTresWithResponse(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	res := &api.SlurmdbV0045PostTresResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostUsersAssociationWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0045PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0045PostUsersAssociationWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	res := &api.SlurmdbV0045PostUsersAssociationResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiUsersAddCondRespStr{},
	}
	return res, nil
}

// SlurmdbV0045PostUsersWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	res := &api.SlurmdbV0045PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostUsersWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostUsersWithResponse(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	res := &api.SlurmdbV0045PostUsersResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostWckeysWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	res := &api.SlurmdbV0045PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostWckeysWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	res := &api.SlurmdbV0045PostWckeysResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiResp{},
	}
	return res, nil
}

// SlurmdbV0045PostJobWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	res := &api.SlurmdbV0045PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0045PostJobWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0045PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	res := &api.SlurmdbV0045PostJobResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0045PostJobsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	res := &api.SlurmdbV0045PostJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobModifyResp{},
	}
	return res, nil
}

// SlurmdbV0045PostJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (f *fakeClient) SlurmdbV0045PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0045PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	res := &api.SlurmdbV0045PostJobsResponse{
		HTTPResponse: &HttpSuccess,
		JSON200:      &api.V0045OpenapiJobModifyResp{},
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
