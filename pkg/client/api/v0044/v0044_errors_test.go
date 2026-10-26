// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"errors"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

func TestGetOpenapiErrorsKeepsNumber(t *testing.T) {
	in := api.V0044OpenapiErrors{{Error: ptr.To("Invalid account"), ErrorNumber: ptr.To(apierrors.ErrInvalidAccount), Source: ptr.To("_job_submit")}}
	errs := getOpenapiErrors(&in)
	var se *apierrors.SlurmError
	if len(errs) != 1 || !errors.As(errors.Join(errs...), &se) || se.Number != apierrors.ErrInvalidAccount || se.Source != "_job_submit" || se.Error() != "Invalid account" {
		t.Fatalf("got %v, %+v", errs, se)
	}
}
