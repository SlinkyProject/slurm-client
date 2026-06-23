// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type AssocInterface interface {
	CreateAssoc(ctx context.Context, req any) error
	DeleteAssoc(ctx context.Context, assoc api.V0045Assoc) error
	ListAssoc(ctx context.Context) (*types.V0045AssocList, error)
}

var _ AssocInterface = &SlurmClient{}

// CreateAssoc implements AssocInterface (POST is upsert).
func (c *SlurmClient) CreateAssoc(ctx context.Context, req any) error {
	r, ok := req.(api.V0045Assoc)
	if !ok {
		return errors.New("expected req to be api.V0045Assoc")
	}
	body := api.SlurmdbV0045PostAssociationsJSONRequestBody{
		Associations: api.V0045AssocList{r},
	}
	res, err := c.SlurmdbV0045PostAssociationsWithResponse(ctx, body)
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

// DeleteAssoc implements AssocInterface using a filter param set.
func (c *SlurmClient) DeleteAssoc(ctx context.Context, assoc api.V0045Assoc) error {
	params := &api.SlurmdbV0045DeleteAssociationParams{
		Account:   assoc.Account,
		User:      ptr.To(assoc.User),
		Cluster:   assoc.Cluster,
		Partition: assoc.Partition,
	}
	res, err := c.SlurmdbV0045DeleteAssociationWithResponse(ctx, params)
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

// ListAssoc implements AssocInterface.
func (c *SlurmClient) ListAssoc(ctx context.Context) (*types.V0045AssocList, error) {
	params := &api.SlurmdbV0045GetAssociationsParams{}
	res, err := c.SlurmdbV0045GetAssociationsWithResponse(ctx, params)
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
	list := &types.V0045AssocList{
		Items: make([]types.V0045Assoc, len(res.JSON200.Associations)),
	}
	for i, item := range res.JSON200.Associations {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
