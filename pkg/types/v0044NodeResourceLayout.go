// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044NodeResourceLayout = "V0044NodeResources"
)

type V0044NodeResourceLayout struct {
	api.V0044NodeResourceLayoutList
}

// GetKey implements Object.
func (o *V0044NodeResourceLayout) GetKey() object.ObjectKey {
	// Invalid JobId because there is no JobId in the payload
	return object.ObjectKey("0")
}

// GetType implements Object.
func (o *V0044NodeResourceLayout) GetType() object.ObjectType {
	return ObjectTypeV0044NodeResourceLayout
}

// DeepCopyObject implements Object.
func (o *V0044NodeResourceLayout) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044NodeResourceLayout) DeepCopy() *V0044NodeResourceLayout {
	out := new(V0044NodeResourceLayout)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044NodeResourceLayoutList struct {
	Items []V0044NodeResourceLayout
}

// GetType implements ObjectList.
func (o *V0044NodeResourceLayoutList) GetType() object.ObjectType {
	return ObjectTypeV0044NodeResourceLayout
}

// GetItems implements ObjectList.
func (o *V0044NodeResourceLayoutList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044NodeResourceLayoutList) AppendItem(object object.Object) {
	out, ok := object.(*V0044NodeResourceLayout)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044NodeResourceLayoutList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044NodeResourceLayoutList)
	out.Items = make([]V0044NodeResourceLayout, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
