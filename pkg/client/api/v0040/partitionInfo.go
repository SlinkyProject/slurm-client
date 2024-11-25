// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0040

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*types.V0040PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*types.V0040PartitionInfoList, error)
}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*types.V0040PartitionInfo, error) {
	params := &api.SlurmV0040GetPartitionParams{}
	res, err := c.SlurmV0040GetPartitionWithResponse(ctx, name, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Partitions) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0040PartitionInfo{}
	utils.RemarshalOrDie(res.JSON200.Partitions[0], out)
	return out, nil
}

// ListPartitionInfo implements ClientInterface
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*types.V0040PartitionInfoList, error) {
	params := &api.SlurmV0040GetPartitionsParams{}
	res, err := c.SlurmV0040GetPartitionsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0040PartitionInfoList{
		Items: make([]types.V0040PartitionInfo, len(res.JSON200.Partitions)),
	}
	for i, item := range res.JSON200.Partitions {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
