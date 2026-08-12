// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Account = "V0042Account"
)

type V0042Account struct {
	api.V0042Account
}

// GetKey implements Object.
func (o *V0042Account) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0042Account) GetType() object.ObjectType {
	return ObjectTypeV0042Account
}

// DeepCopyObject implements Object.
func (o *V0042Account) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042Account) DeepCopy() *V0042Account {
	out := new(V0042Account)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042AccountList struct {
	Items []V0042Account
}

// GetType implements ObjectList.
func (o *V0042AccountList) GetType() object.ObjectType {
	return ObjectTypeV0042Account
}

// GetItems implements ObjectList.
func (o *V0042AccountList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042AccountList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0042Account)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042AccountList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042AccountList)
	out.Items = make([]V0042Account, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
