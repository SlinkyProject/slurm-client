// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Reconfigure = "V0042Reconfigure"
)

type V0042Reconfigure struct{}

// GetKey implements Object.
func (o *V0042Reconfigure) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0042Reconfigure) GetType() object.ObjectType {
	return ObjectTypeV0042Reconfigure
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042Reconfigure) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042Reconfigure) DeepCopy() *V0042Reconfigure {
	return utils.Clone(o)
}

type V0042ReconfigureList struct {
	Items []V0042Reconfigure
}

// GetType implements ObjectList.
func (o *V0042ReconfigureList) GetType() object.ObjectType {
	return ObjectTypeV0042Reconfigure
}

// GetItems implements ObjectList.
func (o *V0042ReconfigureList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042ReconfigureList) AppendItem(object object.Object) {
	item, ok := object.(*V0042Reconfigure)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0042ReconfigureList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0042ReconfigureList) DeepCopy() *V0042ReconfigureList {
	out := new(V0042ReconfigureList)
	out.Items = make([]V0042Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
