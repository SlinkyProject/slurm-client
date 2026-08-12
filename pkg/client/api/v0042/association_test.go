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

func TestAssocInterface_Compiles(t *testing.T) {
	var _ AssocInterface = &SlurmClient{}
}

func TestSlurmClient_ListAssoc(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0042AssocList
		wantErr bool
	}{
		{
			name:    "Empty list",
			client:  fake.NewFakeClient(),
			want:    &types.V0042AssocList{Items: make([]types.V0042Assoc, 0)},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
						return &api.SlurmdbV0042GetAssociationsResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200: &api.V0042OpenapiAssocsResp{
								Associations: api.V0042AssocList{
									{User: "alice", Account: ptr.To("research")},
								},
							},
						}, nil
					},
				}).
				Build(),
			want: &types.V0042AssocList{Items: []types.V0042Assoc{
				{V0042Assoc: api.V0042Assoc{User: "alice", Account: ptr.To("research")}},
			}},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
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
		wantKey string
		wantErr bool
	}{
		{
			name:    "returns composite key from request",
			req:     api.V0042Assoc{Cluster: ptr.To("linux"), Account: ptr.To("research"), User: "alice"},
			wantKey: "linux/research/alice/",
			wantErr: false,
		},
		{name: "invalid type provided", req: api.V0042Account{Name: "research"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClient()}
			gotKey, err := c.CreateAssoc(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateAssoc() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotKey != tt.wantKey {
				t.Errorf("SlurmClient.CreateAssoc() key = %q, want %q", gotKey, tt.wantKey)
			}
		})
	}
}

func TestSlurmClient_UpdateAssoc(t *testing.T) {
	var gotBody api.V0042OpenapiAssocsResp
	c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClientBuilder().
		WithInterceptorFuncs(interceptor.Funcs{
			SlurmdbV0042PostAssociationsWithResponse: func(ctx context.Context, body api.V0042OpenapiAssocsResp, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042PostAssociationsResponse, error) {
				gotBody = body
				return &api.SlurmdbV0042PostAssociationsResponse{HTTPResponse: &fake.HttpSuccess}, nil
			},
		}).
		Build()}

	// req identity fields deliberately differ from key; key must win.
	req := api.V0042Assoc{Cluster: ptr.To("other"), Account: ptr.To("other"), User: "bob"}
	if err := c.UpdateAssoc(context.Background(), "linux/research/alice/", req); err != nil {
		t.Fatalf("SlurmClient.UpdateAssoc() error = %v", err)
	}
	if len(gotBody.Associations) != 1 {
		t.Fatalf("SlurmClient.UpdateAssoc() posted %d associations, want 1", len(gotBody.Associations))
	}
	got := gotBody.Associations[0]
	if ptr.Deref(got.Cluster, "") != "linux" || ptr.Deref(got.Account, "") != "research" || got.User != "alice" || got.Partition != nil {
		t.Errorf("SlurmClient.UpdateAssoc() pinned identity = %+v, want cluster=linux account=research user=alice partition=nil", got)
	}

	// invalid type provided
	if err := c.UpdateAssoc(context.Background(), "linux/research/alice/", api.V0042Account{Name: "x"}); err == nil {
		t.Errorf("SlurmClient.UpdateAssoc() expected error for invalid req type")
	}
}

