// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045Stats_GetKey(t *testing.T) {
	type fields struct {
		V0045Stats api.V0045StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "key",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Stats{
				V0045StatsMsg: tt.fields.V0045Stats,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Stats.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Stats_GetType(t *testing.T) {
	type fields struct {
		V0045Stats api.V0045StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: ObjectTypeV0045Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Stats{
				V0045StatsMsg: tt.fields.V0045Stats,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Stats.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Stats_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045Stats api.V0045StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: &V0045Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: &V0045Stats{api.V0045StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Stats{
				V0045StatsMsg: tt.fields.V0045Stats,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Stats.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Stats_DeepCopy(t *testing.T) {
	type fields struct {
		V0045Stats api.V0045StatsMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045Stats
	}{
		{
			name: "empty",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: &V0045Stats{},
		},
		{
			name: "id",
			fields: fields{
				V0045Stats: api.V0045StatsMsg{},
			},
			want: &V0045Stats{api.V0045StatsMsg{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Stats{
				V0045StatsMsg: tt.fields.V0045Stats,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Stats.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045StatsList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045Stats{},
			},
			want: ObjectTypeV0045Stats,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045StatsList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045StatsList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045StatsList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Stats{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045Stats{
					{V0045StatsMsg: api.V0045StatsMsg{}},
				},
			},
			want: []object.Object{
				&V0045Stats{api.V0045StatsMsg{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045StatsList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045StatsList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045StatsList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045Stats
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
				Items: []V0045Stats{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045Stats{},
			},
			args: args{
				object: &V0045Stats{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Stats{
					{V0045StatsMsg: api.V0045StatsMsg{}},
				},
			},
			args: args{
				object: &V0045Stats{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045StatsList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045StatsList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045StatsList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045Stats
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Stats{},
			},
			want: &V0045StatsList{
				Items: []V0045Stats{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Stats{
					{V0045StatsMsg: api.V0045StatsMsg{}},
				},
			},
			want: &V0045StatsList{
				Items: []V0045Stats{
					{api.V0045StatsMsg{}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045StatsList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045StatsList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
