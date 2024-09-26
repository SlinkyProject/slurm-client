// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/log"

	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

func (c *client) getNode(ctx context.Context, nodeName string) (*slurmtypes.Node, error) {
	var node *slurmtypes.Node
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		node, err = c.clients[i].GetNode(ctx, nodeName)
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
	log.V(1).Info("GetNode()", "node", node)
	return node, nil
}

func (c *client) listNodes(ctx context.Context) (*slurmtypes.NodeList, error) {
	var nodeList *slurmtypes.NodeList
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		nodeList, err = c.clients[i].ListNodes(ctx)
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
	log.V(1).Info("ListNodes()", "size", nodeList.Size(), "nodeList", nodeList)
	return nodeList, nil
}

func (c *client) deleteNode(ctx context.Context, nodeName string) error {
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		err = c.clients[i].DeleteNode(ctx, nodeName)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return err
	}
	log.V(1).Info("DeleteNode()", "nodeName", nodeName)
	return nil
}

func (c *client) updateNode(ctx context.Context, node *slurmtypes.Node, originalNode *slurmtypes.Node) error {
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		err = c.clients[i].UpdateNode(ctx, node, originalNode)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return err
	}
	log.V(1).Info("UpdateNode()", "node", node)
	return nil
}
