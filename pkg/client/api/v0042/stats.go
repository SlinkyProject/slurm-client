// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type StatsInterface interface {
	GetStats(ctx context.Context) (*types.V0042Stats, error)
	ListStats(ctx context.Context) (*types.V0042StatsList, error)
}

// GetStats implements ClientInterface
func (c *SlurmClient) GetStats(ctx context.Context) (*types.V0042Stats, error) {
	res, err := c.SlurmV0042GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	out := &types.V0042Stats{}
	utils.RemarshalOrDie(res.JSON200.Statistics, out)
	return out, nil
}

// ListStats implements ClientInterface
func (c *SlurmClient) ListStats(ctx context.Context) (*types.V0042StatsList, error) {
	res, err := c.GetStats(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0042StatsList{
		Items: []types.V0042Stats{
			*res,
		},
	}
	return list, nil
}
