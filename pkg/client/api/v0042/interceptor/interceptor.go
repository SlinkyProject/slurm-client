// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
)

type Funcs struct {
	SlurmV0042DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error)
	SlurmV0042DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error)
	SlurmV0042DeleteJobsWithResponse                        func(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error)
	SlurmV0042DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error)
	SlurmV0042GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error)
	SlurmV0042GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error)
	SlurmV0042GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error)
	SlurmV0042GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error)
	SlurmV0042GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error)
	SlurmV0042GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error)
	SlurmV0042GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error)
	SlurmV0042GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error)
	SlurmV0042GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error)
	SlurmV0042GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error)
	SlurmV0042GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error)
	SlurmV0042GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error)
	SlurmV0042GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error)
	SlurmV0042GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error)
	SlurmV0042PostJobAllocateWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error)
	SlurmV0042PostJobAllocateWithResponse                   func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error)
	SlurmV0042PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error)
	SlurmV0042PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error)
	SlurmV0042PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error)
	SlurmV0042PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error)
	SlurmV0042PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error)
	SlurmV0042PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error)
	SlurmV0042PostNodesWithBodyWithResponse                 func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error)
	SlurmV0042PostNodesWithResponse                         func(ctx context.Context, body api.SlurmV0042PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error)
	SlurmdbV0042DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error)
	SlurmdbV0042DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error)
	SlurmdbV0042DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error)
	SlurmdbV0042DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error)
	SlurmdbV0042DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error)
	SlurmdbV0042DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error)
	SlurmdbV0042DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error)
	SlurmdbV0042GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error)
	SlurmdbV0042GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error)
	SlurmdbV0042GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error)
	SlurmdbV0042GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error)
	SlurmdbV0042GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error)
	SlurmdbV0042GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error)
	SlurmdbV0042GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error)
	SlurmdbV0042GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error)
	SlurmdbV0042GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error)
	SlurmdbV0042GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error)
	SlurmdbV0042GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error)
	SlurmdbV0042GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error)
	SlurmdbV0042GetPingWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetPingResponse, error)
	SlurmdbV0042GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error)
	SlurmdbV0042GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error)
	SlurmdbV0042GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error)
	SlurmdbV0042GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error)
	SlurmdbV0042GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error)
	SlurmdbV0042GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error)
	SlurmdbV0042GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error)
	SlurmdbV0042PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error)
	SlurmdbV0042PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error)
	SlurmdbV0042PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error)
	SlurmdbV0042PostAccountsWithResponse                    func(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error)
	SlurmdbV0042PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error)
	SlurmdbV0042PostAssociationsWithResponse                func(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error)
	SlurmdbV0042PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error)
	SlurmdbV0042PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error)
	SlurmdbV0042PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error)
	SlurmdbV0042PostConfigWithResponse                      func(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error)
	SlurmdbV0042PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error)
	SlurmdbV0042PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error)
	SlurmdbV0042PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error)
	SlurmdbV0042PostTresWithResponse                        func(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error)
	SlurmdbV0042PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error)
	SlurmdbV0042PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error)
	SlurmdbV0042PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error)
	SlurmdbV0042PostUsersWithResponse                       func(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error)
	SlurmdbV0042PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error)
	SlurmdbV0042PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0042DeleteJobWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
	if i.funcs.SlurmV0042DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0042DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0042DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0042DeleteJobsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	if i.funcs.SlurmV0042DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0042DeleteJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042DeleteJobsWithResponse(ctx context.Context, body api.V0042KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobsResponse, error) {
	if i.funcs.SlurmV0042DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0042DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0042DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0042DeleteNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteNodeResponse, error) {
	if i.funcs.SlurmV0042DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0042DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0042DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0042GetDiagWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
	if i.funcs.SlurmV0042GetDiagWithResponse != nil {
		return i.funcs.SlurmV0042GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0042GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0042GetJobWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
	if i.funcs.SlurmV0042GetJobWithResponse != nil {
		return i.funcs.SlurmV0042GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0042GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0042GetJobsStateWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsStateResponse, error) {
	if i.funcs.SlurmV0042GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0042GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042GetJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
	if i.funcs.SlurmV0042GetJobsWithResponse != nil {
		return i.funcs.SlurmV0042GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042GetLicensesWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetLicensesResponse, error) {
	if i.funcs.SlurmV0042GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0042GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0042GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0042GetNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0042GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodeResponse, error) {
	if i.funcs.SlurmV0042GetNodeWithResponse != nil {
		return i.funcs.SlurmV0042GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0042GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0042GetNodesWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetNodesWithResponse(ctx context.Context, params *api.SlurmV0042GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetNodesResponse, error) {
	if i.funcs.SlurmV0042GetNodesWithResponse != nil {
		return i.funcs.SlurmV0042GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042GetPartitionWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0042GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionResponse, error) {
	if i.funcs.SlurmV0042GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0042GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0042GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0042GetPartitionsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0042GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPartitionsResponse, error) {
	if i.funcs.SlurmV0042GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0042GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042GetPingWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetPingResponse, error) {
	if i.funcs.SlurmV0042GetPingWithResponse != nil {
		return i.funcs.SlurmV0042GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0042GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0042GetReconfigureWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReconfigureResponse, error) {
	if i.funcs.SlurmV0042GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0042GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0042GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0042GetReservationWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0042GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationResponse, error) {
	if i.funcs.SlurmV0042GetReservationWithResponse != nil {
		return i.funcs.SlurmV0042GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0042GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0042GetReservationsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0042GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetReservationsResponse, error) {
	if i.funcs.SlurmV0042GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0042GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042GetSharesWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042GetSharesWithResponse(ctx context.Context, params *api.SlurmV0042GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetSharesResponse, error) {
	if i.funcs.SlurmV0042GetSharesWithResponse != nil {
		return i.funcs.SlurmV0042GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0042GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0042PostJobAllocateWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0042PostJobAllocateWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0042PostJobAllocateWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobAllocateWithResponse(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0042PostJobAllocateWithResponse != nil {
		return i.funcs.SlurmV0042PostJobAllocateWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobAllocateWithResponse(ctx, body, reqEditors...)
}

// SlurmV0042PostJobSubmitWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0042PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0042PostJobSubmitWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobSubmitWithResponse(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0042PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0042PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0042PostJobWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	if i.funcs.SlurmV0042PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0042PostJobWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostJobWithResponse(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
	if i.funcs.SlurmV0042PostJobWithResponse != nil {
		return i.funcs.SlurmV0042PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0042PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0042PostNodeWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	if i.funcs.SlurmV0042PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0042PostNodeWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0042UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodeResponse, error) {
	if i.funcs.SlurmV0042PostNodeWithResponse != nil {
		return i.funcs.SlurmV0042PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0042PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmV0042PostNodesWithBodyWithResponse implements v0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostNodesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	if i.funcs.SlurmV0042PostNodesWithBodyWithResponse != nil {
		return i.funcs.SlurmV0042PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0042PostNodesWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0042PostNodesWithResponse implements v0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0042PostNodesWithResponse(ctx context.Context, body api.SlurmV0042PostNodesJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostNodesResponse, error) {
	if i.funcs.SlurmV0042PostNodesWithResponse != nil {
		return i.funcs.SlurmV0042PostNodesWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0042PostNodesWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042DeleteAccountWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0042DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0042DeleteAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0042DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042DeleteAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0042DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042DeleteClusterWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0042DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0042DeleteSingleQosWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0042DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0042DeleteUserWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0042DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0042DeleteWckeyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0042DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0042DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0042DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0042GetAccountWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
	if i.funcs.SlurmdbV0042GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0042GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0042GetAccountsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0042GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0042GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0042GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0042GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0042GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0042GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetClusterWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0042GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClusterResponse, error) {
	if i.funcs.SlurmdbV0042GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0042GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0042GetClustersWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetClustersResponse, error) {
	if i.funcs.SlurmdbV0042GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0042GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetConfigWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetConfigResponse, error) {
	if i.funcs.SlurmdbV0042GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0042GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0042GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0042GetDiagWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetDiagResponse, error) {
	if i.funcs.SlurmdbV0042GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0042GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0042GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0042GetInstanceWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0042GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0042GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetInstancesWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0042GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0042GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0042GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetJobWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobResponse, error) {
	if i.funcs.SlurmdbV0042GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0042GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0042GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0042GetJobsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetJobsResponse, error) {
	if i.funcs.SlurmdbV0042GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0042GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetPingWithResponse implements v0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetPingResponse, error) {
	if i.funcs.SlurmdbV0042GetPingWithResponse != nil {
		return i.funcs.SlurmdbV0042GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0042GetPingWithResponse(ctx, reqEditors...)
}

// SlurmdbV0042GetQosWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0042GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetQosResponse, error) {
	if i.funcs.SlurmdbV0042GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0042GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetSingleQosWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0042GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0042GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0042GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0042GetTresWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetTresResponse, error) {
	if i.funcs.SlurmdbV0042GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0042GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0042GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0042GetUserWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0042GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUserResponse, error) {
	if i.funcs.SlurmdbV0042GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0042GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0042GetUsersWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0042GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetUsersResponse, error) {
	if i.funcs.SlurmdbV0042GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0042GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042GetWckeyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0042GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0042GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0042GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0042GetWckeysWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0042GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0042GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0042GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0042PostAccountsAssociationWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostAccountsAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAccountsAssociationWithResponse(ctx context.Context, body api.V0042OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0042PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostAccountsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0042PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostAccountsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAccountsWithResponse(ctx context.Context, body api.V0042OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0042PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostAssociationsWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0042PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostAssociationsWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostAssociationsWithResponse(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0042PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0042PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostClustersWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	if i.funcs.SlurmdbV0042PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0042PostClustersWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0042PostClustersParams, body api.V0042OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostClustersResponse, error) {
	if i.funcs.SlurmdbV0042PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0042PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0042PostConfigWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	if i.funcs.SlurmdbV0042PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostConfigWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostConfigWithResponse(ctx context.Context, body api.V0042OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostConfigResponse, error) {
	if i.funcs.SlurmdbV0042PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0042PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostQosWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	if i.funcs.SlurmdbV0042PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0042PostQosWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0042PostQosParams, body api.V0042OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostQosResponse, error) {
	if i.funcs.SlurmdbV0042PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0042PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0042PostTresWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	if i.funcs.SlurmdbV0042PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostTresWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostTresWithResponse(ctx context.Context, body api.V0042OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostTresResponse, error) {
	if i.funcs.SlurmdbV0042PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0042PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostUsersAssociationWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0042PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0042PostUsersAssociationWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0042PostUsersAssociationParams, body api.V0042OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0042PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0042PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0042PostUsersWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	if i.funcs.SlurmdbV0042PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0042PostUsersWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostUsersWithResponse(ctx context.Context, body api.V0042OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostUsersResponse, error) {
	if i.funcs.SlurmdbV0042PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0042PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0042PostWckeysWithBodyWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0042PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0042PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0042PostWckeysWithResponse implements V0042.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0042PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0042PostWckeysParams, body api.V0042OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0042PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0042PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0042PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
