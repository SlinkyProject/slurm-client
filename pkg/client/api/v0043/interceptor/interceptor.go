// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
)

type Funcs struct {
	SlurmV0043DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error)
	SlurmV0043DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error)
	SlurmV0043DeleteJobsWithResponse                        func(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error)
	SlurmV0043DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error)
	SlurmV0043GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error)
	SlurmV0043GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error)
	SlurmV0043GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error)
	SlurmV0043GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error)
	SlurmV0043GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error)
	SlurmV0043GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error)
	SlurmV0043GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error)
	SlurmV0043GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error)
	SlurmV0043GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error)
	SlurmV0043GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error)
	SlurmV0043GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error)
	SlurmV0043PostReservationWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error)
	SlurmV0043PostReservationWithResponse                   func(ctx context.Context, body api.SlurmV0043PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error)
	SlurmV0043DeleteReservationWithResponse                 func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteReservationResponse, error)
	SlurmV0043GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error)
	SlurmV0043GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error)
	SlurmV0043PostReservationsWithBodyWithResponse          func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error)
	SlurmV0043PostReservationsWithResponse                  func(ctx context.Context, body api.SlurmV0043PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error)
	SlurmV0043GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error)
	SlurmV0043PostJobAllocateWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error)
	SlurmV0043PostJobAllocateWithResponse                   func(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error)
	SlurmV0043PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error)
	SlurmV0043PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error)
	SlurmV0043PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error)
	SlurmV0043PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error)
	SlurmV0043PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error)
	SlurmV0043PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error)
	SlurmV0043PostNodesWithBodyWithResponse                 func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error)
	SlurmV0043PostNodesWithResponse                         func(ctx context.Context, body api.SlurmV0043PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error)
	SlurmdbV0043DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error)
	SlurmdbV0043DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error)
	SlurmdbV0043DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error)
	SlurmdbV0043DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error)
	SlurmdbV0043DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error)
	SlurmdbV0043DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error)
	SlurmdbV0043DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error)
	SlurmdbV0043GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error)
	SlurmdbV0043GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error)
	SlurmdbV0043GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error)
	SlurmdbV0043GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error)
	SlurmdbV0043GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error)
	SlurmdbV0043GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error)
	SlurmdbV0043GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error)
	SlurmdbV0043GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error)
	SlurmdbV0043GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error)
	SlurmdbV0043GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error)
	SlurmdbV0043GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error)
	SlurmdbV0043GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error)
	SlurmdbV0043GetPingWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetPingResponse, error)
	SlurmdbV0043GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error)
	SlurmdbV0043GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error)
	SlurmdbV0043GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error)
	SlurmdbV0043GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error)
	SlurmdbV0043GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error)
	SlurmdbV0043GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error)
	SlurmdbV0043GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error)
	SlurmdbV0043PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error)
	SlurmdbV0043PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error)
	SlurmdbV0043PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error)
	SlurmdbV0043PostAccountsWithResponse                    func(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error)
	SlurmdbV0043PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error)
	SlurmdbV0043PostAssociationsWithResponse                func(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error)
	SlurmdbV0043PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error)
	SlurmdbV0043PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error)
	SlurmdbV0043PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error)
	SlurmdbV0043PostConfigWithResponse                      func(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error)
	SlurmdbV0043PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error)
	SlurmdbV0043PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error)
	SlurmdbV0043PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error)
	SlurmdbV0043PostTresWithResponse                        func(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error)
	SlurmdbV0043PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error)
	SlurmdbV0043PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error)
	SlurmdbV0043PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error)
	SlurmdbV0043PostUsersWithResponse                       func(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error)
	SlurmdbV0043PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error)
	SlurmdbV0043PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0043DeleteJobWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
	if i.funcs.SlurmV0043DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0043DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0043DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0043DeleteJobsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	if i.funcs.SlurmV0043DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043DeleteJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043DeleteJobsWithResponse(ctx context.Context, body api.V0043KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobsResponse, error) {
	if i.funcs.SlurmV0043DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0043DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0043DeleteNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteNodeResponse, error) {
	if i.funcs.SlurmV0043DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0043DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0043DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0043GetDiagWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetDiagResponse, error) {
	if i.funcs.SlurmV0043GetDiagWithResponse != nil {
		return i.funcs.SlurmV0043GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0043GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0043GetJobWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
	if i.funcs.SlurmV0043GetJobWithResponse != nil {
		return i.funcs.SlurmV0043GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0043GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0043GetJobsStateWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsStateResponse, error) {
	if i.funcs.SlurmV0043GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0043GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043GetJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
	if i.funcs.SlurmV0043GetJobsWithResponse != nil {
		return i.funcs.SlurmV0043GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043GetLicensesWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetLicensesResponse, error) {
	if i.funcs.SlurmV0043GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0043GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0043GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0043GetNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0043GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodeResponse, error) {
	if i.funcs.SlurmV0043GetNodeWithResponse != nil {
		return i.funcs.SlurmV0043GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0043GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0043GetNodesWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetNodesWithResponse(ctx context.Context, params *api.SlurmV0043GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetNodesResponse, error) {
	if i.funcs.SlurmV0043GetNodesWithResponse != nil {
		return i.funcs.SlurmV0043GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043GetPartitionWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
	if i.funcs.SlurmV0043GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0043GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0043GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0043GetPartitionsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
	if i.funcs.SlurmV0043GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0043GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043GetPingWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
	if i.funcs.SlurmV0043GetPingWithResponse != nil {
		return i.funcs.SlurmV0043GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0043GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0043GetReconfigureWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
	if i.funcs.SlurmV0043GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0043GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0043GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0043PostReservationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	if i.funcs.SlurmV0043PostReservationWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043PostReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostReservationWithResponse(ctx context.Context, body api.SlurmV0043PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationResponse, error) {
	if i.funcs.SlurmV0043PostReservationWithResponse != nil {
		return i.funcs.SlurmV0043PostReservationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043PostReservationWithResponse(ctx, body, reqEditors...)
}

// SlurmV0043DeleteReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteReservationResponse, error) {
	if i.funcs.SlurmV0043DeleteReservationWithResponse != nil {
		return i.funcs.SlurmV0043DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
	}
	return i.client.SlurmV0043DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
}

// SlurmV0043GetReservationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0043GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationResponse, error) {
	if i.funcs.SlurmV0043GetReservationWithResponse != nil {
		return i.funcs.SlurmV0043GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0043GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0043GetReservationsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0043GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReservationsResponse, error) {
	if i.funcs.SlurmV0043GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0043GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043PostReservationsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	if i.funcs.SlurmV0043PostReservationsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043PostReservationsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostReservationsWithResponse(ctx context.Context, body api.SlurmV0043PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostReservationsResponse, error) {
	if i.funcs.SlurmV0043PostReservationsWithResponse != nil {
		return i.funcs.SlurmV0043PostReservationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043PostReservationsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0043GetSharesWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043GetSharesWithResponse(ctx context.Context, params *api.SlurmV0043GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetSharesResponse, error) {
	if i.funcs.SlurmV0043GetSharesWithResponse != nil {
		return i.funcs.SlurmV0043GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0043GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0043PostJobAllocateWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0043PostJobAllocateWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043PostJobAllocateWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobAllocateWithResponse(ctx context.Context, body api.V0043JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0043PostJobAllocateWithResponse != nil {
		return i.funcs.SlurmV0043PostJobAllocateWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobAllocateWithResponse(ctx, body, reqEditors...)
}

// SlurmV0043PostJobSubmitWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0043PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043PostJobSubmitWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobSubmitWithResponse(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0043PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0043PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0043PostJobWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	if i.funcs.SlurmV0043PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0043PostJobWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostJobWithResponse(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
	if i.funcs.SlurmV0043PostJobWithResponse != nil {
		return i.funcs.SlurmV0043PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0043PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0043PostNodeWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	if i.funcs.SlurmV0043PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0043PostNodeWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0043UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodeResponse, error) {
	if i.funcs.SlurmV0043PostNodeWithResponse != nil {
		return i.funcs.SlurmV0043PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0043PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmV0043PostNodesWithBodyWithResponse implements v0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	if i.funcs.SlurmV0043PostNodesWithBodyWithResponse != nil {
		return i.funcs.SlurmV0043PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0043PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0043PostNodesWithResponse implements v0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0043PostNodesWithResponse(ctx context.Context, body api.SlurmV0043PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostNodesResponse, error) {
	if i.funcs.SlurmV0043PostNodesWithResponse != nil {
		return i.funcs.SlurmV0043PostNodesWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0043PostNodesWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043DeleteAccountWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0043DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0043DeleteAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0043DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043DeleteAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0043DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043DeleteClusterWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0043DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0043DeleteSingleQosWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0043DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0043DeleteUserWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0043DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0043DeleteWckeyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0043DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0043DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0043DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0043GetAccountWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0043GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountResponse, error) {
	if i.funcs.SlurmdbV0043GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0043GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0043GetAccountsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0043GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0043GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0043GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0043GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0043GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0043GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetClusterWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0043GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClusterResponse, error) {
	if i.funcs.SlurmdbV0043GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0043GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0043GetClustersWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetClustersResponse, error) {
	if i.funcs.SlurmdbV0043GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0043GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetConfigWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetConfigResponse, error) {
	if i.funcs.SlurmdbV0043GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0043GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0043GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0043GetDiagWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetDiagResponse, error) {
	if i.funcs.SlurmdbV0043GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0043GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0043GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0043GetInstanceWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0043GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0043GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetInstancesWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0043GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0043GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0043GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetJobWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobResponse, error) {
	if i.funcs.SlurmdbV0043GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0043GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0043GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0043GetJobsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetJobsResponse, error) {
	if i.funcs.SlurmdbV0043GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0043GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetPingWithResponse implements v0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetPingResponse, error) {
	if i.funcs.SlurmdbV0043GetPingWithResponse != nil {
		return i.funcs.SlurmdbV0043GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0043GetPingWithResponse(ctx, reqEditors...)
}

// SlurmdbV0043GetQosWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0043GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetQosResponse, error) {
	if i.funcs.SlurmdbV0043GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0043GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetSingleQosWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0043GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0043GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0043GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0043GetTresWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetTresResponse, error) {
	if i.funcs.SlurmdbV0043GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0043GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0043GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0043GetUserWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0043GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUserResponse, error) {
	if i.funcs.SlurmdbV0043GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0043GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0043GetUsersWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0043GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetUsersResponse, error) {
	if i.funcs.SlurmdbV0043GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0043GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043GetWckeyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0043GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0043GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0043GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0043GetWckeysWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0043GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0043GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0043GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0043PostAccountsAssociationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostAccountsAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAccountsAssociationWithResponse(ctx context.Context, body api.V0043OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0043PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostAccountsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0043PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostAccountsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAccountsWithResponse(ctx context.Context, body api.V0043OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0043PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostAssociationsWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0043PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostAssociationsWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostAssociationsWithResponse(ctx context.Context, body api.V0043OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0043PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0043PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostClustersWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	if i.funcs.SlurmdbV0043PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0043PostClustersWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0043PostClustersParams, body api.V0043OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostClustersResponse, error) {
	if i.funcs.SlurmdbV0043PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0043PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0043PostConfigWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	if i.funcs.SlurmdbV0043PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostConfigWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostConfigWithResponse(ctx context.Context, body api.V0043OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostConfigResponse, error) {
	if i.funcs.SlurmdbV0043PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0043PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostQosWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	if i.funcs.SlurmdbV0043PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0043PostQosWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0043PostQosParams, body api.V0043OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostQosResponse, error) {
	if i.funcs.SlurmdbV0043PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0043PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0043PostTresWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	if i.funcs.SlurmdbV0043PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostTresWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostTresWithResponse(ctx context.Context, body api.V0043OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostTresResponse, error) {
	if i.funcs.SlurmdbV0043PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0043PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostUsersAssociationWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0043PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0043PostUsersAssociationWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0043PostUsersAssociationParams, body api.V0043OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0043PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0043PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0043PostUsersWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	if i.funcs.SlurmdbV0043PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0043PostUsersWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostUsersWithResponse(ctx context.Context, body api.V0043OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostUsersResponse, error) {
	if i.funcs.SlurmdbV0043PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0043PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0043PostWckeysWithBodyWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0043PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0043PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0043PostWckeysWithResponse implements V0043.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0043PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0043PostWckeysParams, body api.V0043OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0043PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0043PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0043PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0043PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
