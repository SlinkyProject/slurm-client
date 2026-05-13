// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type JobInfoInterface interface {
	CreateJobInfo(ctx context.Context, req any) (*int32, error)
	DeleteJobInfo(ctx context.Context, jobId string) error
	UpdateJobInfo(ctx context.Context, jobId string, req any) error
	GetJobInfo(ctx context.Context, jobId string) (*types.V0045JobInfo, error)
	ListJobInfo(ctx context.Context) (*types.V0045JobInfoList, error)
}

var _ JobInfoInterface = &SlurmClient{}

// CreateJobInfo implements ClientInterface
func (c *SlurmClient) CreateJobInfo(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0045JobSubmitReq)
	if !ok {
		return nil, errors.New("expected req to be V0045JobSubmitReq")
	}

	body := api.SlurmV0045PostJobSubmitJSONRequestBody(r)
	res, err := c.SlurmV0045PostJobSubmitWithResponse(ctx, body)
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
	params := &api.SlurmV0045DeleteJobParams{}
	res, err := c.SlurmV0045DeleteJobWithResponse(ctx, jobId, params)
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
	r, ok := req.(api.V0045JobDescMsg)
	if !ok {
		return errors.New("expected req to be V0045JobDescMsg")
	}

	body := api.SlurmV0045PostJobJSONRequestBody(r)
	res, err := c.SlurmV0045PostJobWithResponse(ctx, jobId, body)
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
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*types.V0045JobInfo, error) {
	params := &api.SlurmV0045GetJobParams{}
	res, err := c.SlurmV0045GetJobWithResponse(ctx, jobId, params)
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

	out := &types.V0045JobInfo{}
	utils.RemarshalOrDie(res.JSON200.Jobs[0], out)
	return out, nil
}

// ListJobInfo implements ClientInterface
func (c *SlurmClient) ListJobInfo(ctx context.Context) (*types.V0045JobInfoList, error) {
	params := &api.SlurmV0045GetJobsParams{}
	res, err := c.SlurmV0045GetJobsWithResponse(ctx, params)
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

	list := &types.V0045JobInfoList{
		Items: make([]types.V0045JobInfo, len(res.JSON200.Jobs)),
	}
	for i, item := range res.JSON200.Jobs {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
