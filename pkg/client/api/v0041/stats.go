// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type StatsInterface interface {
	GetStats(ctx context.Context) (*types.V0041Stats, error)
	ListStats(ctx context.Context) (*types.V0041StatsList, error)
}

// GetStats implements ClientInterface
func (c *SlurmClient) GetStats(ctx context.Context) (*types.V0041Stats, error) {
	res, err := c.SlurmV0041GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	out := &types.V0041Stats{}
	utils.RemarshalOrDie(res.JSON200.Statistics, out)
	return out, nil
}

// ListStats implements ClientInterface
func (c *SlurmClient) ListStats(ctx context.Context) (*types.V0041StatsList, error) {
	res, err := c.GetStats(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0041StatsList{
		Items: []types.V0041Stats{
			*res,
		},
	}
	return list, nil
}
