// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*types.V0042PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*types.V0042PartitionInfoList, error)
}

var _ PartitionInterface = &SlurmClient{}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*types.V0042PartitionInfo, error) {
	params := &api.SlurmV0042GetPartitionParams{}
	res, err := c.SlurmV0042GetPartitionWithResponse(ctx, name, params)
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

	if len(res.JSON200.Partitions) == 0 {
		return nil, apierrors.ErrNotFound
	}

	out := &types.V0042PartitionInfo{}
	utils.RemarshalOrDie(res.JSON200.Partitions[0], out)
	return out, nil
}

// ListPartitionInfo implements ClientInterface
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*types.V0042PartitionInfoList, error) {
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

	list := &types.V0042PartitionInfoList{
		Items: make([]types.V0042PartitionInfo, len(res.JSON200.Partitions)),
	}
	for i, item := range res.JSON200.Partitions {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
