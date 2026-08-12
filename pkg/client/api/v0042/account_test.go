// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestAccountInterface_Compiles(t *testing.T) {
	var _ AccountInterface = &SlurmClient{}
}

func TestSlurmClient_GetAccount(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0042Account
		wantErr bool
	}{
		{
			name:    "Not Found",
			client:  fake.NewFakeClient(),
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
						return &api.SlurmdbV0042GetAccountResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0042OpenapiAccountsResp{
								Accounts: api.V0042AccountList{{Name: "research"}},
							},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0042Account{V0042Account: api.V0042Account{Name: "research"}},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
						return &api.SlurmdbV0042GetAccountResponse{
							HTTPResponse: &http.Response{
								Status:     http.StatusText(http.StatusInternalServerError),
								StatusCode: http.StatusInternalServerError,
							},
							JSONDefault: &api.V0042OpenapiAccountsResp{
								Errors: &[]api.V0042OpenapiError{{Error: ptr.To("error 1")}},
							},
						}, nil
					},
				}).
				Build(),
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAccountWithResponse: func(ctx context.Context, accountName string, params *api.SlurmdbV0042GetAccountParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountResponse, error) {
						return nil, errors.New(http.StatusText(http.StatusBadGateway))
					},
				}).
				Build(),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			got, err := c.GetAccount(context.Background(), "research")
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
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0042AccountList
		wantErr bool
	}{
		{
			name:    "Empty list",
			client:  fake.NewFakeClient(),
			want:    &types.V0042AccountList{Items: make([]types.V0042Account, 0)},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAccountsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAccountsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAccountsResponse, error) {
						return &api.SlurmdbV0042GetAccountsResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0042OpenapiAccountsResp{
								Accounts: api.V0042AccountList{{Name: "a"}, {Name: "b"}},
							},
						}, nil
					},
				}).
				Build(),
			want: &types.V0042AccountList{Items: []types.V0042Account{
				{V0042Account: api.V0042Account{Name: "a"}},
				{V0042Account: api.V0042Account{Name: "b"}},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
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
		req     any
		want    string
		wantErr bool
	}{
		{name: "default", req: api.V0042Account{Name: "research"}, want: "research", wantErr: false},
		{name: "invalid type provided", req: api.V0042User{Name: "research"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
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
	c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
	if err := c.DeleteAccount(context.Background(), "research"); err != nil {
		t.Errorf("SlurmClient.DeleteAccount() error = %v, wantErr false", err)
	}
}
