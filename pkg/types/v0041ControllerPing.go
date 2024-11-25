// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041ControllerPing = "V0041ControllerPing"
)

const (
	V0041ControllerPingPingedUP   = "UP"
	V0041ControllerPingPingedDOWN = "DOWN"
)

type V0041ControllerPing struct {
	api.V0041ControllerPing
}

// GetKey implements Object.
func (o *V0041ControllerPing) GetKey() object.ObjectKey {
	if o.Hostname == nil {
		panic("Hostname cannot be empty")
	}
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0041ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0041ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0041ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041ControllerPing) DeepCopy() *V0041ControllerPing {
	out := new(V0041ControllerPing)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041ControllerPingList struct {
	Items []V0041ControllerPing
}

// GetType implements ObjectList.
func (o *V0041ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0041ControllerPing
}

// GetItems implements ObjectList.
func (o *V0041ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041ControllerPingList) AppendItem(object object.Object) {
	out := object.(*V0041ControllerPing)
	utils.RemarshalOrDie(object, out)
	o.Items = append(o.Items, *out)
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041ControllerPingList)
	out.Items = make([]V0041ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
