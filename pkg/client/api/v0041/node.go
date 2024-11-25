// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type NodeInterface interface {
	DeleteNode(ctx context.Context, nodeName string) error
	UpdateNode(ctx context.Context, nodeName string, req any) error
	GetNode(ctx context.Context, nodeName string) (*types.V0041Node, error)
	ListNodes(ctx context.Context) (*types.V0041NodeList, error)
}

// DeleteNode implements ClientInterface
func (c *SlurmClient) DeleteNode(ctx context.Context, nodeName string) error {
	res, err := c.SlurmV0041DeleteNodeWithResponse(ctx, nodeName)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		return errors.New(http.StatusText(res.StatusCode()))
	}
	return nil
}

// UpdateNode implements ClientInterface
func (c *SlurmClient) UpdateNode(ctx context.Context, nodeName string, req any) error {
	r, ok := req.(api.V0041UpdateNodeMsg)
	if !ok {
		return errors.New("expected req to be V0041UpdateNodeMsg")
	}
	body := api.SlurmV0041PostNodeJSONRequestBody(r)
	res, err := c.SlurmV0041PostNodeWithResponse(ctx, nodeName, body)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		return errors.New(http.StatusText(res.StatusCode()))
	}
	return nil
}

// GetNode implements ClientInterface
func (c *SlurmClient) GetNode(ctx context.Context, nodeName string) (*types.V0041Node, error) {
	params := &api.SlurmV0041GetNodeParams{}
	res, err := c.SlurmV0041GetNodeWithResponse(ctx, nodeName, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Nodes) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0041Node{}
	utils.RemarshalOrDie(res.JSON200.Nodes[0], out)
	return out, nil
}

// ListNodes implements ClientInterface
func (c *SlurmClient) ListNodes(ctx context.Context) (*types.V0041NodeList, error) {
	params := &api.SlurmV0041GetNodesParams{}
	res, err := c.SlurmV0041GetNodesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0041NodeList{
		Items: make([]types.V0041Node, len(res.JSON200.Nodes)),
	}
	for i, item := range res.JSON200.Nodes {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
