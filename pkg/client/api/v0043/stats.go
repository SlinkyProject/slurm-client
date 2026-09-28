// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
)

type StatsInterface interface {
	GetStats(ctx context.Context) (*api.V0043StatsMsg, error)
	ListStats(ctx context.Context) (*api.V0043StatsMsgObjectList, error)
}

var _ StatsInterface = &SlurmClient{}

// GetStats implements ClientInterface
func (c *SlurmClient) GetStats(ctx context.Context) (*api.V0043StatsMsg, error) {
	res, err := c.SlurmV0043GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}
	out := res.JSON200.Statistics
	return &out, nil
}

// ListStats implements ClientInterface
func (c *SlurmClient) ListStats(ctx context.Context) (*api.V0043StatsMsgObjectList, error) {
	res, err := c.SlurmV0043GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}
	return &api.V0043StatsMsgObjectList{Items: []api.V0043StatsMsg{res.JSON200.Statistics}}, nil
}
