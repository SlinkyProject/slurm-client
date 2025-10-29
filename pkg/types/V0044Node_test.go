// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Node_GetKey(t *testing.T) {
	type fields struct {
		V0044Node api.V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044Node: api.V0044Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0044Node: api.V0044Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Node{
				V0044Node: tt.fields.V0044Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Node_GetType(t *testing.T) {
	type fields struct {
		V0044Node api.V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044Node: api.V0044Node{},
			},
			want: ObjectTypeV0044Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Node{
				V0044Node: tt.fields.V0044Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044Node api.V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044Node: api.V0044Node{},
			},
			want: &V0044Node{},
		},
		{
			name: "id",
			fields: fields{
				V0044Node: api.V0044Node{Name: ptr.To("node-0")},
			},
			want: &V0044Node{api.V0044Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Node{
				V0044Node: tt.fields.V0044Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0044Node api.V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044Node
	}{
		{
			name: "empty",
			fields: fields{
				V0044Node: api.V0044Node{},
			},
			want: &V0044Node{},
		},
		{
			name: "id",
			fields: fields{
				V0044Node: api.V0044Node{Name: ptr.To("node-0")},
			},
			want: &V0044Node{api.V0044Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Node{
				V0044Node: tt.fields.V0044Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0044Node api.V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0044NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0044Node: api.V0044Node{},
			},
			want: set.New[api.V0044NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0044Node: api.V0044Node{
					State: ptr.To([]api.V0044NodeState{api.V0044NodeStateIDLE}),
				},
			},
			want: set.New(api.V0044NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0044Node: api.V0044Node{
					State: ptr.To([]api.V0044NodeState{api.V0044NodeStateIDLE, api.V0044NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0044NodeStateIDLE, api.V0044NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Node{
				V0044Node: tt.fields.V0044Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0044Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044Node{},
			},
			want: ObjectTypeV0044Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044Node{
					{V0044Node: api.V0044Node{Name: ptr.To("node-0")}},
					{V0044Node: api.V0044Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0044Node{api.V0044Node{Name: ptr.To("node-0")}},
				&V0044Node{api.V0044Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044Node
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
				Items: []V0044Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044Node{},
			},
			args: args{
				object: &V0044Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Node{
					{V0044Node: api.V0044Node{Name: ptr.To("node-0")}},
					{V0044Node: api.V0044Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0044Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Node{},
			},
			want: &V0044NodeList{
				Items: []V0044Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Node{
					{V0044Node: api.V0044Node{Name: ptr.To("node-0")}},
					{V0044Node: api.V0044Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0044NodeList{
				Items: []V0044Node{
					{api.V0044Node{Name: ptr.To("node-0")}},
					{api.V0044Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
