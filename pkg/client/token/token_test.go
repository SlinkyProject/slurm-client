// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package token

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestFileProvider(t *testing.T) {
	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("initial-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	provider := FileProvider{Path: tokenPath}

	value, err := provider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if value != "initial-token" {
		t.Fatalf("Token() = %q, want %q", value, "initial-token")
	}

	if err := os.WriteFile(tokenPath, []byte("SLURM_JWT=rotated-token\n"), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	value, err = provider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if value != "rotated-token" {
		t.Fatalf("Token() = %q, want %q", value, "rotated-token")
	}
}

func TestFileProviderErrors(t *testing.T) {
	t.Run("empty path", func(t *testing.T) {
		if _, err := (FileProvider{}).Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		provider := FileProvider{Path: filepath.Join(t.TempDir(), "missing")}
		if _, err := provider.Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("empty file", func(t *testing.T) {
		tokenPath := filepath.Join(t.TempDir(), "token")
		if err := os.WriteFile(tokenPath, nil, 0o600); err != nil {
			t.Fatalf("WriteFile() error = %v", err)
		}
		if _, err := (FileProvider{Path: tokenPath}).Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, err := (FileProvider{Path: "token"}).Token(ctx); !errors.Is(err, context.Canceled) {
			t.Fatalf("Token() error = %v, want %v", err, context.Canceled)
		}
	})
}
