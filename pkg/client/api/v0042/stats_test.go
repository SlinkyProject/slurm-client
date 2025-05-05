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

func TestSlurmClient_GetStats(t *testing.T) {
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
		want    *types.V0042Stats
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
							res := &api.SlurmV0042GetDiagResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0042OpenapiDiagResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0042Stats{
				V0042StatsMsg: api.V0042StatsMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
							res := &api.SlurmV0042GetDiagResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiDiagResp{
									Errors: &[]api.V0042OpenapiError{
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
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
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
			got, err := c.GetStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListStats(t *testing.T) {
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
		want    *types.V0042StatsList
		wantErr bool
	}{
		{
			name: "Fetch",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
							res := &api.SlurmV0042GetDiagResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0042OpenapiDiagResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0042StatsList{
				Items: []types.V0042Stats{
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
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
							res := &api.SlurmV0042GetDiagResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSON200: &api.V0042OpenapiDiagResp{
									Errors: &[]api.V0042OpenapiError{
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
						SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
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
			got, err := c.ListStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
