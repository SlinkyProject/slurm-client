// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043PartitionInfo = "V0043PartitionInfo"
)

type V0043PartitionInfo struct {
	api.V0043PartitionInfo
}

// GetKey implements Object.
func (o *V0043PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0043PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0043PartitionInfo
}

// DeepCopyObject implements Object.
func (o *V0043PartitionInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043PartitionInfo) DeepCopy() *V0043PartitionInfo {
	out := new(V0043PartitionInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0043PartitionInfo) GetStateAsSet() set.Set[api.V0043PartitionInfoPartitionState] {
	out := make(set.Set[api.V0043PartitionInfoPartitionState])
	if o.Partition == nil {
		return out
	}
	states := ptr.Deref(o.Partition.State, []api.V0043PartitionInfoPartitionState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0043PartitionInfoList struct {
	Items []V0043PartitionInfo
}

// GetType implements ObjectList.
func (o *V0043PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0043PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0043PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043PartitionInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0043PartitionInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043PartitionInfoList)
	out.Items = make([]V0043PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
