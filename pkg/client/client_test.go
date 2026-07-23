// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type args struct {
		config *Config
		opts   []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "with token",
			args: args{
				config: &Config{
					Server: "http://bar",
				},
			},
			wantErr: true,
		},
		{
			name: "with server",
			args: args{
				config: &Config{
					AuthToken: "foo",
				},
			},
			wantErr: true,
		},
		{
			name: "valid token",
			args: args{
				config: &Config{
					AuthToken: "foo",
					Server:    "http://bar",
				},
			},
		},
		{
			name: "valid provider",
			args: args{
				config: &Config{
					Server:        "http://bar",
					TokenProvider: StaticTokenProvider("foo"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.config, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetTokenReturnsCachedProviderToken(t *testing.T) {
	const providerToken = "provider-token"

	var providerCalls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	slurmClient, err := NewClient(&Config{
		Server:    server.URL,
		AuthToken: "ignored-static-token",
		TokenProvider: TokenProviderFunc(func(context.Context) (string, error) {
			providerCalls++
			return providerToken, nil
		}),
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if got := slurmClient.GetToken(); got != "" {
		t.Fatalf("GetToken() before a request = %q, want an empty token", got)
	}
	if providerCalls != 0 {
		t.Fatalf("provider calls after GetToken() = %d, want 0", providerCalls)
	}

	internalClient := slurmClient.(*client)
	if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err != nil {
		t.Fatalf("v0045 ping error = %v", err)
	}
	if providerCalls != 1 {
		t.Fatalf("provider calls after a request = %d, want 1", providerCalls)
	}
	if got := slurmClient.GetToken(); got != providerToken {
		t.Fatalf("GetToken() after a request = %q, want %q", got, providerToken)
	}
	if providerCalls != 1 {
		t.Fatalf("provider calls after the second GetToken() = %d, want 1", providerCalls)
	}
}

func TestSetTokenUpdatesSubsequentRequests(t *testing.T) {
	const (
		initialToken = "initial-token"
		rotatedToken = "rotated-token"
		tokenHeader  = "X-SLURM-USER-TOKEN" //nolint:gosec // disable G101
	)

	var (
		mu     sync.Mutex
		tokens []string
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mu.Lock()
		tokens = append(tokens, req.Header.Get(tokenHeader))
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	slurmClient, err := NewClient(&Config{
		Server:    server.URL,
		AuthToken: initialToken,
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	internalClient := slurmClient.(*client)

	v0042Client := internalClient.v0042Client
	v0043Client := internalClient.v0043Client
	v0044Client := internalClient.v0044Client
	v0045Client := internalClient.v0045Client

	makeRequests := func() {
		t.Helper()
		if _, err := internalClient.v0042Client.SlurmV0042GetPingWithResponse(context.Background()); err != nil {
			t.Errorf("v0042 ping error = %v", err)
		}
		if _, err := internalClient.v0043Client.SlurmV0043GetPingWithResponse(context.Background()); err != nil {
			t.Errorf("v0043 ping error = %v", err)
		}
		if _, err := internalClient.v0044Client.SlurmV0044GetPingWithResponse(context.Background()); err != nil {
			t.Errorf("v0044 ping error = %v", err)
		}
		if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err != nil {
			t.Errorf("v0045 ping error = %v", err)
		}
	}

	makeRequests()
	slurmClient.SetToken(rotatedToken)
	makeRequests()

	if got := slurmClient.GetToken(); got != rotatedToken {
		t.Fatalf("GetToken() = %q, want %q", got, rotatedToken)
	}
	if internalClient.v0042Client != v0042Client ||
		internalClient.v0043Client != v0043Client ||
		internalClient.v0044Client != v0044Client ||
		internalClient.v0045Client != v0045Client {
		t.Fatal("SetToken() replaced a versioned API client")
	}

	mu.Lock()
	defer mu.Unlock()
	wantTokens := []string{
		initialToken, initialToken, initialToken, initialToken,
		rotatedToken, rotatedToken, rotatedToken, rotatedToken,
	}
	if len(tokens) != len(wantTokens) {
		t.Fatalf("received %d tokens, want %d: %q", len(tokens), len(wantTokens), tokens)
	}
	for i := range wantTokens {
		if tokens[i] != wantTokens[i] {
			t.Errorf("token[%d] = %q, want %q", i, tokens[i], wantTokens[i])
		}
	}
}

func TestSetTokenWinsOverInFlightProviderResolution(t *testing.T) {
	const (
		providerToken = "provider-token"
		rotatedToken  = "rotated-token"
	)

	started := make(chan struct{})
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	slurmClient, err := NewClient(&Config{
		Server: server.URL,
		TokenProvider: TokenProviderFunc(func(ctx context.Context) (string, error) {
			close(started)
			select {
			case <-release:
				return providerToken, nil
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}),
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	internalClient := slurmClient.(*client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	requestDone := make(chan error, 1)
	go func() {
		_, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(ctx)
		requestDone <- err
	}()

	select {
	case <-started:
	case <-ctx.Done():
		t.Fatalf("provider was not called: %v", ctx.Err())
	}

	slurmClient.SetToken(rotatedToken)
	close(release)
	if err := <-requestDone; err != nil {
		t.Fatalf("v0045 ping error = %v", err)
	}
	if got := slurmClient.GetToken(); got != rotatedToken {
		t.Fatalf("GetToken() = %q, want %q", got, rotatedToken)
	}
}

func TestSetTokenConcurrentWithRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	slurmClient, err := NewClient(&Config{
		Server:    server.URL,
		AuthToken: "initial-token",
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	internalClient := slurmClient.(*client)

	const (
		readers    = 8
		iterations = 50
	)
	start := make(chan struct{})
	errCh := make(chan error, readers)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-start
		for i := 0; i < readers*iterations; i++ {
			slurmClient.SetToken("rotated-token")
		}
	}()

	for range readers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for range iterations {
				if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err != nil {
					errCh <- err
					return
				}
			}
		}()
	}

	close(start)
	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Errorf("v0045 ping error = %v", err)
	}
}

func TestFileTokenProviderUpdatesSubsequentRequests(t *testing.T) {
	const tokenHeader = "X-SLURM-USER-TOKEN" //nolint:gosec // disable G101

	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("initial-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	tokens := make(chan string, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tokens <- req.Header.Get(tokenHeader)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	slurmClient, err := NewClient(&Config{
		Server:        server.URL,
		TokenProvider: FileTokenProvider{Path: tokenPath},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	internalClient := slurmClient.(*client)

	if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err != nil {
		t.Fatalf("initial v0045 ping error = %v", err)
	}
	if got := <-tokens; got != "initial-token" {
		t.Fatalf("initial token = %q, want %q", got, "initial-token")
	}

	if err := os.WriteFile(tokenPath, []byte("rotated-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err != nil {
		t.Fatalf("rotated v0045 ping error = %v", err)
	}
	if got := <-tokens; got != "rotated-token" {
		t.Fatalf("rotated token = %q, want %q", got, "rotated-token")
	}

	if err := os.WriteFile(tokenPath, nil, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if _, err := internalClient.v0045Client.SlurmV0045GetPingWithResponse(context.Background()); err == nil {
		t.Fatal("v0045 ping error = nil for an empty token file")
	}
}
