// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *client) getPing(ctx context.Context, host string) (*types.Ping, error) {
	var ping *types.Ping
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		ping, err = c.clients[i].GetPing(ctx, host)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	log.V(1).Info("GetPing()", "ping", ping)
	return ping, nil
}

func (c *client) listPing(ctx context.Context) (*types.PingList, error) {
	var pingList *types.PingList
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		pingList, err = c.clients[i].ListPing(ctx)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	log.V(1).Info("ListPing()", "size", len(pingList.Items), "pingList", pingList)
	return pingList, nil
}
