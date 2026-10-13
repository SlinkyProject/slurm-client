// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	apierrors "github.com/SlinkyProject/slurm-client/pkg/errors"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

// A nil skips slice exercises the default; multiple values exercise option precedence.
func writeJob(ctx context.Context, c Client, obj object.Object, operation string, skips ...bool) error {
	if operation == "update" {
		opts := make([]UpdateOption, 0, len(skips))
		for _, skip := range skips {
			opts = append(opts, &UpdateOptions{SkipReadAfterUpdate: skip})
		}
		return c.Update(ctx, obj, api.V0044JobDescMsg{Name: ptr.To("updated")}, opts...)
	}
	opts := make([]DeleteOption, 0, len(skips))
	for _, skip := range skips {
		opts = append(opts, &DeleteOptions{SkipReadAfterDelete: skip})
	}
	return c.Delete(ctx, obj, opts...)
}

func TestWriteReadback(t *testing.T) {
	writeErr := errors.New("write failed")
	readErr := errors.New("read failed")
	notFoundErr := fmt.Errorf("job gone: %w", apierrors.ErrNotFound)
	for _, operation := range []string{"update", "delete"} {
		for _, tc := range []struct {
			name     string
			skips    []bool
			writeErr error
			readErr  error
			wantErr  error
			wantRead bool
		}{
			{name: "default", wantRead: true},
			{name: "explicit false", skips: []bool{false}, wantRead: true},
			{name: "skip", skips: []bool{true}},
			{name: "last option wins", skips: []bool{true, false}, wantRead: true},
			{name: "read failure", readErr: readErr, wantErr: readErr, wantRead: true},
			{name: "skip read failure", skips: []bool{true}, readErr: readErr},
			{name: "not found", readErr: notFoundErr, wantErr: notFoundErr, wantRead: true},
			{name: "skip not found", skips: []bool{true}, readErr: notFoundErr},
			{name: "write failure", writeErr: writeErr, wantErr: writeErr},
			{name: "skip write failure", skips: []bool{true}, writeErr: writeErr, wantErr: writeErr},
		} {
			t.Run(operation+"/"+tc.name, func(t *testing.T) {
				obj := &api.V0044JobInfo{JobId: ptr.To(int32(42)), Name: ptr.To("original")}
				writes, reads := 0, 0
				write := func(key string) error {
					writes++
					require.Equal(t, "42", key)
					return tc.writeErr
				}
				c := &client{resources: map[object.ObjectType]resource{
					obj.GetType(): {
						update: func(_ context.Context, key string, req any) error {
							require.Equal(t, "update", operation)
							require.Equal(t, api.V0044JobDescMsg{Name: ptr.To("updated")}, req)
							return write(key)
						},
						delete: func(_ context.Context, key string) error {
							require.Equal(t, "delete", operation)
							return write(key)
						},
						get: func(_ context.Context, key object.ObjectKey, dst object.Object) error {
							reads++
							require.Equal(t, 1, writes, "readback must follow the write")
							require.Equal(t, obj.GetKey(), key)
							if tc.readErr == nil {
								dst.(*api.V0044JobInfo).Name = ptr.To("observed")
							}
							return tc.readErr
						},
					},
				}}
				err := writeJob(t.Context(), c, obj, operation, tc.skips...)
				wantErr := tc.wantErr
				if operation == "delete" && errors.Is(wantErr, apierrors.ErrNotFound) {
					wantErr = nil
				}
				require.ErrorIs(t, err, wantErr)
				require.Equal(t, 1, writes)
				wantReads := 0
				if tc.wantRead {
					wantReads = 1
				}
				require.Equal(t, wantReads, reads)
				wantName := "original"
				if tc.wantRead && tc.readErr == nil {
					wantName = "observed"
				}
				require.Equal(t, &api.V0044JobInfo{JobId: ptr.To(int32(42)), Name: ptr.To(wantName)}, obj)
			})
		}
	}
}

func TestSkipWriteReadbackPreservesCache(t *testing.T) {
	for _, operation := range []string{"update", "delete"} {
		t.Run(operation, func(t *testing.T) {
			obj := &api.V0044JobInfo{JobId: ptr.To(int32(42)), Name: ptr.To("original")}
			// Model a started informer with a populated cache, without running its
			// workers, so any refresh request remains observable in syncObjCh.
			i := newInformer(obj.GetType(), &emptyClient{}, defaultSyncPeriod).(*informerCache)
			i.started, i.dirty = true, false
			i.processObject(obj)
			writes := 0
			c := &client{
				informers: map[object.ObjectType]InformerCache{obj.GetType(): i},
				resources: map[object.ObjectType]resource{
					obj.GetType(): {
						cacheable: true,
						update: func(context.Context, string, any) error {
							writes++
							return nil
						},
						delete: func(context.Context, string) error {
							writes++
							return nil
						},
						get: func(context.Context, object.ObjectKey, object.Object) error {
							t.Fatal("unexpected direct readback")
							return nil
						},
					},
				},
			}
			ctx, cancel := context.WithTimeout(t.Context(), time.Second)
			defer cancel()
			require.NoError(t, writeJob(ctx, c, obj, operation, true))
			require.Equal(t, 1, writes)
			require.Empty(t, i.syncObjCh, "write must not enqueue a cache refresh")
			cached := &api.V0044JobInfo{}
			require.NoError(t, c.Get(ctx, obj.GetKey(), cached))
			require.Equal(t, &api.V0044JobInfo{JobId: ptr.To(int32(42)), Name: ptr.To("original")}, cached)
			require.Equal(t, cached, obj)
		})
	}
}
