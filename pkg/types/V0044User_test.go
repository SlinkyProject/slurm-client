// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044User_GetKey(t *testing.T) {
	user := &V0044User{V0044User: api.V0044User{Name: "alice"}}
	if got := user.GetKey(); got != object.ObjectKey("alice") {
		t.Errorf("GetKey() = %q, want %q", got, "alice")
	}
	if got := user.GetType(); got != ObjectTypeV0044User {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0044User)
	}
}

func TestV0044UserList_GetItems(t *testing.T) {
	list := &V0044UserList{Items: []V0044User{
		{V0044User: api.V0044User{Name: "alice"}},
	}}
	if len(list.GetItems()) != 1 {
		t.Errorf("GetItems() len = %d, want 1", len(list.GetItems()))
	}
}
