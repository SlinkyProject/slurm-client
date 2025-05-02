// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042ControllerPing = "V0042ControllerPing"
)

const (
	V0042ControllerPingPingedUP   = "UP"
	V0042ControllerPingPingedDOWN = "DOWN"
)

type V0042ControllerPing struct {
	api.V0042ControllerPing
}

// GetKey implements Object.
func (o *V0042ControllerPing) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0042ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0042ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0042ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042ControllerPing) DeepCopy() *V0042ControllerPing {
	out := new(V0042ControllerPing)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042ControllerPingList struct {
	Items []V0042ControllerPing
}

// GetType implements ObjectList.
func (o *V0042ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0042ControllerPing
}

// GetItems implements ObjectList.
func (o *V0042ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042ControllerPingList) AppendItem(object object.Object) {
	out, ok := object.(*V0042ControllerPing)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042ControllerPingList)
	out.Items = make([]V0042ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
