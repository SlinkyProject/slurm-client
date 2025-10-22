// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
)

type Funcs struct {
	SlurmV0044DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error)
	SlurmV0044DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error)
	SlurmV0044DeleteJobsWithResponse                        func(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error)
	SlurmV0044DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error)
	SlurmV0044GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error)
	SlurmV0044GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error)
	SlurmV0044GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error)
	SlurmV0044GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error)
	SlurmV0044GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error)
	SlurmV0044GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error)
	SlurmV0044GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error)
	SlurmV0044GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error)
	SlurmV0044GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error)
	SlurmV0044GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error)
	SlurmV0044GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error)
	SlurmV0044PostReservationWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error)
	SlurmV0044PostReservationWithResponse                   func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error)
	SlurmV0044DeleteReservationWithResponse                 func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error)
	SlurmV0044GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error)
	SlurmV0044GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error)
	SlurmV0044PostReservationsWithBodyWithResponse          func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error)
	SlurmV0044PostReservationsWithResponse                  func(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error)
	SlurmV0044GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error)
	SlurmV0044PostJobAllocateWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error)
	SlurmV0044PostJobAllocateWithResponse                   func(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error)
	SlurmV0044PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error)
	SlurmV0044PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error)
	SlurmV0044PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error)
	SlurmV0044PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error)
	SlurmV0044PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error)
	SlurmV0044PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error)
	SlurmV0044PostNodesWithBodyWithResponse                 func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error)
	SlurmV0044PostNodesWithResponse                         func(ctx context.Context, body api.SlurmV0044PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error)
	SlurmV0044PostNewNodeWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error)
	SlurmV0044PostNewNodeWithResponse                       func(ctx context.Context, body api.SlurmV0044PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error)
	SlurmdbV0044DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error)
	SlurmdbV0044DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error)
	SlurmdbV0044DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error)
	SlurmdbV0044DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error)
	SlurmdbV0044DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error)
	SlurmdbV0044DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error)
	SlurmdbV0044DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error)
	SlurmdbV0044GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error)
	SlurmdbV0044GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error)
	SlurmdbV0044GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error)
	SlurmdbV0044GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error)
	SlurmdbV0044GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error)
	SlurmdbV0044GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error)
	SlurmdbV0044GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error)
	SlurmdbV0044GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error)
	SlurmdbV0044GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error)
	SlurmdbV0044GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error)
	SlurmdbV0044GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error)
	SlurmdbV0044GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error)
	SlurmdbV0044GetPingWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetPingResponse, error)
	SlurmdbV0044GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error)
	SlurmdbV0044GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error)
	SlurmdbV0044GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error)
	SlurmdbV0044GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error)
	SlurmdbV0044GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error)
	SlurmdbV0044GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error)
	SlurmdbV0044GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error)
	SlurmdbV0044PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error)
	SlurmdbV0044PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error)
	SlurmdbV0044PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error)
	SlurmdbV0044PostAccountsWithResponse                    func(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error)
	SlurmdbV0044PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error)
	SlurmdbV0044PostAssociationsWithResponse                func(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error)
	SlurmdbV0044PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error)
	SlurmdbV0044PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error)
	SlurmdbV0044PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error)
	SlurmdbV0044PostConfigWithResponse                      func(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error)
	SlurmdbV0044PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error)
	SlurmdbV0044PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error)
	SlurmdbV0044PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error)
	SlurmdbV0044PostTresWithResponse                        func(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error)
	SlurmdbV0044PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error)
	SlurmdbV0044PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error)
	SlurmdbV0044PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error)
	SlurmdbV0044PostUsersWithResponse                       func(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error)
	SlurmdbV0044PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error)
	SlurmdbV0044PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error)
	SlurmdbV0044PostJobWithBodyWithResponse                 func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error)
	SlurmdbV0044PostJobWithResponse                         func(ctx context.Context, jobId string, body api.SlurmdbV0044PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error)
	SlurmdbV0044PostJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error)
	SlurmdbV0044PostJobsWithResponse                        func(ctx context.Context, body api.SlurmdbV0044PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0044DeleteJobWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobResponse, error) {
	if i.funcs.SlurmV0044DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0044DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0044DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0044DeleteJobsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	if i.funcs.SlurmV0044DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044DeleteJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044DeleteJobsWithResponse(ctx context.Context, body api.V0044KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteJobsResponse, error) {
	if i.funcs.SlurmV0044DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0044DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044DeleteNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
	if i.funcs.SlurmV0044DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0044DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0044DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0044GetDiagWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetDiagResponse, error) {
	if i.funcs.SlurmV0044GetDiagWithResponse != nil {
		return i.funcs.SlurmV0044GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0044GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0044GetJobWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0044GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobResponse, error) {
	if i.funcs.SlurmV0044GetJobWithResponse != nil {
		return i.funcs.SlurmV0044GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0044GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0044GetJobsStateWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsStateResponse, error) {
	if i.funcs.SlurmV0044GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0044GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044GetJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetJobsResponse, error) {
	if i.funcs.SlurmV0044GetJobsWithResponse != nil {
		return i.funcs.SlurmV0044GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044GetLicensesWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetLicensesResponse, error) {
	if i.funcs.SlurmV0044GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0044GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0044GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0044GetNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
	if i.funcs.SlurmV0044GetNodeWithResponse != nil {
		return i.funcs.SlurmV0044GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0044GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0044GetNodesWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetNodesWithResponse(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
	if i.funcs.SlurmV0044GetNodesWithResponse != nil {
		return i.funcs.SlurmV0044GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044GetPartitionWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0044GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionResponse, error) {
	if i.funcs.SlurmV0044GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0044GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0044GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0044GetPartitionsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0044GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPartitionsResponse, error) {
	if i.funcs.SlurmV0044GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0044GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044GetPingWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetPingResponse, error) {
	if i.funcs.SlurmV0044GetPingWithResponse != nil {
		return i.funcs.SlurmV0044GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0044GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0044GetReconfigureWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReconfigureResponse, error) {
	if i.funcs.SlurmV0044GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0044GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0044GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0044PostReservationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	if i.funcs.SlurmV0044PostReservationWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostReservationWithResponse(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
	if i.funcs.SlurmV0044PostReservationWithResponse != nil {
		return i.funcs.SlurmV0044PostReservationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostReservationWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044DeleteReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
	if i.funcs.SlurmV0044DeleteReservationWithResponse != nil {
		return i.funcs.SlurmV0044DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
	}
	return i.client.SlurmV0044DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
}

// SlurmV0044GetReservationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
	if i.funcs.SlurmV0044GetReservationWithResponse != nil {
		return i.funcs.SlurmV0044GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0044GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0044GetReservationsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
	if i.funcs.SlurmV0044GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0044GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044PostReservationsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	if i.funcs.SlurmV0044PostReservationsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostReservationsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostReservationsWithResponse(ctx context.Context, body api.SlurmV0044PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationsResponse, error) {
	if i.funcs.SlurmV0044PostReservationsWithResponse != nil {
		return i.funcs.SlurmV0044PostReservationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostReservationsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044GetSharesWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044GetSharesWithResponse(ctx context.Context, params *api.SlurmV0044GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetSharesResponse, error) {
	if i.funcs.SlurmV0044GetSharesWithResponse != nil {
		return i.funcs.SlurmV0044GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0044GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0044PostJobAllocateWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0044PostJobAllocateWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostJobAllocateWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobAllocateWithResponse(ctx context.Context, body api.V0044JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0044PostJobAllocateWithResponse != nil {
		return i.funcs.SlurmV0044PostJobAllocateWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobAllocateWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044PostJobSubmitWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0044PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostJobSubmitWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobSubmitWithResponse(ctx context.Context, body api.V0044JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0044PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0044PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044PostJobWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	if i.funcs.SlurmV0044PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0044PostJobWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostJobWithResponse(ctx context.Context, jobId string, body api.V0044JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostJobResponse, error) {
	if i.funcs.SlurmV0044PostJobWithResponse != nil {
		return i.funcs.SlurmV0044PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0044PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0044PostNodeWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	if i.funcs.SlurmV0044PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0044PostNodeWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
	if i.funcs.SlurmV0044PostNodeWithResponse != nil {
		return i.funcs.SlurmV0044PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmV0044PostNodesWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	if i.funcs.SlurmV0044PostNodesWithBodyWithResponse != nil {
		return i.funcs.SlurmV0044PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostNodesWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNodesWithResponse(ctx context.Context, body api.SlurmV0044PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodesResponse, error) {
	if i.funcs.SlurmV0044PostNodesWithResponse != nil {
		return i.funcs.SlurmV0044PostNodesWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNodesWithResponse(ctx, body, reqEditors...)
}

// SlurmV0044PostNewNodeWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	if i.funcs.SlurmV0044PostNodesWithResponse != nil {
		return i.funcs.SlurmV0044PostNewNodeWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNewNodeWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0044PostNewNodeWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0044PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0044PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNewNodeResponse, error) {
	if i.funcs.SlurmV0044PostNodesWithResponse != nil {
		return i.funcs.SlurmV0044PostNewNodeWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0044PostNewNodeWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044DeleteAccountWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0044DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0044DeleteAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0044DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044DeleteAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0044DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044DeleteClusterWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0044DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0044DeleteSingleQosWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0044DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0044DeleteUserWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0044DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0044DeleteWckeyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0044DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0044DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0044DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0044GetAccountWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
	if i.funcs.SlurmdbV0044GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0044GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0044GetAccountsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0044GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0044GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0044GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0044GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0044GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0044GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetClusterWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0044GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClusterResponse, error) {
	if i.funcs.SlurmdbV0044GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0044GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0044GetClustersWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetClustersResponse, error) {
	if i.funcs.SlurmdbV0044GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0044GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetConfigWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetConfigResponse, error) {
	if i.funcs.SlurmdbV0044GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0044GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0044GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0044GetDiagWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetDiagResponse, error) {
	if i.funcs.SlurmdbV0044GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0044GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0044GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0044GetInstanceWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0044GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0044GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetInstancesWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0044GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0044GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0044GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetJobWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobResponse, error) {
	if i.funcs.SlurmdbV0044GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0044GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0044GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0044GetJobsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0044GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetJobsResponse, error) {
	if i.funcs.SlurmdbV0044GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0044GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetPingWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetPingResponse, error) {
	if i.funcs.SlurmdbV0044GetPingWithResponse != nil {
		return i.funcs.SlurmdbV0044GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0044GetPingWithResponse(ctx, reqEditors...)
}

// SlurmdbV0044GetQosWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
	if i.funcs.SlurmdbV0044GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0044GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetSingleQosWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0044GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0044GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0044GetTresWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetTresResponse, error) {
	if i.funcs.SlurmdbV0044GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0044GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0044GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0044GetUserWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0044GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUserResponse, error) {
	if i.funcs.SlurmdbV0044GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0044GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0044GetUsersWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0044GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetUsersResponse, error) {
	if i.funcs.SlurmdbV0044GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0044GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044GetWckeyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0044GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0044GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0044GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0044GetWckeysWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0044GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0044GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0044GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0044PostAccountsAssociationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostAccountsAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAccountsAssociationWithResponse(ctx context.Context, body api.V0044OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0044PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostAccountsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0044PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostAccountsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAccountsWithResponse(ctx context.Context, body api.V0044OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0044PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostAssociationsWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0044PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostAssociationsWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostAssociationsWithResponse(ctx context.Context, body api.V0044OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0044PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0044PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostClustersWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	if i.funcs.SlurmdbV0044PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0044PostClustersWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0044PostClustersParams, body api.V0044OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostClustersResponse, error) {
	if i.funcs.SlurmdbV0044PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0044PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0044PostConfigWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	if i.funcs.SlurmdbV0044PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostConfigWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostConfigWithResponse(ctx context.Context, body api.V0044OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostConfigResponse, error) {
	if i.funcs.SlurmdbV0044PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0044PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostQosWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	if i.funcs.SlurmdbV0044PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0044PostQosWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0044PostQosParams, body api.V0044OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostQosResponse, error) {
	if i.funcs.SlurmdbV0044PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0044PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0044PostTresWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	if i.funcs.SlurmdbV0044PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostTresWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostTresWithResponse(ctx context.Context, body api.V0044OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostTresResponse, error) {
	if i.funcs.SlurmdbV0044PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0044PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostUsersAssociationWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0044PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0044PostUsersAssociationWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0044PostUsersAssociationParams, body api.V0044OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0044PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0044PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0044PostUsersWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	if i.funcs.SlurmdbV0044PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostUsersWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostUsersWithResponse(ctx context.Context, body api.V0044OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostUsersResponse, error) {
	if i.funcs.SlurmdbV0044PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0044PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0044PostWckeysWithBodyWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0044PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0044PostWckeysWithResponse implements V0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0044PostWckeysParams, body api.V0044OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0044PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0044PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0044PostJobWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	if i.funcs.SlurmdbV0044PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmdbV0044PostJobWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0044PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobResponse, error) {
	if i.funcs.SlurmdbV0044PostJobWithResponse != nil {
		return i.funcs.SlurmdbV0044PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmdbV0044PostJobsWithBodyWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	if i.funcs.SlurmdbV0044PostJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0044PostJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0044PostJobsWithResponse implements v0044.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0044PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0044PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostJobsResponse, error) {
	if i.funcs.SlurmdbV0044PostJobsWithResponse != nil {
		return i.funcs.SlurmdbV0044PostJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0044PostJobsWithResponse(ctx, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
