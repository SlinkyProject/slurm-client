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

func TestSlurmClient_GetReconfigure(t *testing.T) {
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
		want    *types.V0043Reconfigure
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
							res := &api.SlurmV0043GetReconfigureResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0043OpenapiResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    &types.V0043Reconfigure{},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
							res := &api.SlurmV0043GetReconfigureResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiResp{
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
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
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
			got, err := c.GetReconfigure(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetReconfigure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetReconfigure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListReconfigure(t *testing.T) {
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
		want    *types.V0043ReconfigureList
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
							res := &api.SlurmV0043GetReconfigureResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0043OpenapiResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0043ReconfigureList{
				Items: []types.V0043Reconfigure{
					{},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
							res := &api.SlurmV0043GetReconfigureResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSON200: &api.V0043OpenapiResp{
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
						SlurmV0043GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetReconfigureResponse, error) {
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
			got, err := c.ListReconfigure(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListReconfigure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListReconfigure() = %v, want %v", got, tt.want)
			}
		})
	}
}
