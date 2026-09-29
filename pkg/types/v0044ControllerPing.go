// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044ControllerPing = "V0044ControllerPing"
)

const (
	V0044ControllerPingPingedUP   = "UP"
	V0044ControllerPingPingedDOWN = "DOWN"
)

type V0044ControllerPing struct {
	api.V0044ControllerPing
}

// GetKey implements Object.
func (o *V0044ControllerPing) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Hostname, ""))
}

// GetType implements Object.
func (o *V0044ControllerPing) GetType() object.ObjectType {
	return ObjectTypeV0044ControllerPing
}

// DeepCopyObject implements Object.
func (o *V0044ControllerPing) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044ControllerPing) DeepCopy() *V0044ControllerPing {
	return utils.Clone(o)
}

type V0044ControllerPingList struct {
	Items []V0044ControllerPing
}

// GetType implements ObjectList.
func (o *V0044ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeV0044ControllerPing
}

// GetItems implements ObjectList.
func (o *V0044ControllerPingList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044ControllerPingList) AppendItem(object object.Object) {
	item, ok := object.(*V0044ControllerPing)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044ControllerPingList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044ControllerPingList)
	out.Items = make([]V0044ControllerPing, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
