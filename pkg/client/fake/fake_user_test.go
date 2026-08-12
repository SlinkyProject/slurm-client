// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/client"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestFakeClient_V0042UserRoundTrip(t *testing.T) {
	ctx := context.Background()
	user := &types.V0042User{V0042User: apiv0042.V0042User{Name: "alice"}}
	c := NewClientBuilder().WithObjects(user).Build()

	got := &types.V0042User{}
	if err := c.Get(ctx, object.ObjectKey("alice"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "alice" {
		t.Errorf("got user %q, want alice", got.Name)
	}

	list := &types.V0042UserList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0043UserRoundTrip(t *testing.T) {
	ctx := context.Background()
	user := &types.V0043User{V0043User: apiv0043.V0043User{Name: "alice"}}
	c := NewClientBuilder().WithObjects(user).Build()

	got := &types.V0043User{}
	if err := c.Get(ctx, object.ObjectKey("alice"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "alice" {
		t.Errorf("got user %q, want alice", got.Name)
	}

	list := &types.V0043UserList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0044UserRoundTrip(t *testing.T) {
	ctx := context.Background()
	user := &types.V0044User{V0044User: apiv0044.V0044User{Name: "alice"}}
	c := NewClientBuilder().WithObjects(user).Build()

	got := &types.V0044User{}
	if err := c.Get(ctx, object.ObjectKey("alice"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "alice" {
		t.Errorf("got user %q, want alice", got.Name)
	}

	list := &types.V0044UserList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0045UserRoundTrip(t *testing.T) {
	ctx := context.Background()
	user := &types.V0045User{V0045User: apiv0045.V0045User{Name: "alice"}}
	c := NewClientBuilder().WithObjects(user).Build()

	got := &types.V0045User{}
	if err := c.Get(ctx, object.ObjectKey("alice"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "alice" {
		t.Errorf("got user %q, want alice", got.Name)
	}

	list := &types.V0045UserList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}
