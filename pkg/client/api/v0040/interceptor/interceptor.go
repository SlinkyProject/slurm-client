// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package interceptor

import (
	"context"
	"io"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
)

type Funcs struct {
	SlurmV0040DeleteJobWithResponse                         func(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error)
	SlurmV0040DeleteJobsWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error)
	SlurmV0040DeleteJobsWithResponse                        func(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error)
	SlurmV0040DeleteNodeWithResponse                        func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error)
	SlurmV0040GetDiagWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error)
	SlurmV0040GetJobWithResponse                            func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error)
	SlurmV0040GetJobsStateWithResponse                      func(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error)
	SlurmV0040GetJobsWithResponse                           func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error)
	SlurmV0040GetLicensesWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error)
	SlurmV0040GetNodeWithResponse                           func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error)
	SlurmV0040GetNodesWithResponse                          func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error)
	SlurmV0040GetPartitionWithResponse                      func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error)
	SlurmV0040GetPartitionsWithResponse                     func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error)
	SlurmV0040GetPingWithResponse                           func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error)
	SlurmV0040GetReconfigureWithResponse                    func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error)
	SlurmV0040GetReservationWithResponse                    func(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error)
	SlurmV0040GetReservationsWithResponse                   func(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error)
	SlurmV0040GetSharesWithResponse                         func(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error)
	SlurmV0040PostJobSubmitWithBodyWithResponse             func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error)
	SlurmV0040PostJobSubmitWithResponse                     func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error)
	SlurmV0040PostJobWithBodyWithResponse                   func(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error)
	SlurmV0040PostJobWithResponse                           func(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error)
	SlurmV0040PostNodeWithBodyWithResponse                  func(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error)
	SlurmV0040PostNodeWithResponse                          func(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error)
	SlurmdbV0040DeleteAccountWithResponse                   func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error)
	SlurmdbV0040DeleteAssociationWithResponse               func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error)
	SlurmdbV0040DeleteAssociationsWithResponse              func(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error)
	SlurmdbV0040DeleteClusterWithResponse                   func(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error)
	SlurmdbV0040DeleteSingleQosWithResponse                 func(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error)
	SlurmdbV0040DeleteUserWithResponse                      func(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error)
	SlurmdbV0040DeleteWckeyWithResponse                     func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error)
	SlurmdbV0040GetAccountWithResponse                      func(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error)
	SlurmdbV0040GetAccountsWithResponse                     func(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error)
	SlurmdbV0040GetAssociationWithResponse                  func(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error)
	SlurmdbV0040GetAssociationsWithResponse                 func(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error)
	SlurmdbV0040GetClusterWithResponse                      func(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error)
	SlurmdbV0040GetClustersWithResponse                     func(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error)
	SlurmdbV0040GetConfigWithResponse                       func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error)
	SlurmdbV0040GetDiagWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error)
	SlurmdbV0040GetInstanceWithResponse                     func(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error)
	SlurmdbV0040GetInstancesWithResponse                    func(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error)
	SlurmdbV0040GetJobWithResponse                          func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error)
	SlurmdbV0040GetJobsWithResponse                         func(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error)
	SlurmdbV0040GetQosWithResponse                          func(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error)
	SlurmdbV0040GetSingleQosWithResponse                    func(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error)
	SlurmdbV0040GetTresWithResponse                         func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error)
	SlurmdbV0040GetUserWithResponse                         func(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error)
	SlurmdbV0040GetUsersWithResponse                        func(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error)
	SlurmdbV0040GetWckeyWithResponse                        func(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error)
	SlurmdbV0040GetWckeysWithResponse                       func(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error)
	SlurmdbV0040PostAccountsAssociationWithBodyWithResponse func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error)
	SlurmdbV0040PostAccountsAssociationWithResponse         func(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error)
	SlurmdbV0040PostAccountsWithBodyWithResponse            func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error)
	SlurmdbV0040PostAccountsWithResponse                    func(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error)
	SlurmdbV0040PostAssociationsWithBodyWithResponse        func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error)
	SlurmdbV0040PostAssociationsWithResponse                func(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error)
	SlurmdbV0040PostClustersWithBodyWithResponse            func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error)
	SlurmdbV0040PostClustersWithResponse                    func(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error)
	SlurmdbV0040PostConfigWithBodyWithResponse              func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error)
	SlurmdbV0040PostConfigWithResponse                      func(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error)
	SlurmdbV0040PostQosWithBodyWithResponse                 func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error)
	SlurmdbV0040PostQosWithResponse                         func(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error)
	SlurmdbV0040PostTresWithBodyWithResponse                func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error)
	SlurmdbV0040PostTresWithResponse                        func(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error)
	SlurmdbV0040PostUsersAssociationWithBodyWithResponse    func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error)
	SlurmdbV0040PostUsersAssociationWithResponse            func(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error)
	SlurmdbV0040PostUsersWithBodyWithResponse               func(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error)
	SlurmdbV0040PostUsersWithResponse                       func(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error)
	SlurmdbV0040PostWckeysWithBodyWithResponse              func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error)
	SlurmdbV0040PostWckeysWithResponse                      func(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error)
}

