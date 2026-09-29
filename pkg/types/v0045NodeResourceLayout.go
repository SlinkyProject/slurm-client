// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045NodeResourceLayout = "V0045NodeResources"
)

type V0045NodeResourceLayout struct {
	api.V0045NodeResourceLayoutList
}

// GetKey implements Object.
func (o *V0045NodeResourceLayout) GetKey() object.ObjectKey {
	// Invalid JobId because there is no JobId in the payload
	return object.ObjectKey("0")
}

// GetType implements Object.
func (o *V0045NodeResourceLayout) GetType() object.ObjectType {
	return ObjectTypeV0045NodeResourceLayout
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045NodeResourceLayout) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0045NodeResourceLayout) DeepCopy() *V0045NodeResourceLayout {
	return utils.Clone(o)
}

type V0045NodeResourceLayoutList struct {
	Items []V0045NodeResourceLayout
}

// GetType implements ObjectList.
func (o *V0045NodeResourceLayoutList) GetType() object.ObjectType {
	return ObjectTypeV0045NodeResourceLayout
}

// GetItems implements ObjectList.
func (o *V0045NodeResourceLayoutList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045NodeResourceLayoutList) AppendItem(object object.Object) {
	item, ok := object.(*V0045NodeResourceLayout)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045NodeResourceLayoutList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045NodeResourceLayoutList) DeepCopy() *V0045NodeResourceLayoutList {
	out := new(V0045NodeResourceLayoutList)
	out.Items = make([]V0045NodeResourceLayout, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
