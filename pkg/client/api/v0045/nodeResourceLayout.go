// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

type NodeResourceLayoutInterface interface {
	GetNodeResourceLayout(ctx context.Context, jobId string) (*api.V0045NodeResourceLayoutList, error)
}

var _ NodeResourceLayoutInterface = &SlurmClient{}

// GetNodeResourceLayout implements ClientInterface
func (c *SlurmClient) GetNodeResourceLayout(ctx context.Context, jobId string) (*api.V0045NodeResourceLayoutList, error) {
	res, err := c.SlurmV0045GetResourcesWithResponse(ctx, jobId)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{apierrors.NewHTTPError(res.StatusCode())}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	if len(res.JSON200.Nodes) == 0 {
		return nil, apierrors.ErrNotFound
	}

	return &res.JSON200.Nodes, nil
}
