// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package errors_test

import (
	"errors"
	"net/http"
	"testing"

	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

func TestNewHTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       error
	}{
		{
			name:       "Not Found",
			statusCode: http.StatusNotFound,
			want:       apierrors.ErrNotFound,
		},
		{
			name:       "Not Implemented",
			statusCode: http.StatusNotImplemented,
			want:       apierrors.ErrNotImplemented,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := apierrors.NewHTTPError(tt.statusCode)
			if !errors.Is(err, tt.want) {
				t.Errorf("NewHTTPError() error = %v, want errors.Is(error, %v)", err, tt.want)
			}
		})
	}
}

func TestNewHTTPError_UnmappedStatus(t *testing.T) {
	err := apierrors.NewHTTPError(http.StatusInternalServerError)
	if got, want := err.Error(), http.StatusText(http.StatusInternalServerError); got != want {
		t.Errorf("NewHTTPError() error = %q, want %q", got, want)
	}
	if errors.Is(err, apierrors.ErrNotFound) {
		t.Errorf("NewHTTPError() error = %v, unexpectedly matches ErrNotFound", err)
	}
}
