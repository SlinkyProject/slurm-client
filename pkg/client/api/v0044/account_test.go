// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestAccountInterface_Compiles(t *testing.T) {
	var _ AccountInterface = &SlurmClient{}
}

func TestSlurmClient_GetAccount(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0044Account
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "research",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
							res := &api.SlurmdbV0044GetAccountResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiAccountsResp{
									Accounts: api.V0044AccountList{
										{Name: "research"},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "research",
			},
			want: &types.V0044Account{
				V0044Account: api.V0044Account{Name: "research"},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
							res := &api.SlurmdbV0044GetAccountResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiAccountsResp{
									Errors: &[]api.V0044OpenapiError{
										{Error: ptr.To("error 1")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "research",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0044GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "research",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.GetAccount(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListAccount(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    *types.V0044AccountList
		wantErr bool
	}{
		{
			name: "Empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			want: &types.V0044AccountList{
				Items: make([]types.V0044Account, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
							res := &api.SlurmdbV0044GetAccountsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiAccountsResp{
									Accounts: api.V0044AccountList{
										{Name: "a"},
										{Name: "b"},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			want: &types.V0044AccountList{
				Items: []types.V0044Account{
					{V0044Account: api.V0044Account{Name: "a"}},
					{V0044Account: api.V0044Account{Name: "b"}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAccountsResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.ListAccount(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_CreateAccount(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		req     any
		want    string
		wantErr bool
	}{
		{
			name:    "default",
			client:  fake.NewFakeClient(),
			req:     api.V0044Account{Name: "research"},
			want:    "research",
			wantErr: false,
		},
		{
			name:    "invalid type provided",
			client:  fake.NewFakeClient(),
			req:     api.V0044User{Name: "research"},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0044PostAccountsWithResponse: func(ctx context.Context, body api.SlurmdbV0044PostAccountsJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044PostAccountsResponse, error) {
						return nil, errors.New(http.StatusText(http.StatusBadGateway))
					},
				}).
				Build(),
			req:     api.V0044Account{Name: "research"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			got, err := c.CreateAccount(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlurmClient.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_DeleteAccount(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		wantErr bool
	}{
		{
			name:    "account does not exist",
			client:  fake.NewFakeClient(),
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0044DeleteAccountWithResponse: func(ctx context.Context, accountName string, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044DeleteAccountResponse, error) {
						return nil, errors.New(http.StatusText(http.StatusBadGateway))
					},
				}).
				Build(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			err := c.DeleteAccount(context.Background(), "research")
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
