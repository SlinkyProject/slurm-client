// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
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
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
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
