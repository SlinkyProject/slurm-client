// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type createTransport func(*http.Request) (*http.Response, error)

func (f createTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCreateSetKeyError(t *testing.T) {
	for _, skip := range []bool{false, true} {
		t.Run(fmt.Sprintf("skip=%v", skip), func(t *testing.T) {
			obj := &apiv0045.V0045JobInfo{JobId: ptr.To(int32(42))}
			creates := 0
			cl := &client{resources: map[object.ObjectType]resource{
				obj.GetType(): {
					create: func(context.Context, any) (object.ObjectKey, error) {
						creates++
						return "invalid-key", nil
					},
					get: func(context.Context, object.ObjectKey, object.Object) error {
						t.Fatal("read back after SetKey failed")
						return nil
					},
				},
			}}
			err := cl.Create(t.Context(), obj, nil, &CreateOptions{SkipReadAfterCreate: skip})
			require.ErrorIs(t, err, strconv.ErrSyntax)
			require.ErrorContains(t, err, `set key for created object "invalid-key"`)
			require.Equal(t, 1, creates)
			require.Equal(t, object.ObjectKey("42"), obj.GetKey())
		})
	}
}

func TestCreateKeepsAcceptedIdentity(t *testing.T) {
	cases := []struct {
		obj object.Object
		req any
		key object.ObjectKey
	}{
		{
			obj: &apiv0042.V0042JobInfo{},
			req: apiv0042.V0042JobSubmitReq{},
			key: "42",
		},
		{
			obj: &apiv0043.V0043JobInfo{},
			req: apiv0043.V0043JobSubmitReq{},
			key: "42",
		},
		{
			obj: &apiv0044.V0044JobInfo{},
			req: apiv0044.V0044JobSubmitReq{},
			key: "42",
		},
		{
			obj: &apiv0045.V0045JobInfo{},
			req: apiv0045.V0045JobSubmitReq{},
			key: "42",
		},
		{
			obj: &apiv0044.V0044Node{},
			req: apiv0044.V0044OpenapiCreateNodeReq{NodeConf: "NodeName=node-1"},
			key: "node-1",
		},
		{
			obj: &apiv0045.V0045Node{},
			req: apiv0045.V0045OpenapiCreateNodeReq{NodeConf: "NodeName=node-1"},
			key: "node-1",
		},
		{
			obj: &apiv0044.V0044ReservationInfo{},
			req: apiv0044.V0044ReservationDescMsg{Name: ptr.To("reservation-1")},
			key: "reservation-1",
		},
		{
			obj: &apiv0045.V0045ReservationInfo{},
			req: apiv0045.V0045ReservationDescMsg{Name: ptr.To("reservation-1")},
			key: "reservation-1",
		},
	}
	for _, tc := range cases {
		for _, skip := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/skip=%v", tc.obj.GetType(), skip), func(t *testing.T) {
				posts, gets := 0, 0
				transport := createTransport(func(r *http.Request) (*http.Response, error) {
					code, body := http.StatusOK, `{"job_id":42}`
					if r.Method == http.MethodPost {
						posts++
					} else {
						gets++
						code, body = http.StatusServiceUnavailable, `{}`
					}
					return &http.Response{
						StatusCode: code,
						Header:     http.Header{"Content-Type": {"application/json"}},
						Body:       io.NopCloser(strings.NewReader(body)),
						Request:    r,
					}, nil
				})
				cl, err := NewClient(&Config{
					Server:        "http://slurm",
					TokenProvider: token.StaticProvider("test"),
					HTTPClient:    &http.Client{Transport: transport},
				})
				require.NoError(t, err)
				obj := tc.obj.DeepCopyObject().(object.Object)
				err = cl.Create(t.Context(), obj, tc.req, &CreateOptions{SkipReadAfterCreate: skip})
				if skip {
					require.NoError(t, err)
					require.Zero(t, gets)
				} else {
					require.Error(t, err)
					require.Equal(t, 1, gets)
				}
				require.Equal(t, tc.key, obj.GetKey(), "accepted identity")
				require.Equal(t, 1, posts)
			})
		}
	}
}

func TestCreateInvalidJobID(t *testing.T) {
	versions := []struct {
		obj object.Object
		req any
	}{
		{
			obj: &apiv0042.V0042JobInfo{JobId: ptr.To(int32(42))},
			req: apiv0042.V0042JobSubmitReq{},
		},
		{
			obj: &apiv0043.V0043JobInfo{JobId: ptr.To(int32(42))},
			req: apiv0043.V0043JobSubmitReq{},
		},
		{
			obj: &apiv0044.V0044JobInfo{JobId: ptr.To(int32(42))},
			req: apiv0044.V0044JobSubmitReq{},
		},
		{
			obj: &apiv0045.V0045JobInfo{JobId: ptr.To(int32(42))},
			req: apiv0045.V0045JobSubmitReq{},
		},
	}
	responses := []struct {
		name string
		body string
	}{
		{name: "missing", body: `{}`},
		{name: "null", body: `{"job_id":null}`},
		{name: "zero", body: `{"job_id":0}`},
		{name: "negative", body: `{"job_id":-1}`},
	}
	for _, version := range versions {
		for _, response := range responses {
			for _, skip := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/skip=%v", version.obj.GetType(), response.name, skip), func(t *testing.T) {
					posts, gets := 0, 0
					transport := createTransport(func(r *http.Request) (*http.Response, error) {
						if r.Method == http.MethodPost {
							posts++
						} else {
							gets++
						}
						return &http.Response{
							StatusCode: http.StatusOK,
							Header:     http.Header{"Content-Type": {"application/json"}},
							Body:       io.NopCloser(strings.NewReader(response.body)),
							Request:    r,
						}, nil
					})
					cl, err := NewClient(&Config{
						Server:        "http://slurm",
						TokenProvider: token.StaticProvider("test"),
						HTTPClient:    &http.Client{Transport: transport},
					})
					require.NoError(t, err)
					obj := version.obj.DeepCopyObject().(object.Object)
					err = cl.Create(t.Context(), obj, version.req, &CreateOptions{SkipReadAfterCreate: skip})
					wantErr := "slurm acknowledged submission without a valid key"
					require.EqualError(t, err, wantErr)
					require.Equal(t, 1, posts)
					require.Zero(t, gets)
					require.Equal(t, version.obj.GetKey(), obj.GetKey(), "invalid response changed the destination identity")
				})
			}
		}
	}
}
