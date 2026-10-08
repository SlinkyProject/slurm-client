// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
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
