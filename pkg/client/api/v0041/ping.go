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

func parsePing(pingV api.V0041ControllerPing) types.Ping {
	ping := types.Ping{
		Hostname: ptr.Deref(pingV.Hostname, ""),
		Pinged:   ptr.Deref(pingV.Pinged, "") == types.PingPingedUP,
	}
	return ping
}

// GetPing implements SlurmClientInterface
func (c *SlurmClient) GetPing(ctx context.Context, host string) (*types.Ping, error) {
	pingList, err := c.ListPing(ctx)
	if err != nil {
		return nil, err
	}
	ping := &types.Ping{}
	for _, p := range pingList.Items {
		if p.Hostname == host {
			ping = &p
		}
	}
	return ping, nil
}

// ListPing implements SlurmClientInterface
func (c *SlurmClient) ListPing(ctx context.Context) (*types.PingList, error) {
	res, err := c.SlurmV0041GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	pingList := &types.PingList{}
	for _, p := range res.JSON200.Pings {
		ping := parsePing(p)
		pingList.AppendItem(&ping)
	}
	return pingList, nil
}
