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

func TestSlurmClient_CreateJobInfo(t *testing.T) {
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
		want    *int32
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
							res := &api.SlurmV0040PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiJobSubmitResponse{
									JobId:  ptr.To[int32](1),
									StepId: nil,
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0040JobSubmitReq{},
			},
			want:    ptr.To[int32](1),
			wantErr: false,
		},
		{
			name: "Bad request",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
							res := &api.SlurmV0040PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiJobSubmitResponse{
									JobId:  ptr.To[int32](1),
									StepId: nil,
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
							res := &api.SlurmV0040PostJobSubmitResponse{
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
				req: api.V0040JobSubmitReq{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040PostJobSubmitWithResponse: func(ctx context.Context, body api.V0040JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040PostJobSubmitResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0040JobSubmitReq{},
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
			got, err := c.CreateJobInfo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateJobInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.CreateJobInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_DeleteJobInfo(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
							res := &api.SlurmV0040DeleteJobResponse{
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
				ctx:   context.Background(),
				jobId: "1",
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040DeleteJobResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			if err := c.DeleteJobInfo(tt.args.ctx, tt.args.jobId); (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.DeleteJobInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlurmClient_GetJobInfo(t *testing.T) {
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
		want    *types.V0040JobInfo
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
							res := &api.SlurmV0040GetJobResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiJobInfoResp{
									Jobs: []api.V0040JobInfo{
										{JobId: ptr.To[int32](1)},
									},
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
			},
			want: &types.V0040JobInfo{
				V0040JobInfo: api.V0040JobInfo{
					JobId: ptr.To[int32](1),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
							res := &api.SlurmV0040GetJobResponse{
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
				ctx:   context.Background(),
				jobId: "1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0040GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
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
			got, err := c.GetJobInfo(tt.args.ctx, tt.args.jobId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetJobInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetJobInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListJobInfo(t *testing.T) {
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
		want    *types.V0040JobInfoList
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
			want: &types.V0040JobInfoList{
				Items: make([]types.V0040JobInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
							res := &api.SlurmV0040GetJobsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0040OpenapiJobInfoResp{
									Jobs: []api.V0040JobInfo{
										{JobId: ptr.To[int32](1)},
										{JobId: ptr.To[int32](2)},
										{JobId: ptr.To[int32](3)},
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
			want: &types.V0040JobInfoList{
				Items: []types.V0040JobInfo{
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)}},
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](2)}},
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](3)}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
							res := &api.SlurmV0040GetJobsResponse{
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
						SlurmV0040GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0040GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0040GetJobsResponse, error) {
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
			got, err := c.ListJobInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListJobInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListJobInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
