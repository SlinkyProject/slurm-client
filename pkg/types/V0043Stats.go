// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Stats = "V0043Stats"
)

type V0043Stats struct {
	api.V0043StatsMsg
}

// GetKey implements Object.
func (o *V0043Stats) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0043Stats) GetType() object.ObjectType {
	return ObjectTypeV0043Stats
}

// DeepCopyObject implements Object.
func (o *V0043Stats) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043Stats) DeepCopy() *V0043Stats {
	out := new(V0043Stats)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043StatsList struct {
	Items []V0043Stats
}

// GetType implements ObjectList.
func (o *V0043StatsList) GetType() object.ObjectType {
	return ObjectTypeV0043Stats
}

// GetItems implements ObjectList.
func (o *V0043StatsList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043StatsList) AppendItem(object object.Object) {
	out, ok := object.(*V0043Stats)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043StatsList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043StatsList)
	out.Items = make([]V0043Stats, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
