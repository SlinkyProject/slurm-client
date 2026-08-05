// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"

	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type ControllerPingInfoInterface interface {
	GetControllerPing(ctx context.Context, host string) (*types.V0045ControllerPing, error)
	ListControllerPing(ctx context.Context) (*types.V0045ControllerPingList, error)
}

var _ ControllerPingInfoInterface = &SlurmClient{}

// GetControllerPing implements ClientInterface
func (c *SlurmClient) GetControllerPing(ctx context.Context, host string) (*types.V0045ControllerPing, error) {
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
	return nil, apierrors.ErrNotFound
}

// ListControllerPing implements ClientInterface
func (c *SlurmClient) ListControllerPing(ctx context.Context) (*types.V0045ControllerPingList, error) {
	res, err := c.SlurmV0045GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}
	list := &types.V0045ControllerPingList{
		Items: make([]types.V0045ControllerPing, len(res.JSON200.Pings)),
	}
	for i, item := range res.JSON200.Pings {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
