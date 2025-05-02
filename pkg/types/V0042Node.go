// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Node = "V0042Node"
)

type V0042Node struct {
	api.V0042Node
}

// GetKey implements Object.
func (o *V0042Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0042Node) GetType() object.ObjectType {
	return ObjectTypeV0042Node
}

// DeepCopyObject implements Object.
func (o *V0042Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042Node) DeepCopy() *V0042Node {
	out := new(V0042Node)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0042Node) GetStateAsSet() set.Set[api.V0042NodeState] {
	out := make(set.Set[api.V0042NodeState])
	states := ptr.Deref(o.State, []api.V0042NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0042NodeList struct {
	Items []V0042Node
}

// GetType implements ObjectList.
func (o *V0042NodeList) GetType() object.ObjectType {
	return ObjectTypeV0042Node
}

// GetItems implements ObjectList.
func (o *V0042NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042NodeList) AppendItem(object object.Object) {
	out, ok := object.(*V0042Node)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042NodeList)
	out.Items = make([]V0042Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
