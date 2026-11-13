// SPDX-FileCopyrightText: Copyright (C) NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	slurmcache "github.com/SlinkyProject/slurm-client/pkg/cache"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/event"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type orderedReply struct {
	jobs []api.V0044JobInfo
	err  error
}
type orderedReader struct {
	emptyClient
	gets, lists chan chan orderedReply
}

func (r *orderedReader) Get(ctx context.Context, _ object.ObjectKey, obj object.Object, _ ...GetOption) error {
	reply, err := r.read(ctx, r.gets)
	if err != nil {
		return err
	}
	if reply.err == nil {
		*obj.(*api.V0044JobInfo) = reply.jobs[0]
	}
	return reply.err
}
func (r *orderedReader) List(ctx context.Context, obj object.ObjectList, _ ...ListOption) error {
	reply, err := r.read(ctx, r.lists)
	if err != nil {
		return err
	}
	obj.(*api.V0044JobInfoObjectList).Items = reply.jobs
	return reply.err
}
func (*orderedReader) read(ctx context.Context, started chan chan orderedReply) (orderedReply, error) {
	response := make(chan orderedReply, 1)
	select {
	case started <- response:
	case <-ctx.Done():
		return orderedReply{}, ctx.Err()
	}
	select {
	case reply := <-response:
		return reply, nil
	case <-ctx.Done():
		return orderedReply{}, ctx.Err()
	}
}
func receiveOrdered[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case result := <-ch:
		return result
	case <-time.After(5 * time.Second):
		t.Fatal("fetch did not progress")
		var zero T
		return zero
	}
}
func orderingFixture(t *testing.T) (*informerCache, func(bool) (chan orderedReply, <-chan struct{})) {
	t.Helper()
	r := &orderedReader{gets: make(chan chan orderedReply), lists: make(chan chan orderedReply)}
	i := readyIndexInformer()
	i.reader = r
	if err := i.AddIndex("owners", jobOwners); err != nil {
		t.Fatal(err)
	}
	stop := make(chan struct{})
	var completed []<-chan struct{}
	t.Cleanup(func() {
		close(stop)
		for _, done := range completed {
			receiveOrdered(t, done)
		}
	})
	fetch := func(list bool) (chan orderedReply, <-chan struct{}) {
		done := make(chan struct{})
		started := r.gets
		if list {
			started = r.lists
			go func() { i.runListFetch(stop); close(done) }()
		} else {
			go func() { i.runGetFetch(stop, "1"); close(done) }()
		}
		completed = append(completed, done)
		return receiveOrdered(t, started), done
	}
	return i, fetch
}
func assertOrderingSnapshot(t *testing.T, i *informerCache, owner string, present bool) {
	t.Helper()
	jobs := &api.V0044JobInfoObjectList{}
	if err := i.List(t.Context(), jobs); err != nil {
		t.Fatal(err)
	}
	want := 0
	if present {
		want = 1
	}
	if len(jobs.Items) != want {
		t.Fatalf("cached jobs=%d, want %d", len(jobs.Items), want)
	}
	got := &api.V0044JobInfo{}
	err := i.Get(t.Context(), "1", got)
	if present {
		if err != nil || !reflect.DeepEqual(got, &jobs.Items[0]) {
			t.Fatalf("cached GET disagrees with list: job=%v err=%v", got, err)
		}
	} else if !errors.Is(err, apierrors.ErrNotFound) {
		t.Fatalf("cached GET after deletion: %v", err)
	}
	for _, key := range []string{"old", "new"} {
		ids := indexedIDs(t, i, key)
		expected := []int32{}
		if present && key == owner {
			expected = []int32{1}
		}
		if !reflect.DeepEqual(ids, expected) {
			t.Fatalf("index %q=%v, want %v", key, ids, expected)
		}
	}
	if len(i.inflight) != 0 || len(i.deleted) != 0 {
		t.Fatalf("completed requests retained ordering metadata: inflight=%d deleted=%d", len(i.inflight), len(i.deleted))
	}
}
func TestInformerOlderListPreservesReadback(t *testing.T) {
	for _, mode := range []string{"create", "update", "delete"} {
		t.Run(mode, func(t *testing.T) {
			i, fetch := orderingFixture(t)
			old := *indexJob(1, "old")
			fresh := *indexJob(1, "new")
			if mode != "create" {
				i.processObjects(&api.V0044JobInfoObjectList{Items: []api.V0044JobInfo{old}})
			}
			i.handler = slurmcache.ResourceEventHandlerFuncs{}
			list, listDone := fetch(true)
			get, getDone := fetch(false)
			if mode == "delete" {
				get <- orderedReply{err: apierrors.ErrNotFound}
			} else {
				get <- orderedReply{jobs: []api.V0044JobInfo{fresh}}
			}
			receiveOrdered(t, getDone)
			if mode == "create" {
				list <- orderedReply{}
			} else {
				list <- orderedReply{jobs: []api.V0044JobInfo{old}}
			}
			receiveOrdered(t, listDone)
			assertOrderingSnapshot(t, i, "new", mode != "delete")
			want := map[string]event.EventType{"create": event.Added, "update": event.Modified, "delete": event.Deleted}[mode]
			if got := receiveOrdered(t, i.eventCh); got.Type != want {
				t.Fatalf("event=%v, want %v", got.Type, want)
			}
			if len(i.eventCh) != 0 {
				t.Fatal("stale list emitted an event")
			}
			// A later authoritative list must still remove the object and its indexes.
			list, listDone = fetch(true)
			list <- orderedReply{}
			receiveOrdered(t, listDone)
			assertOrderingSnapshot(t, i, "", false)
		})
	}
}
func TestInformerOlderGetPreservesNewerList(t *testing.T) {
	for _, mode := range []string{"update", "delete", "not-found", "error", "failed-list"} {
		t.Run(mode, func(t *testing.T) {
			i, fetch := orderingFixture(t)
			old := *indexJob(1, "old")
			fresh := *indexJob(1, "new")
			i.processObjects(&api.V0044JobInfoObjectList{Items: []api.V0044JobInfo{old}})
			i.handler = slurmcache.ResourceEventHandlerFuncs{}
			get, getDone := fetch(false)
			list, listDone := fetch(true)
			switch mode {
			case "delete":
				list <- orderedReply{}
			case "failed-list":
				list <- orderedReply{err: errors.New("list unavailable")}
			default:
				list <- orderedReply{jobs: []api.V0044JobInfo{fresh}}
			}
			receiveOrdered(t, listDone)
			switch mode {
			case "not-found":
				get <- orderedReply{err: apierrors.ErrNotFound}
			case "error":
				get <- orderedReply{err: errors.New("get unavailable")}
			default:
				get <- orderedReply{jobs: []api.V0044JobInfo{old}}
			}
			receiveOrdered(t, getDone)
			if mode == "failed-list" {
				// The failed list must not suppress a successful older GET.
				got := &api.V0044JobInfo{}
				if err := i.Get(t.Context(), "1", got); err != nil || !reflect.DeepEqual(got, &old) {
					t.Fatalf("older GET after failed list: job=%v err=%v", got, err)
				}
				if len(i.eventCh) != 0 {
					t.Fatal("unchanged GET emitted an event")
				}
			} else {
				assertOrderingSnapshot(t, i, "new", mode != "delete")
				want := event.Modified
				if mode == "delete" {
					want = event.Deleted
				}
				if got := receiveOrdered(t, i.eventCh); got.Type != want {
					t.Fatalf("event=%v, want %v", got.Type, want)
				}
				if len(i.eventCh) != 0 {
					t.Fatal("stale GET emitted an event")
				}
			}
		})
	}
}