func TestSlurmClient_GetAssoc(t *testing.T) {
	matching := func() api.ClientWithResponsesInterface {
		return fake.NewFakeClientBuilder().
			WithInterceptorFuncs(interceptor.Funcs{
				SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
					return &api.SlurmdbV0042GetAssociationsResponse{
						HTTPResponse: &fake.HttpSuccess,
						JSON200: &api.V0042OpenapiAssocsResp{
							Associations: api.V0042AssocList{
								{Cluster: ptr.To("linux"), Account: ptr.To("research"), User: "alice"},
							},
						},
					}, nil
				},
			}).
			Build()
	}
	tests := []struct {
		name    string
		key     string
		client  api.ClientWithResponsesInterface
		want    *types.V0042Assoc
		wantErr bool
	}{
		{
			name:   "found by key",
			key:    "linux/research/alice/",
			client: matching(),
			want:   &types.V0042Assoc{V0042Assoc: api.V0042Assoc{Cluster: ptr.To("linux"), Account: ptr.To("research"), User: "alice"}},
		},
		{
			name:    "not found - empty result",
			key:     "linux/research/bob/",
			client:  fake.NewFakeClient(),
			wantErr: true,
		},
		{
			name:    "not found - no key match",
			key:     "linux/research/bob/",
			client:  matching(),
			wantErr: true,
		},
		{
			name: "HTTP error",
			key:  "linux/research/alice/",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
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
			got, err := c.GetAssoc(context.Background(), tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetAssoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetAssoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_GetAssoc_FiltersServerSide(t *testing.T) {
	var gotParams *api.SlurmdbV0042GetAssociationsParams
	c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClientBuilder().
		WithInterceptorFuncs(interceptor.Funcs{
			SlurmdbV0042GetAssociationsWithResponse: func(ctx context.Context, params *api.SlurmdbV0042GetAssociationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042GetAssociationsResponse, error) {
				gotParams = params
				return &api.SlurmdbV0042GetAssociationsResponse{
					HTTPResponse: &fake.HttpSuccess,
					JSON200: &api.V0042OpenapiAssocsResp{
						Associations: api.V0042AssocList{
							{Cluster: ptr.To("linux"), Account: ptr.To("research"), User: "alice"},
						},
					},
				}, nil
			},
		}).
		Build()}

	if _, err := c.GetAssoc(context.Background(), "linux/research/alice/"); err != nil {
		t.Fatalf("SlurmClient.GetAssoc() error = %v", err)
	}
	if gotParams == nil {
		t.Fatal("SlurmClient.GetAssoc() did not call the associations endpoint")
	}
	if ptr.Deref(gotParams.Cluster, "") != "linux" || ptr.Deref(gotParams.Account, "") != "research" || ptr.Deref(gotParams.User, "") != "alice" {
		t.Errorf("SlurmClient.GetAssoc() filter params = %+v, want cluster=linux account=research user=alice", gotParams)
	}
	if gotParams.Partition != nil {
		t.Errorf("SlurmClient.GetAssoc() partition param = %v, want nil", gotParams.Partition)
	}
}

func TestSlurmClient_DeleteAssoc(t *testing.T) {
	tests := []struct {
		name          string
		assoc         api.V0042Assoc
		wantUserParam *string
	}{
		{
			name:          "with user",
			assoc:         api.V0042Assoc{Account: ptr.To("research"), User: "alice"},
			wantUserParam: ptr.To("alice"),
		},
		{
			name:          "empty user omitted",
			assoc:         api.V0042Assoc{Account: ptr.To("research")},
			wantUserParam: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotParams *api.SlurmdbV0042DeleteAssociationParams
			c := &SlurmClient{ClientWithResponsesInterface: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmdbV0042DeleteAssociationWithResponse: func(ctx context.Context, params *api.SlurmdbV0042DeleteAssociationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0042DeleteAssociationResponse, error) {
						gotParams = params
						return &api.SlurmdbV0042DeleteAssociationResponse{HTTPResponse: &fake.HttpSuccess, JSON200: &api.V0042OpenapiAssocsRemovedResp{}}, nil
					},
				}).
				Build()}
			if err := c.DeleteAssoc(context.Background(), tt.assoc); err != nil {
				t.Fatalf("SlurmClient.DeleteAssoc() error = %v", err)
			}
			if !reflect.DeepEqual(gotParams.User, tt.wantUserParam) {
				t.Errorf("SlurmClient.DeleteAssoc() User param = %v, want %v", ptr.Deref(gotParams.User, "<nil>"), ptr.Deref(tt.wantUserParam, "<nil>"))
			}
		})
	}
}
