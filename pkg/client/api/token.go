// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// TokenProvider returns the authentication token to add to a request. It is
// called for every request and must be safe for concurrent use.
type TokenProvider interface {
	Token(ctx context.Context) (string, error)
}

// TokenProviderFunc adapts a function to a TokenProvider.
type TokenProviderFunc func(ctx context.Context) (string, error)

// Token implements TokenProvider.
func (f TokenProviderFunc) Token(ctx context.Context) (string, error) {
	if f == nil {
		return "", fmt.Errorf("token provider function cannot be nil")
	}
	return f(ctx)
}

type staticTokenProvider string

// Token implements TokenProvider.
func (p staticTokenProvider) Token(context.Context) (string, error) {
	return string(p), nil
}

// StaticTokenProvider returns a provider that always supplies token.
func StaticTokenProvider(token string) TokenProvider {
	return staticTokenProvider(token)
}

// FileTokenProvider opens and reads Path for every request.
type FileTokenProvider struct {
	Path string
}

// Token implements TokenProvider.
func (p FileTokenProvider) Token(ctx context.Context) (string, error) {
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

type clientOptions struct {
	tokenProvider TokenProvider
}

// ClientOption configures a versioned Slurm API client.
type ClientOption interface {
	apply(*clientOptions)
}

type clientOptionFunc func(*clientOptions)

func (f clientOptionFunc) apply(options *clientOptions) {
	f(options)
}

// WithTokenProvider configures a client to resolve its authentication token
// for every request.
func WithTokenProvider(tokenProvider TokenProvider) ClientOption {
	return clientOptionFunc(func(options *clientOptions) {
		options.tokenProvider = tokenProvider
	})
}

// NewTokenProvider returns the configured request-time token provider. Without
// WithTokenProvider, it returns a provider for the fixed token.
func NewTokenProvider(token string, opts ...ClientOption) (TokenProvider, error) {
	options := &clientOptions{
		tokenProvider: StaticTokenProvider(token),
	}
	for _, opt := range opts {
		if opt != nil {
			opt.apply(options)
		}
	}
	if options.tokenProvider == nil {
		return nil, fmt.Errorf("token provider cannot be nil")
	}
	return options.tokenProvider, nil
}
