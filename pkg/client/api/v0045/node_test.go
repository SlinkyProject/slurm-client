// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_CreateNewNode(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx context.Context
		req any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "Create new node",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0045OpenapiCreateNodeReq{
					NodeConf: "NodeName=node-0 CPUs=4 State=EXTERNAL",
				},
			},
			want:    ptr.To("node-0"),
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostNewNodeWithResponse: func(ctx context.Context, body api.V0045OpenapiCreateNodeReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
							res := &api.SlurmV0045PostNewNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiResp{
									Errors: &[]api.V0045OpenapiError{
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
				req: api.V0045OpenapiCreateNodeReq{
					NodeConf: "NodeName=node-0 CPUs=4 State=EXTERNAL",
				},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostNewNodeWithResponse: func(ctx context.Context, body api.V0045OpenapiCreateNodeReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNewNodeResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0045OpenapiCreateNodeReq{
					NodeConf: "NodeName=node-0 CPUs=4 State=EXTERNAL",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid request type",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
				req: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.CreateNewNode(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateNewNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("SlurmClient.CreateNewNode() returned nil, want %v", *tt.want)
				} else if *got != *tt.want {
					t.Errorf("SlurmClient.CreateNewNode() = %v, want %v", *got, *tt.want)
				}
			}
		})
	}
}

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
						SlurmV0045DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
							res := &api.SlurmV0045DeleteNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiResp{
									Errors: &[]api.V0045OpenapiError{
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
						SlurmV0045DeleteNodeWithResponse: func(ctx context.Context, nodeName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteNodeResponse, error) {
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
				req:      api.V0045UpdateNodeMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
							res := &api.SlurmV0045PostNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiResp{
									Errors: &[]api.V0045OpenapiError{
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
				req:      api.V0045UpdateNodeMsg{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostNodeWithResponse: func(ctx context.Context, nodeName string, body api.V0045UpdateNodeMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostNodeResponse, error) {
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
		want    *types.V0045Node
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
						SlurmV0045GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
							res := &api.SlurmV0045GetNodeResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0045OpenapiNodesResp{
									Nodes: []api.V0045Node{
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
			want: &types.V0045Node{
				V0045Node: api.V0045Node{
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
						SlurmV0045GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
							res := &api.SlurmV0045GetNodeResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiNodesResp{
									Errors: &[]api.V0045OpenapiError{
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
						SlurmV0045GetNodeWithResponse: func(ctx context.Context, nodeName string, params *api.SlurmV0045GetNodeParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodeResponse, error) {
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
		want    *types.V0045NodeList
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
			want: &types.V0045NodeList{
				Items: make([]types.V0045Node, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
							res := &api.SlurmV0045GetNodesResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0045OpenapiNodesResp{
									Nodes: []api.V0045Node{
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
			want: &types.V0045NodeList{
				Items: []types.V0045Node{
					{V0045Node: api.V0045Node{Name: ptr.To("node-0")}},
					{V0045Node: api.V0045Node{Name: ptr.To("node-1")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
							res := &api.SlurmV0045GetNodesResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiNodesResp{
									Errors: &[]api.V0045OpenapiError{
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
						SlurmV0045GetNodesWithResponse: func(ctx context.Context, params *api.SlurmV0045GetNodesParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetNodesResponse, error) {
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
