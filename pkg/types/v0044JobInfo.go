// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

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
	return object.ObjectKey(fmt.Sprintf("%d", jobId))
}

// GetType implements Object.
func (o *V0044JobInfo) GetType() object.ObjectType {
	return ObjectTypeV0044JobInfo
}

// DeepCopyObject implements Object.
func (o *V0044JobInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044JobInfo) DeepCopy() *V0044JobInfo {
	out := new(V0044JobInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0044JobInfo) GetStateAsSet() set.Set[api.V0044JobInfoJobState] {
	out := make(set.Set[api.V0044JobInfoJobState])
	states := ptr.Deref(o.JobState, []api.V0044JobInfoJobState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
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
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044JobInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0044JobInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044JobInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044JobInfoList)
	out.Items = make([]V0044JobInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
