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

type UserInterface interface {
	CreateUser(ctx context.Context, req any) (string, error)
	UpdateUser(ctx context.Context, name string, req any) error
	DeleteUser(ctx context.Context, name string) error
	GetUser(ctx context.Context, name string) (*types.V0043User, error)
	ListUser(ctx context.Context) (*types.V0043UserList, error)
}

var _ UserInterface = &SlurmClient{}

// CreateUser implements UserInterface.
func (c *SlurmClient) CreateUser(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0043User)
	if !ok {
		return "", errors.New("expected req to be api.V0043User")
	}
	body := api.SlurmdbV0043PostUsersJSONRequestBody{
		Users: api.V0043UserList{r},
	}
	res, err := c.SlurmdbV0043PostUsersWithResponse(ctx, body)
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
	return r.Name, nil
}

// UpdateUser implements UserInterface. POST is upsert in slurmdbd.
func (c *SlurmClient) UpdateUser(ctx context.Context, name string, req any) error {
	r, ok := req.(api.V0043User)
	if !ok {
		return errors.New("expected req to be api.V0043User")
	}
	r.Name = name
	_, err := c.CreateUser(ctx, r)
	return err
}

// DeleteUser implements UserInterface.
func (c *SlurmClient) DeleteUser(ctx context.Context, name string) error {
	res, err := c.SlurmdbV0043DeleteUserWithResponse(ctx, name)
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

// GetUser implements UserInterface.
func (c *SlurmClient) GetUser(ctx context.Context, name string) (*types.V0043User, error) {
	params := &api.SlurmdbV0043GetUserParams{}
	res, err := c.SlurmdbV0043GetUserWithResponse(ctx, name, params)
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
	if res.JSON200 == nil || len(res.JSON200.Users) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0043User{}
	utils.RemarshalOrDie(res.JSON200.Users[0], out)
	return out, nil
}

// ListUser implements UserInterface.
func (c *SlurmClient) ListUser(ctx context.Context) (*types.V0043UserList, error) {
	params := &api.SlurmdbV0043GetUsersParams{}
	res, err := c.SlurmdbV0043GetUsersWithResponse(ctx, params)
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
	list := &types.V0043UserList{
		Items: make([]types.V0043User, len(res.JSON200.Users)),
	}
	for i, item := range res.JSON200.Users {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
