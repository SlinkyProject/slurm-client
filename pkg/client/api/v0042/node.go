// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

type NodeInterface interface {
	DeleteNode(ctx context.Context, nodeName string) error
	UpdateNode(ctx context.Context, nodeName string, req any) error
	GetNode(ctx context.Context, nodeName string) (*api.V0042Node, error)
	ListNodes(ctx context.Context) (*api.V0042NodeObjectList, error)
}

var _ NodeInterface = &SlurmClient{}

// DeleteNode implements ClientInterface
func (c *SlurmClient) DeleteNode(ctx context.Context, nodeName string) error {
	res, err := c.SlurmV0042DeleteNodeWithResponse(ctx, nodeName)
	if err != nil {
		return err
	} else if res.StatusCode() != http.StatusOK {
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
	r, ok := req.(api.V0042UpdateNodeMsg)
	if !ok {
		return errors.New("expected req to be V0042UpdateNodeMsg")
	}
	body := api.SlurmV0042PostNodeJSONRequestBody(r)
	res, err := c.SlurmV0042PostNodeWithResponse(ctx, nodeName, body)
	if err != nil {
		return err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return errors.Join(errs...)
	}
	return nil
}

// GetNode implements ClientInterface
func (c *SlurmClient) GetNode(ctx context.Context, nodeName string) (*api.V0042Node, error) {
	params := &api.SlurmV0042GetNodeParams{}
	res, err := c.SlurmV0042GetNodeWithResponse(ctx, nodeName, params)
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
func (c *SlurmClient) ListNodes(ctx context.Context) (*api.V0042NodeObjectList, error) {
	params := &api.SlurmV0042GetNodesParams{}
	res, err := c.SlurmV0042GetNodesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}
	return &api.V0042NodeObjectList{Items: res.JSON200.Nodes}, nil
}
