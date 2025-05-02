// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type JobInfoInterface interface {
	CreateJobInfo(ctx context.Context, req any) (*int32, error)
	CreateJobInfoAlloc(ctx context.Context, req any) (*int32, error)
	DeleteJobInfo(ctx context.Context, jobId string) error
	UpdateJobInfo(ctx context.Context, jobId string, req any) error
	GetJobInfo(ctx context.Context, jobId string) (*types.V0042JobInfo, error)
	ListJobInfo(ctx context.Context) (*types.V0042JobInfoList, error)
}

// CreateJobInfo implements ClientInterface
func (c *SlurmClient) CreateJobInfo(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0042JobSubmitReq)
	if !ok {
		return nil, errors.New("expected req to be V0042JobSubmitReq")
	}
	body := api.SlurmV0042PostJobSubmitJSONRequestBody(r)
	res, err := c.SlurmV0042PostJobSubmitWithResponse(ctx, body)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	return res.JSON200.JobId, nil
}

// CreateJobInfoAlloc implements ClientInterface
func (c *SlurmClient) CreateJobInfoAlloc(ctx context.Context, req any) (*int32, error) {
	r, ok := req.(api.V0042JobAllocReq)
	if !ok {
		return nil, errors.New("expected req to be V0042JobAllocReq")
	}
	body := api.SlurmV0042PostJobAllocateJSONRequestBody(r)
	res, err := c.SlurmV0042PostJobAllocateWithResponse(ctx, body)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	return res.JSON200.JobId, nil
}

// DeleteJobInfo implements ClientInterface
func (c *SlurmClient) DeleteJobInfo(ctx context.Context, jobId string) error {
	params := &api.SlurmV0042DeleteJobParams{}
	res, err := c.SlurmV0042DeleteJobWithResponse(ctx, jobId, params)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return utilerrors.NewAggregate(errs)
	}
	return nil
}

// UpdateJobInfo implements ClientInterface
func (c *SlurmClient) UpdateJobInfo(ctx context.Context, jobId string, req any) error {
	r, ok := req.(api.V0042JobDescMsg)
	if !ok {
		return errors.New("expected req to be V0042JobDescMsg")
	}
	body := api.SlurmV0042PostJobJSONRequestBody(r)
	res, err := c.SlurmV0042PostJobWithResponse(ctx, jobId, body)
	if err != nil {
		return err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return utilerrors.NewAggregate(errs)
	}
	return nil
}

// GetJobInfo implements ClientInterface
func (c *SlurmClient) GetJobInfo(ctx context.Context, jobId string) (*types.V0042JobInfo, error) {
	params := &api.SlurmV0042GetJobParams{}
	res, err := c.SlurmV0042GetJobWithResponse(ctx, jobId, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return nil, utilerrors.NewAggregate(errs)
	} else if len(res.JSON200.Jobs) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0042JobInfo{}
	utils.RemarshalOrDie(res.JSON200.Jobs[0], out)
	return out, nil
}

// ListJobInfo implements ClientInterface
func (c *SlurmClient) ListJobInfo(ctx context.Context) (*types.V0042JobInfoList, error) {
	params := &api.SlurmV0042GetJobsParams{}
	res, err := c.SlurmV0042GetJobsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			for _, e := range ptr.Deref(res.JSONDefault.Errors, []api.V0042OpenapiError{}) {
				if e.Error != nil {
					errs = append(errs, errors.New(*e.Error))
				}
			}
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	list := &types.V0042JobInfoList{
		Items: make([]types.V0042JobInfo, len(res.JSON200.Jobs)),
	}
	for i, item := range res.JSON200.Jobs {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
