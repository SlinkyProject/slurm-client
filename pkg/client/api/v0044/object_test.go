// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestGeneratedObjects(t *testing.T) {
	node := &api.V0044Node{Name: ptr.To("node-0"), State: &[]api.V0044NodeState{api.V0044NodeStateIDLE}}

	copied := node.DeepCopy()
	*copied.Name = "other"
	require.Equal(t, "node-0", *node.Name)

	list := &api.V0044NodeObjectList{}
	list.AppendItem(node)
	list.AppendItem(&api.V0044JobInfo{})
	require.Len(t, list.Items, 1)
	require.Equal(t, api.ObjectTypeV0044Node, list.GetType())
	require.Equal(t, object.ObjectKey("node-0"), list.GetItems()[0].GetKey())

	job := &api.V0044JobInfo{JobId: ptr.To[int32](7)}
	require.Equal(t, object.ObjectKey("7"), job.GetKey())

	require.Equal(t, object.ObjectKey(""), (&api.V0044StatsMsg{}).GetKey())
	require.Len(t, (&api.V0044StatsMsgObjectList{Items: []api.V0044StatsMsg{{}}}).GetItems(), 1)

	layout := &api.V0044NodeResourceLayoutListObject{{Node: "node1"}}
	require.Equal(t, object.ObjectKey("0"), layout.GetKey())
	require.Equal(t, api.ObjectTypeV0044NodeResourceLayoutList, layout.GetType())
	layouts := &api.V0044NodeResourceLayoutListObjectList{}
	layouts.AppendItem(layout)
	require.Equal(t, api.ObjectTypeV0044NodeResourceLayoutList, layouts.GetType())
	require.Len(t, layouts.Items, 1)

	for _, v := range []any{&api.V0044OpenapiNodesResp{}, ptr.To(api.V0044AccountFlags(""))} {
		_, ok := v.(object.RuntimeObject)
		require.False(t, ok, "%T", v)
	}

	var (
		_ object.Object     = node
		_ object.ObjectList = list
	)
}
