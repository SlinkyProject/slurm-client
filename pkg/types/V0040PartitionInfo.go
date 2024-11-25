// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"k8s.io/utils/ptr"

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

// DeepEqualObject implements Object.
func (in *V0040PartitionInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(in, object)
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
	out := object.(*V0040PartitionInfo)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
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
