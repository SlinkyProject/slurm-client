// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	v0043 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func init() {
	addResource(api.ObjectTypeV0043ControllerPing, resource{
		newObject: func() object.Object { return &api.V0043ControllerPing{} },
		newList:   func() object.ObjectList { return &api.V0043ControllerPingObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0043JobInfo, resource{
		newObject: func() object.Object { return &api.V0043JobInfo{} },
		newList:   func() object.ObjectList { return &api.V0043JobInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0043Node, resource{
		newObject: func() object.Object { return &api.V0043Node{} },
		newList:   func() object.ObjectList { return &api.V0043NodeObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0043PartitionInfo, resource{
		newObject: func() object.Object { return &api.V0043PartitionInfo{} },
		newList:   func() object.ObjectList { return &api.V0043PartitionInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0043StatsMsg, resource{
		newObject: func() object.Object { return &api.V0043StatsMsg{} },
		newList:   func() object.ObjectList { return &api.V0043StatsMsgObjectList{} },
		cacheable: true,
	})
}

func bindV0043(dst map[object.ObjectType]resource, c v0043.ClientInterface) {
	setOps(dst, api.ObjectTypeV0043ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0043ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0043ControllerPingObjectList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0043JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0043JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0043JobInfoObjectList, error) {
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
	setOps(dst, api.ObjectTypeV0043Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0043Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0043NodeObjectList, error) {
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
	setOps(dst, api.ObjectTypeV0043PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0043PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0043PartitionInfoObjectList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0043StatsMsg, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0043StatsMsg, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0043StatsMsgObjectList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
