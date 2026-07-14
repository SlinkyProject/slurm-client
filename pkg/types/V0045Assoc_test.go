// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045Assoc_GetKey(t *testing.T) {
	assoc := &V0045Assoc{V0045Assoc: api.V0045Assoc{
		Cluster:   ptr.To("linux"),
		Account:   ptr.To("research"),
		User:      "alice",
		Partition: ptr.To("debug"),
	}}
	want := object.ObjectKey("linux/research/alice/debug")
	if got := assoc.GetKey(); got != want {
		t.Errorf("GetKey() = %q, want %q", got, want)
	}
	if got := assoc.GetType(); got != ObjectTypeV0045Assoc {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0045Assoc)
	}
}

func TestV0045AssocList_GetItems(t *testing.T) {
	list := &V0045AssocList{Items: []V0045Assoc{
		{V0045Assoc: api.V0045Assoc{User: "alice"}},
		{V0045Assoc: api.V0045Assoc{User: "bob"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
