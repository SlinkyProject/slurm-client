// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestUserInterface_Compiles(t *testing.T) {
	var _ UserInterface = &SlurmClient{}
}

func TestSlurmClient_GetUser(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0045User
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
					SlurmdbV0045GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
						return &api.SlurmdbV0045GetUserResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0045OpenapiUsersResp{
								Users: api.V0045UserList{{Name: "alice"}},
							},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0045User{V0045User: api.V0045User{Name: "alice"}},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0045GetUserWithResponse: func(ctx context.Context, name string, params *api.SlurmdbV0045GetUserParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUserResponse, error) {
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
			got, err := c.GetUser(context.Background(), "alice")
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListUser(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0045UserList
		wantErr bool
	}{
		{
			name:    "Empty list",
			client:  fake.NewFakeClient(),
			want:    &types.V0045UserList{Items: make([]types.V0045User, 0)},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0045GetUsersWithResponse: func(ctx context.Context, params *api.SlurmdbV0045GetUsersParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0045GetUsersResponse, error) {
						return &api.SlurmdbV0045GetUsersResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0045OpenapiUsersResp{
								Users: api.V0045UserList{{Name: "alice"}, {Name: "bob"}},
							},
						}, nil
					},
				}).
				Build(),
			want: &types.V0045UserList{Items: []types.V0045User{
				{V0045User: api.V0045User{Name: "alice"}},
				{V0045User: api.V0045User{Name: "bob"}},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			got, err := c.ListUser(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		req     any
		want    string
		wantErr bool
	}{
		{name: "default", req: api.V0045User{Name: "alice"}, want: "alice", wantErr: false},
		{name: "invalid type provided", req: api.V0045Account{Name: "alice"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
			got, err := c.CreateUser(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlurmClient.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_DeleteUser(t *testing.T) {
	c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
	if err := c.DeleteUser(context.Background(), "alice"); err != nil {
		t.Errorf("SlurmClient.DeleteUser() error = %v, wantErr false", err)
	}
}
