// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

type sequenceReader struct {
	emptyClient
	calls   atomic.Int32
	entered chan int32
	release [2]chan struct{}
}

func (r *sequenceReader) Get(ctx context.Context, key object.ObjectKey, obj object.Object, _ ...GetOption) error {
	n := r.calls.Add(1)
	r.entered <- n
	if n <= 2 {
		select {
		case <-r.release[n-1]:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	if err := obj.SetKey(key); err != nil {
		return err
	}
	obj.(*api.V0044JobInfo).Name = ptr.To(string(rune('0' + n)))
	return nil
}

func TestSequenceReleasesCompletedBatchWhileNextFetchIsBlocked(t *testing.T) {
	reader := &sequenceReader{entered: make(chan int32, 8), release: [2]chan struct{}{make(chan struct{}), make(chan struct{})}}
	i := newInformer(api.ObjectTypeV0044JobInfo, reader, defaultSyncPeriod).(*informerCache)
	stop := make(chan struct{})
	defer close(stop)
	go i.Run(stop)
	for !i.HasStarted() {
		time.Sleep(time.Millisecond)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()
	type answer struct {
		name string
		err  error
	}
	read := func() chan answer {
		done := make(chan answer, 1)
		go func() {
			job := &api.V0044JobInfo{}
			err := i.Get(ctx, "42", job, &GetOptions{RefreshCache: true})
			done <- answer{ptr.Deref(job.Name, ""), err}
		}()
		return done
	}
	first := read()
	select {
	case <-reader.entered:
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
	second := read()
	for {
		i.mu.RLock()
		pending := i.getSequences["42"].requested == 2 && i.fetching["42"]
		i.mu.RUnlock()
		if pending {
			break
		}
		if ctx.Err() != nil {
			t.Fatal(ctx.Err())
		}
		time.Sleep(time.Millisecond)
	}
	close(reader.release[0])
	select {
	case got := <-first:
		if got.err != nil || got.name != "1" {
			t.Fatalf("first result: %+v", got)
		}
	case <-ctx.Done():
		t.Fatal("first caller waited for a later fetch")
	}
	select {
	case <-reader.entered:
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
	select {
	case got := <-second:
		t.Fatalf("second caller returned before its fetch: %+v", got)
	default:
	}
	close(reader.release[1])
	select {
	case got := <-second:
		if got.err != nil || got.name != "2" {
			t.Fatalf("second result: %+v", got)
		}
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}
}

func TestSequenceStopReleasesCurrentAndPendingWaiters(t *testing.T) {
	reader := &sequenceReader{entered: make(chan int32, 8), release: [2]chan struct{}{make(chan struct{}), make(chan struct{})}}
	i := newInformer(api.ObjectTypeV0044JobInfo, reader, defaultSyncPeriod).(*informerCache)
	stop := make(chan struct{})
	go i.Run(stop)
	for !i.HasStarted() {
		time.Sleep(time.Millisecond)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()
	done := make(chan error, 2)
	read := func() { done <- i.Get(ctx, "42", &api.V0044JobInfo{}, &GetOptions{RefreshCache: true}) }
	go read()
	select {
	case <-reader.entered:
	case <-ctx.Done():
		close(stop)
		t.Fatal(ctx.Err())
	}
	go read()
	for {
		i.mu.RLock()
		pending := i.getSequences["42"].requested == 2
		i.mu.RUnlock()
		if pending {
			break
		}
		if ctx.Err() != nil {
			close(stop)
			t.Fatal(ctx.Err())
		}
		time.Sleep(time.Millisecond)
	}
	close(stop)
	for range 2 {
		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected shutdown cancellation, got %v", err)
			}
		case <-ctx.Done():
			t.Fatal("waiter remained blocked after shutdown")
		}
	}
}
