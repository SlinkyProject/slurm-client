// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041PartitionInfo = "V0041PartitionInfo"
)

type V0041PartitionInfo struct {
	api.V0041PartitionInfo
}

// GetKey implements Object.
func (o *V0041PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0041PartitionInfo) GetType() object.ObjectType {
	return ObjectTypeV0041PartitionInfo
}

// DeepEqualObject implements Object.
func (in *V0041PartitionInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(in, object)
}

// DeepCopyObject implements Object.
func (o *V0041PartitionInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041PartitionInfo) DeepCopy() *V0041PartitionInfo {
	out := new(V0041PartitionInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041PartitionInfoList struct {
	Items []V0041PartitionInfo
}

// GetType implements ObjectList.
func (o *V0041PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypeV0041PartitionInfo
}

// GetItems implements ObjectList.
func (o *V0041PartitionInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// Size implements ObjectList.
func (o *V0041PartitionInfoList) Size() int {
	return len(o.Items)
}

// AppendItem implements ObjectList.
func (o *V0041PartitionInfoList) AppendItem(object object.Object) {
	out := object.(*V0041PartitionInfo)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041PartitionInfoList)
	out.Items = make([]V0041PartitionInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
