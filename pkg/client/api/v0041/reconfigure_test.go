// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0041/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0041/interceptor"
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
		want    *types.V0041Reconfigure
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
							res := &api.SlurmV0041GetReconfigureResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0041OpenapiResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    &types.V0041Reconfigure{},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
							res := &api.SlurmV0041GetReconfigureResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0041OpenapiResp{
									Errors: &[]api.V0041OpenapiError{
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
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
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
		want    *types.V0041ReconfigureList
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
							res := &api.SlurmV0041GetReconfigureResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0041OpenapiResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0041ReconfigureList{
				Items: []types.V0041Reconfigure{
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
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
							res := &api.SlurmV0041GetReconfigureResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSON200: &api.V0041OpenapiResp{
									Errors: &[]api.V0041OpenapiError{
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
						SlurmV0041GetReconfigureWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReconfigureResponse, error) {
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
