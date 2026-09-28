// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	v0045 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func init() {
	addResource(types.ObjectTypeV0045ControllerPing, resource{
		newObject: func() object.Object { return &types.V0045ControllerPing{} },
		newList:   func() object.ObjectList { return &types.V0045ControllerPingList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0045JobInfo, resource{
		newObject: func() object.Object { return &types.V0045JobInfo{} },
		newList:   func() object.ObjectList { return &types.V0045JobInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0045Node, resource{
		newObject: func() object.Object { return &types.V0045Node{} },
		newList:   func() object.ObjectList { return &types.V0045NodeList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0045NodeResourceLayout, resource{
		newObject: func() object.Object { return &types.V0045NodeResourceLayout{} },
	})
	addResource(types.ObjectTypeV0045PartitionInfo, resource{
		newObject: func() object.Object { return &types.V0045PartitionInfo{} },
		newList:   func() object.ObjectList { return &types.V0045PartitionInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0045Reconfigure, resource{
		newObject: func() object.Object { return &types.V0045Reconfigure{} },
	})
	addResource(types.ObjectTypeV0045ReservationInfo, resource{
		newObject: func() object.Object { return &types.V0045ReservationInfo{} },
		newList:   func() object.ObjectList { return &types.V0045ReservationInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0045Stats, resource{
		newObject: func() object.Object { return &types.V0045Stats{} },
		newList:   func() object.ObjectList { return &types.V0045StatsList{} },
		cacheable: true,
	})
}

func bindV0045(dst map[object.ObjectType]resource, c v0045.ClientInterface) {
	setOps(dst, types.ObjectTypeV0045ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045ControllerPingList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0045JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045JobInfoList, error) {
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
	setOps(dst, types.ObjectTypeV0045Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045NodeList, error) {
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
	setOps(dst, types.ObjectTypeV0045NodeResourceLayout, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045NodeResourceLayout, error) {
				return c.GetNodeResourceLayout(ctx, string(key))
			})
		},
	})
	setOps(dst, types.ObjectTypeV0045PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045PartitionInfoList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0045Reconfigure, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045Reconfigure, error) {
				return c.GetReconfigure(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045ReconfigureList, error) {
				return c.ListReconfigure(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0045ReservationInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045ReservationInfo, error) {
				return c.GetReservationInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045ReservationInfoList, error) {
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
	setOps(dst, types.ObjectTypeV0045Stats, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0045Stats, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0045StatsList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
