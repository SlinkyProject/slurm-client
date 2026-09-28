// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*api.V0045PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*api.V0045PartitionInfoObjectList, error)
}

var _ PartitionInterface = &SlurmClient{}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*api.V0045PartitionInfo, error) {
	params := &api.SlurmV0045GetPartitionParams{}
	res, err := c.SlurmV0045GetPartitionWithResponse(ctx, name, params)
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

	if len(res.JSON200.Partitions) == 0 {
		return nil, apierrors.ErrNotFound
	}

	out := res.JSON200.Partitions[0]
	return &out, nil
}

// ListPartitionInfo implements ClientInterface
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*api.V0045PartitionInfoObjectList, error) {
	params := &api.SlurmV0045GetPartitionsParams{}
	res, err := c.SlurmV0045GetPartitionsWithResponse(ctx, params)
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

	return &api.V0045PartitionInfoObjectList{Items: res.JSON200.Partitions}, nil
}
