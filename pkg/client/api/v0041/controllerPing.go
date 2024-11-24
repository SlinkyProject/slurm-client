// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func parseControllerPing(pingV api.V0041ControllerPing) types.ControllerPing {
	ping := types.ControllerPing{
		Hostname: ptr.Deref(pingV.Hostname, ""),
		Pinged:   ptr.Deref(pingV.Pinged, "") == types.ControllerPingPingedUP,
	}
	return ping
}

// GetControllerPing implements SlurmClientInterface
func (c *SlurmClient) GetControllerPing(ctx context.Context, host string) (*types.ControllerPing, error) {
	pingList, err := c.ListControllerPing(ctx)
	if err != nil {
		return nil, err
	}
	ping := &types.ControllerPing{}
	for _, p := range pingList.Items {
		if p.Hostname == host {
			ping = &p
		}
	}
	return ping, nil
}

// ListControllerPing implements SlurmClientInterface
func (c *SlurmClient) ListControllerPing(ctx context.Context) (*types.ControllerPingList, error) {
	res, err := c.SlurmV0041GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	pingList := &types.ControllerPingList{}
	for _, p := range res.JSON200.Pings {
		ping := parseControllerPing(p)
		pingList.AppendItem(&ping)
	}
	return pingList, nil
}
