// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strconv"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042JobInfo = "V0042JobInfo"
)

type V0042JobInfo struct {
	api.V0042JobInfo
}

// GetKey implements Object.
func (o *V0042JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(strconv.Itoa(int(jobId)))
}

// GetType implements Object.
func (o *V0042JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0042JobInfo
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042JobInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042JobInfo) DeepCopy() *V0042JobInfo {
	return utils.Clone(o)
}

type V0042JobInfoList struct {
	Items []V0042JobInfo
}

// GetType implements ObjectList.
func (o *V0042JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0042JobInfo
}

// GetItems implements ObjectList.
func (o *V0042JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042JobInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0042JobInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042JobInfoList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042JobInfoList) DeepCopy() *V0042JobInfoList {
	out := new(V0042JobInfoList)
	out.Items = make([]V0042JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