type interceptor struct {
	client api.ClientWithResponsesInterface
	funcs  Funcs
}

// SlurmV0040DeleteJobWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040DeleteJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
	if i.funcs.SlurmV0040DeleteJobWithResponse != nil {
		return i.funcs.SlurmV0040DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0040DeleteJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0040DeleteJobsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040DeleteJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	if i.funcs.SlurmV0040DeleteJobsWithBodyWithResponse != nil {
		return i.funcs.SlurmV0040DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0040DeleteJobsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0040DeleteJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040DeleteJobsWithResponse(ctx context.Context, body api.V0040KillJobsMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobsResponse, error) {
	if i.funcs.SlurmV0040DeleteJobsWithResponse != nil {
		return i.funcs.SlurmV0040DeleteJobsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0040DeleteJobsWithResponse(ctx, body, reqEditors...)
}

// SlurmV0040DeleteNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040DeleteNodeWithResponse(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
	if i.funcs.SlurmV0040DeleteNodeWithResponse != nil {
		return i.funcs.SlurmV0040DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
	}
	return i.client.SlurmV0040DeleteNodeWithResponse(ctx, nodeName, reqEditors...)
}

// SlurmV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetDiagResponse, error) {
	if i.funcs.SlurmV0040GetDiagWithResponse != nil {
		return i.funcs.SlurmV0040GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0040GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetJobWithResponse(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
	if i.funcs.SlurmV0040GetJobWithResponse != nil {
		return i.funcs.SlurmV0040GetJobWithResponse(ctx, jobId, params, reqEditors...)
	}
	return i.client.SlurmV0040GetJobWithResponse(ctx, jobId, params, reqEditors...)
}

// SlurmV0040GetJobsStateWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetJobsStateWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsStateParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsStateResponse, error) {
	if i.funcs.SlurmV0040GetJobsStateWithResponse != nil {
		return i.funcs.SlurmV0040GetJobsStateWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetJobsStateWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
	if i.funcs.SlurmV0040GetJobsWithResponse != nil {
		return i.funcs.SlurmV0040GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040GetLicensesWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetLicensesWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetLicensesResponse, error) {
	if i.funcs.SlurmV0040GetLicensesWithResponse != nil {
		return i.funcs.SlurmV0040GetLicensesWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0040GetLicensesWithResponse(ctx, reqEditors...)
}

// SlurmV0040GetNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetNodeWithResponse(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
	if i.funcs.SlurmV0040GetNodeWithResponse != nil {
		return i.funcs.SlurmV0040GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
	}
	return i.client.SlurmV0040GetNodeWithResponse(ctx, nodeName, params, reqEditors...)
}

// SlurmV0040GetNodesWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetNodesWithResponse(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
	if i.funcs.SlurmV0040GetNodesWithResponse != nil {
		return i.funcs.SlurmV0040GetNodesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetNodesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040GetPartitionWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetPartitionWithResponse(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
	if i.funcs.SlurmV0040GetPartitionWithResponse != nil {
		return i.funcs.SlurmV0040GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
	}
	return i.client.SlurmV0040GetPartitionWithResponse(ctx, partitionName, params, reqEditors...)
}

// SlurmV0040GetPartitionsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetPartitionsWithResponse(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
	if i.funcs.SlurmV0040GetPartitionsWithResponse != nil {
		return i.funcs.SlurmV0040GetPartitionsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetPartitionsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040GetPingWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetPingWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPingResponse, error) {
	if i.funcs.SlurmV0040GetPingWithResponse != nil {
		return i.funcs.SlurmV0040GetPingWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0040GetPingWithResponse(ctx, reqEditors...)
}

// SlurmV0040GetReconfigureWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetReconfigureWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReconfigureResponse, error) {
	if i.funcs.SlurmV0040GetReconfigureWithResponse != nil {
		return i.funcs.SlurmV0040GetReconfigureWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmV0040GetReconfigureWithResponse(ctx, reqEditors...)
}

// SlurmV0040GetReservationWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetReservationWithResponse(ctx context.Context, reservationName string, params *api.SlurmV0040GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationResponse, error) {
	if i.funcs.SlurmV0040GetReservationWithResponse != nil {
		return i.funcs.SlurmV0040GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
	}
	return i.client.SlurmV0040GetReservationWithResponse(ctx, reservationName, params, reqEditors...)
}

// SlurmV0040GetReservationsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetReservationsWithResponse(ctx context.Context, params *api.SlurmV0040GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetReservationsResponse, error) {
	if i.funcs.SlurmV0040GetReservationsWithResponse != nil {
		return i.funcs.SlurmV0040GetReservationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetReservationsWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040GetSharesWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040GetSharesWithResponse(ctx context.Context, params *api.SlurmV0040GetSharesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetSharesResponse, error) {
	if i.funcs.SlurmV0040GetSharesWithResponse != nil {
		return i.funcs.SlurmV0040GetSharesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmV0040GetSharesWithResponse(ctx, params, reqEditors...)
}

// SlurmV0040PostJobSubmitWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostJobSubmitWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0040PostJobSubmitWithBodyWithResponse != nil {
		return i.funcs.SlurmV0040PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0040PostJobSubmitWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmV0040PostJobSubmitWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostJobSubmitWithResponse(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
	if i.funcs.SlurmV0040PostJobSubmitWithResponse != nil {
		return i.funcs.SlurmV0040PostJobSubmitWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmV0040PostJobSubmitWithResponse(ctx, body, reqEditors...)
}

// SlurmV0040PostJobWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	if i.funcs.SlurmV0040PostJobWithBodyWithResponse != nil {
		return i.funcs.SlurmV0040PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0040PostJobWithBodyWithResponse(ctx, jobId, contentType, body, reqEditors...)
}

// SlurmV0040PostJobWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostJobWithResponse(ctx context.Context, jobId string, body api.V0040JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobResponse, error) {
	if i.funcs.SlurmV0040PostJobWithResponse != nil {
		return i.funcs.SlurmV0040PostJobWithResponse(ctx, jobId, body, reqEditors...)
	}
	return i.client.SlurmV0040PostJobWithResponse(ctx, jobId, body, reqEditors...)
}

// SlurmV0040PostNodeWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostNodeWithBodyWithResponse(ctx context.Context, nodeName string, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	if i.funcs.SlurmV0040PostNodeWithBodyWithResponse != nil {
		return i.funcs.SlurmV0040PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
	}
	return i.client.SlurmV0040PostNodeWithBodyWithResponse(ctx, nodeName, contentType, body, reqEditors...)
}

// SlurmV0040PostNodeWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmV0040PostNodeWithResponse(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
	if i.funcs.SlurmV0040PostNodeWithResponse != nil {
		return i.funcs.SlurmV0040PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
	}
	return i.client.SlurmV0040PostNodeWithResponse(ctx, nodeName, body, reqEditors...)
}

// SlurmdbV0040DeleteAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteAccountWithResponse(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAccountResponse, error) {
	if i.funcs.SlurmdbV0040DeleteAccountWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteAccountWithResponse(ctx, accountName, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteAccountWithResponse(ctx, accountName, reqEditors...)
}

// SlurmdbV0040DeleteAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationResponse, error) {
	if i.funcs.SlurmdbV0040DeleteAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040DeleteAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040DeleteAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteAssociationsResponse, error) {
	if i.funcs.SlurmdbV0040DeleteAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040DeleteClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040DeleteClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteClusterResponse, error) {
	if i.funcs.SlurmdbV0040DeleteClusterWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0040DeleteSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteSingleQosWithResponse(ctx context.Context, qos string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteSingleQosResponse, error) {
	if i.funcs.SlurmdbV0040DeleteSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteSingleQosWithResponse(ctx, qos, reqEditors...)
}

// SlurmdbV0040DeleteUserWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteUserWithResponse(ctx context.Context, name string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteUserResponse, error) {
	if i.funcs.SlurmdbV0040DeleteUserWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteUserWithResponse(ctx, name, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteUserWithResponse(ctx, name, reqEditors...)
}

// SlurmdbV0040DeleteWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040DeleteWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040DeleteWckeyResponse, error) {
	if i.funcs.SlurmdbV0040DeleteWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0040DeleteWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0040DeleteWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0040GetAccountWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetAccountWithResponse(ctx context.Context, accountName string, params *api.SlurmdbV0040GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountResponse, error) {
	if i.funcs.SlurmdbV0040GetAccountWithResponse != nil {
		return i.funcs.SlurmdbV0040GetAccountWithResponse(ctx, accountName, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetAccountWithResponse(ctx, accountName, params, reqEditors...)
}

// SlurmdbV0040GetAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetAccountsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAccountsResponse, error) {
	if i.funcs.SlurmdbV0040GetAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0040GetAccountsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetAccountsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationResponse, error) {
	if i.funcs.SlurmdbV0040GetAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0040GetAssociationWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetAssociationWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetAssociationsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetAssociationsResponse, error) {
	if i.funcs.SlurmdbV0040GetAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0040GetAssociationsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetAssociationsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetClusterWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetClusterWithResponse(ctx context.Context, clusterName string, params *api.SlurmdbV0040GetClusterParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClusterResponse, error) {
	if i.funcs.SlurmdbV0040GetClusterWithResponse != nil {
		return i.funcs.SlurmdbV0040GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetClusterWithResponse(ctx, clusterName, params, reqEditors...)
}

// SlurmdbV0040GetClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetClustersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetClustersResponse, error) {
	if i.funcs.SlurmdbV0040GetClustersWithResponse != nil {
		return i.funcs.SlurmdbV0040GetClustersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetClustersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetConfigWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetConfigResponse, error) {
	if i.funcs.SlurmdbV0040GetConfigWithResponse != nil {
		return i.funcs.SlurmdbV0040GetConfigWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0040GetConfigWithResponse(ctx, reqEditors...)
}

// SlurmdbV0040GetDiagWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetDiagWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetDiagResponse, error) {
	if i.funcs.SlurmdbV0040GetDiagWithResponse != nil {
		return i.funcs.SlurmdbV0040GetDiagWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0040GetDiagWithResponse(ctx, reqEditors...)
}

// SlurmdbV0040GetInstanceWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetInstanceWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstanceParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstanceResponse, error) {
	if i.funcs.SlurmdbV0040GetInstanceWithResponse != nil {
		return i.funcs.SlurmdbV0040GetInstanceWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetInstanceWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetInstancesWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetInstancesWithResponse(ctx context.Context, params *api.SlurmdbV0040GetInstancesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetInstancesResponse, error) {
	if i.funcs.SlurmdbV0040GetInstancesWithResponse != nil {
		return i.funcs.SlurmdbV0040GetInstancesWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetInstancesWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetJobWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetJobWithResponse(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobResponse, error) {
	if i.funcs.SlurmdbV0040GetJobWithResponse != nil {
		return i.funcs.SlurmdbV0040GetJobWithResponse(ctx, jobId, reqEditors...)
	}
	return i.client.SlurmdbV0040GetJobWithResponse(ctx, jobId, reqEditors...)
}

// SlurmdbV0040GetJobsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetJobsWithResponse(ctx context.Context, params *api.SlurmdbV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetJobsResponse, error) {
	if i.funcs.SlurmdbV0040GetJobsWithResponse != nil {
		return i.funcs.SlurmdbV0040GetJobsWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetJobsWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetQosWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetQosWithResponse(ctx context.Context, params *api.SlurmdbV0040GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetQosResponse, error) {
	if i.funcs.SlurmdbV0040GetQosWithResponse != nil {
		return i.funcs.SlurmdbV0040GetQosWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetQosWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetSingleQosWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetSingleQosWithResponse(ctx context.Context, qos string, params *api.SlurmdbV0040GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetSingleQosResponse, error) {
	if i.funcs.SlurmdbV0040GetSingleQosWithResponse != nil {
		return i.funcs.SlurmdbV0040GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetSingleQosWithResponse(ctx, qos, params, reqEditors...)
}

// SlurmdbV0040GetTresWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetTresWithResponse(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetTresResponse, error) {
	if i.funcs.SlurmdbV0040GetTresWithResponse != nil {
		return i.funcs.SlurmdbV0040GetTresWithResponse(ctx, reqEditors...)
	}
	return i.client.SlurmdbV0040GetTresWithResponse(ctx, reqEditors...)
}

// SlurmdbV0040GetUserWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetUserWithResponse(ctx context.Context, name string, params *api.SlurmdbV0040GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUserResponse, error) {
	if i.funcs.SlurmdbV0040GetUserWithResponse != nil {
		return i.funcs.SlurmdbV0040GetUserWithResponse(ctx, name, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetUserWithResponse(ctx, name, params, reqEditors...)
}

// SlurmdbV0040GetUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetUsersWithResponse(ctx context.Context, params *api.SlurmdbV0040GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetUsersResponse, error) {
	if i.funcs.SlurmdbV0040GetUsersWithResponse != nil {
		return i.funcs.SlurmdbV0040GetUsersWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetUsersWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040GetWckeyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetWckeyWithResponse(ctx context.Context, id string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeyResponse, error) {
	if i.funcs.SlurmdbV0040GetWckeyWithResponse != nil {
		return i.funcs.SlurmdbV0040GetWckeyWithResponse(ctx, id, reqEditors...)
	}
	return i.client.SlurmdbV0040GetWckeyWithResponse(ctx, id, reqEditors...)
}

// SlurmdbV0040GetWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040GetWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040GetWckeysParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040GetWckeysResponse, error) {
	if i.funcs.SlurmdbV0040GetWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0040GetWckeysWithResponse(ctx, params, reqEditors...)
	}
	return i.client.SlurmdbV0040GetWckeysWithResponse(ctx, params, reqEditors...)
}

// SlurmdbV0040PostAccountsAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAccountsAssociationWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostAccountsAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAccountsAssociationWithResponse(ctx context.Context, body api.V0040OpenapiAccountsAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsAssociationResponse, error) {
	if i.funcs.SlurmdbV0040PostAccountsAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAccountsAssociationWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostAccountsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAccountsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0040PostAccountsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAccountsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostAccountsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAccountsWithResponse(ctx context.Context, body api.V0040OpenapiAccountsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAccountsResponse, error) {
	if i.funcs.SlurmdbV0040PostAccountsWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAccountsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAccountsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostAssociationsWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0040PostAssociationsWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAssociationsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostAssociationsWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostAssociationsWithResponse(ctx context.Context, body api.V0040OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostAssociationsResponse, error) {
	if i.funcs.SlurmdbV0040PostAssociationsWithResponse != nil {
		return i.funcs.SlurmdbV0040PostAssociationsWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostAssociationsWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostClustersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostClustersWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	if i.funcs.SlurmdbV0040PostClustersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostClustersWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0040PostClustersWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostClustersWithResponse(ctx context.Context, params *api.SlurmdbV0040PostClustersParams, body api.V0040OpenapiClustersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostClustersResponse, error) {
	if i.funcs.SlurmdbV0040PostClustersWithResponse != nil {
		return i.funcs.SlurmdbV0040PostClustersWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostClustersWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0040PostConfigWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostConfigWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	if i.funcs.SlurmdbV0040PostConfigWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostConfigWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostConfigWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostConfigWithResponse(ctx context.Context, body api.V0040OpenapiSlurmdbdConfigResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostConfigResponse, error) {
	if i.funcs.SlurmdbV0040PostConfigWithResponse != nil {
		return i.funcs.SlurmdbV0040PostConfigWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostConfigWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostQosWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostQosWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	if i.funcs.SlurmdbV0040PostQosWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostQosWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0040PostQosWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostQosWithResponse(ctx context.Context, params *api.SlurmdbV0040PostQosParams, body api.V0040OpenapiSlurmdbdQosResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostQosResponse, error) {
	if i.funcs.SlurmdbV0040PostQosWithResponse != nil {
		return i.funcs.SlurmdbV0040PostQosWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostQosWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0040PostTresWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostTresWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	if i.funcs.SlurmdbV0040PostTresWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostTresWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostTresWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostTresWithResponse(ctx context.Context, body api.V0040OpenapiTresResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostTresResponse, error) {
	if i.funcs.SlurmdbV0040PostTresWithResponse != nil {
		return i.funcs.SlurmdbV0040PostTresWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostTresWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostUsersAssociationWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0040PostUsersAssociationWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostUsersAssociationWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0040PostUsersAssociationWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostUsersAssociationWithResponse(ctx context.Context, params *api.SlurmdbV0040PostUsersAssociationParams, body api.V0040OpenapiUsersAddCondResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersAssociationResponse, error) {
	if i.funcs.SlurmdbV0040PostUsersAssociationWithResponse != nil {
		return i.funcs.SlurmdbV0040PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostUsersAssociationWithResponse(ctx, params, body, reqEditors...)
}

// SlurmdbV0040PostUsersWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	if i.funcs.SlurmdbV0040PostUsersWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostUsersWithBodyWithResponse(ctx, contentType, body, reqEditors...)
}

// SlurmdbV0040PostUsersWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostUsersWithResponse(ctx context.Context, body api.V0040OpenapiUsersResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostUsersResponse, error) {
	if i.funcs.SlurmdbV0040PostUsersWithResponse != nil {
		return i.funcs.SlurmdbV0040PostUsersWithResponse(ctx, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostUsersWithResponse(ctx, body, reqEditors...)
}

// SlurmdbV0040PostWckeysWithBodyWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostWckeysWithBodyWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, contentType string, body io.Reader, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0040PostWckeysWithBodyWithResponse != nil {
		return i.funcs.SlurmdbV0040PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostWckeysWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
}

// SlurmdbV0040PostWckeysWithResponse implements v0040.ClientWithResponsesInterface.
func (i *interceptor) SlurmdbV0040PostWckeysWithResponse(ctx context.Context, params *api.SlurmdbV0040PostWckeysParams, body api.V0040OpenapiWckeyResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0040PostWckeysResponse, error) {
	if i.funcs.SlurmdbV0040PostWckeysWithResponse != nil {
		return i.funcs.SlurmdbV0040PostWckeysWithResponse(ctx, params, body, reqEditors...)
	}
	return i.client.SlurmdbV0040PostWckeysWithResponse(ctx, params, body, reqEditors...)
}

var _ api.ClientWithResponsesInterface = &interceptor{}

func NewClient(interceptedClient api.ClientWithResponsesInterface, funcs Funcs) api.ClientWithResponsesInterface {
	return &interceptor{
		client: interceptedClient,
		funcs:  funcs,
	}
}
