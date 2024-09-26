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
	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

func parsePartitionInfo(partitionInfoV api.V0040PartitionInfo) slurmtypes.PartitionInfo {
	partitionInfo := slurmtypes.PartitionInfo{
		Name:           ptr.Deref(partitionInfoV.Name, ""),
		NodeSets:       ptr.Deref(partitionInfoV.NodeSets, ""),
		Nodes:          ptr.Deref(partitionInfoV.Nodes.Total, 0),
		Cpus:           ptr.Deref(partitionInfoV.Cpus.Total, 0),
		PartitionState: make(set.Set[slurmtypes.PartitionInfoPartitionState], 0),
	}
	states := ptr.Deref(partitionInfoV.Partition.State, []api.V0040PartitionInfoPartitionState{})
	for _, state := range states {
		partitionInfo.PartitionState.Insert(slurmtypes.PartitionInfoPartitionState(state))
	}
	return partitionInfo
}

// GetPartitionInfo implements SlurmClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*slurmtypes.PartitionInfo, error) {
	params := &api.SlurmV0040GetPartitionParams{}
	res, err := c.SlurmV0040GetPartitionWithResponse(ctx, name, params)
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
func (c *SlurmClient) ListPartitionInfos(ctx context.Context) (*slurmtypes.PartitionInfoList, error) {
	params := &api.SlurmV0040GetPartitionsParams{}
	res, err := c.SlurmV0040GetPartitionsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	partitionInfoList := &slurmtypes.PartitionInfoList{}
	for _, partitionInfoV := range res.JSON200.Partitions {
		partitionInfo := parsePartitionInfo(partitionInfoV)
		partitionInfoList.Items = append(partitionInfoList.Items, partitionInfo)
	}
	return partitionInfoList, nil
}
