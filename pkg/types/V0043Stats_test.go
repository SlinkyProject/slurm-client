// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043Stats_GetKey(t *testing.T) {
	type fields struct {
		V0043Stats api.V0043StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "key",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Stats{
				V0043StatsMsg: tt.fields.V0043Stats,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043Stats_GetType(t *testing.T) {
	type fields struct {
		V0043Stats api.V0043StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: ObjectTypeV0043Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Stats{
				V0043StatsMsg: tt.fields.V0043Stats,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043Stats_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0043Stats api.V0043StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: &V0043Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: &V0043Stats{api.V0043StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Stats{
				V0043StatsMsg: tt.fields.V0043Stats,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043Stats_DeepCopy(t *testing.T) {
	type fields struct {
		V0043Stats api.V0043StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0043Stats
	}{
		{
			name: "empty",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: &V0043Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0043Stats: api.V0043StatsMsg{},
			},
			want: &V0043Stats{api.V0043StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Stats{
				V0043StatsMsg: tt.fields.V0043Stats,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043StatsList_GetType(t *testing.T) {
	type fields struct {
		Items []V0043Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0043Stats{},
			},
			want: ObjectTypeV0043Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043StatsList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0043Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043Stats{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0043Stats{
					{V0043StatsMsg: api.V0043StatsMsg{}},
				},
			},
			want: []object.Object{
				&V0043Stats{api.V0043StatsMsg{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043StatsList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0043Stats
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
				Items: []V0043Stats{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0043Stats{},
			},
			args: args{
				object: &V0043Stats{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043Stats{
					{V0043StatsMsg: api.V0043StatsMsg{}},
				},
			},
			args: args{
				object: &V0043Stats{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043StatsList{
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

func TestV0043StatsList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0043Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043Stats{},
			},
			want: &V0043StatsList{
				Items: []V0043Stats{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043Stats{
					{V0043StatsMsg: api.V0043StatsMsg{}},
				},
			},
			want: &V0043StatsList{
				Items: []V0043Stats{
					{api.V0043StatsMsg{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043StatsList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
