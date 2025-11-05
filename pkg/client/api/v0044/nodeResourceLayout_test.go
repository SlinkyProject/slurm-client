// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"reflect"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_GetNodeResourceLayout(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx   context.Context
		jobId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0044NodeResourceLayout
		wantErr bool
	}{
		{
			name: "Not found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
							res := &api.SlurmV0044GetResourcesResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200:      &api.V0044OpenapiResourceLayoutResp{},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:   ctx,
				jobId: "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nodes defined",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetResourcesWithResponse: func(ctx context.Context, jobId string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetResourcesResponse, error) {
							res := &api.SlurmV0044GetResourcesResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiResourceLayoutResp{
									Nodes: api.V0044NodeResourceLayoutList{
										{Node: "node1"},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:   ctx,
				jobId: "1",
			},
			want: &types.V0044NodeResourceLayout{
				V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{
					{Node: "node1"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.GetNodeResourceLayout(tt.args.ctx, tt.args.jobId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetNodeResourceLayout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetNodeResourceLayout() = %v, want %v", got, tt.want)
			}
		})
	}
}
