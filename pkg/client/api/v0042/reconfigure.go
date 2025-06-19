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
)

type ReconfigureInterface interface {
	GetReconfigure(ctx context.Context) (*types.V0042Reconfigure, error)
	ListReconfigure(ctx context.Context) (*types.V0042ReconfigureList, error)
}

// GetReconfigure implements ClientInterface
func (c *SlurmClient) GetReconfigure(ctx context.Context) (*types.V0042Reconfigure, error) {
	res, err := c.SlurmV0042GetReconfigureWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
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
	out := &types.V0042Reconfigure{}
	return out, nil
}

// ListReconfigure implements ClientInterface
func (c *SlurmClient) ListReconfigure(ctx context.Context) (*types.V0042ReconfigureList, error) {
	res, err := c.GetReconfigure(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0042ReconfigureList{
		Items: []types.V0042Reconfigure{
			*res,
		},
	}
	return list, nil
}
