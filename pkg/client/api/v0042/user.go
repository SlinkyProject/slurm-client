// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type UserInterface interface {
	CreateUser(ctx context.Context, req any) (string, error)
	UpdateUser(ctx context.Context, name string, req any) error
	DeleteUser(ctx context.Context, name string) error
	GetUser(ctx context.Context, name string) (*types.V0042User, error)
	ListUser(ctx context.Context) (*types.V0042UserList, error)
}

var _ UserInterface = &SlurmClient{}

// CreateUser implements UserInterface.
func (c *SlurmClient) CreateUser(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0042User)
	if !ok {
		return "", errors.New("expected req to be api.V0042User")
	}
	body := api.SlurmdbV0042PostUsersJSONRequestBody{
		Users: api.V0042UserList{r},
	}
	res, err := c.SlurmdbV0042PostUsersWithResponse(ctx, body)
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
	r, ok := req.(api.V0042User)
	if !ok {
		return errors.New("expected req to be api.V0042User")
	}
	r.Name = name
	_, err := c.CreateUser(ctx, r)
	return err
}

// DeleteUser implements UserInterface.
func (c *SlurmClient) DeleteUser(ctx context.Context, name string) error {
	res, err := c.SlurmdbV0042DeleteUserWithResponse(ctx, name)
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
func (c *SlurmClient) GetUser(ctx context.Context, name string) (*types.V0042User, error) {
	params := &api.SlurmdbV0042GetUserParams{}
	res, err := c.SlurmdbV0042GetUserWithResponse(ctx, name, params)
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
	out := &types.V0042User{}
	utils.RemarshalOrDie(res.JSON200.Users[0], out)
	return out, nil
}

// ListUser implements UserInterface.
func (c *SlurmClient) ListUser(ctx context.Context) (*types.V0042UserList, error) {
	params := &api.SlurmdbV0042GetUsersParams{}
	res, err := c.SlurmdbV0042GetUsersWithResponse(ctx, params)
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
	list := &types.V0042UserList{
		Items: make([]types.V0042User, len(res.JSON200.Users)),
	}
	for i, item := range res.JSON200.Users {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
