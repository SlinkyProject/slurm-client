// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0040ControllerPing = "V0040ControllerPing"
)

const (
	V0040ControllerPingPingedUP   = "UP"
	V0040ControllerPingPingedDOWN = "DOWN"
)

type V0040ControllerPing struct {
	api.V0040ControllerPing
}

// GetKey implements Object.
func (o *V0040ControllerPing) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0040ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0040ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0040ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0040ControllerPing) DeepCopy() *V0040ControllerPing {
	out := new(V0040ControllerPing)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0040ControllerPingList struct {
	Items []V0040ControllerPing `json:"items,omitempty"`
}

// GetType implements ObjectList.
func (o *V0040ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0040ControllerPing
}

// GetItems implements ObjectList.
func (o *V0040ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0040ControllerPingList) AppendItem(object object.Object) {
	out, ok := object.(*V0040ControllerPing)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0040ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0040ControllerPingList)
	out.Items = make([]V0040ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
