// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041Reconfigure = "V0041Reconfigure"
)

type V0041Reconfigure struct{}

// GetKey implements Object.
func (o *V0041Reconfigure) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0041Reconfigure) GetType() object.ObjectType {
	return ObjectTypeV0041Reconfigure
}

// DeepCopyObject implements Object.
func (o *V0041Reconfigure) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041Reconfigure) DeepCopy() *V0041Reconfigure {
	out := new(V0041Reconfigure)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041ReconfigureList struct {
	Items []V0041Reconfigure
}

// GetType implements ObjectList.
func (o *V0041ReconfigureList) GetType() object.ObjectType {
	return ObjectTypeV0041Reconfigure
}

// GetItems implements ObjectList.
func (o *V0041ReconfigureList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041ReconfigureList) AppendItem(object object.Object) {
	out, ok := object.(*V0041Reconfigure)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041ReconfigureList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041ReconfigureList)
	out.Items = make([]V0041Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
