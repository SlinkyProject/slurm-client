// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045Stats = "V0045Stats"
)

type V0045Stats struct {
	api.V0045StatsMsg
}

// GetKey implements Object.
func (o *V0045Stats) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0045Stats) GetType() object.ObjectType {
	return ObjectTypeV0045Stats
}

// DeepCopyObject implements Object.
func (o *V0045Stats) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045Stats) DeepCopy() *V0045Stats {
	out := new(V0045Stats)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045StatsList struct {
	Items []V0045Stats
}

// GetType implements ObjectList.
func (o *V0045StatsList) GetType() object.ObjectType {
	return ObjectTypeV0045Stats
}

// GetItems implements ObjectList.
func (o *V0045StatsList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045StatsList) AppendItem(object object.Object) {
	out, ok := object.(*V0045Stats)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045StatsList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045StatsList)
	out.Items = make([]V0045Stats, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
