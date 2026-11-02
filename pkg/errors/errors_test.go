// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package errors_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

func TestNewHTTPError_NotFound(t *testing.T) {
	if err := apierrors.NewHTTPError(http.StatusNotFound); !errors.Is(err, apierrors.ErrNotFound) {
		t.Errorf("NewHTTPError(404) error = %v, want ErrNotFound", err)
	}
}

func TestNewHTTPError_UnmappedStatus(t *testing.T) {
	for _, statusCode := range []int{http.StatusInternalServerError, http.StatusNotImplemented} {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			err := apierrors.NewHTTPError(statusCode)
			if got, want := err.Error(), http.StatusText(statusCode); got != want {
				t.Errorf("NewHTTPError() error = %q, want %q", got, want)
			}
			if errors.Is(err, apierrors.ErrNotFound) {
				t.Errorf("NewHTTPError() error = %v, unexpectedly matches ErrNotFound", err)
			}
			if errors.Is(err, apierrors.ErrNotImplemented) {
				t.Errorf("NewHTTPError() error = %v, unexpectedly matches ErrNotImplemented", err)
			}
		})
	}
}

func TestSlurmErrorAs(t *testing.T) {
	err := fmt.Errorf("submit: %w", errors.Join(errors.New("Bad Request"), &apierrors.SlurmError{Number: apierrors.ErrInvalidTimeLimit, Message: "Invalid time limit"}))
	var se *apierrors.SlurmError
	if !errors.As(err, &se) || se.Number != apierrors.ErrInvalidTimeLimit || se.Error() != "Invalid time limit" {
		t.Errorf("errors.As = %v, %+v", errors.As(err, &se), se)
	}
}
