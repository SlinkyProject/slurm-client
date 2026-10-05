// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045JobInfo = "V0045JobInfo"
)

type V0045JobInfo struct {
	api.V0045JobInfo
}

// GetKey implements Object.
func (o *V0045JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(fmt.Sprintf("%d", jobId))
}

// GetType implements Object.
func (o *V0045JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0045JobInfo
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045JobInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0045JobInfo) DeepCopy() *V0045JobInfo {
	return utils.Clone(o)
}

type V0045JobInfoList struct {
	Items []V0045JobInfo
}

// GetType implements ObjectList.
func (o *V0045JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0045JobInfo
}

// GetItems implements ObjectList.
func (o *V0045JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045JobInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0045JobInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045JobInfoList) DeepCopyObject() object.RuntimeObject {
	out := new(V0045JobInfoList)
	out.Items = make([]V0045JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
