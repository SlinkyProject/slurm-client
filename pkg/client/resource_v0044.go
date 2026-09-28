// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	v0044 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func init() {
	addResource(types.ObjectTypeV0044ControllerPing, resource{
		newObject: func() object.Object { return &types.V0044ControllerPing{} },
		newList:   func() object.ObjectList { return &types.V0044ControllerPingList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0044JobInfo, resource{
		newObject: func() object.Object { return &types.V0044JobInfo{} },
		newList:   func() object.ObjectList { return &types.V0044JobInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0044Node, resource{
		newObject: func() object.Object { return &types.V0044Node{} },
		newList:   func() object.ObjectList { return &types.V0044NodeList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0044NodeResourceLayout, resource{
		newObject: func() object.Object { return &types.V0044NodeResourceLayout{} },
	})
	addResource(types.ObjectTypeV0044PartitionInfo, resource{
		newObject: func() object.Object { return &types.V0044PartitionInfo{} },
		newList:   func() object.ObjectList { return &types.V0044PartitionInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0044Reconfigure, resource{
		newObject: func() object.Object { return &types.V0044Reconfigure{} },
	})
	addResource(types.ObjectTypeV0044ReservationInfo, resource{
		newObject: func() object.Object { return &types.V0044ReservationInfo{} },
		newList:   func() object.ObjectList { return &types.V0044ReservationInfoList{} },
		cacheable: true,
	})
	addResource(types.ObjectTypeV0044Stats, resource{
		newObject: func() object.Object { return &types.V0044Stats{} },
		newList:   func() object.ObjectList { return &types.V0044StatsList{} },
		cacheable: true,
	})
}

func bindV0044(dst map[object.ObjectType]resource, c v0044.ClientInterface) {
	setOps(dst, types.ObjectTypeV0044ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044ControllerPingList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0044JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044JobInfoList, error) {
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
	setOps(dst, types.ObjectTypeV0044Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044NodeList, error) {
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
	setOps(dst, types.ObjectTypeV0044NodeResourceLayout, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044NodeResourceLayout, error) {
				return c.GetNodeResourceLayout(ctx, string(key))
			})
		},
	})
	setOps(dst, types.ObjectTypeV0044PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044PartitionInfoList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0044Reconfigure, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044Reconfigure, error) {
				return c.GetReconfigure(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044ReconfigureList, error) {
				return c.ListReconfigure(ctx)
			})
		},
	})
	setOps(dst, types.ObjectTypeV0044ReservationInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044ReservationInfo, error) {
				return c.GetReservationInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044ReservationInfoList, error) {
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
	setOps(dst, types.ObjectTypeV0044Stats, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*types.V0044Stats, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*types.V0044StatsList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
