// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
)

type headerCapture struct {
	header http.Header
}

func (h *headerCapture) RoundTrip(req *http.Request) (*http.Response, error) {
	h.header = req.Header.Clone()
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader("{}")),
		Request:    req,
	}, nil
}

func TestNewSlurmClientPostRequestHasSingleContentTypeHeader(t *testing.T) {
	const contentTypeHeader = "Content-Type"

	capture := &headerCapture{}
	client, err := NewSlurmClient("http://slurm.example", "test-token", &http.Client{Transport: capture})
	if err != nil {
		t.Fatalf("NewSlurmClient() error = %v", err)
	}

	_, err = client.SlurmV0042PostNodeWithResponse(context.Background(), "test", api.V0042UpdateNodeMsg{})
	if err != nil {
		t.Fatalf("SlurmV0042PostNodeWithResponse() error = %v", err)
	}

	contentTypes := capture.header.Values(contentTypeHeader)
	if len(contentTypes) != 1 {
		t.Fatalf("Content-Type header count = %d, want 1: %v", len(contentTypes), contentTypes)
	}
	if contentTypes[0] != "application/json" {
		t.Fatalf("Content-Type = %q, want %q", contentTypes[0], "application/json")
	}
}
