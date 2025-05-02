// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042PartitionInfo = "V0042PartitionInfo"
)

type V0042PartitionInfo struct {
	api.V0042PartitionInfo
}

// GetKey implements Object.
func (o *V0042PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0042PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0042PartitionInfo
}

// DeepCopyObject implements Object.
func (o *V0042PartitionInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042PartitionInfo) DeepCopy() *V0042PartitionInfo {
	out := new(V0042PartitionInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0042PartitionInfo) GetStateAsSet() set.Set[api.V0042PartitionInfoPartitionState] {
	out := make(set.Set[api.V0042PartitionInfoPartitionState])
	if o.Partition == nil {
		return out
	}
	states := ptr.Deref(o.Partition.State, []api.V0042PartitionInfoPartitionState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0042PartitionInfoList struct {
	Items []V0042PartitionInfo
}

// GetType implements ObjectList.
func (o *V0042PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0042PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0042PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042PartitionInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0042PartitionInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042PartitionInfoList)
	out.Items = make([]V0042PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
