// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045Reconfigure = "V0045Reconfigure"
)

type V0045Reconfigure struct{}

// GetKey implements Object.
func (o *V0045Reconfigure) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0045Reconfigure) GetType() object.ObjectType {
	return ObjectTypeV0045Reconfigure
}

// DeepCopyObject implements Object.
func (o *V0045Reconfigure) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045Reconfigure) DeepCopy() *V0045Reconfigure {
	out := new(V0045Reconfigure)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045ReconfigureList struct {
	Items []V0045Reconfigure
}

// GetType implements ObjectList.
func (o *V0045ReconfigureList) GetType() object.ObjectType {
	return ObjectTypeV0045Reconfigure
}

// GetItems implements ObjectList.
func (o *V0045ReconfigureList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045ReconfigureList) AppendItem(object object.Object) {
	out, ok := object.(*V0045Reconfigure)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045ReconfigureList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045ReconfigureList)
	out.Items = make([]V0045Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
