// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

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

// DeepCopyObject implements Object.
func (o *V0043JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043JobInfo) DeepCopy() *V0043JobInfo {
	out := new(V0043JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0043JobInfo) GetStateAsSet() set.Set[api.V0043JobInfoJobState] {
	out := make(set.Set[api.V0043JobInfoJobState])
	states := ptr.Deref(o.JobState, []api.V0043JobInfoJobState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
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
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043JobInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0043JobInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043JobInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043JobInfoList)
	out.Items = make([]V0043JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
