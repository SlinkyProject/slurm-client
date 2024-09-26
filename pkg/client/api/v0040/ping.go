// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0040

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

func parsePing(pingV api.V0040ControllerPing) slurmtypes.Ping {
	ping := slurmtypes.Ping{
		Hostname: ptr.Deref(pingV.Hostname, ""),
		Pinged:   ptr.Deref(pingV.Pinged, "") == slurmtypes.PingPingedUP,
	}
	return ping
}

// GetPing implements SlurmClientInterface
func (c *SlurmClient) GetPing(ctx context.Context, host string) (*slurmtypes.Ping, error) {
	pingList, err := c.ListPing(ctx)
	if err != nil {
		return nil, err
	}
	ping := &slurmtypes.Ping{}
	for _, p := range pingList.Items {
		if p.Hostname == host {
			ping = &p
		}
	}
	return ping, nil
}

// ListPing implements SlurmClientInterface
func (c *SlurmClient) ListPing(ctx context.Context) (*slurmtypes.PingList, error) {
	res, err := c.SlurmV0040GetPingWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	pingList := &slurmtypes.PingList{}
	for _, p := range res.JSON200.Pings {
		ping := parsePing(p)
		pingList.AppendItem(&ping)
	}
	return pingList, nil
}
