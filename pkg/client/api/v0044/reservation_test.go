// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0044/interceptor"
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
		want    *types.V0044ReservationInfo
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
						SlurmV0044GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
							res := &api.SlurmV0044GetReservationResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiReservationResp{
									Reservations: []api.V0044ReservationInfo{
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
			want: &types.V0044ReservationInfo{
				V0044ReservationInfo: api.V0044ReservationInfo{
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
						SlurmV0044GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
							res := &api.SlurmV0044GetReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiReservationResp{
									Errors: &[]api.V0044OpenapiError{
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
						SlurmV0044GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0044GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationResponse, error) {
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
		want    *types.V0044ReservationInfoList
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
			want: &types.V0044ReservationInfoList{
				Items: make([]types.V0044ReservationInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
							res := &api.SlurmV0044GetReservationsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0044OpenapiReservationResp{
									Reservations: []api.V0044ReservationInfo{
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
			want: &types.V0044ReservationInfoList{
				Items: []types.V0044ReservationInfo{
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("reservation-0")}},
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("reservation-1")}},
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("reservation-2")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
							res := &api.SlurmV0044GetReservationsResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiReservationResp{
									Errors: &[]api.V0044OpenapiError{
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
						SlurmV0044GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0044GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044GetReservationsResponse, error) {
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
			req: api.V0044ReservationDescMsg{
				Name:      ptr.To("slurm"),
				StartTime: ptr.To(api.V0044Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
				Duration:  ptr.To(api.V0044Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
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
			req: api.V0044Account{
				Name: "slurm",
			},
			wantErr: true,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
							res := &api.SlurmV0044PostReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiReservationModResp{
									Errors: &[]api.V0044OpenapiError{
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
			req: api.V0044ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0044ReservationDescMsg{
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
		reservation api.V0044ReservationDescMsg
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
						SlurmV0044DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
							res := &api.SlurmV0044DeleteReservationResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSONDefault: &api.V0044OpenapiResp{
									Errors: &[]api.V0044OpenapiError{
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
				reservation: api.V0044ReservationDescMsg{
					Name:      ptr.To("slurm"),
					StartTime: ptr.To(api.V0044Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
					Duration:  ptr.To(api.V0044Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
							res := &api.SlurmV0044DeleteReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiResp{
									Errors: &[]api.V0044OpenapiError{
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
						SlurmV0044DeleteReservationWithResponse: func(ctx context.Context, reservationName string, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044DeleteReservationResponse, error) {
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
			var emptyReservation api.V0044ReservationDescMsg
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
			req: api.V0044ReservationDescMsg{
				Name:      ptr.To("slurm"),
				StartTime: ptr.To(api.V0044Uint64NoValStruct{Number: ptr.To(time.Now().Unix())}),
				Duration:  ptr.To(api.V0044Uint32NoValStruct{Number: ptr.To(int32(oneHour.Seconds()))}),
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
			req: api.V0044Account{
				Name: "slurm",
			},
			wantErr: true,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
							res := &api.SlurmV0044PostReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0044OpenapiReservationModResp{
									Errors: &[]api.V0044OpenapiError{
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
			req: api.V0044ReservationDescMsg{
				Name: ptr.To("slurm"),
			},
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0044PostReservationWithResponse: func(ctx context.Context, body api.SlurmV0044PostReservationJSONRequestBody, reqEditors ...api.RequestEditorFn) (*api.SlurmV0044PostReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			req: api.V0044ReservationDescMsg{
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
