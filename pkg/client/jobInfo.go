// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GetJobInfo implements SlurmClientInterface
func (c *client) getJobInfo(ctx context.Context, jobId string) (*types.JobInfo, error) {
	var jobInfo *types.JobInfo
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		jobInfo, err = c.clients[i].GetJobInfo(ctx, jobId)
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
	log.V(1).Info("GetJobInfo()", "job", jobInfo)
	return jobInfo, nil
}

// ListJobInfo implements SlurmClientInterface
func (c *client) listJobInfos(ctx context.Context) (*types.JobInfoList, error) {
	var jobInfoList *types.JobInfoList
	var err error
	log := log.FromContext(ctx)
	for i := c.clientUse; i < len(c.clients); i++ {
		jobInfoList, err = c.clients[i].ListJobInfo(ctx)
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
	log.V(1).Info("ListJobInfo()", "size", jobInfoList.Size(), "jobList", jobInfoList)
	return jobInfoList, nil
}
