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
	ObjectTypeV0045ReservationInfo = "V0045ReservationInfo"
)

type V0045ReservationInfo struct {
	api.V0045ReservationInfo
}

// GetKey implements Object.
func (o *V0045ReservationInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0045ReservationInfo) GetType() object.ObjectType {
	return ObjectTypeV0045ReservationInfo
}

// DeepCopyObject implements Object.
func (o *V0045ReservationInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0045ReservationInfo) DeepCopy() *V0045ReservationInfo {
	return utils.Clone(o)
}

type V0045ReservationInfoList struct {
	Items []V0045ReservationInfo
}

// GetType implements ObjectList.
func (o *V0045ReservationInfoList) GetType() object.ObjectType {
	return ObjectTypeV0045ReservationInfo
}

// GetItems implements ObjectList.
func (o *V0045ReservationInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0045ReservationInfoList) AppendItem(object object.Object) {
	item, ok := object.(*V0045ReservationInfo)
	if !ok {
		return
	}
	o.Items = append(o.Items, *item.DeepCopy())
}

// DeepCopyObjectList implements ObjectList.
func (o *V0045ReservationInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0045ReservationInfoList)
	out.Items = make([]V0045ReservationInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
