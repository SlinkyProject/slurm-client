// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0045

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestGeneratedObjects(t *testing.T) {
	node := &api.V0045Node{Name: ptr.To("node-0"), State: &[]api.V0045NodeState{api.V0045NodeStateIDLE}}

	copied := node.DeepCopy()
	*copied.Name = "other"
	require.Equal(t, "node-0", *node.Name)

	list := &api.V0045NodeObjectList{}
	list.AppendItem(node)
	list.AppendItem(&api.V0045JobInfo{})
	require.Len(t, list.Items, 1)
	require.Equal(t, api.ObjectTypeV0045Node, list.GetType())
	require.Equal(t, object.ObjectKey("node-0"), list.GetItems()[0].GetKey())

	job := &api.V0045JobInfo{JobId: ptr.To[int32](7)}
	require.Equal(t, object.ObjectKey("7"), job.GetKey())

	require.Equal(t, object.ObjectKey(""), (&api.V0045StatsMsg{}).GetKey())
	require.Len(t, (&api.V0045StatsMsgObjectList{Items: []api.V0045StatsMsg{{}}}).GetItems(), 1)

	layout := &api.V0045NodeResourceLayoutListObject{{Node: "node1"}}
	require.Equal(t, object.ObjectKey("0"), layout.GetKey())
	require.Equal(t, api.ObjectTypeV0045NodeResourceLayoutList, layout.GetType())
	layouts := &api.V0045NodeResourceLayoutListObjectList{}
	layouts.AppendItem(layout)
	require.Equal(t, api.ObjectTypeV0045NodeResourceLayoutList, layouts.GetType())
	require.Len(t, layouts.Items, 1)

	for _, v := range []any{&api.V0045OpenapiNodesResp{}, ptr.To(api.V0045AccountFlags("")), ptr.To(api.V0045PartitionInfoFlags(""))} {
		_, ok := v.(object.RuntimeObject)
		require.False(t, ok, "%T", v)
	}

	var (
		_ object.Object     = node
		_ object.ObjectList = list
	)
}
