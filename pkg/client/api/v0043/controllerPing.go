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
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type ControllerPingInfoInterface interface {
	GetControllerPing(ctx context.Context, host string) (*types.V0043ControllerPing, error)
	ListControllerPing(ctx context.Context) (*types.V0043ControllerPingList, error)
}

// GetControllerPing implements ClientInterface
func (c *SlurmClient) GetControllerPing(ctx context.Context, host string) (*types.V0043ControllerPing, error) {
	list, err := c.ListControllerPing(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		pingHost := ptr.Deref(item.Hostname, "")
		if pingHost == host {
			return &item, nil
		}
	}
	return nil, errors.New(http.StatusText(http.StatusNotFound))
}

// ListControllerPing implements ClientInterface
func (c *SlurmClient) ListControllerPing(ctx context.Context) (*types.V0043ControllerPingList, error) {
	res, err := c.SlurmV0043GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
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
	list := &types.V0043ControllerPingList{
		Items: make([]types.V0043ControllerPing, len(res.JSON200.Pings)),
	}
	for i, item := range res.JSON200.Pings {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
