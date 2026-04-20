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

func TestSlurmClient_GetQos(t *testing.T) {
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
		want    *types.V0044Qos
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "normal",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
							res := &api.SlurmdbV0044GetSingleQosResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiSlurmdbdQosResp{
									Qos: []api.V0044Qos{
										{Name: ptr.To("normal")},
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
				name: "normal",
			},
			want: &types.V0044Qos{
				V0044Qos: api.V0044Qos{
					Name: ptr.To("normal"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
							res := &api.SlurmdbV0044GetSingleQosResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiSlurmdbdQosResp{
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
				ctx:  context.Background(),
				name: "normal",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetSingleQosWithResponse: func(ctx context.Context, qos string, params *api.SlurmdbV0044GetSingleQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetSingleQosResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "normal",
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
			got, err := c.GetQos(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetQos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetQos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListQos(t *testing.T) {
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
		want    *types.V0044QosList
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
			want: &types.V0044QosList{
				Items: make([]types.V0044Qos, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
							res := &api.SlurmdbV0044GetQosResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiSlurmdbdQosResp{
									Qos: []api.V0044Qos{
										{Name: ptr.To("normal")},
										{Name: ptr.To("short")},
										{Name: ptr.To("free")},
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
			want: &types.V0044QosList{
				Items: []types.V0044Qos{
					{V0044Qos: api.V0044Qos{Name: ptr.To("normal")}},
					{V0044Qos: api.V0044Qos{Name: ptr.To("short")}},
					{V0044Qos: api.V0044Qos{Name: ptr.To("free")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmdbV0044GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
							res := &api.SlurmdbV0044GetQosResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiSlurmdbdQosResp{
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
						SlurmdbV0044GetQosWithResponse: func(ctx context.Context, params *api.SlurmdbV0044GetQosParams, reqEditors ...api.RequestEditorFn) (*api.SlurmdbV0044GetQosResponse, error) {
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
			got, err := c.ListQos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListQos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListQos() = %v, want %v", got, tt.want)
			}
		})
	}
}
