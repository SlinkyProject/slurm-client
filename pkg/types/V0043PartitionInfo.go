// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

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

// DeepCopyObject implements RuntimeObject.
func (o *V0043PartitionInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043PartitionInfo) DeepCopy() *V0043PartitionInfo {
	return utils.Clone(o)
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
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043PartitionInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0043PartitionInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0043PartitionInfoList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043PartitionInfoList) DeepCopy() *V0043PartitionInfoList {
	out := new(V0043PartitionInfoList)
	out.Items = make([]V0043PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
