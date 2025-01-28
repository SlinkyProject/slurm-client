// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0041Node_GetKey(t *testing.T) {
	type fields struct {
		V0041Node api.V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0041Node: api.V0041Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0041Node: api.V0041Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Node{
				V0041Node: tt.fields.V0041Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Node_GetType(t *testing.T) {
	type fields struct {
		V0041Node api.V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0041Node: api.V0041Node{},
			},
			want: ObjectTypeV0041Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Node{
				V0041Node: tt.fields.V0041Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0041Node api.V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0041Node: api.V0041Node{},
			},
			want: &V0041Node{},
		},
		{
			name: "id",
			fields: fields{
				V0041Node: api.V0041Node{Name: ptr.To("node-0")},
			},
			want: &V0041Node{api.V0041Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Node{
				V0041Node: tt.fields.V0041Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0041Node api.V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0041Node
	}{
		{
			name: "empty",
			fields: fields{
				V0041Node: api.V0041Node{},
			},
			want: &V0041Node{},
		},
		{
			name: "id",
			fields: fields{
				V0041Node: api.V0041Node{Name: ptr.To("node-0")},
			},
			want: &V0041Node{api.V0041Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Node{
				V0041Node: tt.fields.V0041Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0041Node api.V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0041NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0041Node: api.V0041Node{},
			},
			want: set.New[api.V0041NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0041Node: api.V0041Node{
					State: ptr.To([]api.V0041NodeState{api.V0041NodeStateIDLE}),
				},
			},
			want: set.New(api.V0041NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0041Node: api.V0041Node{
					State: ptr.To([]api.V0041NodeState{api.V0041NodeStateIDLE, api.V0041NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0041NodeStateIDLE, api.V0041NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Node{
				V0041Node: tt.fields.V0041Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0041Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0041Node{},
			},
			want: ObjectTypeV0041Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0041Node{
					{V0041Node: api.V0041Node{Name: ptr.To("node-0")}},
					{V0041Node: api.V0041Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0041Node{api.V0041Node{Name: ptr.To("node-0")}},
				&V0041Node{api.V0041Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0041Node
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
				Items: []V0041Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0041Node{},
			},
			args: args{
				object: &V0041Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041Node{
					{V0041Node: api.V0041Node{Name: ptr.To("node-0")}},
					{V0041Node: api.V0041Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0041Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0041NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0041Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041Node{},
			},
			want: &V0041NodeList{
				Items: []V0041Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041Node{
					{V0041Node: api.V0041Node{Name: ptr.To("node-0")}},
					{V0041Node: api.V0041Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0041NodeList{
				Items: []V0041Node{
					{api.V0041Node{Name: ptr.To("node-0")}},
					{api.V0041Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
