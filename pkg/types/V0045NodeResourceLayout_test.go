// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045NodeResource_GetKey(t *testing.T) {
	type fields struct {
		V0045NodeResourceLayout api.V0045NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045NodeResourceLayout: []api.V0045NodeResourceLayout{},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayout{
				V0045NodeResourceLayoutList: tt.fields.V0045NodeResourceLayout,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResource_GetType(t *testing.T) {
	type fields struct {
		V0045NodeResourceLayout api.V0045NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045NodeResourceLayout: []api.V0045NodeResourceLayout{},
			},
			want: ObjectTypeV0045NodeResourceLayout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayout{
				V0045NodeResourceLayoutList: tt.fields.V0045NodeResourceLayout,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResource_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045NodeResourceLayout api.V0045NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045NodeResourceLayout: nil,
			},
			want: &V0045NodeResourceLayout{},
		},
		{
			name: "id",
			fields: fields{
				V0045NodeResourceLayout: []api.V0045NodeResourceLayout{{Node: "node1"}},
			},
			want: &V0045NodeResourceLayout{[]api.V0045NodeResourceLayout{{Node: "node1"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayout{
				V0045NodeResourceLayoutList: tt.fields.V0045NodeResourceLayout,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResource_DeepCopy(t *testing.T) {
	type fields struct {
		V0045NodeResourceLayout api.V0045NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045NodeResourceLayout
	}{
		{
			name: "empty",
			fields: fields{
				V0045NodeResourceLayout: nil,
			},
			want: &V0045NodeResourceLayout{},
		},
		{
			name: "id",
			fields: fields{
				V0045NodeResourceLayout: []api.V0045NodeResourceLayout{{Node: "node1"}},
			},
			want: &V0045NodeResourceLayout{[]api.V0045NodeResourceLayout{{Node: "node1"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayout{
				V0045NodeResourceLayoutList: tt.fields.V0045NodeResourceLayout,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResourceList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045NodeResourceLayout{},
			},
			want: ObjectTypeV0045NodeResourceLayout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResourceList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045NodeResourceLayout{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045NodeResourceLayout{
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node1"}}},
				},
			},
			want: []object.Object{
				&V0045NodeResourceLayout{
					[]api.V0045NodeResourceLayout{
						{Node: "node1"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045NodeResourceList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045NodeResourceLayout
	}
	type args struct {
		object object.Object
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantAppend bool
	}{
		{
			name: "nil",
			fields: fields{
				Items: []V0045NodeResourceLayout{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045NodeResourceLayout{},
			},
			args: args{
				object: &V0045NodeResourceLayout{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045NodeResourceLayout{
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node1"}}},
				},
			},
			args: args{
				object: &V0045NodeResourceLayout{
					V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node2"}},
				},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			require.Equal(t, want, got)
		})
	}
}

func TestV0045NodeResourceList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045NodeResourceLayout{},
			},
			want: &V0045NodeResourceLayoutList{
				Items: []V0045NodeResourceLayout{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045NodeResourceLayout{
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node1"}}},
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node2"}}},
				},
			},
			want: &V0045NodeResourceLayoutList{
				Items: []V0045NodeResourceLayout{
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node1"}}},
					{V0045NodeResourceLayoutList: api.V0045NodeResourceLayoutList{{Node: "node2"}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
