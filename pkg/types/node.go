// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

// +kubebuilder:object:generate=true
package types

import (
	"bytes"
	"encoding/json"
	"reflect"

	"k8s.io/utils/set"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	ObjectTypeNode = "node"
)

// Represents the Node State
// +enum
type NodeState string

const (
	// Base States: mutually exclusive, must only have one
	NodeStateALLOCATED NodeState = "ALLOCATED"
	NodeStateDOWN      NodeState = "DOWN"
	NodeStateERROR     NodeState = "ERROR"
	NodeStateFUTURE    NodeState = "FUTURE"
	NodeStateIDLE      NodeState = "IDLE"
	NodeStateMIXED     NodeState = "MIXED"
	NodeStateUNKNOWN   NodeState = "UNKNOWN"

	// Flag States: not mutually exclusive, can have zero or more
	NodeStateCLOUD           NodeState = "CLOUD"
	NodeStateCOMPLETING      NodeState = "COMPLETING"
	NodeStateDRAIN           NodeState = "DRAIN"
	NodeStateDYNAMICFUTURE   NodeState = "DYNAMIC_FUTURE"
	NodeStateDYNAMICNORM     NodeState = "DYNAMIC_NORM"
	NodeStateFAIL            NodeState = "FAIL"
	NodeStateINVALID         NodeState = "INVALID"
	NodeStateINVALIDREG      NodeState = "INVALID_REG"
	NodeStateMAINTENANCE     NodeState = "MAINTENANCE"
	NodeStateNOTRESPONDING   NodeState = "NOT_RESPONDING"
	NodeStatePERFCTRS        NodeState = "PERFCTRS"
	NodeStatePLANNED         NodeState = "PLANNED"
	NodeStatePOWERDOWN       NodeState = "POWER_DOWN"
	NodeStatePOWERDRAIN      NodeState = "POWER_DRAIN"
	NodeStatePOWEREDDOWN     NodeState = "POWERED_DOWN"
	NodeStatePOWERINGDOWN    NodeState = "POWERING_DOWN"
	NodeStatePOWERINGUP      NodeState = "POWERING_UP"
	NodeStatePOWERUP         NodeState = "POWER_UP"
	NodeStateREBOOTCANCELED  NodeState = "REBOOT_CANCELED"
	NodeStateREBOOTISSUED    NodeState = "REBOOT_ISSUED"
	NodeStateREBOOTREQUESTED NodeState = "REBOOT_REQUESTED"
	NodeStateRESERVED        NodeState = "RESERVED"
	NodeStateRESUME          NodeState = "RESUME"
	NodeStateUNDRAIN         NodeState = "UNDRAIN"
)

// Node represents a Slurm compute node.
type Node struct {
	Address       string
	Comment       string
	Extra         string
	Name          string
	Partitions    set.Set[string]
	State         set.Set[NodeState]
	Cpus          int32
	AllocCpus     int32
	AllocIdleCpus int32
}

type NodeInfo struct {
	Namespace string `json:"namespace"`
	PodName   string `json:"podName"`
}

func (nodeInfo *NodeInfo) Equal(cmp NodeInfo) bool {
	a, _ := json.Marshal(nodeInfo)
	b, _ := json.Marshal(cmp)
	return bytes.Equal(a, b)
}

func (nodeInfo *NodeInfo) ToString() string {
	b, _ := json.Marshal(nodeInfo)
	return string(b)
}

func NodeInfoParse(str string, out *NodeInfo) error {
	return json.Unmarshal([]byte(str), &out)
}

// GetKey implements Object.
func (n *Node) GetKey() object.ObjectKey {
	if n.Name == "" {
		panic("Name cannot be empty")
	}
	return object.ObjectKey(n.Name)
}

// GetType implements Object.
func (n *Node) GetType() object.ObjectType {
	return ObjectTypeNode
}

// DeepEqualObject implements Object.
func (n *Node) DeepEqualObject(object object.Object) bool {
	return reflect.DeepEqual(n, object)
}

// DeepCopyObject implements Object.
func (n *Node) DeepCopyObject() object.Object {
	return n.DeepCopy()
}

// NodeList represents a list of Slurm compute nodes.
type NodeList struct {
	Items []Node
}

// GetType implements ObjectList.
func (n *NodeList) GetType() object.ObjectType {
	return ObjectTypeNode
}

// GetItems implements ObjectList.
func (n *NodeList) GetItems() []object.Object {
	nodes := make([]object.Object, 0)
	for _, node := range n.Items {
		nodes = append(nodes, node.DeepCopyObject())
	}
	return nodes
}

// Size implements ObjectList.
func (n *NodeList) Size() int {
	return len(n.Items)
}

// AppendItem implements ObjectList.
func (n *NodeList) AppendItem(object object.Object) {
	node := object.(*Node)
	node = node.DeepCopy()
	n.Items = append(n.Items, *node)
}

// DeepCopyObjectList implements ObjectList.
func (n *NodeList) DeepCopyObjectList() object.ObjectList {
	return n.DeepCopy()
}
