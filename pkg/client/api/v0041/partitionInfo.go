// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func parsePartitionInfo(partitionInfoV api.V0041PartitionInfo) types.PartitionInfo {
	partitionInfo := types.PartitionInfo{
		Name:           ptr.Deref(partitionInfoV.Name, ""),
		NodeSets:       ptr.Deref(partitionInfoV.NodeSets, ""),
		Nodes:          ptr.Deref(partitionInfoV.Nodes.Total, 0),
		Cpus:           ptr.Deref(partitionInfoV.Cpus.Total, 0),
		PartitionState: make(set.Set[types.PartitionInfoPartitionState], 0),
	}
	states := ptr.Deref(partitionInfoV.Partition.State, []api.V0041PartitionInfoPartitionState{})
	for _, state := range states {
		partitionInfo.PartitionState.Insert(types.PartitionInfoPartitionState(state))
	}
	return partitionInfo
}

// GetPartitionInfo implements SlurmClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*types.PartitionInfo, error) {
	params := &api.SlurmV0041GetPartitionParams{}
	res, err := c.SlurmV0041GetPartitionWithResponse(ctx, name, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Partitions) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	partitionInfoV := res.JSON200.Partitions[0]
	partitionInfo := parsePartitionInfo(partitionInfoV)
	return &partitionInfo, nil
}

// ListPartitionInfos implements SlurmClientInterface
func (c *SlurmClient) ListPartitionInfos(ctx context.Context) (*types.PartitionInfoList, error) {
	params := &api.SlurmV0041GetPartitionsParams{}
	res, err := c.SlurmV0041GetPartitionsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	partitionInfoList := &types.PartitionInfoList{}
	for _, partitionInfoV := range res.JSON200.Partitions {
		partitionInfo := parsePartitionInfo(partitionInfoV)
		partitionInfoList.Items = append(partitionInfoList.Items, partitionInfo)
	}
	return partitionInfoList, nil
}
