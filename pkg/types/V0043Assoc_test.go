// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043Assoc_GetKey(t *testing.T) {
	assoc := &V0043Assoc{V0043Assoc: api.V0043Assoc{
		Cluster:   ptr.To("linux"),
		Account:   ptr.To("research"),
		User:      "alice",
		Partition: ptr.To("debug"),
	}}
	want := object.ObjectKey("linux/research/alice/debug")
	if got := assoc.GetKey(); got != want {
		t.Errorf("GetKey() = %q, want %q", got, want)
	}
	if got := assoc.GetType(); got != ObjectTypeV0043Assoc {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0043Assoc)
	}
}

func TestV0043AssocList_GetItems(t *testing.T) {
	list := &V0043AssocList{Items: []V0043Assoc{
		{V0043Assoc: api.V0043Assoc{User: "alice"}},
		{V0043Assoc: api.V0043Assoc{User: "bob"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
