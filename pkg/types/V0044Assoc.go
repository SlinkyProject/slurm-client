// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Assoc = "V0044Assoc"
)

type V0044Assoc struct {
	api.V0044Assoc
}

// GetKey implements Object. Key is "cluster/account/user/partition".
func (o *V0044Assoc) GetKey() object.ObjectKey {
	return object.ObjectKey(fmt.Sprintf("%s/%s/%s/%s",
		ptr.Deref(o.Cluster, ""),
		ptr.Deref(o.Account, ""),
		o.User,
		ptr.Deref(o.Partition, ""),
	))
}

// GetType implements Object.
func (o *V0044Assoc) GetType() object.ObjectType {
	return ObjectTypeV0044Assoc
}

// DeepCopyObject implements Object.
func (o *V0044Assoc) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Assoc) DeepCopy() *V0044Assoc {
	out := new(V0044Assoc)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044AssocList struct {
	Items []V0044Assoc
}

// GetType implements ObjectList.
func (o *V0044AssocList) GetType() object.ObjectType {
	return ObjectTypeV0044Assoc
}

// GetItems implements ObjectList.
func (o *V0044AssocList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044AssocList) AppendItem(obj object.Object) {
	out, ok := obj.(*V0044Assoc)
	if ok {
		utils.RemarshalOrDie(obj, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044AssocList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044AssocList)
	out.Items = make([]V0044Assoc, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
