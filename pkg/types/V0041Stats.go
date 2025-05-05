// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041Stats = "V0041Stats"
)

type V0041Stats struct {
	api.V0041StatsMsg
}

// GetKey implements Object.
func (o *V0041Stats) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0041Stats) GetType() object.ObjectType {
	return ObjectTypeV0041Stats
}

// DeepCopyObject implements Object.
func (o *V0041Stats) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041Stats) DeepCopy() *V0041Stats {
	out := new(V0041Stats)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041StatsList struct {
	Items []V0041Stats
}

// GetType implements ObjectList.
func (o *V0041StatsList) GetType() object.ObjectType {
	return ObjectTypeV0041Stats
}

// GetItems implements ObjectList.
func (o *V0041StatsList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041StatsList) AppendItem(object object.Object) {
	out, ok := object.(*V0041Stats)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041StatsList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041StatsList)
	out.Items = make([]V0041Stats, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
