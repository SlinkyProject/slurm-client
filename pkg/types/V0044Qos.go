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
	ObjectTypeV0044Qos = "V0044Qos"
)

type V0044Qos struct {
	api.V0044Qos
}

// GetKey implements Object.
func (o *V0044Qos) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0044Qos) GetType() object.ObjectType {
	return ObjectTypeV0044Qos
}

// DeepCopyObject implements Object.
func (o *V0044Qos) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Qos) DeepCopy() *V0044Qos {
	out := new(V0044Qos)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044QosList struct {
	Items []V0044Qos
}

// GetType implements ObjectList.
func (o *V0044QosList) GetType() object.ObjectType {
	return ObjectTypeV0044Qos
}

// GetItems implements ObjectList.
func (o *V0044QosList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044QosList) AppendItem(object object.Object) {
	out, ok := object.(*V0044Qos)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044QosList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044QosList)
	out.Items = make([]V0044Qos, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
