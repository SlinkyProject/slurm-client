// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	apiv0042 "github.com/SlinkyProject/slurm-client/api/v0042"
	apiv0043 "github.com/SlinkyProject/slurm-client/api/v0043"
	apiv0044 "github.com/SlinkyProject/slurm-client/api/v0044"
	apiv0045 "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestSetKey(t *testing.T) {
	cases := []struct {
		objects []object.Object
		keys    []object.ObjectKey
		invalid []object.ObjectKey
	}{
		{
			objects: []object.Object{
				&apiv0042.V0042JobInfo{}, &apiv0043.V0043JobInfo{},
				&apiv0044.V0044JobInfo{}, &apiv0045.V0045JobInfo{},
			},
			keys:    []object.ObjectKey{"0", "-2147483648", "2147483647", "42"},
			invalid: []object.ObjectKey{"", "not-a-job-id", "2147483648", "-2147483649"},
		},
		{
			objects: []object.Object{
				&apiv0042.V0042ControllerPing{}, &apiv0043.V0043ControllerPing{},
				&apiv0044.V0044ControllerPing{}, &apiv0045.V0045ControllerPing{},
				&apiv0042.V0042Node{}, &apiv0043.V0043Node{},
				&apiv0044.V0044Node{}, &apiv0045.V0045Node{},
				&apiv0042.V0042PartitionInfo{}, &apiv0043.V0043PartitionInfo{},
				&apiv0044.V0044PartitionInfo{}, &apiv0045.V0045PartitionInfo{},
				&apiv0042.V0042ReservationInfo{}, &apiv0043.V0043ReservationInfo{},
				&apiv0044.V0044ReservationInfo{}, &apiv0045.V0045ReservationInfo{},
			},
			keys: []object.ObjectKey{"", "42", "name-1"},
		},
		{
			objects: []object.Object{
				&apiv0042.V0042StatsMsg{}, &apiv0043.V0043StatsMsg{},
				&apiv0044.V0044StatsMsg{}, &apiv0045.V0045StatsMsg{},
			},
			keys:    []object.ObjectKey{""},
			invalid: []object.ObjectKey{"name-1"},
		},
		{
			objects: []object.Object{
				&apiv0044.V0044NodeResourceLayoutListObject{},
				&apiv0045.V0045NodeResourceLayoutListObject{},
			},
			keys:    []object.ObjectKey{"0"},
			invalid: []object.ObjectKey{"", "name-1"},
		},
	}
	for _, tc := range cases {
		for _, obj := range tc.objects {
			t.Run(string(obj.GetType()), func(t *testing.T) {
				for _, key := range tc.keys {
					require.NoError(t, obj.SetKey(key))
					require.Equal(t, key, obj.GetKey())
				}
				before := obj.DeepCopyObject()
				for _, key := range tc.invalid {
					require.Error(t, obj.SetKey(key), "key %q", key)
					require.Equal(t, before, obj, "invalid key changed the object")
				}
			})
		}
	}
}
