// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type AssocInterface interface {
	CreateAssoc(ctx context.Context, req any) error
	DeleteAssoc(ctx context.Context, assoc api.V0043Assoc) error
	ListAssoc(ctx context.Context) (*types.V0043AssocList, error)
}

var _ AssocInterface = &SlurmClient{}

// CreateAssoc implements AssocInterface (POST is upsert).
func (c *SlurmClient) CreateAssoc(ctx context.Context, req any) error {
	r, ok := req.(api.V0043Assoc)
	if !ok {
		return errors.New("expected req to be api.V0043Assoc")
	}
	body := api.SlurmdbV0043PostAssociationsJSONRequestBody{
		Associations: api.V0043AssocList{r},
	}
	res, err := c.SlurmdbV0043PostAssociationsWithResponse(ctx, body)
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
func (c *SlurmClient) DeleteAssoc(ctx context.Context, assoc api.V0043Assoc) error {
	params := &api.SlurmdbV0043DeleteAssociationParams{
		Account:   assoc.Account,
		User:      ptr.To(assoc.User),
		Cluster:   assoc.Cluster,
		Partition: assoc.Partition,
	}
	res, err := c.SlurmdbV0043DeleteAssociationWithResponse(ctx, params)
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
func (c *SlurmClient) ListAssoc(ctx context.Context) (*types.V0043AssocList, error) {
	params := &api.SlurmdbV0043GetAssociationsParams{}
	res, err := c.SlurmdbV0043GetAssociationsWithResponse(ctx, params)
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
	list := &types.V0043AssocList{
		Items: make([]types.V0043Assoc, len(res.JSON200.Associations)),
	}
	for i, item := range res.JSON200.Associations {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
