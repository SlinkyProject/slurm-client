// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

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

// DeepCopyObject implements Object.
func (o *V0040JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0040JobInfo) DeepCopy() *V0040JobInfo {
	out := new(V0040JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0040JobInfo) GetStateAsSet() set.Set[api.V0040JobInfoJobState] {
	out := make(set.Set[api.V0040JobInfoJobState])
	states := ptr.Deref(o.JobState, []api.V0040JobInfoJobState{})
	for _, s := range states {
		out.Insert(s)
	}
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

// AppendItem implements ObjectList.
func (o *V0040JobInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0040JobInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
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
