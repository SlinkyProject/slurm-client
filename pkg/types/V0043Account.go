// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Account = "V0043Account"
)

type V0043Account struct {
	api.V0043Account
}

// GetKey implements Object.
func (o *V0043Account) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0043Account) GetType() object.ObjectType {
	return ObjectTypeV0043Account
}

// DeepCopyObject implements Object.
func (o *V0043Account) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043Account) DeepCopy() *V0043Account {
	out := new(V0043Account)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043AccountList struct {
	Items []V0043Account
}

// GetType implements ObjectList.
func (o *V0043AccountList) GetType() object.ObjectType {
	return ObjectTypeV0043Account
}

// GetItems implements ObjectList.
func (o *V0043AccountList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043AccountList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0043Account)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043AccountList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043AccountList)
	out.Items = make([]V0043Account, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
