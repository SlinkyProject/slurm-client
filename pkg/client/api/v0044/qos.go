// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type QosInterface interface {
	CreateQos(ctx context.Context, req any) (string, error)
	UpdateQos(ctx context.Context, name string, req any) error
	DeleteQos(ctx context.Context, name string) error
	GetQos(ctx context.Context, name string) (*types.V0044Qos, error)
	ListQos(ctx context.Context) (*types.V0044QosList, error)
}

var _ QosInterface = &SlurmClient{}

// CreateQos implements ClientInterface
func (c *SlurmClient) CreateQos(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0044OpenapiSlurmdbdQosResp)
	if !ok {
		return "", errors.New("expected req to be V0044OpenapiSlurmdbdQosResp")
	}
	if len(r.Qos) == 0 {
		return "", errors.New("expected req to contain at least one QOS entry")
	}

	params := &api.SlurmdbV0044PostQosParams{}
	body := api.SlurmdbV0044PostQosJSONRequestBody(r)
	res, err := c.SlurmdbV0044PostQosWithResponse(ctx, params, body)
	if err != nil {
		return "", err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return "", utilerrors.NewAggregate(errs)
	}

	var name string
	if len(r.Qos) > 0 {
		name = ptr.Deref(r.Qos[0].Name, "")
	}
	return name, nil
}

// DeleteQos implements ClientInterface
func (c *SlurmClient) DeleteQos(ctx context.Context, name string) error {
	res, err := c.SlurmdbV0044DeleteSingleQosWithResponse(ctx, name)
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

// UpdateQos implements ClientInterface
func (c *SlurmClient) UpdateQos(ctx context.Context, name string, req any) error {
	r, ok := req.(api.V0044OpenapiSlurmdbdQosResp)
	if !ok {
		return errors.New("expected req to be V0044OpenapiSlurmdbdQosResp")
	}
	if len(r.Qos) == 0 {
		return errors.New("expected req to contain at least one QOS entry")
	}

	r.Qos[0].Name = &name

	params := &api.SlurmdbV0044PostQosParams{}
	body := api.SlurmdbV0044PostQosJSONRequestBody(r)
	res, err := c.SlurmdbV0044PostQosWithResponse(ctx, params, body)
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

// GetQos implements ClientInterface
func (c *SlurmClient) GetQos(ctx context.Context, name string) (*types.V0044Qos, error) {
	params := &api.SlurmdbV0044GetSingleQosParams{}
	res, err := c.SlurmdbV0044GetSingleQosWithResponse(ctx, name, params)
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

	if len(res.JSON200.Qos) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0044Qos{}
	utils.RemarshalOrDie(res.JSON200.Qos[0], out)
	return out, nil
}

// ListQos implements ClientInterface
func (c *SlurmClient) ListQos(ctx context.Context) (*types.V0044QosList, error) {
	params := &api.SlurmdbV0044GetQosParams{}
	res, err := c.SlurmdbV0044GetQosWithResponse(ctx, params)
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

	list := &types.V0044QosList{
		Items: make([]types.V0044Qos, len(res.JSON200.Qos)),
	}
	for i, item := range res.JSON200.Qos {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
