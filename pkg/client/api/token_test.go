// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"errors"
	"testing"

	tokenprovider "github.com/SlinkyProject/slurm-client/pkg/client/token"
)

func TestNewTokenProvider(t *testing.T) {
	t.Run("static", func(t *testing.T) {
		provider, err := NewTokenProvider("static-token")
		if err != nil {
			t.Fatalf("NewTokenProvider() error = %v", err)
		}
		got, err := provider.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error = %v", err)
		}
		if got != "static-token" {
			t.Fatalf("provider() = %q, want %q", got, "static-token")
		}
	})

	t.Run("configured", func(t *testing.T) {
		token := "initial-token"
		provider, err := NewTokenProvider("static-token", WithTokenProvider(tokenprovider.ProviderFunc(func(context.Context) (string, error) {
			return token, nil
		})))
		if err != nil {
			t.Fatalf("NewTokenProvider() error = %v", err)
		}

		got, err := provider.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error = %v", err)
		}
		if got != "initial-token" {
			t.Fatalf("provider() = %q, want %q", got, "initial-token")
		}
		token = "rotated-token"
		got, err = provider.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error = %v", err)
		}
		if got != "rotated-token" {
			t.Fatalf("provider() = %q, want %q", got, "rotated-token")
		}
	})

	t.Run("error", func(t *testing.T) {
		wantErr := context.Canceled
		provider, err := NewTokenProvider("static-token", WithTokenProvider(tokenprovider.ProviderFunc(func(context.Context) (string, error) {
			return "", wantErr
		})))
		if err != nil {
			t.Fatalf("NewTokenProvider() error = %v", err)
		}
		if _, err := provider.Token(context.Background()); !errors.Is(err, wantErr) {
			t.Fatalf("Token() error = %v, want %v", err, wantErr)
		}
	})

	t.Run("nil", func(t *testing.T) {
		_, err := NewTokenProvider("static-token", WithTokenProvider(nil))
		if err == nil {
			t.Fatal("NewTokenProvider() error = nil, want an error")
		}
	})

	t.Run("nil function", func(t *testing.T) {
		provider, err := NewTokenProvider("static-token", WithTokenProvider(tokenprovider.ProviderFunc(nil)))
		if err != nil {
			t.Fatalf("NewTokenProvider() error = %v", err)
		}
		if _, err := provider.Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})
}
