// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041Node = "V0041Node"
)

type V0041Node struct {
	api.V0041Node
}

func (o *V0041Node) GetStateAsSet() set.Set[api.V0041NodeState] {
	out := make(set.Set[api.V0041NodeState])
	states := ptr.Deref(o.State, []api.V0041NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

// GetKey implements Object.
func (o *V0041Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0041Node) GetType() object.ObjectType {
	return ObjectTypeV0041Node
}

// DeepEqualObject implements Object.
func (o *V0041Node) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(o, object)
}

// DeepCopyObject implements Object.
func (o *V0041Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041Node) DeepCopy() *V0041Node {
	out := new(V0041Node)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041NodeList struct {
	Items []V0041Node
}

// GetType implements ObjectList.
func (o *V0041NodeList) GetType() object.ObjectType {
	return ObjectTypeV0041Node
}

// GetItems implements ObjectList.
func (o *V0041NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// Size implements ObjectList.
func (o *V0041NodeList) Size() int {
	return len(o.Items)
}

// AppendItem implements ObjectList.
func (o *V0041NodeList) AppendItem(object object.Object) {
	out := object.(*V0041Node)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041NodeList)
	out.Items = make([]V0041Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
