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

func TestAssocInterface_Compiles(t *testing.T) {
	var _ AssocInterface = &SlurmClient{}
}

func TestSlurmClient_ListAssoc(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0044AssocList
		wantErr bool
	}{
		{
			name:    "Empty list",
			client:  fake.NewFakeClient(),
			want:    &types.V0044AssocList{Items: make([]types.V0044Assoc, 0)},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0044GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
						return &api.SlurmdbV0044GetAssociationsResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0044OpenapiAssocsResp{
								Associations: api.V0044AssocList{
									{User: "alice", Account: ptr.To("research")},
								},
							},
						}, nil
					},
				}).
				Build(),
			want: &types.V0044AssocList{Items: []types.V0044Assoc{
				{V0044Assoc: api.V0044Assoc{User: "alice", Account: ptr.To("research")}},
			}},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0044GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetAssociationsResponse, error) {
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
			got, err := c.ListAssoc(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListAssoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListAssoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_CreateAssoc(t *testing.T) {
	tests := []struct {
		name    string
		req     any
		wantErr bool
	}{
		{name: "default", req: api.V0044Assoc{User: "alice", Account: ptr.To("research")}, wantErr: false},
		{name: "invalid type provided", req: api.V0044Account{Name: "research"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
			err := c.CreateAssoc(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateAssoc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlurmClient_DeleteAssoc(t *testing.T) {
	c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
	err := c.DeleteAssoc(context.Background(), api.V0044Assoc{User: "alice", Account: ptr.To("research")})
	if err != nil {
		t.Errorf("SlurmClient.DeleteAssoc() error = %v, wantErr false", err)
	}
}
