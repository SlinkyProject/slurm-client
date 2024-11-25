// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"bytes"
	"encoding/json"
	"reflect"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0040Node = "V0040Node"
)

type V0040Node struct {
	api.V0040Node
}

func (o *V0040Node) GetStateAsSet() set.Set[api.V0040NodeState] {
	out := make(set.Set[api.V0040NodeState])
	states := ptr.Deref(o.State, []api.V0040NodeState{})
	for _, s := range states {
		out.Insert(s)
	}
	return out
}

type V0040NodeInfo struct {
	Namespace string `json:"namespace"`
	PodName   string `json:"podName"`
}

func (nodeInfo *V0040NodeInfo) Equal(cmp V0040NodeInfo) bool {
	a, _ := json.Marshal(nodeInfo)
	b, _ := json.Marshal(cmp)
	return bytes.Equal(a, b)
}

func (nodeInfo *V0040NodeInfo) ToString() string {
	b, _ := json.Marshal(nodeInfo)
	return string(b)
}

func V0040NodeInfoParse(str string, out *V0040NodeInfo) error {
	return json.Unmarshal([]byte(str), &out)
}

// GetKey implements Object.
func (o *V0040Node) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0040Node) GetType() object.ObjectType {
	return ObjectTypeV0040Node
}

// DeepEqualObject implements Object.
func (o *V0040Node) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(o, object)
}

// DeepCopyObject implements Object.
func (o *V0040Node) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0040Node) DeepCopy() *V0040Node {
	out := new(V0040Node)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0040NodeList struct {
	Items []V0040Node
}

// GetType implements ObjectList.
func (o *V0040NodeList) GetType() object.ObjectType {
	return ObjectTypeV0040Node
}

// GetItems implements ObjectList.
func (o *V0040NodeList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// Size implements ObjectList.
func (o *V0040NodeList) Size() int {
	return len(o.Items)
}

// AppendItem implements ObjectList.
func (o *V0040NodeList) AppendItem(object object.Object) {
	out := object.(*V0040Node)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0040NodeList) DeepCopyObjectList() object.ObjectList {
	out := new(V0040NodeList)
	out.Items = make([]V0040Node, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
