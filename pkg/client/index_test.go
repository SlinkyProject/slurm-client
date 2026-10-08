// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/cache"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func indexJob(id int32, owners string) *api.V0044JobInfo {
	return &api.V0044JobInfo{
		JobId:        ptr.To(id),
		AdminComment: ptr.To(owners),
	}
}
func jobOwners(obj object.Object) []string {
	value := ptr.Deref(obj.(*api.V0044JobInfo).AdminComment, "")
	if value == "" {
		return nil
	}
	return strings.Split(value, ",")
}
func readyIndexInformer() *informerCache {
	i := newInformer(api.ObjectTypeV0044JobInfo, &emptyClient{}, defaultSyncPeriod).(*informerCache)
	i.started, i.dirty = true, false
	return i
}
func indexedIDs(t *testing.T, i *informerCache, owner string) []int32 {
	t.Helper()
	jobs := &api.V0044JobInfoObjectList{}
	if err := i.ByIndex(context.Background(), "owners", owner, jobs); err != nil {
		t.Fatal(err)
	}
	ids := []int32{}
	for _, j := range jobs.Items {
		ids = append(ids, *j.JobId)
	}
	slices.Sort(ids)
	return ids
}
func TestInformerIndexLifecycle(t *testing.T) {
	i := readyIndexInformer()
	i.processObject(indexJob(1, "a,b,a"))
	calls := 0
	if err := i.AddIndex("owners", func(o object.Object) []string { calls++; return jobOwners(o) }); err != nil {
		t.Fatal(err)
	}
	check := func(value string, want ...int32) {
		t.Helper()
		if got := indexedIDs(t, i, value); !slices.Equal(got, want) {
			t.Fatalf("%q: got %v want %v", value, got, want)
		}
	}
	check("a", 1)
	check("b", 1)
	check("missing")
	if calls != 1 {
		t.Fatalf("backfill and reads extracted %d times", calls)
	}
	i.processObject(indexJob(1, "a,b,a"))
	if calls != 1 {
		t.Fatal("unchanged object reindexed")
	}
	i.processObject(indexJob(2, "a"))
	check("a", 1, 2)
	i.processObject(indexJob(1, "c"))
	check("a", 2)
	check("b")
	check("c", 1)
	i.processObject(indexJob(1, ""))
	check("c")
	i.processObjects(&api.V0044JobInfoObjectList{Items: []api.V0044JobInfo{*indexJob(1, "d")}})
	check("a")
	check("d", 1)
	i.processObjects(&api.V0044JobInfoObjectList{})
	check("d")
	if len(i.indexes["owners"].values) != 0 || len(i.indexes["owners"].keys) != 0 {
		t.Fatal("deleted objects left index entries")
	}
}
func TestInformerIndexValidationAndCopies(t *testing.T) {
	i := readyIndexInformer()
	for _, test := range []struct {
		name string
		fn   IndexFunc
	}{{"", jobOwners}, {"bad", nil}} {
		if err := i.AddIndex(test.name, test.fn); err == nil {
			t.Fatal("invalid index accepted")
		}
	}
	if err := i.AddIndex("owners", jobOwners); err != nil {
		t.Fatal(err)
	}
	if err := i.AddIndex("owners", jobOwners); err == nil {
		t.Fatal("duplicate index accepted")
	}
	if err := i.ByIndex(t.Context(), "unknown", "a", &api.V0044JobInfoObjectList{}); err == nil {
		t.Fatal("unknown index accepted")
	}
	if err := i.ByIndex(t.Context(), "owners", "a", &api.V0044NodeObjectList{}); err == nil {
		t.Fatal("wrong list type accepted")
	}
	source := indexJob(1, "a")
	i.processObject(source)
	*source.AdminComment = "source mutation"
	jobs := &api.V0044JobInfoObjectList{}
	if err := i.ByIndex(t.Context(), "owners", "a", jobs); err != nil {
		t.Fatal(err)
	}
	*jobs.Items[0].AdminComment = "indexed mutation"
	got := &api.V0044JobInfo{}
	if err := i.Get(t.Context(), "1", got); err != nil {
		t.Fatal(err)
	}
	if *got.AdminComment != "a" {
		t.Fatal("indexed result aliases cache")
	}
	*got.AdminComment = "get mutation"
	all := &api.V0044JobInfoObjectList{}
	if err := i.List(t.Context(), all); err != nil {
		t.Fatal(err)
	}
	if *all.Items[0].AdminComment != "a" {
		t.Fatal("Get result aliases cache")
	}
	*all.Items[0].AdminComment = "list mutation"
	again := &api.V0044JobInfoObjectList{}
	if err := i.ByIndex(t.Context(), "owners", "a", again); err != nil {
		t.Fatal(err)
	}
	if *again.Items[0].AdminComment != "a" {
		t.Fatal("List result aliases cache")
	}
	i.started = false
	if err := i.ByIndex(t.Context(), "owners", "a", again); err == nil {
		t.Fatal("unsynced index accepted")
	}
	i.started = true
	i.dirty = true
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := i.ByIndex(ctx, "owners", "a", again); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation: %v", err)
	}
}

