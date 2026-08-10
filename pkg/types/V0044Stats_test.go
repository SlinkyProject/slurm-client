// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Stats_GetKey(t *testing.T) {
	type fields struct {
		V0044Stats api.V0044StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "key",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Stats{
				V0044StatsMsg: tt.fields.V0044Stats,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044Stats_GetType(t *testing.T) {
	type fields struct {
		V0044Stats api.V0044StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: ObjectTypeV0044Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Stats{
				V0044StatsMsg: tt.fields.V0044Stats,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044Stats_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044Stats api.V0044StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: &V0044Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: &V0044Stats{api.V0044StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Stats{
				V0044StatsMsg: tt.fields.V0044Stats,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044Stats_DeepCopy(t *testing.T) {
	type fields struct {
		V0044Stats api.V0044StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044Stats
	}{
		{
			name: "empty",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: &V0044Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0044Stats: api.V0044StatsMsg{},
			},
			want: &V0044Stats{api.V0044StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Stats{
				V0044StatsMsg: tt.fields.V0044Stats,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044StatsList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044Stats{},
			},
			want: ObjectTypeV0044Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044StatsList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Stats{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044Stats{
					{V0044StatsMsg: api.V0044StatsMsg{}},
				},
			},
			want: []object.Object{
				&V0044Stats{api.V0044StatsMsg{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044StatsList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0044StatsList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044Stats
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
				Items: []V0044Stats{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044Stats{},
			},
			args: args{
				object: &V0044Stats{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Stats{
					{V0044StatsMsg: api.V0044StatsMsg{}},
				},
			},
			args: args{
				object: &V0044Stats{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044StatsList{
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

func TestV0044StatsList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Stats{},
			},
			want: &V0044StatsList{
				Items: []V0044Stats{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Stats{
					{V0044StatsMsg: api.V0044StatsMsg{}},
				},
			},
			want: &V0044StatsList{
				Items: []V0044Stats{
					{api.V0044StatsMsg{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044StatsList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
