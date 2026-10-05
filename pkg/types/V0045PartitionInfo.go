// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045PartitionInfo = "V0045PartitionInfo"
)

type V0045PartitionInfo struct {
	api.V0045PartitionInfo
}

// GetKey implements Object.
func (o *V0045PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0045PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0045PartitionInfo
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045PartitionInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0045PartitionInfo) DeepCopy() *V0045PartitionInfo {
	return utils.Clone(o)
}

type V0045PartitionInfoList struct {
	Items []V0045PartitionInfo
}

// GetType implements ObjectList.
func (o *V0045PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0045PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0045PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045PartitionInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0045PartitionInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045PartitionInfoList) DeepCopyObject() object.RuntimeObject {
	out := new(V0045PartitionInfoList)
	out.Items = make([]V0045PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
