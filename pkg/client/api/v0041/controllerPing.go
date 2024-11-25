// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type ControllerPingInfoInterface interface {
	GetControllerPing(ctx context.Context, host string) (*types.V0041ControllerPing, error)
	ListControllerPing(ctx context.Context) (*types.V0041ControllerPingList, error)
}

// GetControllerPing implements ClientInterface
func (c *SlurmClient) GetControllerPing(ctx context.Context, host string) (*types.V0041ControllerPing, error) {
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
func (c *SlurmClient) ListControllerPing(ctx context.Context) (*types.V0041ControllerPingList, error) {
	res, err := c.SlurmV0041GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0041ControllerPingList{
		Items: make([]types.V0041ControllerPing, len(res.JSON200.Pings)),
	}
	for i, item := range res.JSON200.Pings {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
