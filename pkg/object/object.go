// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package object

type ObjectType string

type ObjectKey string

type RuntimeObject interface {
	GetType() ObjectType
}

type Object interface {
	RuntimeObject
	GetKey() ObjectKey
	DeepCopyObject() Object
}

type ObjectList interface {
	RuntimeObject
	GetItems() []Object
	AppendItem(Object)
	DeepCopyObjectList() ObjectList
}
