// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Node = "V0044Node"
)

type V0044Node struct {
	api.V0044Node
}

// GetKey implements Object.
func (o *V0044Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0044Node) GetType() object.ObjectType {
	return ObjectTypeV0044Node
}

// DeepCopyObject implements Object.
func (o *V0044Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Node) DeepCopy() *V0044Node {
	out := new(V0044Node)
	utils.RemarshalOrDie(o, out)
	return out
}

func (o *V0044Node) GetStateAsSet() set.Set[api.V0044NodeState] {
	out := make(set.Set[api.V0044NodeState])
	states := ptr.Deref(o.State, []api.V0044NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0044NodeList struct {
	Items []V0044Node
}

// GetType implements ObjectList.
func (o *V0044NodeList) GetType() object.ObjectType {
	return ObjectTypeV0044Node
}

// GetItems implements ObjectList.
func (o *V0044NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044NodeList) AppendItem(object object.Object) {
	out, ok := object.(*V0044Node)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044NodeList)
	out.Items = make([]V0044Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
