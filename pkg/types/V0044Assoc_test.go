// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Assoc_GetKey(t *testing.T) {
	assoc := &V0044Assoc{V0044Assoc: api.V0044Assoc{
		Cluster:   ptr.To("linux"),
		Account:   ptr.To("research"),
		User:      "alice",
		Partition: ptr.To("debug"),
	}}
	want := object.ObjectKey("linux/research/alice/debug")
	if got := assoc.GetKey(); got != want {
		t.Errorf("GetKey() = %q, want %q", got, want)
	}
	if got := assoc.GetType(); got != ObjectTypeV0044Assoc {
		t.Errorf("GetType() = %q, want %q", got, ObjectTypeV0044Assoc)
	}
}

func TestV0044AssocList_GetItems(t *testing.T) {
	list := &V0044AssocList{Items: []V0044Assoc{
		{V0044Assoc: api.V0044Assoc{User: "alice"}},
		{V0044Assoc: api.V0044Assoc{User: "bob"}},
	}}
	if len(list.GetItems()) != 2 {
		t.Errorf("GetItems() len = %d, want 2", len(list.GetItems()))
	}
}
