// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

// +kubebuilder:object:generate=true
package types

import (
	"reflect"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	ObjectTypeControllerPing = "controllerPing"
)

const (
	ControllerPingPingedUP   = "UP"
	ControllerPingPingedDOWN = "DOWN"
)

// ControllerPing represents a Slurm slurmctld ping.
type ControllerPing struct {
	Hostname string
	Pinged   bool
}

// GetKey implements Object.
func (p *ControllerPing) GetKey() object.ObjectKey {
	if p.Hostname == "" {
		panic("Hostname cannot be empty")
	}
	return object.ObjectKey(p.Hostname)
}

// GetType implements Object.
func (p *ControllerPing) GetType() object.ObjectType {
	return ObjectTypeControllerPing
}

// DeepEqualObject implements Object.
func (in *ControllerPing) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(in, object)
}

// DeepCopyObject implements Object.
func (p *ControllerPing) DeepCopyObject() object.Object {
	return p.DeepCopy()
}

// ControllerPingList represents a list of Slurm slurmctld pings.
type ControllerPingList struct {
	Items []ControllerPing
}

// GetType implements ObjectList.
func (p *ControllerPingList) GetType() object.ObjectType {
	return ObjectTypeControllerPing
}

// GetItems implements ObjectList.
func (p *ControllerPingList) GetItems() []object.Object {
	pings := make([]object.Object, 0)
	for _, ping := range p.Items {
		pings = append(pings, ping.DeepCopyObject())
	}
	return pings
}

// Size implements ObjectList.
func (p *ControllerPingList) Size() int {
	return len(p.Items)
}

// AppendItem implements ObjectList.
func (p *ControllerPingList) AppendItem(object object.Object) {
	ping := object.(*ControllerPing)
	ping = ping.DeepCopy()
	p.Items = append(p.Items, *ping)
}

// DeepCopyObjectList implements ObjectList.
func (in *ControllerPingList) DeepCopyObjectList() object.ObjectList {
	return in.DeepCopy()
}
