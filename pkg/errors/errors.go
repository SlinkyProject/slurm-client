// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound       = errors.New(http.StatusText(http.StatusNotFound))
	ErrNotImplemented = errors.New(http.StatusText(http.StatusNotImplemented))
)

// NewHTTPError returns the canonical client error for an HTTP status code.
func NewHTTPError(statusCode int) error {
	switch statusCode {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusNotImplemented:
		return ErrNotImplemented
	default:
		return errors.New(http.StatusText(statusCode))
	}
}
