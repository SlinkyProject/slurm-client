// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"net/http"
)

var (
	// ErrNotFound indicates that a lookup returned no matching object.
	ErrNotFound = errors.New("not found")

	// ErrNotImplemented indicates an unsupported object type or operation in this client.
	ErrNotImplemented = errors.New("not implemented")
)

// NewHTTPError returns the client error for an HTTP status code.
func NewHTTPError(statusCode int) error {
	if statusCode == http.StatusNotFound {
		return ErrNotFound
	}
	return errors.New(http.StatusText(statusCode))
}
