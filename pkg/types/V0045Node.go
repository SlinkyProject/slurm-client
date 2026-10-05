// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045Node = "V0045Node"
)

type V0045Node struct {
	api.V0045Node
}

// GetKey implements Object.
func (o *V0045Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0045Node) GetType() object.ObjectType {
	return ObjectTypeV0045Node
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045Node) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0045Node) DeepCopy() *V0045Node {
	return utils.Clone(o)
}

type V0045NodeList struct {
	Items []V0045Node
}

// GetType implements ObjectList.
func (o *V0045NodeList) GetType() object.ObjectType {
	return ObjectTypeV0045Node
}

// GetItems implements ObjectList.
func (o *V0045NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045NodeList) AppendItem(object object.Object) {
	item, ok := object.(*V0045Node)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObject implements RuntimeObject.
func (o *V0045NodeList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0045NodeList) DeepCopy() *V0045NodeList {
	out := new(V0045NodeList)
	out.Items = make([]V0045Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
