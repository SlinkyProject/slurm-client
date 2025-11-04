// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044NodeResource_GetKey(t *testing.T) {
	type fields struct {
		V0044NodeResourceLayout api.V0044NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayout{
				V0044NodeResourceLayoutList: tt.fields.V0044NodeResourceLayout,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayout.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResource_GetType(t *testing.T) {
	type fields struct {
		V0044NodeResourceLayout api.V0044NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{},
			},
			want: ObjectTypeV0044NodeResourceLayout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayout{
				V0044NodeResourceLayoutList: tt.fields.V0044NodeResourceLayout,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayout.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResource_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044NodeResourceLayout api.V0044NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{},
			},
			want: &V0044NodeResourceLayout{},
		},
		{
			name: "id",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{{Node: "node1"}},
			},
			want: &V0044NodeResourceLayout{[]api.V0044NodeResourceLayout{{Node: "node1"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayout{
				V0044NodeResourceLayoutList: tt.fields.V0044NodeResourceLayout,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayout.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResource_DeepCopy(t *testing.T) {
	type fields struct {
		V0044NodeResourceLayout api.V0044NodeResourceLayoutList
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044NodeResourceLayout
	}{
		{
			name: "empty",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{},
			},
			want: &V0044NodeResourceLayout{},
		},
		{
			name: "id",
			fields: fields{
				V0044NodeResourceLayout: []api.V0044NodeResourceLayout{{Node: "node1"}},
			},
			want: &V0044NodeResourceLayout{[]api.V0044NodeResourceLayout{{Node: "node1"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayout{
				V0044NodeResourceLayoutList: tt.fields.V0044NodeResourceLayout,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayout.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResourceList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044NodeResourceLayout{},
			},
			want: ObjectTypeV0044NodeResourceLayout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayoutList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResourceList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044NodeResourceLayout{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044NodeResourceLayout{
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node1"}}},
				},
			},
			want: []object.Object{
				&V0044NodeResourceLayout{
					[]api.V0044NodeResourceLayout{
						{Node: "node1"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayoutList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeResourceList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044NodeResourceLayout
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
				Items: []V0044NodeResourceLayout{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044NodeResourceLayout{},
			},
			args: args{
				object: &V0044NodeResourceLayout{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044NodeResourceLayout{
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node1"}}},
				},
			},
			args: args{
				object: &V0044NodeResourceLayout{
					V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node2"}},
				},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044NodeResourceLayoutList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044NodeResourceList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044NodeResourceLayout
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044NodeResourceLayout{},
			},
			want: &V0044NodeResourceLayoutList{
				Items: []V0044NodeResourceLayout{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044NodeResourceLayout{
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node1"}}},
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node2"}}},
				},
			},
			want: &V0044NodeResourceLayoutList{
				Items: []V0044NodeResourceLayout{
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node1"}}},
					{V0044NodeResourceLayoutList: api.V0044NodeResourceLayoutList{{Node: "node2"}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeResourceLayoutList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeResourceLayoutList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
