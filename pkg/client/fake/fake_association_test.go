// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	"k8s.io/utils/ptr"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestFakeClient_V0042AssocRoundTrip(t *testing.T) {
	ctx := context.Background()
	assoc := &types.V0042Assoc{V0042Assoc: apiv0042.V0042Assoc{
		Cluster: ptr.To("linux"),
		Account: ptr.To("research"),
		User:    "alice",
	}}
	c := NewClientBuilder().WithObjects(assoc).Build()

	got := &types.V0042Assoc{}
	if err := c.Get(ctx, assoc.GetKey(), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.User != "alice" {
		t.Errorf("got assoc user %q, want alice", got.User)
	}
}

func TestFakeClient_V0043AssocRoundTrip(t *testing.T) {
	ctx := context.Background()
	assoc := &types.V0043Assoc{V0043Assoc: apiv0043.V0043Assoc{
		Cluster: ptr.To("linux"),
		Account: ptr.To("research"),
		User:    "alice",
	}}
	c := NewClientBuilder().WithObjects(assoc).Build()

	got := &types.V0043Assoc{}
	if err := c.Get(ctx, assoc.GetKey(), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.User != "alice" {
		t.Errorf("got assoc user %q, want alice", got.User)
	}
}

func TestFakeClient_V0044AssocRoundTrip(t *testing.T) {
	ctx := context.Background()
	assoc := &types.V0044Assoc{V0044Assoc: apiv0044.V0044Assoc{
		Cluster: ptr.To("linux"),
		Account: ptr.To("research"),
		User:    "alice",
	}}
	c := NewClientBuilder().WithObjects(assoc).Build()

	got := &types.V0044Assoc{}
	if err := c.Get(ctx, assoc.GetKey(), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.User != "alice" {
		t.Errorf("got assoc user %q, want alice", got.User)
	}
}

func TestFakeClient_V0045AssocRoundTrip(t *testing.T) {
	ctx := context.Background()
	assoc := &types.V0045Assoc{V0045Assoc: apiv0045.V0045Assoc{
		Cluster: ptr.To("linux"),
		Account: ptr.To("research"),
		User:    "alice",
	}}
	c := NewClientBuilder().WithObjects(assoc).Build()

	got := &types.V0045Assoc{}
	if err := c.Get(ctx, assoc.GetKey(), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.User != "alice" {
		t.Errorf("got assoc user %q, want alice", got.User)
	}
}
