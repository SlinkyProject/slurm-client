// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/client/token"
)

type createTransport func(*http.Request) (*http.Response, error)

func (f createTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCreateKeepsAcceptedIdentity(t *testing.T) {
	for _, skip := range []bool{false, true} {
		t.Run(fmt.Sprint(skip), func(t *testing.T) {
			posts, gets := 0, 0
			cl, err := NewClient(&Config{Server: "http://slurm", TokenProvider: token.StaticProvider("test"), HTTPClient: &http.Client{Transport: createTransport(func(r *http.Request) (*http.Response, error) {
				code, body := 200, `{"job_id":42}`
				if r.Method == http.MethodPost {
					posts++
				} else {
					gets++
					code, body = 503, `{}`
				}
				return &http.Response{StatusCode: code, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
			})}})
			if err != nil {
				t.Fatal(err)
			}
			job := &api.V0044JobInfo{}
			err = cl.Create(t.Context(), job, api.V0044JobSubmitReq{}, &CreateOptions{SkipReadAfterCreate: skip})
			if skip && err != nil || !skip && err == nil {
				t.Fatalf("skip=%v: %v", skip, err)
			}
			if job.JobId == nil || *job.JobId != 42 {
				t.Fatalf("lost accepted ID: %#v", job.JobId)
			}
			if posts != 1 || skip && gets != 0 || !skip && gets != 1 {
				t.Fatalf("POST=%d GET=%d", posts, gets)
			}
		})
	}
}
