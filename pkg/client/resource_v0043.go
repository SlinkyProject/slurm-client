// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	v0043 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func init() {
	addResource(types.ObjectTypeV0043ControllerPing, resource{
		newObject: func() object.Object { return &types.V0043ControllerPing{} },
		newList:   func() object.ObjectList { return &types.V0043ControllerPingList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0043JobInfo, resource{
		newObject: func() object.Object { return &types.V0043JobInfo{} },
		newList:   func() object.ObjectList { return &types.V0043JobInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0043Node, resource{
		newObject: func() object.Object { return &types.V0043Node{} },
		newList:   func() object.ObjectList { return &types.V0043NodeList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0043PartitionInfo, resource{
		newObject: func() object.Object { return &types.V0043PartitionInfo{} },
		newList:   func() object.ObjectList { return &types.V0043PartitionInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0043Reconfigure, resource{
		newObject: func() object.Object { return &types.V0043Reconfigure{} },
	})
	addResource(types.ObjectTypeV0043Stats, resource{
		newObject: func() object.Object { return &types.V0043Stats{} },
		newList:   func() object.ObjectList { return &types.V0043StatsList{} },
		cacheable: true,
	})
}

func bindV0043(dst map[object.ObjectType]resource, c v0043.ClientInterface) {
	setOps(dst, types.ObjectTypeV0043ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043ControllerPingList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0043JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043JobInfoList, error) {
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
	setOps(dst, types.ObjectTypeV0043Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043NodeList, error) {
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
	setOps(dst, types.ObjectTypeV0043PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043PartitionInfoList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0043Reconfigure, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043Reconfigure, error) {
				return c.GetReconfigure(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043ReconfigureList, error) {
				return c.ListReconfigure(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0043Stats, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0043Stats, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0043StatsList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
