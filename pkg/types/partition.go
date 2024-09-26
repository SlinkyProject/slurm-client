// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"

	"k8s.io/utils/set"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	ObjectTypePartitionInfo = "partition"
)

// Represents the PartitionInfo Partition State
// +enum
type PartitionInfoPartitionState string

const (
	PartitionInfoPartitionStateDOWN     PartitionInfoPartitionState = "DOWN"
	PartitionInfoPartitionStateDRAIN    PartitionInfoPartitionState = "DRAIN"
	PartitionInfoPartitionStateINACTIVE PartitionInfoPartitionState = "INACTIVE"
	PartitionInfoPartitionStateUNKNOWN  PartitionInfoPartitionState = "UNKNOWN"
	PartitionInfoPartitionStateUP       PartitionInfoPartitionState = "UP"
)

// PartitionInfo represents a Slurm partition.
type PartitionInfo struct {
	Name           string
	NodeSets       string
	Nodes          int32
	PartitionState set.Set[PartitionInfoPartitionState]
	Cpus           int32
}

//+kubebuilder:object:root=true
//+kubebuilder:object:generate=false

// GetKey implements Object.
func (p *PartitionInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(p.Name)
}

// GetType implements Object.
func (p *PartitionInfo) GetType() object.ObjectType {
	return ObjectTypePartitionInfo
}

// DeepEqualObject implements Object.
func (in *PartitionInfo) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(in, object)
}

// DeepCopyObject implements Object.
func (p *PartitionInfo) DeepCopyObject() object.Object {
	return p.DeepCopy()
}

// PartitionInfoList represents a list of Slurm partitions.
type PartitionInfoList struct {
	Items []PartitionInfo
}

// GetType implements ObjectList.
func (p *PartitionInfoList) GetType() object.ObjectType {
	return ObjectTypePartitionInfo
}

// GetItems implements ObjectList.
func (p *PartitionInfoList) GetItems() []object.Object {
	partitionInfos := make([]object.Object, 0)
	for _, partitionInfo := range p.Items {
		partitionInfos = append(partitionInfos, partitionInfo.DeepCopyObject())
	}
	return partitionInfos
}

// Size implements ObjectList.
func (p *PartitionInfoList) Size() int {
	return len(p.Items)
}

// AppendItem implements ObjectList.
func (p *PartitionInfoList) AppendItem(object object.Object) {
	partitionInfo := object.(*PartitionInfo)
	partitionInfo = partitionInfo.DeepCopy()
	p.Items = append(p.Items, *partitionInfo)
}

// DeepCopyObjectList implements ObjectList.
func (in *PartitionInfoList) DeepCopyObjectList() object.ObjectList {
	return in.DeepCopy()
}
