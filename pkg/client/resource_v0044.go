// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	v0044 "github.com/SlinkyProject/slurm-client/pkg/client/api/v0044"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func init() {
	addResource(api.ObjectTypeV0044ControllerPing, resource{
		newObject: func() object.Object { return &api.V0044ControllerPing{} },
		newList:   func() object.ObjectList { return &api.V0044ControllerPingObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0044JobInfo, resource{
		newObject: func() object.Object { return &api.V0044JobInfo{} },
		newList:   func() object.ObjectList { return &api.V0044JobInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0044Node, resource{
		newObject: func() object.Object { return &api.V0044Node{} },
		newList:   func() object.ObjectList { return &api.V0044NodeObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0044NodeResourceLayoutList, resource{
		newObject: func() object.Object { return &api.V0044NodeResourceLayoutListObject{} },
	})
	addResource(api.ObjectTypeV0044PartitionInfo, resource{
		newObject: func() object.Object { return &api.V0044PartitionInfo{} },
		newList:   func() object.ObjectList { return &api.V0044PartitionInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0044ReservationInfo, resource{
		newObject: func() object.Object { return &api.V0044ReservationInfo{} },
		newList:   func() object.ObjectList { return &api.V0044ReservationInfoObjectList{} },
		cacheable: true,
	})
	addResource(api.ObjectTypeV0044StatsMsg, resource{
		newObject: func() object.Object { return &api.V0044StatsMsg{} },
		newList:   func() object.ObjectList { return &api.V0044StatsMsgObjectList{} },
		cacheable: true,
	})
}

func bindV0044(dst map[object.ObjectType]resource, c v0044.ClientInterface) {
	setOps(dst, api.ObjectTypeV0044ControllerPing, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044ControllerPing, error) {
				return c.GetControllerPing(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044ControllerPingObjectList, error) {
				return c.ListControllerPing(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0044JobInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044JobInfo, error) {
				return c.GetJobInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044JobInfoObjectList, error) {
				return c.ListJobInfo(ctx)
			})
		},
		create: func(ctx context.Context, obj object.Object, req any) (object.ObjectKey, error) {
			dst, ok := obj.(*api.V0044JobInfo)
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
	setOps(dst, api.ObjectTypeV0044Node, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044Node, error) {
				return c.GetNode(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044NodeObjectList, error) {
				return c.ListNodes(ctx)
			})
		},
		create: func(ctx context.Context, obj object.Object, req any) (object.ObjectKey, error) {
			dst, ok := obj.(*api.V0044Node)
			if !ok {
				return "", apierrors.ErrNotImplemented
			}
			identity, err := c.CreateNewNode(ctx, req)
			key, err := nodeKey(identity, err)
			if err == nil {
				dst.Name = identity
			}
			return key, err
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateNode(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteNode(ctx, key)
		},
	})
	setOps(dst, api.ObjectTypeV0044NodeResourceLayoutList, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044NodeResourceLayoutListObject, error) {
				nodes, err := c.GetNodeResourceLayout(ctx, string(key))
				if err != nil {
					return nil, err
				}
				out := api.V0044NodeResourceLayoutListObject(*nodes)
				return &out, nil
			})
		},
	})
	setOps(dst, api.ObjectTypeV0044PartitionInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044PartitionInfo, error) {
				return c.GetPartitionInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044PartitionInfoObjectList, error) {
				return c.ListPartitionInfo(ctx)
			})
		},
	})
	setOps(dst, api.ObjectTypeV0044ReservationInfo, resource{
		get: func(ctx context.Context, key object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044ReservationInfo, error) {
				return c.GetReservationInfo(ctx, string(key))
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044ReservationInfoObjectList, error) {
				return c.ListReservationInfo(ctx)
			})
		},
		create: func(ctx context.Context, obj object.Object, req any) (object.ObjectKey, error) {
			dst, ok := obj.(*api.V0044ReservationInfo)
			if !ok {
				return "", apierrors.ErrNotImplemented
			}
			identity, err := c.CreateReservationInfo(ctx, req)
			key, err := nameKey(identity, err)
			if err == nil {
				dst.Name = ptr.To(identity)
			}
			return key, err
		},
		update: func(ctx context.Context, key string, req any) error {
			return c.UpdateReservationInfo(ctx, key, req)
		},
		delete: func(ctx context.Context, key string) error {
			return c.DeleteReservationInfo(ctx, key)
		},
	})
	setOps(dst, api.ObjectTypeV0044StatsMsg, resource{
		get: func(ctx context.Context, _ object.ObjectKey, obj object.Object) error {
			return getInto(obj, func() (*api.V0044StatsMsg, error) {
				return c.GetStats(ctx)
			})
		},
		list: func(ctx context.Context, list object.ObjectList) error {
			return listInto(list, func() (*api.V0044StatsMsgObjectList, error) {
				return c.ListStats(ctx)
			})
		},
	})
}
