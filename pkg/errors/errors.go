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

// SlurmError is an error reported by slurmrestd. Use errors.As to read the
// Slurm error number; Error returns the short message alone.
type SlurmError struct {
	Number      int32
	Source      string
	Message     string
	Description string
}

func (e *SlurmError) Error() string { return e.Message }

// Slurm error numbers from slurm_errno.h, for matching SlurmError.Number.
const (
	ErrInvalidPartitionName int32 = 2000
	ErrInvalidNodeCount     int32 = 2006
	ErrMaxJobCount          int32 = 2007 // Slurm 26.05.2 on; earlier releases return EAGAIN instead.
	ErrBadTaskCount         int32 = 2025
	ErrInvalidFeature       int32 = 2029
	ErrInvalidAccount       int32 = 2045
	ErrInvalidLicenses      int32 = 2048
	ErrAccountingPolicy     int32 = 2050
	ErrInvalidTimeLimit     int32 = 2051
	ErrInvalidWCKey         int32 = 2057
	ErrInvalidQOS           int32 = 2066
	ErrPartitionNotAvail    int32 = 2070
	ErrInvalidGRES          int32 = 2072
	ErrPartitionDown        int32 = 2101
	ErrInvalidTRES          int32 = 2115
	ErrUserIDInvalid        int32 = 2165
	ErrUserIDUnknown        int32 = 2166
)
