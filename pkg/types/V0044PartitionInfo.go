// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044PartitionInfo = "V0044PartitionInfo"
)

type V0044PartitionInfo struct {
	api.V0044PartitionInfo
}

// GetKey implements Object.
func (o *V0044PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0044PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0044PartitionInfo
}

// DeepCopyObject implements Object.
func (o *V0044PartitionInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044PartitionInfo) DeepCopy() *V0044PartitionInfo {
	out := new(V0044PartitionInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0044PartitionInfo) GetStateAsSet() set.Set[api.V0044PartitionInfoPartitionState] {
	out := make(set.Set[api.V0044PartitionInfoPartitionState])
	if o.Partition == nil {
		return out
	}
	states := ptr.Deref(o.Partition.State, []api.V0044PartitionInfoPartitionState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0044PartitionInfoList struct {
	Items []V0044PartitionInfo
}

// GetType implements ObjectList.
func (o *V0044PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0044PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0044PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044PartitionInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0044PartitionInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044PartitionInfoList)
	out.Items = make([]V0044PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
