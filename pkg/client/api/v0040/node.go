// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0040

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func parseNode(nodeV api.V0040Node) types.Node {
	node := types.Node{
		Address:       ptr.Deref(nodeV.Address, ""),
		Comment:       ptr.Deref(nodeV.Comment, ""),
		Extra:         ptr.Deref(nodeV.Extra, ""),
		Name:          ptr.Deref(nodeV.Name, ""),
		Partitions:    make(set.Set[string], 0),
		State:         make(set.Set[types.NodeState], 0),
		Cpus:          ptr.Deref(nodeV.Cpus, 0),
		AllocCpus:     ptr.Deref(nodeV.AllocCpus, 0),
		AllocIdleCpus: ptr.Deref(nodeV.AllocIdleCpus, 0),
	}
	partitions := ptr.Deref(nodeV.Partitions, api.V0040CsvString{})
	node.Partitions.Insert(partitions...)
	states := ptr.Deref(nodeV.State, []api.V0040NodeState{})
	for _, state := range states {
		node.State.Insert(types.NodeState(state))
	}
	return node
}

// DeleteNode implements SlurmClientInterface
func (c *SlurmClient) DeleteNode(ctx context.Context, nodeName string) error {
	res, err := c.SlurmV0040DeleteNodeWithResponse(ctx, nodeName)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		return errors.New(http.StatusText(res.StatusCode()))
	}
	return nil
}

// GetNode implements SlurmClientInterface
func (c *SlurmClient) GetNode(ctx context.Context, nodeName string) (*types.Node, error) {
	params := &api.SlurmV0040GetNodeParams{}
	res, err := c.SlurmV0040GetNodeWithResponse(ctx, nodeName, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Nodes) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	nodeV := res.JSON200.Nodes[0]
	node := parseNode(nodeV)
	return &node, nil
}

// ListNodes implements SlurmClientInterface
func (c *SlurmClient) ListNodes(ctx context.Context) (*types.NodeList, error) {
	params := &api.SlurmV0040GetNodesParams{}
	res, err := c.SlurmV0040GetNodesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	nodeList := &types.NodeList{}
	for _, nodeV := range res.JSON200.Nodes {
		node := parseNode(nodeV)
		nodeList.Items = append(nodeList.Items, node)
	}
	return nodeList, nil
}

// UpdateNode implements SlurmClientInterface
func (c *SlurmClient) UpdateNode(ctx context.Context, node *types.Node, originalNode *types.Node) error {
	body := api.SlurmV0040PostNodeJSONRequestBody{
		// No way to know what the node name changed from.
		// Name: utils.Reference([]string{node.Name}),
	}

	if node.Comment != originalNode.Comment {
		body.Comment = ptr.To(node.Comment)
	}

	if node.Extra != originalNode.Extra {
		body.Extra = ptr.To(node.Extra)
	}

	if !node.State.Equal(originalNode.State) {
		diffStates := node.State.Difference(originalNode.State)
		states := []api.V0040UpdateNodeMsgState{}
		for _, state := range diffStates.UnsortedList() {
			states = append(states, api.V0040UpdateNodeMsgState(state))
		}
		body.State = ptr.To(states)
	}

	res, err := c.SlurmV0040PostNodeWithResponse(ctx, node.Name, body)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		return errors.New(http.StatusText(res.StatusCode()))
	}
	return nil
}
