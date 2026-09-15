// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045Account_GetKey(t *testing.T) {
	acct := &V0045Account{V0045Account: api.V0045Account{Name: "research"}}
	if got := acct.GetKey(); got != object.ObjectKey("research") {
		t.Errorf("GetKey() = %q, want %q", got, "research")
	}
	if got := acct.GetType(); got != ObjectTypeV0045Account {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0045Account)
	}
}

func TestV0045AccountList_GetItems(t *testing.T) {
	list := &V0045AccountList{Items: []V0045Account{
		{V0045Account: api.V0045Account{Name: "a"}},
		{V0045Account: api.V0045Account{Name: "b"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
