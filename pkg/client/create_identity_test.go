// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"k8s.io/utils/ptr"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type createTransport func(*http.Request) (*http.Response, error)

func (f createTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

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
				if err != nil {
					t.Fatal(err)
				}
				obj := tc.obj.DeepCopyObject().(object.Object)
				err = cl.Create(t.Context(), obj, tc.req, &CreateOptions{SkipReadAfterCreate: skip})
				if skip && err != nil || !skip && err == nil {
					t.Fatalf("skip=%v: %v", skip, err)
				}
				if key := obj.GetKey(); key != tc.key {
					t.Fatalf("accepted identity = %q, want %q", key, tc.key)
				}
				if posts != 1 || skip && gets != 0 || !skip && gets != 1 {
					t.Fatalf("POST=%d GET=%d", posts, gets)
				}
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
					if err != nil {
						t.Fatal(err)
					}
					obj := version.obj.DeepCopyObject().(object.Object)
					err = cl.Create(t.Context(), obj, version.req, &CreateOptions{SkipReadAfterCreate: skip})
					if !errors.Is(err, apierrors.ErrInvalidJobID) {
						t.Fatalf("Create() error = %v, want ErrInvalidJobID", err)
					}
					if posts != 1 || gets != 0 {
						t.Fatalf("POST=%d GET=%d, want POST=1 GET=0", posts, gets)
					}
					if key := obj.GetKey(); key != version.obj.GetKey() {
						t.Fatalf("invalid response changed the destination identity to %q", key)
					}
				})
			}
		}
	}
}
