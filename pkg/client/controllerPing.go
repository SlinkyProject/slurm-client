// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *client) getControllerPing(ctx context.Context, host string) (*types.ControllerPing, error) {
	var ping *types.ControllerPing
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		ping, err = c.clients[i].GetControllerPing(ctx, host)
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
	log.V(1).Info("GetControllerPing()", "ping", ping)
	return ping, nil
}

func (c *client) listControllerPing(ctx context.Context) (*types.ControllerPingList, error) {
	var pingList *types.ControllerPingList
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		pingList, err = c.clients[i].ListControllerPing(ctx)
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
	log.V(1).Info("ListControllerPing()", "size", len(pingList.Items), "controllerPingList", pingList)
	return pingList, nil
}
