// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0045ControllerPing = "V0045ControllerPing"
)

const (
	V0045ControllerPingPingedUP   = "UP"
	V0045ControllerPingPingedDOWN = "DOWN"
)

type V0045ControllerPing struct {
	api.V0045ControllerPing
}

// GetKey implements Object.
func (o *V0045ControllerPing) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0045ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0045ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0045ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045ControllerPing) DeepCopy() *V0045ControllerPing {
	out := new(V0045ControllerPing)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0045ControllerPingList struct {
	Items []V0045ControllerPing
}

// GetType implements ObjectList.
func (o *V0045ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0045ControllerPing
}

// GetItems implements ObjectList.
func (o *V0045ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045ControllerPingList) AppendItem(object object.Object) {
	out, ok := object.(*V0045ControllerPing)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045ControllerPingList)
	out.Items = make([]V0045ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
