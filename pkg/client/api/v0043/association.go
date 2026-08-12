// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0043

import (
	"context"
	"errors"
	"net/http"
	"strings"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type AssocInterface interface {
	CreateAssoc(ctx context.Context, req any) (string, error)
	UpdateAssoc(ctx context.Context, key string, req any) error
	DeleteAssoc(ctx context.Context, assoc api.V0043Assoc) error
	GetAssoc(ctx context.Context, key string) (*types.V0043Assoc, error)
	ListAssoc(ctx context.Context) (*types.V0043AssocList, error)
}

var _ AssocInterface = &SlurmClient{}

// parseAssocKey splits a "cluster/account/user/partition" association key.
func parseAssocKey(key string) (cluster, account, user, partition string) {
	parts := strings.SplitN(key, "/", 4)
	for len(parts) < 4 {
		parts = append(parts, "")
	}
	return parts[0], parts[1], parts[2], parts[3]
}

func nonEmptyPtr(s string) *string {
	if s == "" {
		return nil
	}
	return ptr.To(s)
}

// CreateAssoc implements AssocInterface (POST is upsert). It returns the
// composite key derived from the request so callers can refresh by identity.
func (c *SlurmClient) CreateAssoc(ctx context.Context, req any) (string, error) {
	r, ok := req.(api.V0043Assoc)
	if !ok {
		return "", errors.New("expected req to be api.V0043Assoc")
	}
	body := api.SlurmdbV0043PostAssociationsJSONRequestBody{
		Associations: api.V0043AssocList{r},
	}
	res, err := c.SlurmdbV0043PostAssociationsWithResponse(ctx, body)
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
	assoc := &types.V0043Assoc{V0043Assoc: r}
	return string(assoc.GetKey()), nil
}

// UpdateAssoc implements AssocInterface. POST is upsert in slurmdbd; the
// association identity is pinned from key so the update targets the intended
// association regardless of the identity fields set on req.
func (c *SlurmClient) UpdateAssoc(ctx context.Context, key string, req any) error {
	r, ok := req.(api.V0043Assoc)
	if !ok {
		return errors.New("expected req to be api.V0043Assoc")
	}
	cluster, account, user, partition := parseAssocKey(key)
	r.Cluster = nonEmptyPtr(cluster)
	r.Account = nonEmptyPtr(account)
	r.User = user
	r.Partition = nonEmptyPtr(partition)
	_, err := c.CreateAssoc(ctx, r)
	return err
}

// DeleteAssoc implements AssocInterface using a filter param set.
func (c *SlurmClient) DeleteAssoc(ctx context.Context, assoc api.V0043Assoc) error {
	params := &api.SlurmdbV0043DeleteAssociationParams{
		Account:   assoc.Account,
		User:      nonEmptyPtr(assoc.User),
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

// GetAssoc implements AssocInterface. There is no get-by-key endpoint, so the
// composite key is decomposed into server-side filter params to avoid listing
// every association.
func (c *SlurmClient) GetAssoc(ctx context.Context, key string) (*types.V0043Assoc, error) {
	cluster, account, user, partition := parseAssocKey(key)
	params := &api.SlurmdbV0043GetAssociationsParams{
		Account:   nonEmptyPtr(account),
		Cluster:   nonEmptyPtr(cluster),
		User:      nonEmptyPtr(user),
		Partition: nonEmptyPtr(partition),
	}
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
	if res.JSON200 == nil {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	for _, item := range res.JSON200.Associations {
		out := &types.V0043Assoc{}
		utils.RemarshalOrDie(item, out)
		if string(out.GetKey()) == key {
			return out, nil
		}
	}
	return nil, errors.New(http.StatusText(http.StatusNotFound))
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
