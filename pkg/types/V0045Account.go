// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045Account = "V0045Account"
)

type V0045Account struct {
	api.V0045Account
}

// GetKey implements Object.
func (o *V0045Account) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0045Account) GetType() object.ObjectType {
	return ObjectTypeV0045Account
}

// DeepCopyObject implements Object.
func (o *V0045Account) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045Account) DeepCopy() *V0045Account {
	out := new(V0045Account)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045AccountList struct {
	Items []V0045Account
}

// GetType implements ObjectList.
func (o *V0045AccountList) GetType() object.ObjectType {
	return ObjectTypeV0045Account
}

// GetItems implements ObjectList.
func (o *V0045AccountList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045AccountList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0045Account)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045AccountList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045AccountList)
	out.Items = make([]V0045Account, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
