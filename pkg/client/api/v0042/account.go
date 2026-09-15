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

type AccountInterface interface {
	CreateAccount(ctx context.Context, req any) (string, error)
	UpdateAccount(ctx context.Context, name string, req any) error
	DeleteAccount(ctx context.Context, name string) error
	GetAccount(ctx context.Context, name string) (*types.V0042Account, error)
	ListAccount(ctx context.Context) (*types.V0042AccountList, error)
}

var _ AccountInterface = &SlurmClient{}

// CreateAccount implements AccountInterface.
func (c *SlurmClient) CreateAccount(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0042Account)
	if !ok {
		return "", errors.New("expected req to be api.V0042Account")
	}
	body := api.SlurmdbV0042PostAccountsJSONRequestBody{
		Accounts: api.V0042AccountList{r},
	}
	res, err := c.SlurmdbV0042PostAccountsWithResponse(ctx, body)
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

// UpdateAccount implements AccountInterface. POST is upsert in slurmdbd.
func (c *SlurmClient) UpdateAccount(ctx context.Context, name string, req any) error {
	r, ok := req.(api.V0042Account)
	if !ok {
		return errors.New("expected req to be api.V0042Account")
	}
	r.Name = name
	_, err := c.CreateAccount(ctx, r)
	return err
}

// DeleteAccount implements AccountInterface.
func (c *SlurmClient) DeleteAccount(ctx context.Context, name string) error {
	res, err := c.SlurmdbV0042DeleteAccountWithResponse(ctx, name)
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

// GetAccount implements AccountInterface.
func (c *SlurmClient) GetAccount(ctx context.Context, name string) (*types.V0042Account, error) {
	params := &api.SlurmdbV0042GetAccountParams{}
	res, err := c.SlurmdbV0042GetAccountWithResponse(ctx, name, params)
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
	if res.JSON200 == nil || len(res.JSON200.Accounts) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0042Account{}
	utils.RemarshalOrDie(res.JSON200.Accounts[0], out)
	return out, nil
}

// ListAccount implements AccountInterface.
func (c *SlurmClient) ListAccount(ctx context.Context) (*types.V0042AccountList, error) {
	params := &api.SlurmdbV0042GetAccountsParams{}
	res, err := c.SlurmdbV0042GetAccountsWithResponse(ctx, params)
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
	list := &types.V0042AccountList{
		Items: make([]types.V0042Account, len(res.JSON200.Accounts)),
	}
	for i, item := range res.JSON200.Accounts {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
