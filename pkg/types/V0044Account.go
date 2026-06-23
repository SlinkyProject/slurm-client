// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Account = "V0044Account"
)

type V0044Account struct {
	api.V0044Account
}

// GetKey implements Object.
func (o *V0044Account) GetKey() object.ObjectKey {
	return object.ObjectKey(o.Name)
}

// GetType implements Object.
func (o *V0044Account) GetType() object.ObjectType {
	return ObjectTypeV0044Account
}

// DeepCopyObject implements Object.
func (o *V0044Account) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Account) DeepCopy() *V0044Account {
	out := new(V0044Account)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044AccountList struct {
	Items []V0044Account
}

// GetType implements ObjectList.
func (o *V0044AccountList) GetType() object.ObjectType {
	return ObjectTypeV0044Account
}

// GetItems implements ObjectList.
func (o *V0044AccountList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044AccountList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0044Account)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044AccountList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044AccountList)
	out.Items = make([]V0044Account, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
