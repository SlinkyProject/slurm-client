// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0040Node_GetKey(t *testing.T) {
	type fields struct {
		V0040Node api.V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0040Node: api.V0040Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0040Node: api.V0040Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040Node{
				V0040Node: tt.fields.V0040Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040Node_GetType(t *testing.T) {
	type fields struct {
		V0040Node api.V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0040Node: api.V0040Node{},
			},
			want: ObjectTypeV0040Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040Node{
				V0040Node: tt.fields.V0040Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0040Node api.V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0040Node: api.V0040Node{},
			},
			want: &V0040Node{},
		},
		{
			name: "id",
			fields: fields{
				V0040Node: api.V0040Node{Name: ptr.To("node-0")},
			},
			want: &V0040Node{api.V0040Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040Node{
				V0040Node: tt.fields.V0040Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0040Node api.V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0040Node
	}{
		{
			name: "empty",
			fields: fields{
				V0040Node: api.V0040Node{},
			},
			want: &V0040Node{},
		},
		{
			name: "id",
			fields: fields{
				V0040Node: api.V0040Node{Name: ptr.To("node-0")},
			},
			want: &V0040Node{api.V0040Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040Node{
				V0040Node: tt.fields.V0040Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0040Node api.V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0040NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0040Node: api.V0040Node{},
			},
			want: set.New[api.V0040NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0040Node: api.V0040Node{
					State: ptr.To([]api.V0040NodeState{api.V0040NodeStateIDLE}),
				},
			},
			want: set.New(api.V0040NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0040Node: api.V0040Node{
					State: ptr.To([]api.V0040NodeState{api.V0040NodeStateIDLE, api.V0040NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0040NodeStateIDLE, api.V0040NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040Node{
				V0040Node: tt.fields.V0040Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0040Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0040Node{},
			},
			want: ObjectTypeV0040Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0040Node{
					{V0040Node: api.V0040Node{Name: ptr.To("node-0")}},
					{V0040Node: api.V0040Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0040Node{api.V0040Node{Name: ptr.To("node-0")}},
				&V0040Node{api.V0040Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0040Node
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
				Items: []V0040Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0040Node{},
			},
			args: args{
				object: &V0040Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040Node{
					{V0040Node: api.V0040Node{Name: ptr.To("node-0")}},
					{V0040Node: api.V0040Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0040Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0040NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0040NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0040Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040Node{},
			},
			want: &V0040NodeList{
				Items: []V0040Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040Node{
					{V0040Node: api.V0040Node{Name: ptr.To("node-0")}},
					{V0040Node: api.V0040Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0040NodeList{
				Items: []V0040Node{
					{api.V0040Node{Name: ptr.To("node-0")}},
					{api.V0040Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
