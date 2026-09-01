// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
)

func TestNewSlurmClientPostRequestHasSingleContentTypeHeader(t *testing.T) {
	const contentTypeHeader = "Content-Type"

	var (
		mu           sync.Mutex
		contentTypes []string
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mu.Lock()
		contentTypes = req.Header.Values(contentTypeHeader)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer server.Close()

	client, err := NewSlurmClient(server.URL, "test-token", server.Client())
	if err != nil {
		t.Fatalf("NewSlurmClient() error = %v", err)
	}

	_, err = client.SlurmV0045PostReservationWithResponse(context.Background(), api.V0045ReservationDescMsg{
		Name: ptr.To("test"),
	})
	if err != nil {
		t.Fatalf("SlurmV0045PostReservationWithResponse() error = %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(contentTypes) != 1 {
		t.Fatalf("Content-Type header count = %d, want 1: %v", len(contentTypes), contentTypes)
	}
	if contentTypes[0] != "application/json" {
		t.Fatalf("Content-Type = %q, want %q", contentTypes[0], "application/json")
	}
}
