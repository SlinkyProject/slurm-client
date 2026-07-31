// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

// Package token supplies Slurm authentication tokens to clients.
package token

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// Provider returns the authentication token to add to a request. It is called
// for every request and must be safe for concurrent use.
type Provider interface {
	Token(ctx context.Context) (string, error)
}

// ProviderFunc adapts a function to a Provider.
type ProviderFunc func(ctx context.Context) (string, error)

// Token implements Provider.
func (f ProviderFunc) Token(ctx context.Context) (string, error) {
	if f == nil {
		return "", fmt.Errorf("token provider function cannot be nil")
	}
	return f(ctx)
}

type staticProvider string

// Token implements Provider.
func (p staticProvider) Token(context.Context) (string, error) {
	return string(p), nil
}

// StaticProvider returns a provider that always supplies token.
func StaticProvider(token string) Provider {
	return staticProvider(token)
}

// FileProvider opens and reads Path for every request.
type FileProvider struct {
	Path string
}

// Token implements Provider.
func (p FileProvider) Token(ctx context.Context) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if p.Path == "" {
		return "", fmt.Errorf("token file path cannot be empty")
	}

	data, err := os.ReadFile(p.Path)
	if err != nil {
		return "", fmt.Errorf("unable to read token file %q: %w", p.Path, err)
	}

	token := strings.TrimSpace(string(data))
	token = strings.TrimSpace(strings.TrimPrefix(token, "SLURM_JWT="))
	if token == "" {
		return "", fmt.Errorf("token file %q is empty", p.Path)
	}
	return token, nil
}
