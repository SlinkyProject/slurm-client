// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045Assoc = "V0045Assoc"
)

type V0045Assoc struct {
	api.V0045Assoc
}

// GetKey implements Object. Key is "cluster/account/user/partition".
func (o *V0045Assoc) GetKey() object.ObjectKey {
	return object.ObjectKey(fmt.Sprintf("%s/%s/%s/%s",
		ptr.Deref(o.Cluster, ""),
		ptr.Deref(o.Account, ""),
		o.User,
		ptr.Deref(o.Partition, ""),
	))
}

// GetType implements Object.
func (o *V0045Assoc) GetType() object.ObjectType {
	return ObjectTypeV0045Assoc
}

// DeepCopyObject implements Object.
func (o *V0045Assoc) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045Assoc) DeepCopy() *V0045Assoc {
	out := new(V0045Assoc)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045AssocList struct {
	Items []V0045Assoc
}

// GetType implements ObjectList.
func (o *V0045AssocList) GetType() object.ObjectType {
	return ObjectTypeV0045Assoc
}

// GetItems implements ObjectList.
func (o *V0045AssocList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045AssocList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0045Assoc)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045AssocList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045AssocList)
	out.Items = make([]V0045Assoc, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
