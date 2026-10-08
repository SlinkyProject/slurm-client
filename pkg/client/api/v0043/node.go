// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

type NodeInterface interface {
	DeleteNode(ctx context.Context, nodeName string) error
	UpdateNode(ctx context.Context, nodeName string, req any) error
	GetNode(ctx context.Context, nodeName string) (*api.V0043Node, error)
	ListNodes(ctx context.Context) (*api.V0043NodeObjectList, error)
}

var _ NodeInterface = &SlurmClient{}

// DeleteNode implements ClientInterface
func (c *SlurmClient) DeleteNode(ctx context.Context, nodeName string) error {
	res, err := c.SlurmV0043DeleteNodeWithResponse(ctx, nodeName)
	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return errors.Join(errs...)
	}

	return nil
}

// UpdateNode implements ClientInterface
func (c *SlurmClient) UpdateNode(ctx context.Context, nodeName string, req any) error {
	r, ok := req.(api.V0043UpdateNodeMsg)
	if !ok {
		return errors.New("expected req to be V0043UpdateNodeMsg")
	}
	body := api.SlurmV0043PostNodeJSONRequestBody(r)
	res, err := c.SlurmV0043PostNodeWithResponse(ctx, nodeName, body)
	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return errors.Join(errs...)
	}

	return nil
}

// GetNode implements ClientInterface
func (c *SlurmClient) GetNode(ctx context.Context, nodeName string) (*api.V0043Node, error) {
	params := &api.SlurmV0043GetNodeParams{}
	res, err := c.SlurmV0043GetNodeWithResponse(ctx, nodeName, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{apierrors.NewHTTPError(res.StatusCode())}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	if len(res.JSON200.Nodes) == 0 {
		return nil, apierrors.ErrNotFound
	}

	out := res.JSON200.Nodes[0]
	return &out, nil
}

// ListNodes implements ClientInterface
func (c *SlurmClient) ListNodes(ctx context.Context) (*api.V0043NodeObjectList, error) {
	params := &api.SlurmV0043GetNodesParams{}
	res, err := c.SlurmV0043GetNodesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	return &api.V0043NodeObjectList{Items: res.JSON200.Nodes}, nil
}
