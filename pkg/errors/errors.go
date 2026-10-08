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

	// ErrInvalidJobID means Slurm acknowledged a job submission without a positive
	// job ID. The job may have been accepted, so callers must not automatically
	// resubmit it in response to this error.
	ErrInvalidJobID = errors.New("slurm acknowledged submission without a valid job ID")
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
