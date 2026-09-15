// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045User_GetKey(t *testing.T) {
	user := &V0045User{V0045User: api.V0045User{Name: "alice"}}
	if got := user.GetKey(); got != object.ObjectKey("alice") {
		t.Errorf("GetKey() = %q, want %q", got, "alice")
	}
	if got := user.GetType(); got != ObjectTypeV0045User {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0045User)
	}
}

func TestV0045UserList_GetItems(t *testing.T) {
	list := &V0045UserList{Items: []V0045User{
		{V0045User: api.V0045User{Name: "alice"}},
	}}
	if len(list.GetItems()) != 1 {
		t.Errorf("GetItems() len = %d, want 1", len(list.GetItems()))
	}
}
