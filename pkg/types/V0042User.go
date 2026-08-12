// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042User = "V0042User"
)

type V0042User struct {
	api.V0042User
}

// GetKey implements Object.
func (o *V0042User) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0042User) GetType() object.ObjectType {
	return ObjectTypeV0042User
}

// DeepCopyObject implements Object.
func (o *V0042User) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042User) DeepCopy() *V0042User {
	out := new(V0042User)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042UserList struct {
	Items []V0042User
}

// GetType implements ObjectList.
func (o *V0042UserList) GetType() object.ObjectType {
	return ObjectTypeV0042User
}

// GetItems implements ObjectList.
func (o *V0042UserList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042UserList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0042User)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042UserList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042UserList)
	out.Items = make([]V0042User, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
