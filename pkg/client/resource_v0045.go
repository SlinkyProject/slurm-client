// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	v0045 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func init() {
	addResource(api.ObjectTypeV0045ControllerPing, resource{
		newObject: func() object.Object { return &api.V0045ControllerPing{} },
		newList:   func() object.ObjectList { return &api.V0045ControllerPingObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0045JobInfo, resource{
		newObject: func() object.Object { return &api.V0045JobInfo{} },
		newList:   func() object.ObjectList { return &api.V0045JobInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0045Node, resource{
		newObject: func() object.Object { return &api.V0045Node{} },
		newList:   func() object.ObjectList { return &api.V0045NodeObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0045NodeResourceLayoutList, resource{
		newObject: func() object.Object { return &api.V0045NodeResourceLayoutListObject{} },
	})
	addResource(api.ObjectTypeV0045PartitionInfo, resource{
		newObject: func() object.Object { return &api.V0045PartitionInfo{} },
		newList:   func() object.ObjectList { return &api.V0045PartitionInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0045ReservationInfo, resource{
		newObject: func() object.Object { return &api.V0045ReservationInfo{} },
		newList:   func() object.ObjectList { return &api.V0045ReservationInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0045StatsMsg, resource{
		newObject: func() object.Object { return &api.V0045StatsMsg{} },
		newList:   func() object.ObjectList { return &api.V0045StatsMsgObjectList{} },
		cacheable: true,
	})
}

func bindV0045(dst map[object.ObjectType]resource, c v0045.ClientInterface) {
	setOps(dst, api.ObjectTypeV0045ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045ControllerPingObjectList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0045JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045JobInfoObjectList, error) {
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
	setOps(dst, api.ObjectTypeV0045Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045NodeObjectList, error) {
				return c.ListNodes(ctx)
			})
		},
		create: func(ctx context.Context, req any) (object.ObjectKey, error) {
			return nodeKey(c.CreateNewNode(ctx, req))
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateNode(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteNode(ctx, key)
		},
	})
	setOps(dst, api.ObjectTypeV0045NodeResourceLayoutList, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045NodeResourceLayoutListObject, error) {
				nodes, err := c.GetNodeResourceLayout(ctx, string(key))
				if err != nil {
					return nil, err
				}
				out := api.V0045NodeResourceLayoutListObject(*nodes)
				return &out, nil
			})
		},
	})
	setOps(dst, api.ObjectTypeV0045PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045PartitionInfoObjectList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0045ReservationInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045ReservationInfo, error) {
				return c.GetReservationInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045ReservationInfoObjectList, error) {
				return c.ListReservationInfo(ctx)
			})
		},
		create: func(ctx context.Context, req any) (object.ObjectKey, error) {
			return nameKey(c.CreateReservationInfo(ctx, req))
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateReservationInfo(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteReservationInfo(ctx, key)
		},
	})
	setOps(dst, api.ObjectTypeV0045StatsMsg, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0045StatsMsg, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0045StatsMsgObjectList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
