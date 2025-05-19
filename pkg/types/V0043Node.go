// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Node = "V0043Node"
)

type V0043Node struct {
	api.V0043Node
}

// GetKey implements Object.
func (o *V0043Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0043Node) GetType() object.ObjectType {
	return ObjectTypeV0043Node
}

// DeepCopyObject implements Object.
func (o *V0043Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043Node) DeepCopy() *V0043Node {
	out := new(V0043Node)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0043Node) GetStateAsSet() set.Set[api.V0043NodeState] {
	out := make(set.Set[api.V0043NodeState])
	states := ptr.Deref(o.State, []api.V0043NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0043NodeList struct {
	Items []V0043Node
}

// GetType implements ObjectList.
func (o *V0043NodeList) GetType() object.ObjectType {
	return ObjectTypeV0043Node
}

// GetItems implements ObjectList.
func (o *V0043NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043NodeList) AppendItem(object object.Object) {
	out, ok := object.(*V0043Node)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043NodeList)
	out.Items = make([]V0043Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
