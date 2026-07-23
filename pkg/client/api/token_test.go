// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
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
		provider, err := NewTokenProvider("static-token", WithTokenProvider(TokenProviderFunc(func(context.Context) (string, error) {
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
		provider, err := NewTokenProvider("static-token", WithTokenProvider(TokenProviderFunc(func(context.Context) (string, error) {
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
		provider, err := NewTokenProvider("static-token", WithTokenProvider(TokenProviderFunc(nil)))
		if err != nil {
			t.Fatalf("NewTokenProvider() error = %v", err)
		}
		if _, err := provider.Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})
}

func TestFileTokenProvider(t *testing.T) {
	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("initial-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	provider := FileTokenProvider{Path: tokenPath}

	token, err := provider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if token != "initial-token" {
		t.Fatalf("Token() = %q, want %q", token, "initial-token")
	}

	if err := os.WriteFile(tokenPath, []byte("SLURM_JWT=rotated-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	token, err = provider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if token != "rotated-token" {
		t.Fatalf("Token() = %q, want %q", token, "rotated-token")
	}
}

func TestFileTokenProviderErrors(t *testing.T) {
	t.Run("empty path", func(t *testing.T) {
		if _, err := (FileTokenProvider{}).Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		provider := FileTokenProvider{Path: filepath.Join(t.TempDir(), "missing")}
		if _, err := provider.Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("empty file", func(t *testing.T) {
		tokenPath := filepath.Join(t.TempDir(), "token")
		if err := os.WriteFile(tokenPath, nil, 0o600); err != nil {
			t.Fatalf("WriteFile() error = %v", err)
		}
		if _, err := (FileTokenProvider{Path: tokenPath}).Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, err := (FileTokenProvider{Path: "token"}).Token(ctx); !errors.Is(err, context.Canceled) {
			t.Fatalf("Token() error = %v, want %v", err, context.Canceled)
		}
	})
}
