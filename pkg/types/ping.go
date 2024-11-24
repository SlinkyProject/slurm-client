// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

// +kubebuilder:object:generate=true
package types

import (
	"reflect"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	ObjectTypePing = "ping"
)

const (
	PingPingedUP   = "UP"
	PingPingedDOWN = "DOWN"
)

// Ping represents a Slurm slurmctld ping.
type Ping struct {
	Hostname string
	Pinged   bool
}

// GetKey implements Object.
func (p *Ping) GetKey() object.ObjectKey {
	if p.Hostname == "" {
		panic("Hostname cannot be empty")
	}
	return object.ObjectKey(p.Hostname)
}

// GetType implements Object.
func (p *Ping) GetType() object.ObjectType {
	return ObjectTypePing
}

// DeepEqualObject implements Object.
func (in *Ping) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(in, object)
}

// DeepCopyObject implements Object.
func (p *Ping) DeepCopyObject() object.Object {
	return p.DeepCopy()
}

// PingList represents a list of Slurm slurmctld pings.
type PingList struct {
	Items []Ping
}

// GetType implements ObjectList.
func (p *PingList) GetType() object.ObjectType {
	return ObjectTypePing
}

// GetItems implements ObjectList.
func (p *PingList) GetItems() []object.Object {
	pings := make([]object.Object, 0)
	for _, ping := range p.Items {
		pings = append(pings, ping.DeepCopyObject())
	}
	return pings
}

// Size implements ObjectList.
func (p *PingList) Size() int {
	return len(p.Items)
}

// AppendItem implements ObjectList.
func (p *PingList) AppendItem(object object.Object) {
	ping := object.(*Ping)
	ping = ping.DeepCopy()
	p.Items = append(p.Items, *ping)
}

// DeepCopyObjectList implements ObjectList.
func (in *PingList) DeepCopyObjectList() object.ObjectList {
	return in.DeepCopy()
}
