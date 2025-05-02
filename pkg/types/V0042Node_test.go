// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042Node_GetKey(t *testing.T) {
	type fields struct {
		V0042Node api.V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0042Node: api.V0042Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0042Node: api.V0042Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Node{
				V0042Node: tt.fields.V0042Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Node_GetType(t *testing.T) {
	type fields struct {
		V0042Node api.V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0042Node: api.V0042Node{},
			},
			want: ObjectTypeV0042Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Node{
				V0042Node: tt.fields.V0042Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0042Node api.V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0042Node: api.V0042Node{},
			},
			want: &V0042Node{},
		},
		{
			name: "id",
			fields: fields{
				V0042Node: api.V0042Node{Name: ptr.To("node-0")},
			},
			want: &V0042Node{api.V0042Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Node{
				V0042Node: tt.fields.V0042Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0042Node api.V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0042Node
	}{
		{
			name: "empty",
			fields: fields{
				V0042Node: api.V0042Node{},
			},
			want: &V0042Node{},
		},
		{
			name: "id",
			fields: fields{
				V0042Node: api.V0042Node{Name: ptr.To("node-0")},
			},
			want: &V0042Node{api.V0042Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Node{
				V0042Node: tt.fields.V0042Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0042Node api.V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0042NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0042Node: api.V0042Node{},
			},
			want: set.New[api.V0042NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0042Node: api.V0042Node{
					State: ptr.To([]api.V0042NodeState{api.V0042NodeStateIDLE}),
				},
			},
			want: set.New(api.V0042NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0042Node: api.V0042Node{
					State: ptr.To([]api.V0042NodeState{api.V0042NodeStateIDLE, api.V0042NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0042NodeStateIDLE, api.V0042NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Node{
				V0042Node: tt.fields.V0042Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0042Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0042Node{},
			},
			want: ObjectTypeV0042Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0042Node{
					{V0042Node: api.V0042Node{Name: ptr.To("node-0")}},
					{V0042Node: api.V0042Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0042Node{api.V0042Node{Name: ptr.To("node-0")}},
				&V0042Node{api.V0042Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0042Node
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
				Items: []V0042Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0042Node{},
			},
			args: args{
				object: &V0042Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042Node{
					{V0042Node: api.V0042Node{Name: ptr.To("node-0")}},
					{V0042Node: api.V0042Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0042Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0042NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0042Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042Node{},
			},
			want: &V0042NodeList{
				Items: []V0042Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042Node{
					{V0042Node: api.V0042Node{Name: ptr.To("node-0")}},
					{V0042Node: api.V0042Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0042NodeList{
				Items: []V0042Node{
					{api.V0042Node{Name: ptr.To("node-0")}},
					{api.V0042Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
