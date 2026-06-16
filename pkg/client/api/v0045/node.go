// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type NodeInterface interface {
	CreateNewNode(ctx context.Context, req any) (*string, error)
	DeleteNode(ctx context.Context, nodeName string) error
	UpdateNode(ctx context.Context, nodeName string, req any) error
	GetNode(ctx context.Context, nodeName string) (*types.V0045Node, error)
	ListNodes(ctx context.Context) (*types.V0045NodeList, error)
}

var _ NodeInterface = &SlurmClient{}

// CreateNewNode implements ClientInterface
func (c *SlurmClient) CreateNewNode(ctx context.Context, req any) (*string, error) {
	r, ok := req.(api.V0045OpenapiCreateNodeReq)
	if !ok {
		return nil, errors.New("expected req to be V0045OpenapiCreateNodeReq")
	}
	body := api.SlurmV0045PostNewNodeJSONRequestBody(r)
	res, err := c.SlurmV0045PostNewNodeWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	nodeName, err := utils.ParseNodeName(r.NodeConf)
	if err != nil {
		return nil, err
	}

	return &nodeName, nil
}

// DeleteNode implements ClientInterface
func (c *SlurmClient) DeleteNode(ctx context.Context, nodeName string) error {
	res, err := c.SlurmV0045DeleteNodeWithResponse(ctx, nodeName)
	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return utilerrors.NewAggregate(errs)
	}

	return nil
}

// UpdateNode implements ClientInterface
func (c *SlurmClient) UpdateNode(ctx context.Context, nodeName string, req any) error {
	r, ok := req.(api.V0045UpdateNodeMsg)
	if !ok {
		return errors.New("expected req to be V0045UpdateNodeMsg")
	}
	body := api.SlurmV0045PostNodeJSONRequestBody(r)
	res, err := c.SlurmV0045PostNodeWithResponse(ctx, nodeName, body)
	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return utilerrors.NewAggregate(errs)
	}

	return nil
}

// GetNode implements ClientInterface
func (c *SlurmClient) GetNode(ctx context.Context, nodeName string) (*types.V0045Node, error) {
	params := &api.SlurmV0045GetNodeParams{}
	res, err := c.SlurmV0045GetNodeWithResponse(ctx, nodeName, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	if len(res.JSON200.Nodes) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0045Node{}
	utils.RemarshalOrDie(res.JSON200.Nodes[0], out)
	return out, nil
}

// ListNodes implements ClientInterface
func (c *SlurmClient) ListNodes(ctx context.Context) (*types.V0045NodeList, error) {
	params := &api.SlurmV0045GetNodesParams{}
	res, err := c.SlurmV0045GetNodesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	list := &types.V0045NodeList{
		Items: make([]types.V0045Node, len(res.JSON200.Nodes)),
	}
	for i, item := range res.JSON200.Nodes {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
