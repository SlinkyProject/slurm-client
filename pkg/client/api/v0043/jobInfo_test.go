// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0043/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0043/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"k8s.io/utils/ptr"
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
						SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
							res := &api.SlurmV0043PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiJobSubmitResponse{
									JobId: ptr.To[int32](1),
								},
							}
							return res, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0043JobSubmitReq{},
			},
			want:    ptr.To[int32](1),
			wantErr: false,
		},
		{
			name: "Bad request",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
							res := &api.SlurmV0043PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiJobSubmitResponse{
									JobId: ptr.To[int32](1),
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
						SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
							res := &api.SlurmV0043PostJobSubmitResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiJobSubmitResponse{
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
				req: api.V0043JobSubmitReq{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043PostJobSubmitWithResponse: func(ctx context.Context, body api.V0043JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobSubmitResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0043JobSubmitReq{},
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
						SlurmV0043DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
							res := &api.SlurmV0043DeleteJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiKillJobResp{
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
						SlurmV0043DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043DeleteJobResponse, error) {
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

func TestSlurmClient_UpdateJobInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx   context.Context
		jobId string
		req   any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update existing job",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "1",
				req:   api.V0043JobDescMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
							res := &api.SlurmV0043PostJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiJobPostResponse{
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
				ctx:   context.Background(),
				jobId: "1",
				req:   api.V0043JobDescMsg{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0043JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043PostJobResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:   context.Background(),
				jobId: "0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			if err := c.UpdateJobInfo(tt.args.ctx, tt.args.jobId, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.UpdateJobInfo() error = %v, wantErr %v", err, tt.wantErr)
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
		want    *types.V0043JobInfo
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
						SlurmV0043GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
							jid, err := strconv.Atoi(jobId)
							if err != nil {
								t.Fatal(err)
							}
							res := &api.SlurmV0043GetJobResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiJobInfoResp{
									Jobs: []api.V0043JobInfo{
										{JobId: ptr.To(int32(jid))}, //nolint:gosec // disable G109
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
			want: &types.V0043JobInfo{
				V0043JobInfo: api.V0043JobInfo{
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
						SlurmV0043GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
							res := &api.SlurmV0043GetJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0043OpenapiJobInfoResp{
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
						SlurmV0043GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0043GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobResponse, error) {
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
		want    *types.V0043JobInfoList
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
			want: &types.V0043JobInfoList{
				Items: make([]types.V0043JobInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
							res := &api.SlurmV0043GetJobsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0043OpenapiJobInfoResp{
									Jobs: []api.V0043JobInfo{
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
			want: &types.V0043JobInfoList{
				Items: []types.V0043JobInfo{
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)}},
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](2)}},
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](3)}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
							res := &api.SlurmV0043GetJobsResponse{
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
						SlurmV0043GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0043GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0043GetJobsResponse, error) {
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
