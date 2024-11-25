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
	ObjectTypeV0041JobInfo = "V0041JobInfo"
)

type V0041JobInfo struct {
	api.V0041JobInfo
}

// GetKey implements Object.
func (o *V0041JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(utils.NumberToString(jobId))
}

// GetType implements Object.
func (o *V0041JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0041JobInfo
}

// DeepEqualObject implements Object.
func (o *V0041JobInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(o, object)
}

// DeepCopyObject implements Object.
func (o *V0041JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041JobInfo) DeepCopy() *V0041JobInfo {
	out := new(V0041JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041JobInfoList struct {
	Items []V0041JobInfo
}

// GetType implements ObjectList.
func (o *V0041JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0041JobInfo
}

// GetItems implements ObjectList.
func (o *V0041JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041JobInfoList) AppendItem(object object.Object) {
	out := object.(*V0041JobInfo)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041JobInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041JobInfoList)
	out.Items = make([]V0041JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
