// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

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

// DeepCopyObject implements Object.
func (o *V0045Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045Node) DeepCopy() *V0045Node {
	out := new(V0045Node)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0045Node) GetStateAsSet() set.Set[api.V0045NodeState] {
	out := make(set.Set[api.V0045NodeState])
	states := ptr.Deref(o.State, []api.V0045NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
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
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045NodeList) AppendItem(object object.Object) {
	out, ok := object.(*V0045Node)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045NodeList)
	out.Items = make([]V0045Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
