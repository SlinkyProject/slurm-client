// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045Node_GetKey(t *testing.T) {
	type fields struct {
		V0045Node api.V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045Node: api.V0045Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0045Node: api.V0045Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Node{
				V0045Node: tt.fields.V0045Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Node_GetType(t *testing.T) {
	type fields struct {
		V0045Node api.V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045Node: api.V0045Node{},
			},
			want: ObjectTypeV0045Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Node{
				V0045Node: tt.fields.V0045Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045Node api.V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045Node: api.V0045Node{},
			},
			want: &V0045Node{},
		},
		{
			name: "id",
			fields: fields{
				V0045Node: api.V0045Node{Name: ptr.To("node-0")},
			},
			want: &V0045Node{api.V0045Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Node{
				V0045Node: tt.fields.V0045Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0045Node api.V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045Node
	}{
		{
			name: "empty",
			fields: fields{
				V0045Node: api.V0045Node{},
			},
			want: &V0045Node{},
		},
		{
			name: "id",
			fields: fields{
				V0045Node: api.V0045Node{Name: ptr.To("node-0")},
			},
			want: &V0045Node{api.V0045Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Node{
				V0045Node: tt.fields.V0045Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0045Node api.V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0045NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0045Node: api.V0045Node{},
			},
			want: set.New[api.V0045NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0045Node: api.V0045Node{
					State: ptr.To([]api.V0045NodeState{api.V0045NodeStateIDLE}),
				},
			},
			want: set.New(api.V0045NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0045Node: api.V0045Node{
					State: ptr.To([]api.V0045NodeState{api.V0045NodeStateIDLE, api.V0045NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0045NodeStateIDLE, api.V0045NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Node{
				V0045Node: tt.fields.V0045Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0045Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045Node{},
			},
			want: ObjectTypeV0045Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045Node{
					{V0045Node: api.V0045Node{Name: ptr.To("node-0")}},
					{V0045Node: api.V0045Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0045Node{api.V0045Node{Name: ptr.To("node-0")}},
				&V0045Node{api.V0045Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045Node
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
				Items: []V0045Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045Node{},
			},
			args: args{
				object: &V0045Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Node{
					{V0045Node: api.V0045Node{Name: ptr.To("node-0")}},
					{V0045Node: api.V0045Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0045Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Node{},
			},
			want: &V0045NodeList{
				Items: []V0045Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Node{
					{V0045Node: api.V0045Node{Name: ptr.To("node-0")}},
					{V0045Node: api.V0045Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0045NodeList{
				Items: []V0045Node{
					{api.V0045Node{Name: ptr.To("node-0")}},
					{api.V0045Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
