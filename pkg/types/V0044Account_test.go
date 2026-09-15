// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Account_GetKey(t *testing.T) {
	acct := &V0044Account{V0044Account: api.V0044Account{Name: "research"}}
	if got := acct.GetKey(); got != object.ObjectKey("research") {
		t.Errorf("GetKey() = %q, want %q", got, "research")
	}
	if got := acct.GetType(); got != ObjectTypeV0044Account {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0044Account)
	}
}

func TestV0044AccountList_GetItems(t *testing.T) {
	list := &V0044AccountList{Items: []V0044Account{
		{V0044Account: api.V0044Account{Name: "a"}},
		{V0044Account: api.V0044Account{Name: "b"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
