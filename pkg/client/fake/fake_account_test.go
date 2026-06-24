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

func TestFakeClient_V0042AccountRoundTrip(t *testing.T) {
	ctx := context.Background()
	acct := &types.V0042Account{V0042Account: apiv0042.V0042Account{Name: "research"}}
	c := NewClientBuilder().WithObjects(acct).Build()

	got := &types.V0042Account{}
	if err := c.Get(ctx, object.ObjectKey("research"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "research" {
		t.Errorf("got account %q, want research", got.Name)
	}

	list := &types.V0042AccountList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0043AccountRoundTrip(t *testing.T) {
	ctx := context.Background()
	acct := &types.V0043Account{V0043Account: apiv0043.V0043Account{Name: "research"}}
	c := NewClientBuilder().WithObjects(acct).Build()

	got := &types.V0043Account{}
	if err := c.Get(ctx, object.ObjectKey("research"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "research" {
		t.Errorf("got account %q, want research", got.Name)
	}

	list := &types.V0043AccountList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0044AccountRoundTrip(t *testing.T) {
	ctx := context.Background()
	acct := &types.V0044Account{V0044Account: apiv0044.V0044Account{Name: "research"}}
	c := NewClientBuilder().WithObjects(acct).Build()

	got := &types.V0044Account{}
	if err := c.Get(ctx, object.ObjectKey("research"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "research" {
		t.Errorf("got account %q, want research", got.Name)
	}

	list := &types.V0044AccountList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}

func TestFakeClient_V0045AccountRoundTrip(t *testing.T) {
	ctx := context.Background()
	acct := &types.V0045Account{V0045Account: apiv0045.V0045Account{Name: "research"}}
	c := NewClientBuilder().WithObjects(acct).Build()

	got := &types.V0045Account{}
	if err := c.Get(ctx, object.ObjectKey("research"), got, &client.GetOptions{SkipCache: true}); err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.Name != "research" {
		t.Errorf("got account %q, want research", got.Name)
	}

	list := &types.V0045AccountList{}
	if err := c.List(ctx, list); err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(list.Items) != 1 {
		t.Errorf("List() len = %d, want 1", len(list.Items))
	}
}
