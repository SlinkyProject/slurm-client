// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

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
	return object.ObjectKey(fmt.Sprintf("%d", jobId))
}

// GetType implements Object.
func (o *V0042JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0042JobInfo
}

// DeepCopyObject implements Object.
func (o *V0042JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042JobInfo) DeepCopy() *V0042JobInfo {
	out := new(V0042JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0042JobInfo) GetStateAsSet() set.Set[api.V0042JobInfoJobState] {
	out := make(set.Set[api.V0042JobInfoJobState])
	states := ptr.Deref(o.JobState, []api.V0042JobInfoJobState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
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
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042JobInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0042JobInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042JobInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042JobInfoList)
	out.Items = make([]V0042JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
