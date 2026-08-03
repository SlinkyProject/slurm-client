// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"

	tokenprovider "github.com/SlinkyProject/slurm-client/pkg/client/token"
)

type clientOptions struct {
	tokenProvider tokenprovider.Provider
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
func WithTokenProvider(tokenProvider tokenprovider.Provider) ClientOption {
	return clientOptionFunc(func(options *clientOptions) {
		options.tokenProvider = tokenProvider
	})
}

// NewTokenProvider returns the configured request-time token provider. Without
// WithTokenProvider, it returns a provider for the fixed token.
func NewTokenProvider(token string, opts ...ClientOption) (tokenprovider.Provider, error) {
	options := &clientOptions{
		tokenProvider: tokenprovider.StaticProvider(token),
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
