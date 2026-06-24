// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Assoc = "V0043Assoc"
)

type V0043Assoc struct {
	api.V0043Assoc
}

// GetKey implements Object. Key is "cluster/account/user/partition".
func (o *V0043Assoc) GetKey() object.ObjectKey {
	return object.ObjectKey(fmt.Sprintf("%s/%s/%s/%s",
		ptr.Deref(o.Cluster, ""),
		ptr.Deref(o.Account, ""),
		o.User,
		ptr.Deref(o.Partition, ""),
	))
}

// GetType implements Object.
func (o *V0043Assoc) GetType() object.ObjectType {
	return ObjectTypeV0043Assoc
}

// DeepCopyObject implements Object.
func (o *V0043Assoc) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043Assoc) DeepCopy() *V0043Assoc {
	out := new(V0043Assoc)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043AssocList struct {
	Items []V0043Assoc
}

// GetType implements ObjectList.
func (o *V0043AssocList) GetType() object.ObjectType {
	return ObjectTypeV0043Assoc
}

// GetItems implements ObjectList.
func (o *V0043AssocList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043AssocList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0043Assoc)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043AssocList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043AssocList)
	out.Items = make([]V0043Assoc, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
