// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Assoc = "V0042Assoc"
)

type V0042Assoc struct {
	api.V0042Assoc
}

// GetKey implements Object. Key is "cluster/account/user/partition".
func (o *V0042Assoc) GetKey() object.ObjectKey {
	return object.ObjectKey(fmt.Sprintf("%s/%s/%s/%s",
		ptr.Deref(o.Cluster, ""),
		ptr.Deref(o.Account, ""),
		o.User,
		ptr.Deref(o.Partition, ""),
	))
}

// GetType implements Object.
func (o *V0042Assoc) GetType() object.ObjectType {
	return ObjectTypeV0042Assoc
}

// DeepCopyObject implements Object.
func (o *V0042Assoc) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042Assoc) DeepCopy() *V0042Assoc {
	out := new(V0042Assoc)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042AssocList struct {
	Items []V0042Assoc
}

// GetType implements ObjectList.
func (o *V0042AssocList) GetType() object.ObjectType {
	return ObjectTypeV0042Assoc
}

// GetItems implements ObjectList.
func (o *V0042AssocList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042AssocList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0042Assoc)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042AssocList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042AssocList)
	out.Items = make([]V0042Assoc, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
