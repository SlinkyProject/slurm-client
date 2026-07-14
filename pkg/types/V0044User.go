// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044User = "V0044User"
)

type V0044User struct {
	api.V0044User
}

// GetKey implements Object.
func (o *V0044User) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0044User) GetType() object.ObjectType {
	return ObjectTypeV0044User
}

// DeepCopyObject implements Object.
func (o *V0044User) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044User) DeepCopy() *V0044User {
	out := new(V0044User)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044UserList struct {
	Items []V0044User
}

// GetType implements ObjectList.
func (o *V0044UserList) GetType() object.ObjectType {
	return ObjectTypeV0044User
}

// GetItems implements ObjectList.
func (o *V0044UserList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044UserList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0044User)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044UserList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044UserList)
	out.Items = make([]V0044User, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
