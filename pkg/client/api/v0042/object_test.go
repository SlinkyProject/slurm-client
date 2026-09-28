// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestGeneratedObjects(t *testing.T) {
	node := &api.V0042Node{Name: ptr.To("node-0"), State: &[]api.V0042NodeState{api.V0042NodeStateIDLE}}
	copied := node.DeepCopy()
	*copied.Name = "other"
	require.Equal(t, "node-0", *node.Name)

	list := &api.V0042NodeObjectList{}
	list.AppendItem(node)
	list.AppendItem(&api.V0042JobInfo{})
	require.Len(t, list.Items, 1)
	require.Equal(t, api.ObjectTypeV0042Node, list.GetType())
	require.Equal(t, object.ObjectKey("node-0"), list.GetItems()[0].GetKey())

	job := &api.V0042JobInfo{JobId: ptr.To[int32](7)}
	require.Equal(t, object.ObjectKey("7"), job.GetKey())

	require.Equal(t, object.ObjectKey(""), (&api.V0042StatsMsg{}).GetKey())
	require.Len(t, (&api.V0042StatsMsgObjectList{Items: []api.V0042StatsMsg{{}}}).GetItems(), 1)

	for _, v := range []any{&api.V0042OpenapiNodesResp{}, ptr.To(api.V0042AccountFlags(""))} {
		_, ok := v.(object.RuntimeObject)
		require.False(t, ok, "%T", v)
	}

	var (
		_ object.Object     = node
		_ object.ObjectList = list
	)
}
