// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type NodeResourceLayoutInterface interface {
	GetNodeResourceLayout(ctx context.Context, jobId string) (*types.V0045NodeResourceLayout, error)
}

var _ NodeResourceLayoutInterface = &SlurmClient{}

// GetNodeResourceLayout implements ClientInterface
func (c *SlurmClient) GetNodeResourceLayout(ctx context.Context, jobId string) (*types.V0045NodeResourceLayout, error) {
	res, err := c.SlurmV0045GetResourcesWithResponse(ctx, jobId)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	if len(res.JSON200.Nodes) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0045NodeResourceLayout{
		V0045NodeResourceLayoutList: make([]api.V0045NodeResourceLayout, len(res.JSON200.Nodes)),
	}
	for i, item := range res.JSON200.Nodes {
		utils.RemarshalOrDie(item, &out.V0045NodeResourceLayoutList[i])
	}
	return out, nil
}
