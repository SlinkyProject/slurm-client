// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042Account_GetKey(t *testing.T) {
	acct := &V0042Account{V0042Account: api.V0042Account{Name: "research"}}
	if got := acct.GetKey(); got != object.ObjectKey("research") {
		t.Errorf("GetKey() = %q, want %q", got, "research")
	}
	if got := acct.GetType(); got != ObjectTypeV0042Account {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0042Account)
	}
}

func TestV0042AccountList_GetItems(t *testing.T) {
	list := &V0042AccountList{Items: []V0042Account{
		{V0042Account: api.V0042Account{Name: "a"}},
		{V0042Account: api.V0042Account{Name: "b"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
