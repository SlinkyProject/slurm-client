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

type PartitionInterface interface {
	GetPartitionInfo(ctx context.Context, name string) (*types.V0042PartitionInfo, error)
	ListPartitionInfo(ctx context.Context) (*types.V0042PartitionInfoList, error)
}

// GetPartitionInfo implements ClientInterface
func (c *SlurmClient) GetPartitionInfo(ctx context.Context, name string) (*types.V0042PartitionInfo, error) {
	params := &api.SlurmV0042GetPartitionParams{}
	res, err := c.SlurmV0042GetPartitionWithResponse(ctx, name, params)
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
	} else if len(res.JSON200.Partitions) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	out := &types.V0042PartitionInfo{}
	utils.RemarshalOrDie(res.JSON200.Partitions[0], out)
	return out, nil
}

// ListPartitionInfo implements ClientInterface
func (c *SlurmClient) ListPartitionInfo(ctx context.Context) (*types.V0042PartitionInfoList, error) {
	params := &api.SlurmV0042GetPartitionsParams{}
	res, err := c.SlurmV0042GetPartitionsWithResponse(ctx, params)
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
	list := &types.V0042PartitionInfoList{
		Items: make([]types.V0042PartitionInfo, len(res.JSON200.Partitions)),
	}
	for i, item := range res.JSON200.Partitions {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
