// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042Stats_GetKey(t *testing.T) {
	type fields struct {
		V0042Stats api.V0042StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "key",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Stats{
				V0042StatsMsg: tt.fields.V0042Stats,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042Stats_GetType(t *testing.T) {
	type fields struct {
		V0042Stats api.V0042StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: ObjectTypeV0042Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Stats{
				V0042StatsMsg: tt.fields.V0042Stats,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042Stats_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0042Stats api.V0042StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: &V0042Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: &V0042Stats{api.V0042StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Stats{
				V0042StatsMsg: tt.fields.V0042Stats,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042Stats_DeepCopy(t *testing.T) {
	type fields struct {
		V0042Stats api.V0042StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0042Stats
	}{
		{
			name: "empty",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: &V0042Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0042Stats: api.V0042StatsMsg{},
			},
			want: &V0042Stats{api.V0042StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Stats{
				V0042StatsMsg: tt.fields.V0042Stats,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042StatsList_GetType(t *testing.T) {
	type fields struct {
		Items []V0042Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0042Stats{},
			},
			want: ObjectTypeV0042Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042StatsList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0042Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042Stats{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0042Stats{
					{V0042StatsMsg: api.V0042StatsMsg{}},
				},
			},
			want: []object.Object{
				&V0042Stats{api.V0042StatsMsg{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0042StatsList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0042Stats
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
				Items: []V0042Stats{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0042Stats{},
			},
			args: args{
				object: &V0042Stats{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042Stats{
					{V0042StatsMsg: api.V0042StatsMsg{}},
				},
			},
			args: args{
				object: &V0042Stats{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042StatsList{
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

func TestV0042StatsList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0042Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042Stats{},
			},
			want: &V0042StatsList{
				Items: []V0042Stats{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042Stats{
					{V0042StatsMsg: api.V0042StatsMsg{}},
				},
			},
			want: &V0042StatsList{
				Items: []V0042Stats{
					{api.V0042StatsMsg{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042StatsList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
