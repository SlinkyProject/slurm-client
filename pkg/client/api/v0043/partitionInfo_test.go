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

func TestSlurmClient_GetPartitionInfo(t *testing.T) {
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
		want    *types.V0043PartitionInfo
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "partition-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
							res := &api.SlurmV0043GetPartitionResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiPartitionResp{
									Partitions: []api.V0043PartitionInfo{
										{Name: ptr.To("partition-0")},
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
				name: "partition-0",
			},
			want: &types.V0043PartitionInfo{
				V0043PartitionInfo: api.V0043PartitionInfo{
					Name: ptr.To("partition-0"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
							res := &api.SlurmV0043GetPartitionResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiPartitionResp{
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
				name: "partition-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0043GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "partition-0",
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
			got, err := c.GetPartitionInfo(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetPartitionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetPartitionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListPartitionInfo(t *testing.T) {
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
		want    *types.V0043PartitionInfoList
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
			want: &types.V0043PartitionInfoList{
				Items: make([]types.V0043PartitionInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
							res := &api.SlurmV0043GetPartitionsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiPartitionResp{
									Partitions: []api.V0043PartitionInfo{
										{Name: ptr.To("partition-0")},
										{Name: ptr.To("partition-1")},
										{Name: ptr.To("partition-2")},
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
			want: &types.V0043PartitionInfoList{
				Items: []types.V0043PartitionInfo{
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("partition-0")}},
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("partition-1")}},
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("partition-2")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
							res := &api.SlurmV0043GetPartitionsResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiPartitionResp{
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
						SlurmV0043GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetPartitionsResponse, error) {
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
			got, err := c.ListPartitionInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListPartitionInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListPartitionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
