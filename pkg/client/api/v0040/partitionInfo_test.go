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
		want    *types.V0040PartitionInfo
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
						SlurmV0040GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
							res := &api.SlurmV0040GetPartitionResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiPartitionResp{
									Partitions: []api.V0040PartitionInfo{
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
			want: &types.V0040PartitionInfo{
				V0040PartitionInfo: api.V0040PartitionInfo{
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
						SlurmV0040GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
							res := &api.SlurmV0040GetPartitionResponse{
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
						SlurmV0040GetPartitionWithResponse: func(ctx context.Context, partitionName string, params *api.SlurmV0040GetPartitionParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionResponse, error) {
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
		want    *types.V0040PartitionInfoList
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
			want: &types.V0040PartitionInfoList{
				Items: make([]types.V0040PartitionInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
							res := &api.SlurmV0040GetPartitionsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiPartitionResp{
									Partitions: []api.V0040PartitionInfo{
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
			want: &types.V0040PartitionInfoList{
				Items: []types.V0040PartitionInfo{
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("partition-0")}},
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("partition-1")}},
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("partition-2")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
							res := &api.SlurmV0040GetPartitionsResponse{
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
						SlurmV0040GetPartitionsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetPartitionsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetPartitionsResponse, error) {
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
