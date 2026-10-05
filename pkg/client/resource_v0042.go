// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	api "github.com/SlinkyProject/slurm-client/pkg/client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func init() {
	addResource(types.ObjectTypeV0042ControllerPing, resource{
		newObject: func() object.Object { return &types.V0042ControllerPing{} },
		newList:   func() object.ObjectList { return &types.V0042ControllerPingList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0042JobInfo, resource{
		newObject: func() object.Object { return &types.V0042JobInfo{} },
		newList:   func() object.ObjectList { return &types.V0042JobInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0042Node, resource{
		newObject: func() object.Object { return &types.V0042Node{} },
		newList:   func() object.ObjectList { return &types.V0042NodeList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0042PartitionInfo, resource{
		newObject: func() object.Object { return &types.V0042PartitionInfo{} },
		newList:   func() object.ObjectList { return &types.V0042PartitionInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0042Stats, resource{
		newObject: func() object.Object { return &types.V0042Stats{} },
		newList:   func() object.ObjectList { return &types.V0042StatsList{} },
		cacheable: true,
	})
}

func bindV0042(dst map[object.ObjectType]resource, c api.ClientInterface) {
	setOps(dst, types.ObjectTypeV0042ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0042ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0042ControllerPingList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0042JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0042JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0042JobInfoList, error) {
				return c.ListJobInfo(ctx)
			})
		},
		create: func(ctx context.Context, req any) (object.ObjectKey, error) {
			return jobKey(c.CreateJobInfo(ctx, req))
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateJobInfo(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteJobInfo(ctx, key)
		},
	})
	setOps(dst, types.ObjectTypeV0042Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0042Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0042NodeList, error) {
				return c.ListNodes(ctx)
			})
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateNode(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteNode(ctx, key)
		},
	})
	setOps(dst, types.ObjectTypeV0042PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0042PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0042PartitionInfoList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0042Stats, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0042Stats, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0042StatsList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
