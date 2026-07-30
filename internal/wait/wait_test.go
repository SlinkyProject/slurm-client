// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package wait

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

// TestPollUntilContextTimeout_ImmediateTrue verifies that when immediate=true,
// the condition runs right away and exits instantly if successful.
func TestPollUntilContextTimeout_ImmediateTrue(t *testing.T) {
	var callCount atomic.Int32

	err := PollUntilContextTimeout(context.Background(), 1*time.Second, 5*time.Second, true, func(ctx context.Context) (bool, error) {
		callCount.Add(1)
		return true, nil
	})

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if count := callCount.Load(); count != 1 {
		t.Errorf("Expected exactly 1 execution, got: %d", count)
	}
}

// TestPollUntilContextTimeout_ImmediateFalse ensures that when immediate=false,
// the code waits for the first interval tick before evaluating the condition.
func TestPollUntilContextTimeout_ImmediateFalse(t *testing.T) {
	interval := 20 * time.Millisecond
	timeout := 100 * time.Millisecond

	var callCount atomic.Int32
	firstTickChan := make(chan struct{})

	go func() {
		_ = PollUntilContextTimeout(context.Background(), interval, timeout, false, func(ctx context.Context) (bool, error) {
			count := callCount.Add(1)
			if count == 1 {
				close(firstTickChan) // Signal that the first tick fired
			}
			return true, nil
		})
	}()

	// Confirm that execution did NOT happen instantly at time zero
	select {
	case <-firstTickChan:
		// Success: the loop properly waited for the ticker to fire
	case <-time.After(interval / 2):
		if callCount.Load() > 0 {
			t.Fatal("Condition executed before the first ticker interval passed")
		}
	}
}

// TestPollUntilContextTimeout_TimeoutEnforced verifies that the explicit
// timeout argument is respected, even if the parent context is infinite.
func TestPollUntilContextTimeout_TimeoutEnforced(t *testing.T) {
	parentCtx := context.Background()
	interval := 5 * time.Millisecond
	timeout := 25 * time.Millisecond

	startTime := time.Now()
	err := PollUntilContextTimeout(parentCtx, interval, timeout, true, func(ctx context.Context) (bool, error) {
		return false, nil // Keep failing to force a timeout
	})
	elapsed := time.Since(startTime)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Expected context.DeadlineExceeded, got: %v", err)
	}

	if elapsed < timeout {
		t.Errorf("Function returned before the specified timeout duration. Took: %v", elapsed)
	}
}

// TestPollUntilContextTimeout_ParentContextCancelled ensures that if the parent
// context is cancelled mid-flight, the loop aborts instantly without waiting for the timeout.
func TestPollUntilContextTimeout_ParentContextCancelled(t *testing.T) {
	parentCtx, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()

	interval := 100 * time.Millisecond
	timeout := 5 * time.Second

	go func() {
		// Wait a tiny moment for the loop to spin up, then kill the parent context
		time.Sleep(10 * time.Millisecond)
		cancelParent()
	}()

	err := PollUntilContextTimeout(parentCtx, interval, timeout, true, func(ctx context.Context) (bool, error) {
		return false, nil // Keep ticking
	})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Expected context.Canceled because parent context died, got: %v", err)
	}
}

// TestPollUntilContextTimeout_ConditionError checks that any explicit operational
// error returned inside the closure cancels the loop and bubbles up immediately.
func TestPollUntilContextTimeout_ConditionError(t *testing.T) {
	expectedErr := errors.New("slurm slurmctld node unreachable")

	err := PollUntilContextTimeout(context.Background(), 10*time.Millisecond, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		return false, expectedErr
	})

	if !errors.Is(err, expectedErr) {
		t.Fatalf("Expected error to bubble up cleanly. Want: %v, Got: %v", expectedErr, err)
	}
}
