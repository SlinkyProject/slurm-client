// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"strconv"
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
