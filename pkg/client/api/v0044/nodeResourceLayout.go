// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type NodeResourceLayoutInterface interface {
	GetNodeResourceLayout(ctx context.Context, jobId string) (*types.V0044NodeResourceLayout, error)
}

var _ NodeResourceLayoutInterface = &SlurmClient{}

// GetNodeResourceLayout implements ClientInterface
func (c *SlurmClient) GetNodeResourceLayout(ctx context.Context, jobId string) (*types.V0044NodeResourceLayout, error) {
	res, err := c.SlurmV0044GetResourcesWithResponse(ctx, jobId)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	if len(res.JSON200.Nodes) == 0 {
		return nil, apierrors.ErrObjectNotFound
	}

	out := &types.V0044NodeResourceLayout{
		V0044NodeResourceLayoutList: make([]api.V0044NodeResourceLayout, len(res.JSON200.Nodes)),
	}
	for i, item := range res.JSON200.Nodes {
		utils.RemarshalOrDie(item, &out.V0044NodeResourceLayoutList[i])
	}
	return out, nil
}
