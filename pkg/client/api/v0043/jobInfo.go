// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type JobInfoInterface interface {
	CreateJobInfo(ctx context.Context, req any) (*int32, error)
	CreateJobInfoAlloc(ctx context.Context, req any) (*int32, error)
	DeleteJobInfo(ctx context.Context, jobId string) error
	UpdateJobInfo(ctx context.Context, jobId string, req any) error
	GetJobInfo(ctx context.Context, jobId string) (*types.V0043JobInfo, error)
	ListJobInfo(ctx context.Context) (*types.V0043JobInfoList, error)
}

// CreateJobInfo implements ClientInterface
func (c *SlurmClient) CreateJobInfo(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0043JobSubmitReq)
	if !ok {
		return nil, errors.New("expected req to be V0043JobSubmitReq")
	}

	body := api.SlurmV0043PostJobSubmitJSONRequestBody(r)
	res, err := c.SlurmV0043PostJobSubmitWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	return res.JSON200.JobId, nil
}

// CreateJobInfoAlloc implements ClientInterface
func (c *SlurmClient) CreateJobInfoAlloc(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0043JobAllocReq)
	if !ok {
		return nil, errors.New("expected req to be V0043JobAllocReq")
	}

	body := api.SlurmV0043PostJobAllocateJSONRequestBody(r)
	res, err := c.SlurmV0043PostJobAllocateWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	return res.JSON200.JobId, nil
}

// DeleteJobInfo implements ClientInterface
func (c *SlurmClient) DeleteJobInfo(ctx context.Context, jobId string) error {
	params := &api.SlurmV0043DeleteJobParams{}
	res, err := c.SlurmV0043DeleteJobWithResponse(ctx, jobId, params)
	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return utilerrors.NewAggregate(errs)
	}

	return nil
}

// UpdateJobInfo implements ClientInterface
func (c *SlurmClient) UpdateJobInfo(ctx context.Context, jobId string, req any) error {
	r, ok := req.(api.V0043JobDescMsg)
	if !ok {
		return errors.New("expected req to be V0043JobDescMsg")
	}

	body := api.SlurmV0043PostJobJSONRequestBody(r)
	res, err := c.SlurmV0043PostJobWithResponse(ctx, jobId, body)
	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return utilerrors.NewAggregate(errs)
	}

	return nil
}

// GetJobInfo implements ClientInterface
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*types.V0043JobInfo, error) {
	params := &api.SlurmV0043GetJobParams{}
	res, err := c.SlurmV0043GetJobWithResponse(ctx, jobId, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	if len(res.JSON200.Jobs) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0043JobInfo{}
	utils.RemarshalOrDie(res.JSON200.Jobs[0], out)
	return out, nil
}

// ListJobInfo implements ClientInterface
func (c *SlurmClient) ListJobInfo(ctx context.Context) (*types.V0043JobInfoList, error) {
	params := &api.SlurmV0043GetJobsParams{}
	res, err := c.SlurmV0043GetJobsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	list := &types.V0043JobInfoList{
		Items: make([]types.V0043JobInfo, len(res.JSON200.Jobs)),
	}
	for i, item := range res.JSON200.Jobs {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
