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
	ObjectTypeV0044ReservationInfo = "V0044ReservationInfo"
)

type V0044ReservationInfo struct {
	api.V0044ReservationInfo
}

// GetKey implements Object.
func (o *V0044ReservationInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0044ReservationInfo) GetType() object.ObjectType {
	return ObjectTypeV0044ReservationInfo
}

// DeepCopyObject implements Object.
func (o *V0044ReservationInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044ReservationInfo) DeepCopy() *V0044ReservationInfo {
	out := new(V0044ReservationInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044ReservationInfoList struct {
	Items []V0044ReservationInfo
}

// GetType implements ObjectList.
func (o *V0044ReservationInfoList) GetType() object.ObjectType {
	return ObjectTypeV0044ReservationInfo
}

// GetItems implements ObjectList.
func (o *V0044ReservationInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044ReservationInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0044ReservationInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044ReservationInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044ReservationInfoList)
	out.Items = make([]V0044ReservationInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
