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

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*api.V0042PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*api.V0042PartitionInfoObjectList, error)
}

var _ PartitionInterface = &SlurmClient{}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*api.V0042PartitionInfo, error) {
	params := &api.SlurmV0042GetPartitionParams{}
	res, err := c.SlurmV0042GetPartitionWithResponse(ctx, name, params)
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
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*api.V0042PartitionInfoObjectList, error) {
	params := &api.SlurmV0042GetPartitionsParams{}
	res, err := c.SlurmV0042GetPartitionsWithResponse(ctx, params)
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

	return &api.V0042PartitionInfoObjectList{Items: res.JSON200.Partitions}, nil
}
