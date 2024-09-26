// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

const (
	DefaultChanSize int32 = 100
)

// EventType defines the possible types of events.
type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
)

// Event represents a single event to a watched resource.
type Event struct {
	Type EventType

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before deletion.
	Object object.Object

	// ObjectOld is:
	//  * If Type is Modified: the old state of the object.
	//  * If Type is Added or Deleted: nil.
	ObjectOld object.Object
}
