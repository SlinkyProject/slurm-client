// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/interceptor"
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
						SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
							res := &api.SlurmV0042PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobSubmitResponse{
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
				req: api.V0042JobSubmitReq{},
			},
			want:    ptr.To[int32](1),
			wantErr: false,
		},
		{
			name: "Bad request",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
							res := &api.SlurmV0042PostJobSubmitResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobSubmitResponse{
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
						SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
							res := &api.SlurmV0042PostJobSubmitResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiJobSubmitResponse{
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
				req: api.V0042JobSubmitReq{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobSubmitWithResponse: func(ctx context.Context, body api.V0042JobSubmitReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobSubmitResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0042JobSubmitReq{},
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
func TestSlurmClient_CreateJobInfoAlloc(t *testing.T) {
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
						SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
							res := &api.SlurmV0042PostJobAllocateResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobAllocResp{
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
				req: api.V0042JobAllocReq{},
			},
			want:    ptr.To[int32](1),
			wantErr: false,
		},
		{
			name: "Bad request",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
							res := &api.SlurmV0042PostJobAllocateResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobAllocResp{
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
						SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
							res := &api.SlurmV0042PostJobAllocateResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiJobAllocResp{
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
				req: api.V0042JobAllocReq{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobAllocateWithResponse: func(ctx context.Context, body api.V0042JobAllocReq, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobAllocateResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
				req: api.V0042JobAllocReq{},
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
			got, err := c.CreateJobInfoAlloc(tt.args.ctx, tt.args.req)
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
						SlurmV0042DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
							res := &api.SlurmV0042DeleteJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiKillJobResp{
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
						SlurmV0042DeleteJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042DeleteJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042DeleteJobResponse, error) {
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
				req:   api.V0042JobDescMsg{},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
							res := &api.SlurmV0042PostJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiJobPostResponse{
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
				ctx:   context.Background(),
				jobId: "1",
				req:   api.V0042JobDescMsg{},
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042PostJobWithResponse: func(ctx context.Context, jobId string, body api.V0042JobDescMsg, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042PostJobResponse, error) {
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
		want    *types.V0042JobInfo
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
						SlurmV0042GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
							jid, err := strconv.Atoi(jobId)
							if err != nil {
								t.Fatal(err)
							}
							res := &api.SlurmV0042GetJobResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobInfoResp{
									Jobs: []api.V0042JobInfo{
										{JobId: ptr.To(int32(jid))},
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
			want: &types.V0042JobInfo{
				V0042JobInfo: api.V0042JobInfo{
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
						SlurmV0042GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
							res := &api.SlurmV0042GetJobResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0042OpenapiJobInfoResp{
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
						SlurmV0042GetJobWithResponse: func(ctx context.Context, jobId string, params *api.SlurmV0042GetJobParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobResponse, error) {
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
		want    *types.V0042JobInfoList
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
			want: &types.V0042JobInfoList{
				Items: make([]types.V0042JobInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
							res := &api.SlurmV0042GetJobsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0042OpenapiJobInfoResp{
									Jobs: []api.V0042JobInfo{
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
			want: &types.V0042JobInfoList{
				Items: []types.V0042JobInfo{
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)}},
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](2)}},
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](3)}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
							res := &api.SlurmV0042GetJobsResponse{
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
						SlurmV0042GetJobsWithResponse: func(ctx context.Context, params *api.SlurmV0042GetJobsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetJobsResponse, error) {
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
