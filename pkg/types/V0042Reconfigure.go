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

// DeepCopyObject implements Object.
func (o *V0042Reconfigure) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042Reconfigure) DeepCopy() *V0042Reconfigure {
	out := new(V0042Reconfigure)
	utils.RemarshalOrDie(o, out)
	return out
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
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042ReconfigureList) AppendItem(object object.Object) {
	out, ok := object.(*V0042Reconfigure)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042ReconfigureList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042ReconfigureList)
	out.Items = make([]V0042Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
