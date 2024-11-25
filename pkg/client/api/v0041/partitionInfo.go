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

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*types.V0041PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*types.V0041PartitionInfoList, error)
}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*types.V0041PartitionInfo, error) {
	params := &api.SlurmV0041GetPartitionParams{}
	res, err := c.SlurmV0041GetPartitionWithResponse(ctx, name, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Partitions) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0041PartitionInfo{}
	utils.RemarshalOrDie(res.JSON200.Partitions[0], out)
	return out, nil
}

// ListPartitionInfo implements ClientInterface
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*types.V0041PartitionInfoList, error) {
	params := &api.SlurmV0041GetPartitionsParams{}
	res, err := c.SlurmV0041GetPartitionsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0041PartitionInfoList{
		Items: make([]types.V0041PartitionInfo, len(res.JSON200.Partitions)),
	}
	for i, item := range res.JSON200.Partitions {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
