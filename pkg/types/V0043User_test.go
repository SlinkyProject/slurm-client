// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043User_GetKey(t *testing.T) {
	user := &V0043User{V0043User: api.V0043User{Name: "alice"}}
	if got := user.GetKey(); got != object.ObjectKey("alice") {
		t.Errorf("GetKey() = %q, want %q", got, "alice")
	}
	if got := user.GetType(); got != ObjectTypeV0043User {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0043User)
	}
}

func TestV0043UserList_GetItems(t *testing.T) {
	list := &V0043UserList{Items: []V0043User{
		{V0043User: api.V0043User{Name: "alice"}},
	}}
	if len(list.GetItems()) != 1 {
		t.Errorf("GetItems() len = %d, want 1", len(list.GetItems()))
	}
}
