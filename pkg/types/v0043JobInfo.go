// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043JobInfo = "V0043JobInfo"
)

type V0043JobInfo struct {
	api.V0043JobInfo
}

// GetKey implements Object.
func (o *V0043JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(fmt.Sprintf("%d", jobId))
}

// GetType implements Object.
func (o *V0043JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0043JobInfo
}

// DeepCopyObject implements RuntimeObject.
func (o *V0043JobInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043JobInfo) DeepCopy() *V0043JobInfo {
	return utils.Clone(o)
}

type V0043JobInfoList struct {
	Items []V0043JobInfo
}

// GetType implements ObjectList.
func (o *V0043JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0043JobInfo
}

// GetItems implements ObjectList.
func (o *V0043JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043JobInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0043JobInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0043JobInfoList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043JobInfoList) DeepCopy() *V0043JobInfoList {
	out := new(V0043JobInfoList)
	out.Items = make([]V0043JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
