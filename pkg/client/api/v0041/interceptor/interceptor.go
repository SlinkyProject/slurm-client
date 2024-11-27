// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
)

type Funcs struct {
	SlurmV0041DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error)
	SlurmV0041DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error)
	SlurmV0041DeleteJobsWithResponse                        func(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error)
	SlurmV0041DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error)
	SlurmV0041GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error)
	SlurmV0041GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error)
	SlurmV0041GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error)
	SlurmV0041GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error)
	SlurmV0041GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error)
	SlurmV0041GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error)
	SlurmV0041GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error)
	SlurmV0041GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error)
	SlurmV0041GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error)
	SlurmV0041GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error)
	SlurmV0041GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error)
	SlurmV0041GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error)
	SlurmV0041GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error)
	SlurmV0041GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error)
	SlurmV0041PostJobAllocateWithBodyWithResponse           func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error)
	SlurmV0041PostJobAllocateWithResponse                   func(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error)
	SlurmV0041PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error)
	SlurmV0041PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error)
	SlurmV0041PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error)
	SlurmV0041PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error)
	SlurmV0041PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error)
	SlurmV0041PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error)
	SlurmdbV0041DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error)
	SlurmdbV0041DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error)
	SlurmdbV0041DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error)
	SlurmdbV0041DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error)
	SlurmdbV0041DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error)
	SlurmdbV0041DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error)
	SlurmdbV0041DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error)
	SlurmdbV0041GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error)
	SlurmdbV0041GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error)
	SlurmdbV0041GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error)
	SlurmdbV0041GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error)
	SlurmdbV0041GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error)
	SlurmdbV0041GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error)
	SlurmdbV0041GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error)
	SlurmdbV0041GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error)
	SlurmdbV0041GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error)
	SlurmdbV0041GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error)
	SlurmdbV0041GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error)
	SlurmdbV0041GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error)
	SlurmdbV0041GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error)
	SlurmdbV0041GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error)
	SlurmdbV0041GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error)
	SlurmdbV0041GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error)
	SlurmdbV0041GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error)
	SlurmdbV0041GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error)
	SlurmdbV0041GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error)
	SlurmdbV0041PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error)
	SlurmdbV0041PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error)
	SlurmdbV0041PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error)
	SlurmdbV0041PostAccountsWithResponse                    func(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error)
	SlurmdbV0041PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error)
	SlurmdbV0041PostAssociationsWithResponse                func(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error)
	SlurmdbV0041PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error)
	SlurmdbV0041PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error)
	SlurmdbV0041PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error)
	SlurmdbV0041PostConfigWithResponse                      func(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error)
	SlurmdbV0041PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error)
	SlurmdbV0041PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error)
	SlurmdbV0041PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error)
	SlurmdbV0041PostTresWithResponse                        func(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error)
	SlurmdbV0041PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error)
	SlurmdbV0041PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error)
	SlurmdbV0041PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error)
	SlurmdbV0041PostUsersWithResponse                       func(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error)
	SlurmdbV0041PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error)
	SlurmdbV0041PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0041DeleteJobWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobResponse, error) {
	if i.funcs.SlurmV0041DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0041DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0041DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0041DeleteJobsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	if i.funcs.SlurmV0041DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0041DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0041DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0041DeleteJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041DeleteJobsWithResponse(ctx context.Context, body api.V0041KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteJobsResponse, error) {
	if i.funcs.SlurmV0041DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0041DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0041DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0041DeleteNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041DeleteNodeResponse, error) {
	if i.funcs.SlurmV0041DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0041DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0041DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
	if i.funcs.SlurmV0041GetDiagWithResponse != nil {
		return i.funcs.SlurmV0041GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0041GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0041GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobResponse, error) {
	if i.funcs.SlurmV0041GetJobWithResponse != nil {
		return i.funcs.SlurmV0041GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0041GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0041GetJobsStateWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsStateResponse, error) {
	if i.funcs.SlurmV0041GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0041GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetJobsResponse, error) {
	if i.funcs.SlurmV0041GetJobsWithResponse != nil {
		return i.funcs.SlurmV0041GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041GetLicensesWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetLicensesResponse, error) {
	if i.funcs.SlurmV0041GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0041GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0041GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0041GetNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0041GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodeResponse, error) {
	if i.funcs.SlurmV0041GetNodeWithResponse != nil {
		return i.funcs.SlurmV0041GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0041GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0041GetNodesWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetNodesWithResponse(ctx context.Context, params *api.SlurmV0041GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetNodesResponse, error) {
	if i.funcs.SlurmV0041GetNodesWithResponse != nil {
		return i.funcs.SlurmV0041GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041GetPartitionWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0041GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionResponse, error) {
	if i.funcs.SlurmV0041GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0041GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0041GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0041GetPartitionsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0041GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPartitionsResponse, error) {
	if i.funcs.SlurmV0041GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0041GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041GetPingWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetPingResponse, error) {
	if i.funcs.SlurmV0041GetPingWithResponse != nil {
		return i.funcs.SlurmV0041GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0041GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0041GetReconfigureWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
	if i.funcs.SlurmV0041GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0041GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0041GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0041GetReservationWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
	if i.funcs.SlurmV0041GetReservationWithResponse != nil {
		return i.funcs.SlurmV0041GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0041GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0041GetReservationsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
	if i.funcs.SlurmV0041GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0041GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041GetSharesWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041GetSharesWithResponse(ctx context.Context, params *api.SlurmV0041GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetSharesResponse, error) {
	if i.funcs.SlurmV0041GetSharesWithResponse != nil {
		return i.funcs.SlurmV0041GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0041GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0041PostJobAllocateWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobAllocateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0041PostJobAllocateWithBodyWithResponse != nil {
		return i.funcs.SlurmV0041PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobAllocateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0041PostJobAllocateWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobAllocateWithResponse(ctx context.Context, body api.V0041JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobAllocateResponse, error) {
	if i.funcs.SlurmV0041PostJobAllocateWithResponse != nil {
		return i.funcs.SlurmV0041PostJobAllocateWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobAllocateWithResponse(ctx, body, reqEditors...)
}

// SlurmV0041PostJobSubmitWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0041PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0041PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0041PostJobSubmitWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobSubmitWithResponse(ctx context.Context, body api.V0041JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0041PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0041PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0041PostJobWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	if i.funcs.SlurmV0041PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0041PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0041PostJobWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostJobWithResponse(ctx context.Context, jobId string, body api.V0041JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostJobResponse, error) {
	if i.funcs.SlurmV0041PostJobWithResponse != nil {
		return i.funcs.SlurmV0041PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0041PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0041PostNodeWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	if i.funcs.SlurmV0041PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0041PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0041PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0041PostNodeWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0041PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0041UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041PostNodeResponse, error) {
	if i.funcs.SlurmV0041PostNodeWithResponse != nil {
		return i.funcs.SlurmV0041PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0041PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmdbV0041DeleteAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0041DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0041DeleteAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0041DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041DeleteAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0041DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041DeleteClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0041DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0041DeleteSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0041DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0041DeleteUserWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0041DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0041DeleteWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0041DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0041DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0041DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0041GetAccountWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0041GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountResponse, error) {
	if i.funcs.SlurmdbV0041GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0041GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0041GetAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0041GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0041GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0041GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0041GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0041GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0041GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetClusterWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0041GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClusterResponse, error) {
	if i.funcs.SlurmdbV0041GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0041GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0041GetClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetClustersResponse, error) {
	if i.funcs.SlurmdbV0041GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0041GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetConfigResponse, error) {
	if i.funcs.SlurmdbV0041GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0041GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0041GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0041GetDiagWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetDiagResponse, error) {
	if i.funcs.SlurmdbV0041GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0041GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0041GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0041GetInstanceWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0041GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0041GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetInstancesWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0041GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0041GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0041GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetJobWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobResponse, error) {
	if i.funcs.SlurmdbV0041GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0041GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0041GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0041GetJobsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0041GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetJobsResponse, error) {
	if i.funcs.SlurmdbV0041GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0041GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetQosWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0041GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetQosResponse, error) {
	if i.funcs.SlurmdbV0041GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0041GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetSingleQosWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0041GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0041GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0041GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0041GetTresWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetTresResponse, error) {
	if i.funcs.SlurmdbV0041GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0041GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0041GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0041GetUserWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0041GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUserResponse, error) {
	if i.funcs.SlurmdbV0041GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0041GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0041GetUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0041GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetUsersResponse, error) {
	if i.funcs.SlurmdbV0041GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0041GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041GetWckeyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0041GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0041GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0041GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0041GetWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0041GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0041GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0041GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0041PostAccountsAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0041PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostAccountsAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAccountsAssociationWithResponse(ctx context.Context, body api.V0041OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0041PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostAccountsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0041PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostAccountsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAccountsWithResponse(ctx context.Context, body api.V0041OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0041PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostAssociationsWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0041PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostAssociationsWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostAssociationsWithResponse(ctx context.Context, body api.V0041OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0041PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0041PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostClustersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	if i.funcs.SlurmdbV0041PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0041PostClustersWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0041PostClustersParams, body api.V0041OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostClustersResponse, error) {
	if i.funcs.SlurmdbV0041PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0041PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0041PostConfigWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	if i.funcs.SlurmdbV0041PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostConfigWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostConfigWithResponse(ctx context.Context, body api.V0041OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostConfigResponse, error) {
	if i.funcs.SlurmdbV0041PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0041PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostQosWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	if i.funcs.SlurmdbV0041PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0041PostQosWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0041PostQosParams, body api.V0041OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostQosResponse, error) {
	if i.funcs.SlurmdbV0041PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0041PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0041PostTresWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	if i.funcs.SlurmdbV0041PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostTresWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostTresWithResponse(ctx context.Context, body api.V0041OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostTresResponse, error) {
	if i.funcs.SlurmdbV0041PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0041PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostUsersAssociationWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0041PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0041PostUsersAssociationWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0041PostUsersAssociationParams, body api.V0041OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0041PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0041PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0041PostUsersWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	if i.funcs.SlurmdbV0041PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0041PostUsersWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostUsersWithResponse(ctx context.Context, body api.V0041OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostUsersResponse, error) {
	if i.funcs.SlurmdbV0041PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0041PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0041PostWckeysWithBodyWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0041PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0041PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0041PostWckeysWithResponse implements v0041.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0041PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0041PostWckeysParams, body api.V0041OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0041PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0041PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0041PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0041PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
