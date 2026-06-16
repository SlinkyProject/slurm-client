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

type ReservationInterface interface {
	CreateReservationInfo(ctx context.Context, req any) (string, error)
	UpdateReservationInfo(ctx context.Context, name string, req any) error
	DeleteReservationInfo(ctx context.Context, name string) error
	GetReservationInfo(ctx context.Context, name string) (*types.V0045ReservationInfo, error)
	ListReservationInfo(ctx context.Context) (*types.V0045ReservationInfoList, error)
}

var _ ReservationInterface = &SlurmClient{}

// CreateReservationInfo implements ClientInterface
func (c *SlurmClient) CreateReservationInfo(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0045ReservationDescMsg)
	if !ok {
		return "", errors.New("expected req to be V0045ReservationDescMsg")
	}

	body := api.SlurmV0045PostReservationJSONRequestBody(r)
	res, err := c.SlurmV0045PostReservationWithResponse(ctx, body)
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

	return *r.Name, nil
}

// DeleteReservationInfo implements ClientInterface
func (c *SlurmClient) DeleteReservationInfo(ctx context.Context, name string) error {
	res, err := c.SlurmV0045DeleteReservationWithResponse(ctx, name)
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

// UpdateReservationInfo implements ClientInterface
func (c *SlurmClient) UpdateReservationInfo(ctx context.Context, name string, req any) error {
	r, ok := req.(api.V0045ReservationDescMsg)
	if !ok {
		return errors.New("expected req to be V0045ReservationDescMsg")
	}

	// endpoint does not use ID parameter, but make it uniform with the rest that do
	r.Name = &name

	body := api.SlurmV0045PostReservationJSONRequestBody(r)
	res, err := c.SlurmV0045PostReservationWithResponse(ctx, body)
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

// GetReservationInfo implements ClientInterface
func (c *SlurmClient) GetReservationInfo(ctx context.Context, name string) (*types.V0045ReservationInfo, error) {
	params := &api.SlurmV0045GetReservationParams{}
	res, err := c.SlurmV0045GetReservationWithResponse(ctx, name, params)
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

	if len(res.JSON200.Reservations) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0045ReservationInfo{}
	utils.RemarshalOrDie(res.JSON200.Reservations[0], out)
	return out, nil
}

// ListReservationInfo implements ClientInterface
func (c *SlurmClient) ListReservationInfo(ctx context.Context) (*types.V0045ReservationInfoList, error) {
	params := &api.SlurmV0045GetReservationsParams{}
	res, err := c.SlurmV0045GetReservationsWithResponse(ctx, params)
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

	list := &types.V0045ReservationInfoList{
		Items: make([]types.V0045ReservationInfo, len(res.JSON200.Reservations)),
	}
	for i, item := range res.JSON200.Reservations {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
