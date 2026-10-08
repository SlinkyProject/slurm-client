// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
)

type StatsInterface interface {
	GetStats(ctx context.Context) (*api.V0045StatsMsg, error)
	ListStats(ctx context.Context) (*api.V0045StatsMsgObjectList, error)
}

var _ StatsInterface = &SlurmClient{}

// GetStats implements ClientInterface
func (c *SlurmClient) GetStats(ctx context.Context) (*api.V0045StatsMsg, error) {
	res, err := c.SlurmV0045GetDiagWithResponse(ctx)
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
func (c *SlurmClient) ListStats(ctx context.Context) (*api.V0045StatsMsgObjectList, error) {
	res, err := c.SlurmV0045GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}
	return &api.V0045StatsMsgObjectList{Items: []api.V0045StatsMsg{res.JSON200.Statistics}}, nil
}
