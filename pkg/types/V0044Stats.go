// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Stats = "V0044Stats"
)

type V0044Stats struct {
	api.V0044StatsMsg
}

// GetKey implements Object.
func (o *V0044Stats) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0044Stats) GetType() object.ObjectType {
	return ObjectTypeV0044Stats
}

// DeepCopyObject implements Object.
func (o *V0044Stats) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Stats) DeepCopy() *V0044Stats {
	out := new(V0044Stats)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044StatsList struct {
	Items []V0044Stats
}

// GetType implements ObjectList.
func (o *V0044StatsList) GetType() object.ObjectType {
	return ObjectTypeV0044Stats
}

// GetItems implements ObjectList.
func (o *V0044StatsList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044StatsList) AppendItem(object object.Object) {
	out, ok := object.(*V0044Stats)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044StatsList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044StatsList)
	out.Items = make([]V0044Stats, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
