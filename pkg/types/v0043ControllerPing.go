// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043ControllerPing = "V0043ControllerPing"
)

const (
	V0043ControllerPingPingedUP   = "UP"
	V0043ControllerPingPingedDOWN = "DOWN"
)

type V0043ControllerPing struct {
	api.V0043ControllerPing
}

// GetKey implements Object.
func (o *V0043ControllerPing) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0043ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0043ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0043ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043ControllerPing) DeepCopy() *V0043ControllerPing {
	out := new(V0043ControllerPing)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043ControllerPingList struct {
	Items []V0043ControllerPing
}

// GetType implements ObjectList.
func (o *V0043ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0043ControllerPing
}

// GetItems implements ObjectList.
func (o *V0043ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043ControllerPingList) AppendItem(object object.Object) {
	out, ok := object.(*V0043ControllerPing)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043ControllerPingList)
	out.Items = make([]V0043ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
