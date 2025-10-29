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

func TestSlurmClient_DeleteNode(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx      context.Context
		nodeName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete existing node",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
							res := &api.SlurmV0044DeleteNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiResp{
									Errors: &[]api.V0044OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteNodeResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			if err := c.DeleteNode(tt.args.ctx, tt.args.nodeName); (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.DeleteNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlurmClient_UpdateNode(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx      context.Context
		nodeName string
		req      any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update existing node",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
				req:      api.V0044UpdateNodeMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
							res := &api.SlurmV0044PostNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiResp{
									Errors: &[]api.V0044OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
				req:      api.V0044UpdateNodeMsg{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0044UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostNodeResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			if err := c.UpdateNode(tt.args.ctx, tt.args.nodeName, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.UpdateNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlurmClient_GetNode(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx      context.Context
		nodeName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0044Node
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
							res := &api.SlurmV0044GetNodeResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiNodesResp{
									Nodes: []api.V0044Node{
										{Name: ptr.To("node-0")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			want: &types.V0044Node{
				V0044Node: api.V0044Node{
					Name: ptr.To("node-0"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
							res := &api.SlurmV0044GetNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiNodesResp{
									Errors: &[]api.V0044OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0044GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodeResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:      context.Background(),
				nodeName: "node-0",
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
			got, err := c.GetNode(tt.args.ctx, tt.args.nodeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListNodes(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0044NodeList
		wantErr bool
	}{
		{
			name: "Empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0044NodeList{
				Items: make([]types.V0044Node, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
							res := &api.SlurmV0044GetNodesResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiNodesResp{
									Nodes: []api.V0044Node{
										{Name: ptr.To("node-0")},
										{Name: ptr.To("node-1")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0044NodeList{
				Items: []types.V0044Node{
					{V0044Node: api.V0044Node{Name: ptr.To("node-0")}},
					{V0044Node: api.V0044Node{Name: ptr.To("node-1")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
							res := &api.SlurmV0044GetNodesResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiNodesResp{
									Errors: &[]api.V0044OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0044GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetNodesResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
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
			got, err := c.ListNodes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
