// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0043/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0043/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_GetControllerPing(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx  context.Context
		host string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0043ControllerPing
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				host: "controller-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
							res := &api.SlurmV0043GetPingResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiPingArrayResp{
									Pings: []api.V0043ControllerPing{
										{Hostname: ptr.To("controller-0")},
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
				host: "controller-0",
			},
			want: &types.V0043ControllerPing{
				V0043ControllerPing: api.V0043ControllerPing{
					Hostname: ptr.To("controller-0"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
							res := &api.SlurmV0043GetPingResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiPingArrayResp{
									Errors: &[]api.V0043OpenapiError{
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
				ctx:  context.Background(),
				host: "controller-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				host: "controller-0",
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
			got, err := c.GetControllerPing(tt.args.ctx, tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetControllerPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetControllerPing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListControllerPing(t *testing.T) {
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
		want    *types.V0043ControllerPingList
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
			want: &types.V0043ControllerPingList{
				Items: make([]types.V0043ControllerPing, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
							res := &api.SlurmV0043GetPingResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiPingArrayResp{
									Pings: []api.V0043ControllerPing{
										{Hostname: ptr.To("controller-0")},
										{Hostname: ptr.To("controller-1")},
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
			want: &types.V0043ControllerPingList{
				Items: []types.V0043ControllerPing{
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
							res := &api.SlurmV0043GetPingResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiPingArrayResp{
									Errors: &[]api.V0043OpenapiError{
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
						SlurmV0043GetPingWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPingResponse, error) {
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
			got, err := c.ListControllerPing(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListControllerPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListControllerPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
