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
	slurmtypes "github.com/SlinkyProject/slurm-client/pkg/types"
)

func parseJobInfo(jobInfoV api.V0041JobInfo) slurmtypes.JobInfo {
	jobInfo := slurmtypes.JobInfo{
		JobId:        ptr.Deref(jobInfoV.JobId, 0),
		Partition:    ptr.Deref(jobInfoV.Partition, ""),
		Uid:          ptr.Deref(jobInfoV.UserId, 0),
		UserName:     ptr.Deref(jobInfoV.UserName, ""),
		JobState:     make(set.Set[slurmtypes.JobInfoJobState], 0),
		Cpus:         int64(ptr.Deref(jobInfoV.Cpus.Number, 0)),
		NodeCount:    int64(ptr.Deref(jobInfoV.NodeCount.Number, 0)),
		Hold:         ptr.Deref(jobInfoV.Hold, false),
		EligibleTime: ptr.Deref(jobInfoV.EligibleTime.Number, 0),
	}
	states := ptr.Deref(jobInfoV.JobState, []api.V0041JobInfoJobState{})
	for _, state := range states {
		jobInfo.JobState.Insert(slurmtypes.JobInfoJobState(state))
	}
	return jobInfo
}

// GetJobInfo implements SlurmClientInterface
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*slurmtypes.JobInfo, error) {
	params := &api.SlurmV0041GetJobParams{}
	res, err := c.SlurmV0041GetJobWithResponse(ctx, jobId, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Jobs) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	jobInfoV := res.JSON200.Jobs[0]
	jobInfo := parseJobInfo(jobInfoV)
	return &jobInfo, nil
}

// ListJobInfos implements SlurmClientInterface
func (c *SlurmClient) ListJobInfos(ctx context.Context) (*slurmtypes.JobInfoList, error) {
	params := &api.SlurmV0041GetJobsParams{}
	res, err := c.SlurmV0041GetJobsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	jobInfoList := &slurmtypes.JobInfoList{}
	for _, jobInfoV := range res.JSON200.Jobs {
		jobInfo := parseJobInfo(jobInfoV)
		jobInfoList.Items = append(jobInfoList.Items, jobInfo)
	}
	return jobInfoList, nil
}
