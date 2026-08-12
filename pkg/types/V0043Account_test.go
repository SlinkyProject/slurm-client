// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043Account_GetKey(t *testing.T) {
	acct := &V0043Account{V0043Account: api.V0043Account{Name: "research"}}
	if got := acct.GetKey(); got != object.ObjectKey("research") {
		t.Errorf("GetKey() = %q, want %q", got, "research")
	}
	if got := acct.GetType(); got != ObjectTypeV0043Account {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0043Account)
	}
}

func TestV0043AccountList_GetItems(t *testing.T) {
	list := &V0043AccountList{Items: []V0043Account{
		{V0043Account: api.V0043Account{Name: "a"}},
		{V0043Account: api.V0043Account{Name: "b"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
