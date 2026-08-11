// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-FileCopyrightText: Copyright 2023 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package wait

import (
	"context"
	"time"
)

type ConditionWithContextFunc func(ctx context.Context) (done bool, err error)

// PollUntilContextTimeout will terminate polling after timeout duration by setting a context timeout.
func PollUntilContextTimeout(ctx context.Context, interval, timeout time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	deadlineCtx, deadlineCancel := context.WithTimeout(ctx, timeout)
	defer deadlineCancel()
	return loopConditionUntilContext(deadlineCtx, interval, immediate, false, condition)
}

// loopConditionUntilContext executes the provided condition at intervals defined by
// the provided timer until the provided context is canceled, the condition returns
// true, or the condition returns an error. If sliding is true, the period is computed
// after condition runs. If it is false then period includes the runtime for condition.
// If immediate is false the first delay happens before any call to condition, if
// immediate is true the condition will be invoked before waiting and guarantees that
// the condition is invoked at least once, regardless of whether the context has been
// canceled. The returned error is the error returned by the last condition or the
// context error if the context was terminated.
//
// This is the common loop construct for all polling in the wait package.
func loopConditionUntilContext(ctx context.Context, interval time.Duration, immediate, sliding bool, condition ConditionWithContextFunc) error {
	var ticker *time.Ticker
	defer func() {
		if ticker != nil {
			ticker.Stop()
		}
	}()
	c := func() <-chan time.Time {
		if ticker == nil {
			ticker = time.NewTicker(interval)
		}
		return ticker.C
	}

	var timeCh <-chan time.Time
	doneCh := ctx.Done()

	if !sliding {
		timeCh = c()
	}

	// if immediate is true the condition is guaranteed to be executed at least once,
	// if we haven't requested immediate execution, delay once.
	if immediate {
		if ok, err := condition(ctx); err != nil || ok {
			return err
		}
	}

	if sliding {
		timeCh = c()
	}

	for {
		// Wait for either the context to be canceled or the next invocation be called.
		select {
		case <-doneCh:
			return ctx.Err()
		case <-timeCh:
		}

		// IMPORTANT: Because there is no channel priority selection in golang
		// it is possible for very short timers to "win" the race in the previous select
		// repeatedly even when the context has been canceled. We therefore must
		// explicitly check for context cancellation on every loop and exit if true to
		// guarantee that we don't invoke condition more than once after context has
		// been canceled.
		if err := ctx.Err(); err != nil {
			return err
		}

		if ok, err := condition(ctx); err != nil || ok {
			return err
		}
	}
}
