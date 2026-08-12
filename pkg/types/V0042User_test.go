// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042User_GetKey(t *testing.T) {
	user := &V0042User{V0042User: api.V0042User{Name: "alice"}}
	if got := user.GetKey(); got != object.ObjectKey("alice") {
		t.Errorf("GetKey() = %q, want %q", got, "alice")
	}
	if got := user.GetType(); got != ObjectTypeV0042User {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0042User)
	}
}

func TestV0042UserList_GetItems(t *testing.T) {
	list := &V0042UserList{Items: []V0042User{
		{V0042User: api.V0042User{Name: "alice"}},
	}}
	if len(list.GetItems()) != 1 {
		t.Errorf("GetItems() len = %d, want 1", len(list.GetItems()))
	}
}
