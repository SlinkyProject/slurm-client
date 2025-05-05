// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0041Stats_GetKey(t *testing.T) {
	type fields struct {
		V0041Stats api.V0041StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "key",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Stats{
				V0041StatsMsg: tt.fields.V0041Stats,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Stats.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Stats_GetType(t *testing.T) {
	type fields struct {
		V0041Stats api.V0041StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: ObjectTypeV0041Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Stats{
				V0041StatsMsg: tt.fields.V0041Stats,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Stats.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Stats_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0041Stats api.V0041StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: &V0041Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: &V0041Stats{api.V0041StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Stats{
				V0041StatsMsg: tt.fields.V0041Stats,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Stats.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Stats_DeepCopy(t *testing.T) {
	type fields struct {
		V0041Stats api.V0041StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0041Stats
	}{
		{
			name: "empty",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: &V0041Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0041Stats: api.V0041StatsMsg{},
			},
			want: &V0041Stats{api.V0041StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Stats{
				V0041StatsMsg: tt.fields.V0041Stats,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Stats.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041StatsList_GetType(t *testing.T) {
	type fields struct {
		Items []V0041Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0041Stats{},
			},
			want: ObjectTypeV0041Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041StatsList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041StatsList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041StatsList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0041Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041Stats{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0041Stats{
					{V0041StatsMsg: api.V0041StatsMsg{}},
				},
			},
			want: []object.Object{
				&V0041Stats{api.V0041StatsMsg{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041StatsList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041StatsList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041StatsList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0041Stats
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
				Items: []V0041Stats{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0041Stats{},
			},
			args: args{
				object: &V0041Stats{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041Stats{
					{V0041StatsMsg: api.V0041StatsMsg{}},
				},
			},
			args: args{
				object: &V0041Stats{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041StatsList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0041StatsList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041StatsList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0041Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041Stats{},
			},
			want: &V0041StatsList{
				Items: []V0041Stats{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041Stats{
					{V0041StatsMsg: api.V0041StatsMsg{}},
				},
			},
			want: &V0041StatsList{
				Items: []V0041Stats{
					{api.V0041StatsMsg{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041StatsList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041StatsList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
