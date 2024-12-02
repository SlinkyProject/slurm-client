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

type JobInfoInterface interface {
	CreateJobInfo(ctx context.Context, req any) (*int32, error)
	DeleteJobInfo(ctx context.Context, jobId string) error
	GetJobInfo(ctx context.Context, jobId string) (*types.V0040JobInfo, error)
	ListJobInfo(ctx context.Context) (*types.V0040JobInfoList, error)
}

// CreateJobInfo implements ClientInterface
func (c *SlurmClient) CreateJobInfo(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0040JobSubmitReq)
	if !ok {
		return nil, errors.New("expected req to be V0040JobSubmitReq")
	}
	body := api.SlurmV0040PostJobSubmitJSONRequestBody(r)
	res, err := c.SlurmV0040PostJobSubmitWithResponse(ctx, body)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	return res.JSON200.JobId, nil
}

// DeleteJobInfo implements ClientInterface
func (c *SlurmClient) DeleteJobInfo(ctx context.Context, jobId string) error {
	params := &api.SlurmV0040DeleteJobParams{}
	res, err := c.SlurmV0040DeleteJobWithResponse(ctx, jobId, params)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		return errors.New(http.StatusText(res.StatusCode()))
	}
	return nil
}

// GetJobInfo implements ClientInterface
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*types.V0040JobInfo, error) {
	params := &api.SlurmV0040GetJobParams{}
	res, err := c.SlurmV0040GetJobWithResponse(ctx, jobId, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	} else if len(res.JSON200.Jobs) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0040JobInfo{}
	utils.RemarshalOrDie(res.JSON200.Jobs[0], out)
	return out, nil
}

// ListJobInfo implements ClientInterface
func (c *SlurmClient) ListJobInfo(ctx context.Context) (*types.V0040JobInfoList, error) {
	params := &api.SlurmV0040GetJobsParams{}
	res, err := c.SlurmV0040GetJobsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode()))
	}
	list := &types.V0040JobInfoList{
		Items: make([]types.V0040JobInfo, len(res.JSON200.Jobs)),
	}
	for i, item := range res.JSON200.Jobs {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
