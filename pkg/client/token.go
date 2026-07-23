// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import clientapi "github.com/SlinkyProject/slurm-client/pkg/client/api"

// TokenProvider returns the authentication token to add to a request.
type TokenProvider = clientapi.TokenProvider

// TokenProviderFunc adapts a function to a TokenProvider.
type TokenProviderFunc = clientapi.TokenProviderFunc

// FileTokenProvider opens and reads Path for every request.
type FileTokenProvider = clientapi.FileTokenProvider

// StaticTokenProvider returns a provider that always supplies token.
func StaticTokenProvider(token string) TokenProvider {
	return clientapi.StaticTokenProvider(token)
}
