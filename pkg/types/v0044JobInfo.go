// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strconv"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044JobInfo = "V0044JobInfo"
)

type V0044JobInfo struct {
	api.V0044JobInfo
}

// GetKey implements Object.
func (o *V0044JobInfo) GetKey() object.ObjectKey {
	jobId := ptr.Deref(o.JobId, 0)
	return object.ObjectKey(strconv.Itoa(int(jobId)))
}

// GetType implements Object.
func (o *V0044JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0044JobInfo
}

// DeepCopyObject implements RuntimeObject.
func (o *V0044JobInfo) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0044JobInfo) DeepCopy() *V0044JobInfo {
	return utils.Clone(o)
}

type V0044JobInfoList struct {
	Items []V0044JobInfo
}

// GetType implements ObjectList.
func (o *V0044JobInfoList) GetType() object.ObjectType {
	return ObjectTypeV0044JobInfo
}

// GetItems implements ObjectList.
func (o *V0044JobInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044JobInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0044JobInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0044JobInfoList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0044JobInfoList) DeepCopy() *V0044JobInfoList {
	out := new(V0044JobInfoList)
	out.Items = make([]V0044JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
