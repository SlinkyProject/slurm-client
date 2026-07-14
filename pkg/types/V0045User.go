// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045User = "V0045User"
)

type V0045User struct {
	api.V0045User
}

// GetKey implements Object.
func (o *V0045User) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0045User) GetType() object.ObjectType {
	return ObjectTypeV0045User
}

// DeepCopyObject implements Object.
func (o *V0045User) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045User) DeepCopy() *V0045User {
	out := new(V0045User)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045UserList struct {
	Items []V0045User
}

// GetType implements ObjectList.
func (o *V0045UserList) GetType() object.ObjectType {
	return ObjectTypeV0045User
}

// GetItems implements ObjectList.
func (o *V0045UserList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045UserList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0045User)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045UserList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045UserList)
	out.Items = make([]V0045User, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
