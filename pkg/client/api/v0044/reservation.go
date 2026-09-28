// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
)

type ReservationInterface interface {
	CreateReservationInfo(ctx context.Context, req any) (string, error)
	UpdateReservationInfo(ctx context.Context, name string, req any) error
	DeleteReservationInfo(ctx context.Context, name string) error
	GetReservationInfo(ctx context.Context, name string) (*api.V0044ReservationInfo, error)
	ListReservationInfo(ctx context.Context) (*api.V0044ReservationInfoObjectList, error)
}

var _ ReservationInterface = &SlurmClient{}

// CreateReservationInfo implements ClientInterface
func (c *SlurmClient) CreateReservationInfo(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0044ReservationDescMsg)
	if !ok {
		return "", errors.New("expected req to be V0044ReservationDescMsg")
	}

	body := api.SlurmV0044PostReservationJSONRequestBody(r)
	res, err := c.SlurmV0044PostReservationWithResponse(ctx, body)
	if err != nil {
		return "", err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return "", errors.Join(errs...)
	}

	return *r.Name, nil
}

// DeleteReservationInfo implements ClientInterface
func (c *SlurmClient) DeleteReservationInfo(ctx context.Context, name string) error {
	res, err := c.SlurmV0044DeleteReservationWithResponse(ctx, name)
	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return errors.Join(errs...)
	}

	return nil
}

// UpdateReservationInfo implements ClientInterface
func (c *SlurmClient) UpdateReservationInfo(ctx context.Context, name string, req any) error {
	r, ok := req.(api.V0044ReservationDescMsg)
	if !ok {
		return errors.New("expected req to be V0044ReservationDescMsg")
	}

	// endpoint does not use ID parameter, but make it uniform with the rest that do
	r.Name = &name

	body := api.SlurmV0044PostReservationJSONRequestBody(r)
	res, err := c.SlurmV0044PostReservationWithResponse(ctx, body)
	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return errors.Join(errs...)
	}

	return nil
}

// GetReservationInfo implements ClientInterface
func (c *SlurmClient) GetReservationInfo(ctx context.Context, name string) (*api.V0044ReservationInfo, error) {
	params := &api.SlurmV0044GetReservationParams{}
	res, err := c.SlurmV0044GetReservationWithResponse(ctx, name, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{apierrors.NewHTTPError(res.StatusCode())}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	if len(res.JSON200.Reservations) == 0 {
		return nil, apierrors.ErrNotFound
	}

	out := res.JSON200.Reservations[0]
	return &out, nil
}

// ListReservationInfo implements ClientInterface
func (c *SlurmClient) ListReservationInfo(ctx context.Context) (*api.V0044ReservationInfoObjectList, error) {
	params := &api.SlurmV0044GetReservationsParams{}
	res, err := c.SlurmV0044GetReservationsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, errors.Join(errs...)
	}

	return &api.V0044ReservationInfoObjectList{Items: res.JSON200.Reservations}, nil
}
