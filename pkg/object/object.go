// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package object

type ObjectType string

type ObjectKey string

type Object interface {
	GetKey() ObjectKey
	GetType() ObjectType
	DeepEqualObject(Object) bool
	DeepCopyObject() Object
}

type ObjectList interface {
	GetType() ObjectType
	GetItems() []Object
	AppendItem(Object)
	DeepCopyObjectList() ObjectList
}
