// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Reconfigure = "V0043Reconfigure"
)

type V0043Reconfigure struct{}

// GetKey implements Object.
func (o *V0043Reconfigure) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0043Reconfigure) GetType() object.ObjectType {
	return ObjectTypeV0043Reconfigure
}

// DeepCopyObject implements RuntimeObject.
func (o *V0043Reconfigure) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043Reconfigure) DeepCopy() *V0043Reconfigure {
	return utils.Clone(o)
}

type V0043ReconfigureList struct {
	Items []V0043Reconfigure
}

// GetType implements ObjectList.
func (o *V0043ReconfigureList) GetType() object.ObjectType {
	return ObjectTypeV0043Reconfigure
}

// GetItems implements ObjectList.
func (o *V0043ReconfigureList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043ReconfigureList) AppendItem(object object.Object) {
	item, ok := object.(*V0043Reconfigure)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0043ReconfigureList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0043ReconfigureList) DeepCopy() *V0043ReconfigureList {
	out := new(V0043ReconfigureList)
	out.Items = make([]V0043Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
