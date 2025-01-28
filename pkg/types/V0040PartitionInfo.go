// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0040PartitionInfo = "V0040PartitionInfo"
)

type V0040PartitionInfo struct {
	api.V0040PartitionInfo
}

// GetKey implements Object.
func (o *V0040PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0040PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0040PartitionInfo
}

// DeepCopyObject implements Object.
func (o *V0040PartitionInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0040PartitionInfo) DeepCopy() *V0040PartitionInfo {
	out := new(V0040PartitionInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0040PartitionInfo) GetStateAsSet() set.Set[api.V0040PartitionInfoPartitionState] {
	out := make(set.Set[api.V0040PartitionInfoPartitionState])
	if o.Partition == nil {
		return out
	}
	states := ptr.Deref(o.Partition.State, []api.V0040PartitionInfoPartitionState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0040PartitionInfoList struct {
	Items []V0040PartitionInfo
}

// GetType implements ObjectList.
func (o *V0040PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0040PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0040PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0040PartitionInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0040PartitionInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0040PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0040PartitionInfoList)
	out.Items = make([]V0040PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
