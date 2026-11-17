// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type slowListClient struct {
	emptyClient
	calls atomic.Int32
}

func (r *slowListClient) List(ctx context.Context, list object.ObjectList, _ ...ListOption) error {
	n := r.calls.Add(1)
	select {
	case <-time.After(150 * time.Millisecond):
		list.AppendItem(&api.V0044JobInfo{JobId: ptr.To(int32(42)), Name: ptr.To(strconv.Itoa(int(n)))})
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestExplicitListRefreshStillWaitsForSubsequentResponse(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		i := newInformer(api.ObjectTypeV0044JobInfo, &slowListClient{}, 100*time.Millisecond).(*informerCache)
		stop := make(chan struct{})
		defer close(stop)
		go i.Run(stop)
		// The first list started at 100ms and completes at 250ms. A refresh
		// requested now must wait for the second response, at 400ms.
		time.Sleep(230 * time.Millisecond)
		jobs := &api.V0044JobInfoObjectList{}
		if err := i.List(t.Context(), jobs, &ListOptions{RefreshCache: true}); err != nil {
			t.Fatal(err)
		}
		if len(jobs.Items) != 1 || ptr.Deref(jobs.Items[0].Name, "") != "2" {
			t.Fatalf("explicit refresh returned an earlier response: %v", jobs.Items)
		}
	})
}

func TestPeriodicRefreshDoesNotStarveIndexReaders(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		i := newInformer(api.ObjectTypeV0044JobInfo, &slowListClient{}, 100*time.Millisecond).(*informerCache)
		if err := i.AddIndex("all", func(object.Object) []string { return []string{"all"} }); err != nil {
			t.Fatal(err)
		}
		stop := make(chan struct{})
		defer close(stop)
		go i.Run(stop)
		// The first periodic list is in flight. Its response takes longer than
		// the polling interval, but less than the cache read timeout.
		time.Sleep(160 * time.Millisecond)
		jobs := &api.V0044JobInfoObjectList{}
		if err := i.ByIndex(t.Context(), "all", "all", jobs); err != nil {
			t.Fatalf("periodic refresh starved the cached read: %v", err)
		}
		if len(jobs.Items) != 1 || ptr.Deref(jobs.Items[0].JobId, 0) != 42 {
			t.Fatalf("unexpected cached jobs: %v", jobs.Items)
		}
	})
}

func TestPeriodicRefreshDoesNotBlockReaders(t *testing.T) {
	for _, method := range []string{"List", "ByIndex"} {
		t.Run(method, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				i := newInformer(api.ObjectTypeV0044JobInfo, &slowListClient{}, 100*time.Millisecond).(*informerCache)
				if err := i.AddIndex("all", func(object.Object) []string { return []string{"all"} }); err != nil {
					t.Fatal(err)
				}
				stop := make(chan struct{})
				defer close(stop)
				go i.Run(stop)

				read := func(wantName string) {
					t.Helper()
					ctx, cancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
					defer cancel()
					jobs := &api.V0044JobInfoObjectList{}
					var err error
					if method == "List" {
						err = i.List(ctx, jobs)
					} else {
						err = i.ByIndex(ctx, "all", "all", jobs)
					}
					if err != nil {
						t.Fatalf("cached read waited for a periodic refresh: %v", err)
					}
					if len(jobs.Items) != 1 || ptr.Deref(jobs.Items[0].JobId, 0) != 42 || ptr.Deref(jobs.Items[0].Name, "") != wantName {
						t.Fatalf("cached jobs = %v, want job 42 from response %s", jobs.Items, wantName)
					}
				}

				// The first response landed at 250ms. The second list started
				// at 300ms and will not finish until 450ms, beyond the read deadline.
				time.Sleep(360 * time.Millisecond)
				read("1")
				// Once the background list lands, readers see its updated snapshot.
				time.Sleep(100 * time.Millisecond)
				read("2")
			})
		})
	}
}

type blockedItemsList struct {
	api.V0044JobInfoObjectList
	entered chan struct{}
	proceed chan struct{}
}

func (l *blockedItemsList) GetItems() []object.Object {
	close(l.entered)
	<-l.proceed
	return l.V0044JobInfoObjectList.GetItems()
}

// Periodic refresh preparation must leave the previous snapshot available.
// Block the copy explicitly so this tests locking, not machine-dependent speed.
func TestPeriodicListPreparationDoesNotBlockReaders(t *testing.T) {
	i := readyIndexInformer()
	i.processObject(indexJob(1, "old"))
	if err := i.AddIndex("owners", jobOwners); err != nil {
		t.Fatal(err)
	}
	list := &blockedItemsList{
		V0044JobInfoObjectList: api.V0044JobInfoObjectList{Items: []api.V0044JobInfo{*indexJob(1, "new")}},
		entered:                make(chan struct{}), proceed: make(chan struct{}),
	}
	original := resourceCatalog[api.ObjectTypeV0044JobInfo]
	r := original
	r.newList = func() object.ObjectList { return list }
	resourceCatalog[api.ObjectTypeV0044JobInfo] = r
	t.Cleanup(func() { resourceCatalog[api.ObjectTypeV0044JobInfo] = original })
	unblock := sync.OnceFunc(func() { close(list.proceed) })
	done := make(chan struct{})
	stop := make(chan struct{})
	go func() { i.runListFetch(stop); close(done) }()
	t.Cleanup(func() { unblock(); <-done; close(stop) })
	select {
	case <-list.entered:
	case <-time.After(time.Second):
		t.Fatal("refresh did not prepare list")
	}
	read := make(chan error, 1)
	jobs := &api.V0044JobInfoObjectList{}
	go func() { read <- i.ByIndex(t.Context(), "owners", "old", jobs) }()
	select {
	case err := <-read:
		if err != nil || len(jobs.Items) != 1 {
			t.Fatalf("old snapshot unavailable during preparation: jobs=%d, err=%v", len(jobs.Items), err)
		}
	case <-time.After(time.Second):
		unblock()
		<-read
		t.Fatal("list preparation held the cache lock")
	}
	unblock()
	<-done
	if got := indexedIDs(t, i, "new"); len(got) != 1 || got[0] != 1 {
		t.Fatalf("new snapshot was not published: %v", got)
	}
	if got := indexedIDs(t, i, "old"); len(got) != 0 {
		t.Fatalf("old index survived publication: %v", got)
	}
}