type missingIndexReader struct{ emptyClient }

func (*missingIndexReader) Get(context.Context, object.ObjectKey, object.Object, ...GetOption) error {
	return apierrors.ErrNotFound
}
func TestInformerIndexGetNotFoundRemovesObject(t *testing.T) {
	i := readyIndexInformer()
	i.reader = &missingIndexReader{}
	if err := i.AddIndex("owners", jobOwners); err != nil {
		t.Fatal(err)
	}
	i.processObject(indexJob(1, "a"))
	i.doGetInformer("1")
	if got := indexedIDs(t, i, "a"); len(got) != 0 {
		t.Fatalf("deleted object still indexed: %v", got)
	}
	if len(i.cache) != 0 {
		t.Fatalf("NotFound added an empty object: %v", i.cache)
	}
}
func TestInformerIndexConcurrentUpdates(t *testing.T) {
	i := readyIndexInformer()
	if err := i.AddIndex("owners", jobOwners); err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for worker := range 5 {
		wg.Go(func() {
			for n := range 300 {
				if worker == 0 {
					i.mu.Lock()
					if n%3 == 0 {
						i.processObjects(&api.V0044JobInfoObjectList{})
					} else {
						i.processObject(indexJob(1, fmt.Sprint(n%2)))
					}
					i.mu.Unlock()
				} else {
					value := fmt.Sprint(n % 2)
					jobs := &api.V0044JobInfoObjectList{}
					if err := i.ByIndex(t.Context(), "owners", value, jobs); err != nil {
						t.Error(err)
						return
					}
					for _, j := range jobs.Items {
						if *j.AdminComment != value {
							t.Errorf("object/index mismatch: %v", j)
						}
					}
				}
			}
		})
	}
	wg.Wait()
}
func TestInformerUpdateEventOldAndNew(t *testing.T) {
	i := readyIndexInformer()
	i.processObject(indexJob(1, "old"))
	i.SetEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: func(oldObj, newObj any) {
		if *oldObj.(*api.V0044JobInfo).AdminComment != "old" || *newObj.(*api.V0044JobInfo).AdminComment != "new" {
			t.Fatal("incorrect update event objects")
		}
	}})
	i.processObject(indexJob(1, "new"))
	i.doHandler(<-i.eventCh)
}

func BenchmarkInformerByIndex(b *testing.B) {
	for _, size := range []int{800, 8000, 80000} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			i := readyIndexInformer()
			for n := range size {
				i.processObject(indexJob(int32(n+1), fmt.Sprint(n)))
			}
			if err := i.AddIndex("owners", jobOwners); err != nil {
				b.Fatal(err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				jobs := &api.V0044JobInfoObjectList{}
				if err := i.ByIndex(b.Context(), "owners", "0", jobs); err != nil {
					b.Fatal(err)
				}
				if len(jobs.Items) != 1 {
					b.Fatal("missing match")
				}
			}
		})
	}
}
