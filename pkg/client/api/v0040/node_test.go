// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0040

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0040/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0040/interceptor"
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
						SlurmV0040DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
							res := &api.SlurmV0040DeleteNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
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
						SlurmV0040DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteNodeResponse, error) {
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
				req:      api.V0040UpdateNodeMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
							res := &api.SlurmV0040PostNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
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
				req:      api.V0040UpdateNodeMsg{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0040UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostNodeResponse, error) {
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
		want    *types.V0040Node
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
						SlurmV0040GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
							res := &api.SlurmV0040GetNodeResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiNodesResp{
									Nodes: []api.V0040Node{
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
			want: &types.V0040Node{
				V0040Node: api.V0040Node{
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
						SlurmV0040GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
							res := &api.SlurmV0040GetNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
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
						SlurmV0040GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0040GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodeResponse, error) {
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
		want    *types.V0040NodeList
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
			want: &types.V0040NodeList{
				Items: make([]types.V0040Node, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
							res := &api.SlurmV0040GetNodesResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiNodesResp{
									Nodes: []api.V0040Node{
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
			want: &types.V0040NodeList{
				Items: []types.V0040Node{
					{V0040Node: api.V0040Node{Name: ptr.To("node-0")}},
					{V0040Node: api.V0040Node{Name: ptr.To("node-1")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
							res := &api.SlurmV0040GetNodesResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
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
						SlurmV0040GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0040GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetNodesResponse, error) {
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
