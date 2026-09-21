// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Reconfigure = "V0044Reconfigure"
)

type V0044Reconfigure struct{}

// GetKey implements Object.
func (o *V0044Reconfigure) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0044Reconfigure) GetType() object.ObjectType {
	return ObjectTypeV0044Reconfigure
}

// DeepCopyObject implements RuntimeObject.
func (o *V0044Reconfigure) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0044Reconfigure) DeepCopy() *V0044Reconfigure {
	out := new(V0044Reconfigure)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044ReconfigureList struct {
	Items []V0044Reconfigure
}

// GetType implements ObjectList.
func (o *V0044ReconfigureList) GetType() object.ObjectType {
	return ObjectTypeV0044Reconfigure
}

// GetItems implements ObjectList.
func (o *V0044ReconfigureList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopy()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044ReconfigureList) AppendItem(object object.Object) {
	out, ok := object.(*V0044Reconfigure)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObject implements RuntimeObject.
func (o *V0044ReconfigureList) DeepCopyObject() object.RuntimeObject {
	return o.DeepCopy()
}

func (o *V0044ReconfigureList) DeepCopy() *V0044ReconfigureList {
	out := new(V0044ReconfigureList)
	out.Items = make([]V0044Reconfigure, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
