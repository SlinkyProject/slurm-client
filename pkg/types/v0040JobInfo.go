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
	ObjectTypeV0040JobInfo = "V0040JobInfo"
)

type V0040JobInfo struct {
	api.V0040JobInfo
}

// GetKey implements Object.
func (o *V0040JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(utils.NumberToString(jobId))
}

// GetType implements Object.
func (o *V0040JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0040JobInfo
}

// DeepEqualObject implements Object.
func (o *V0040JobInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(o, object)
}

// DeepCopyObject implements Object.
func (o *V0040JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0040JobInfo) DeepCopy() *V0040JobInfo {
	out := new(V0040JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0040JobInfoList struct {
	Items []V0040JobInfo
}

// GetType implements ObjectList.
func (o *V0040JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0040JobInfo
}

// GetItems implements ObjectList.
func (o *V0040JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// Size implements ObjectList.
func (o *V0040JobInfoList) Size() int {
	return len(o.Items)
}

// AppendItem implements ObjectList.
func (o *V0040JobInfoList) AppendItem(object object.Object) {
	out := object.(*V0040JobInfo)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0040JobInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0040JobInfoList)
	out.Items = make([]V0040JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
