// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/log"

	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

func (c *client) getPartitionInfo(ctx context.Context, name string) (*slurmtypes.PartitionInfo, error) {
	var partitionInfo *slurmtypes.PartitionInfo
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		partitionInfo, err = c.clients[i].GetPartitionInfo(ctx, name)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	log.V(1).Info("GetPartitionInfo()", "partition", partitionInfo)
	return partitionInfo, nil
}

func (c *client) listPartitionInfos(ctx context.Context) (*slurmtypes.PartitionInfoList, error) {
	var partitionInfoList *slurmtypes.PartitionInfoList
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		partitionInfoList, err = c.clients[i].ListPartitionInfos(ctx)
		if err == nil {
			c.clientUse = i
			break
		} else if tolerateError(err) {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	log.V(1).Info("ListPartitionInfos()", "size", partitionInfoList.Size(), "partitionList", partitionInfoList)
	return partitionInfoList, nil
}
