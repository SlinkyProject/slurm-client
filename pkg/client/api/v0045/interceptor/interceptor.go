// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
)

type Funcs struct {
	SlurmV0045DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error)
	SlurmV0045DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error)
	SlurmV0045DeleteJobsWithResponse                        func(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error)
	SlurmV0045DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error)
	SlurmV0045DeletePartitionWithResponse                   func(ctx context.Context, partitionName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeletePartitionResponse, error)
	SlurmV0045GetConfWithResponse                           func(ctx context.Context, params *api.SlurmV0045GetConfParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetConfResponse, error)
	SlurmV0045GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error)
	SlurmV0045GetJobRequeueWithResponse                     func(ctx context.Context, jobId string, params *api.SlurmV0045GetJobRequeueParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobRequeueResponse, error)
	SlurmV0045GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error)
	SlurmV0045GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error)
	SlurmV0045GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error)
	SlurmV0045GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error)
	SlurmV0045GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error)
	SlurmV0045GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error)
	SlurmV0045GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error)
	SlurmV0045GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error)
	SlurmV0045GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error)
	SlurmV0045GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error)
	SlurmV0045PostReservationWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error)
	SlurmV0045PostReservationWithResponse                   func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error)
	SlurmV0045DeleteReservationWithResponse                 func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error)
	SlurmV0045GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error)
	SlurmV0045GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error)
	SlurmV0045PostReservationsWithBodyWithResponse          func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error)
	SlurmV0045PostReservationsWithResponse                  func(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error)
	SlurmV0045GetResourcesWithResponse                      func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error)
	SlurmV0045GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error)
	SlurmV0045PostJobAllocateWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error)
	SlurmV0045PostJobAllocateWithResponse                   func(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error)
	SlurmV0045PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error)
	SlurmV0045PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error)
	SlurmV0045PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error)
	SlurmV0045PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error)
	SlurmV0045PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error)
	SlurmV0045PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error)
	SlurmV0045PostNodesWithBodyWithResponse                 func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error)
	SlurmV0045PostNodesWithResponse                         func(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error)
	SlurmV0045PostJobsRequeueWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error)
	SlurmV0045PostJobsRequeueWithResponse                   func(ctx context.Context, body api.SlurmV0045PostJobsRequeueJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error)
	SlurmV0045PostPartitionsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error)
	SlurmV0045PostPartitionsWithResponse                    func(ctx context.Context, body api.SlurmV0045PostPartitionsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error)
	SlurmV0045PostNewNodeWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error)
	SlurmV0045PostNewNodeWithResponse                       func(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error)
	SlurmdbV0045DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error)
	SlurmdbV0045DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error)
	SlurmdbV0045DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error)
	SlurmdbV0045DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error)
	SlurmdbV0045DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error)
	SlurmdbV0045DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error)
	SlurmdbV0045DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error)
	SlurmdbV0045GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error)
	SlurmdbV0045GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error)
	SlurmdbV0045GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error)
	SlurmdbV0045GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error)
	SlurmdbV0045GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error)
	SlurmdbV0045GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error)
	SlurmdbV0045GetConfWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfResponse, error)
	SlurmdbV0045GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error)
	SlurmdbV0045GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error)
	SlurmdbV0045GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error)
	SlurmdbV0045GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error)
	SlurmdbV0045GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error)
	SlurmdbV0045GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error)
	SlurmdbV0045GetPingWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetPingResponse, error)
	SlurmdbV0045GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error)
	SlurmdbV0045GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error)
	SlurmdbV0045GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error)
	SlurmdbV0045GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error)
	SlurmdbV0045GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error)
	SlurmdbV0045GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error)
	SlurmdbV0045GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error)
	SlurmdbV0045PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error)
	SlurmdbV0045PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error)
	SlurmdbV0045PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error)
	SlurmdbV0045PostAccountsWithResponse                    func(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error)
	SlurmdbV0045PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error)
	SlurmdbV0045PostAssociationsWithResponse                func(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error)
	SlurmdbV0045PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error)
	SlurmdbV0045PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error)
	SlurmdbV0045PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error)
	SlurmdbV0045PostConfigWithResponse                      func(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error)
	SlurmdbV0045PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error)
	SlurmdbV0045PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error)
	SlurmdbV0045PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error)
	SlurmdbV0045PostTresWithResponse                        func(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error)
	SlurmdbV0045PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error)
	SlurmdbV0045PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error)
	SlurmdbV0045PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error)
	SlurmdbV0045PostUsersWithResponse                       func(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error)
	SlurmdbV0045PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error)
	SlurmdbV0045PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error)
	SlurmdbV0045PostJobWithBodyWithResponse                 func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error)
	SlurmdbV0045PostJobWithResponse                         func(ctx context.Context, jobId string, body api.SlurmdbV0045PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error)
	SlurmdbV0045PostJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error)
	SlurmdbV0045PostJobsWithResponse                        func(ctx context.Context, body api.SlurmdbV0045PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0045DeleteJobWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobResponse, error) {
	if i.funcs.SlurmV0045DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0045DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0045DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0045DeleteJobsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	if i.funcs.SlurmV0045DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045DeleteJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeleteJobsWithResponse(ctx context.Context, body api.V0045KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteJobsResponse, error) {
	if i.funcs.SlurmV0045DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0045DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045DeleteNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
	if i.funcs.SlurmV0045DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0045DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0045DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0045DeletePartitionWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeletePartitionWithResponse(ctx context.Context, partitionName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeletePartitionResponse, error) {
	if i.funcs.SlurmV0045DeletePartitionWithResponse != nil {
		return i.funcs.SlurmV0045DeletePartitionWithResponse(ctx, partitionName, reqEditors...)
	}
	return i.client.SlurmV0045DeletePartitionWithResponse(ctx, partitionName, reqEditors...)
}

// SlurmV0045GetConfWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetConfWithResponse(ctx context.Context, params *api.SlurmV0045GetConfParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetConfResponse, error) {
	if i.funcs.SlurmV0045GetConfWithResponse != nil {
		return i.funcs.SlurmV0045GetConfWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetConfWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045GetDiagWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetDiagResponse, error) {
	if i.funcs.SlurmV0045GetDiagWithResponse != nil {
		return i.funcs.SlurmV0045GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0045GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0045GetJobRequeueWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetJobRequeueWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobRequeueParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobRequeueResponse, error) {
	if i.funcs.SlurmV0045GetJobRequeueWithResponse != nil {
		return i.funcs.SlurmV0045GetJobRequeueWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0045GetJobRequeueWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0045GetJobWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0045GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobResponse, error) {
	if i.funcs.SlurmV0045GetJobWithResponse != nil {
		return i.funcs.SlurmV0045GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0045GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0045GetJobsStateWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsStateResponse, error) {
	if i.funcs.SlurmV0045GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0045GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045GetJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetJobsResponse, error) {
	if i.funcs.SlurmV0045GetJobsWithResponse != nil {
		return i.funcs.SlurmV0045GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045GetLicensesWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetLicensesResponse, error) {
	if i.funcs.SlurmV0045GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0045GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0045GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0045GetNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
	if i.funcs.SlurmV0045GetNodeWithResponse != nil {
		return i.funcs.SlurmV0045GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0045GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0045GetNodesWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetNodesWithResponse(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
	if i.funcs.SlurmV0045GetNodesWithResponse != nil {
		return i.funcs.SlurmV0045GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045GetPartitionWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0045GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionResponse, error) {
	if i.funcs.SlurmV0045GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0045GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0045GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0045GetPartitionsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0045GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPartitionsResponse, error) {
	if i.funcs.SlurmV0045GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0045GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045GetPingWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetPingResponse, error) {
	if i.funcs.SlurmV0045GetPingWithResponse != nil {
		return i.funcs.SlurmV0045GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0045GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0045GetReconfigureWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReconfigureResponse, error) {
	if i.funcs.SlurmV0045GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0045GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0045GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0045PostReservationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostReservationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	if i.funcs.SlurmV0045PostReservationWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostReservationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostReservationWithResponse(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
	if i.funcs.SlurmV0045PostReservationWithResponse != nil {
		return i.funcs.SlurmV0045PostReservationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostReservationWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045DeleteReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045DeleteReservationWithResponse(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
	if i.funcs.SlurmV0045DeleteReservationWithResponse != nil {
		return i.funcs.SlurmV0045DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
	}
	return i.client.SlurmV0045DeleteReservationWithResponse(ctx, reservationName, reqEditors...)
}

// SlurmV0045GetReservationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
	if i.funcs.SlurmV0045GetReservationWithResponse != nil {
		return i.funcs.SlurmV0045GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0045GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0045GetReservationsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
	if i.funcs.SlurmV0045GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0045GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045PostReservationsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostReservationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	if i.funcs.SlurmV0045PostReservationsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostReservationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostReservationsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostReservationsWithResponse(ctx context.Context, body api.SlurmV0045PostReservationsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationsResponse, error) {
	if i.funcs.SlurmV0045PostReservationsWithResponse != nil {
		return i.funcs.SlurmV0045PostReservationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostReservationsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045GetResourcesWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetResourcesWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetResourcesResponse, error) {
	if i.funcs.SlurmV0045GetResourcesWithResponse != nil {
		return i.funcs.SlurmV0045GetResourcesWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmV0045GetResourcesWithResponse(ctx, jobId, reqEditors...)
}

// SlurmV0045GetSharesWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045GetSharesWithResponse(ctx context.Context, params *api.SlurmV0045GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetSharesResponse, error) {
	if i.funcs.SlurmV0045GetSharesWithResponse != nil {
		return i.funcs.SlurmV0045GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0045GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0045PostJobAllocateWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0045PostJobAllocateWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostJobAllocateWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobAllocateWithResponse(ctx context.Context, body api.V0045JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0045PostJobAllocateWithResponse != nil {
		return i.funcs.SlurmV0045PostJobAllocateWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobAllocateWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045PostJobSubmitWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0045PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostJobSubmitWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobSubmitWithResponse(ctx context.Context, body api.V0045JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0045PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0045PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045PostJobWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	if i.funcs.SlurmV0045PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0045PostJobWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobWithResponse(ctx context.Context, jobId string, body api.V0045JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobResponse, error) {
	if i.funcs.SlurmV0045PostJobWithResponse != nil {
		return i.funcs.SlurmV0045PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0045PostNodeWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	if i.funcs.SlurmV0045PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0045PostNodeWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
	if i.funcs.SlurmV0045PostNodeWithResponse != nil {
		return i.funcs.SlurmV0045PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmV0045PostNodesWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	if i.funcs.SlurmV0045PostNodesWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostNodesWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNodesWithResponse(ctx context.Context, body api.SlurmV0045PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodesResponse, error) {
	if i.funcs.SlurmV0045PostNodesWithResponse != nil {
		return i.funcs.SlurmV0045PostNodesWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNodesWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045PostJobsRequeueWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobsRequeueWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	if i.funcs.SlurmV0045PostJobsRequeueWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostJobsRequeueWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobsRequeueWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostJobsRequeueWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostJobsRequeueWithResponse(ctx context.Context, body api.SlurmV0045PostJobsRequeueJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostJobsRequeueResponse, error) {
	if i.funcs.SlurmV0045PostJobsRequeueWithResponse != nil {
		return i.funcs.SlurmV0045PostJobsRequeueWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostJobsRequeueWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045PostPartitionsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostPartitionsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	if i.funcs.SlurmV0045PostPartitionsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostPartitionsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostPartitionsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostPartitionsWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostPartitionsWithResponse(ctx context.Context, body api.SlurmV0045PostPartitionsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostPartitionsResponse, error) {
	if i.funcs.SlurmV0045PostPartitionsWithResponse != nil {
		return i.funcs.SlurmV0045PostPartitionsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostPartitionsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0045PostNewNodeWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNewNodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	if i.funcs.SlurmV0045PostNewNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0045PostNewNodeWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNewNodeWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0045PostNewNodeWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0045PostNewNodeWithResponse(ctx context.Context, body api.SlurmV0045PostNewNodeJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
	if i.funcs.SlurmV0045PostNewNodeWithResponse != nil {
		return i.funcs.SlurmV0045PostNewNodeWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0045PostNewNodeWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045DeleteAccountWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0045DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0045DeleteAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0045DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045DeleteAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0045DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045DeleteClusterWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0045DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0045DeleteSingleQosWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0045DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0045DeleteUserWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0045DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0045DeleteWckeyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0045DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0045DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0045DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0045GetAccountWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0045GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountResponse, error) {
	if i.funcs.SlurmdbV0045GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0045GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0045GetAccountsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0045GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0045GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0045GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0045GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0045GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0045GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetClusterWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0045GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClusterResponse, error) {
	if i.funcs.SlurmdbV0045GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0045GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0045GetClustersWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetClustersResponse, error) {
	if i.funcs.SlurmdbV0045GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0045GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetConfWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetConfWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfResponse, error) {
	if i.funcs.SlurmdbV0045GetConfWithResponse != nil {
		return i.funcs.SlurmdbV0045GetConfWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0045GetConfWithResponse(ctx, reqEditors...)
}

// SlurmdbV0045GetConfigWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetConfigResponse, error) {
	if i.funcs.SlurmdbV0045GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0045GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0045GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0045GetDiagWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetDiagResponse, error) {
	if i.funcs.SlurmdbV0045GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0045GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0045GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0045GetInstanceWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0045GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0045GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetInstancesWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0045GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0045GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0045GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetJobWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobResponse, error) {
	if i.funcs.SlurmdbV0045GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0045GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0045GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0045GetJobsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0045GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetJobsResponse, error) {
	if i.funcs.SlurmdbV0045GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0045GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetPingWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetPingResponse, error) {
	if i.funcs.SlurmdbV0045GetPingWithResponse != nil {
		return i.funcs.SlurmdbV0045GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0045GetPingWithResponse(ctx, reqEditors...)
}

// SlurmdbV0045GetQosWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0045GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetQosResponse, error) {
	if i.funcs.SlurmdbV0045GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0045GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetSingleQosWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0045GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0045GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0045GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0045GetTresWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetTresResponse, error) {
	if i.funcs.SlurmdbV0045GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0045GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0045GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0045GetUserWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
	if i.funcs.SlurmdbV0045GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0045GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0045GetUsersWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
	if i.funcs.SlurmdbV0045GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0045GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045GetWckeyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0045GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0045GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0045GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0045GetWckeysWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0045GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0045GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0045GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0045PostAccountsAssociationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostAccountsAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAccountsAssociationWithResponse(ctx context.Context, body api.V0045OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0045PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostAccountsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0045PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostAccountsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAccountsWithResponse(ctx context.Context, body api.V0045OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0045PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostAssociationsWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0045PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostAssociationsWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostAssociationsWithResponse(ctx context.Context, body api.V0045OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0045PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0045PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostClustersWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	if i.funcs.SlurmdbV0045PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0045PostClustersWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0045PostClustersParams, body api.V0045OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostClustersResponse, error) {
	if i.funcs.SlurmdbV0045PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0045PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0045PostConfigWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	if i.funcs.SlurmdbV0045PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostConfigWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostConfigWithResponse(ctx context.Context, body api.V0045OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostConfigResponse, error) {
	if i.funcs.SlurmdbV0045PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0045PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostQosWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	if i.funcs.SlurmdbV0045PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0045PostQosWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0045PostQosParams, body api.V0045OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostQosResponse, error) {
	if i.funcs.SlurmdbV0045PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0045PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0045PostTresWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	if i.funcs.SlurmdbV0045PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostTresWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostTresWithResponse(ctx context.Context, body api.V0045OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostTresResponse, error) {
	if i.funcs.SlurmdbV0045PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0045PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostUsersAssociationWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0045PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0045PostUsersAssociationWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0045PostUsersAssociationParams, body api.V0045OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0045PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0045PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0045PostUsersWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	if i.funcs.SlurmdbV0045PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostUsersWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostUsersWithResponse(ctx context.Context, body api.V0045OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostUsersResponse, error) {
	if i.funcs.SlurmdbV0045PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0045PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0045PostWckeysWithBodyWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0045PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0045PostWckeysWithResponse implements V0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0045PostWckeysParams, body api.V0045OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0045PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0045PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0045PostJobWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	if i.funcs.SlurmdbV0045PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmdbV0045PostJobWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostJobWithResponse(ctx context.Context, jobId string, body api.SlurmdbV0045PostJobJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobResponse, error) {
	if i.funcs.SlurmdbV0045PostJobWithResponse != nil {
		return i.funcs.SlurmdbV0045PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmdbV0045PostJobsWithBodyWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	if i.funcs.SlurmdbV0045PostJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0045PostJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0045PostJobsWithResponse implements v0045.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0045PostJobsWithResponse(ctx context.Context, body api.SlurmdbV0045PostJobsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045PostJobsResponse, error) {
	if i.funcs.SlurmdbV0045PostJobsWithResponse != nil {
		return i.funcs.SlurmdbV0045PostJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0045PostJobsWithResponse(ctx, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
