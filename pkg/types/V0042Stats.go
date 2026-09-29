// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Stats = "V0042Stats"
)

type V0042Stats struct {
	api.V0042StatsMsg
}

// GetKey implements Object.
func (o *V0042Stats) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0042Stats) GetType() object.ObjectType {
	return ObjectTypeV0042Stats
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042Stats) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042Stats) DeepCopy() *V0042Stats {
	return utils.Clone(o)
}

type V0042StatsList struct {
	Items []V0042Stats
}

// GetType implements ObjectList.
func (o *V0042StatsList) GetType() object.ObjectType {
	return ObjectTypeV0042Stats
}

// GetItems implements ObjectList.
func (o *V0042StatsList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042StatsList) AppendItem(object object.Object) {
	item, ok := object.(*V0042Stats)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042StatsList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042StatsList) DeepCopy() *V0042StatsList {
	out := new(V0042StatsList)
	out.Items = make([]V0042Stats, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
