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

type JobInfoInterface interface {
	GetJobInfo(ctx context.Context, jobId string) (*types.V0041JobInfo, error)
	ListJobInfo(ctx context.Context) (*types.V0041JobInfoList, error)
}

// GetJobInfo implements ClientInterface
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*types.V0041JobInfo, error) {
	params := &api.SlurmV0041GetJobParams{}
	res, err := c.SlurmV0041GetJobWithResponse(ctx, jobId, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Jobs) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0041JobInfo{}
	utils.RemarshalOrDie(res.JSON200.Jobs[0], out)
	return out, nil
}

// ListJobInfo implements ClientInterface
func (c *SlurmClient) ListJobInfo(ctx context.Context) (*types.V0041JobInfoList, error) {
	params := &api.SlurmV0041GetJobsParams{}
	res, err := c.SlurmV0041GetJobsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0041JobInfoList{
		Items: make([]types.V0041JobInfo, len(res.JSON200.Jobs)),
	}
	for i, item := range res.JSON200.Jobs {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
