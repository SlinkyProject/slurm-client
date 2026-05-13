// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0045/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_GetReservationInfo(t *testing.T) {
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
		want    *types.V0045ReservationInfo
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
							res := &api.SlurmV0045GetReservationResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0045OpenapiReservationResp{
									Reservations: []api.V0045ReservationInfo{
										{Name: ptr.To("reservation-0")},
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
				name: "reservation-0",
			},
			want: &types.V0045ReservationInfo{
				V0045ReservationInfo: api.V0045ReservationInfo{
					Name: ptr.To("reservation-0"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
							res := &api.SlurmV0045GetReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiReservationResp{
									Errors: &[]api.V0045OpenapiError{
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
				name: "reservation-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0045GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
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
			got, err := c.GetReservationInfo(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetReservationInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListReservationInfo(t *testing.T) {
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
		want    *types.V0045ReservationInfoList
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
			want: &types.V0045ReservationInfoList{
				Items: make([]types.V0045ReservationInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
							res := &api.SlurmV0045GetReservationsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0045OpenapiReservationResp{
									Reservations: []api.V0045ReservationInfo{
										{Name: ptr.To("reservation-0")},
										{Name: ptr.To("reservation-1")},
										{Name: ptr.To("reservation-2")},
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
			want: &types.V0045ReservationInfoList{
				Items: []types.V0045ReservationInfo{
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("reservation-0")}},
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("reservation-1")}},
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("reservation-2")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
							res := &api.SlurmV0045GetReservationsResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiReservationResp{
									Errors: &[]api.V0045OpenapiError{
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
						SlurmV0045GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0045GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045GetReservationsResponse, error) {
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
			got, err := c.ListReservationInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListReservationInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_CreateReservationInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx context.Context
	}
	oneHour, err := time.ParseDuration("1h")
	if err != nil {
		t.Fatalf("failed to parse duration value when constructing test. err: %v", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		req     any
		want    string
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0045ReservationDescMsg{
				Name:      ptr.To("slurm"),
				StartTime: ptr.To(api.V0045Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
				Duration:  ptr.To(api.V0045Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
			},
			wantErr: false,
			want:    "slurm",
		},
		{
			name: "invalid type provided",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0045Account{
				Name: "slurm",
			},
			wantErr: true,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
							res := &api.SlurmV0045PostReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiReservationModResp{
									Errors: &[]api.V0045OpenapiError{
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
			req: api.V0045ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0045ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.CreateReservationInfo(tt.args.ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.CreateReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.CreateReservationInfo() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestSlurmClient_DeleteReservationInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx         context.Context
		name        string
		reservation api.V0045ReservationDescMsg
	}

	oneHour, err := time.ParseDuration("1h")
	if err != nil {
		t.Fatalf("failed to parse duration value when constructing test. err: %v", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "reservation does not exist",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "slurm",
			},
			wantErr: false,
		},
		{
			name: "reservation exists",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
							res := &api.SlurmV0045DeleteReservationResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSONDefault: &api.V0045OpenapiResp{
									Errors: &[]api.V0045OpenapiError{
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
				name: "slurm",
				reservation: api.V0045ReservationDescMsg{
					Name:      ptr.To("slurm"),
					StartTime: ptr.To(api.V0045Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
					Duration:  ptr.To(api.V0045Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
							res := &api.SlurmV0045DeleteReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiResp{
									Errors: &[]api.V0045OpenapiError{
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
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045DeleteReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			var emptyReservation api.V0045ReservationDescMsg
			if tt.args.reservation != emptyReservation {
				got, err := c.CreateReservationInfo(tt.args.ctx, tt.args.reservation)
				if err != nil {
					t.Errorf("Failed to Create reservation when configuring test: %v", err)
				}
				if got != *tt.args.reservation.Name {
					t.Errorf("Failed to Create reservation when configuring test: got: %v, want: %v", got, tt.args.reservation.Name)
				}
			}
			err := c.DeleteReservationInfo(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.DeleteReservationInfo() \"%v\" error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}

func TestSlurmClient_UpdateReservationInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx  context.Context
		name string
	}
	oneHour, err := time.ParseDuration("1h")
	if err != nil {
		t.Fatalf("failed to parse duration value when constructing test. err: %v", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		req     any
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "slurm",
			},
			req: api.V0045ReservationDescMsg{
				Name:      ptr.To("slurm"),
				StartTime: ptr.To(api.V0045Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
				Duration:  ptr.To(api.V0045Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
			},
			wantErr: false,
		},
		{
			name: "invalid type provided",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0045Account{
				Name: "slurm",
			},
			wantErr: true,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
							res := &api.SlurmV0045PostReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0045OpenapiReservationModResp{
									Errors: &[]api.V0045OpenapiError{
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
			req: api.V0045ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0045PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0045PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0045PostReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0045ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			err := c.UpdateReservationInfo(tt.args.ctx, tt.args.name, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.UpdateReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
