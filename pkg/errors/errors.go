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
