// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/SlinkyProject/slurm-client/pkg/types"
)

type ReconfigureInterface interface {
	GetReconfigure(ctx context.Context) (*types.V0044Reconfigure, error)
	ListReconfigure(ctx context.Context) (*types.V0044ReconfigureList, error)
}

// GetReconfigure implements ClientInterface
func (c *SlurmClient) GetReconfigure(ctx context.Context) (*types.V0044Reconfigure, error) {
	res, err := c.SlurmV0044GetReconfigureWithResponse(ctx)
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
	out := &types.V0044Reconfigure{}
	return out, nil
}

// ListReconfigure implements ClientInterface
func (c *SlurmClient) ListReconfigure(ctx context.Context) (*types.V0044ReconfigureList, error) {
	res, err := c.GetReconfigure(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0044ReconfigureList{
		Items: []types.V0044Reconfigure{
			*res,
		},
	}
	return list, nil
}
