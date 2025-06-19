// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

type ReconfigureInterface interface {
	GetReconfigure(ctx context.Context) (*types.V0043Reconfigure, error)
	ListReconfigure(ctx context.Context) (*types.V0043ReconfigureList, error)
}

// GetReconfigure implements ClientInterface
func (c *SlurmClient) GetReconfigure(ctx context.Context) (*types.V0043Reconfigure, error) {
	res, err := c.SlurmV0043GetReconfigureWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0043OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	out := &types.V0043Reconfigure{}
	return out, nil
}

// ListReconfigure implements ClientInterface
func (c *SlurmClient) ListReconfigure(ctx context.Context) (*types.V0043ReconfigureList, error) {
	res, err := c.GetReconfigure(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0043ReconfigureList{
		Items: []types.V0043Reconfigure{
			*res,
		},
	}
	return list, nil
}
