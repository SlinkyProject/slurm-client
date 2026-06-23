// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043User = "V0043User"
)

type V0043User struct {
	api.V0043User
}

// GetKey implements Object.
func (o *V0043User) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0043User) GetType() object.ObjectType {
	return ObjectTypeV0043User
}

// DeepCopyObject implements Object.
func (o *V0043User) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043User) DeepCopy() *V0043User {
	out := new(V0043User)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043UserList struct {
	Items []V0043User
}

// GetType implements ObjectList.
func (o *V0043UserList) GetType() object.ObjectType {
	return ObjectTypeV0043User
}

// GetItems implements ObjectList.
func (o *V0043UserList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043UserList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0043User)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043UserList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043UserList)
	out.Items = make([]V0043User, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
