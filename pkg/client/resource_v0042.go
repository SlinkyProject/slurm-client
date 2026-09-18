// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	v0042 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0042"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func init() {
	addResource(api.ObjectTypeV0042ControllerPing, resource{
		newObject: func() object.Object { return &api.V0042ControllerPing{} },
		newList:   func() object.ObjectList { return &api.V0042ControllerPingObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0042JobInfo, resource{
		newObject: func() object.Object { return &api.V0042JobInfo{} },
		newList:   func() object.ObjectList { return &api.V0042JobInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0042Node, resource{
		newObject: func() object.Object { return &api.V0042Node{} },
		newList:   func() object.ObjectList { return &api.V0042NodeObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0042PartitionInfo, resource{
		newObject: func() object.Object { return &api.V0042PartitionInfo{} },
		newList:   func() object.ObjectList { return &api.V0042PartitionInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0042StatsMsg, resource{
		newObject: func() object.Object { return &api.V0042StatsMsg{} },
		newList:   func() object.ObjectList { return &api.V0042StatsMsgObjectList{} },
		cacheable: true,
	})
}

func bindV0042(dst map[object.ObjectType]resource, c v0042.ClientInterface) {
	setOps(dst, api.ObjectTypeV0042ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0042ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0042ControllerPingObjectList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0042JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0042JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0042JobInfoObjectList, error) {
				return c.ListJobInfo(ctx)
			})
		},
		create: func(ctx context.Context, obj object.Object, req any) (object.ObjectKey, error) {
			dst, ok := obj.(*api.V0042JobInfo)
			if !ok {
				return "", apierrors.ErrNotImplemented
			}
			identity, err := c.CreateJobInfo(ctx, req)
			key, err := jobKey(identity, err)
			if err == nil {
				dst.JobId = identity
			}
			return key, err
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateJobInfo(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteJobInfo(ctx, key)
		},
	})
	setOps(dst, api.ObjectTypeV0042Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0042Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0042NodeObjectList, error) {
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
	setOps(dst, api.ObjectTypeV0042PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0042PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0042PartitionInfoObjectList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0042StatsMsg, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0042StatsMsg, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0042StatsMsgObjectList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
