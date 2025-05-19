// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043Node_GetKey(t *testing.T) {
	type fields struct {
		V0043Node api.V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0043Node: api.V0043Node{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0043Node: api.V0043Node{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Node{
				V0043Node: tt.fields.V0043Node,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Node.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Node_GetType(t *testing.T) {
	type fields struct {
		V0043Node api.V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0043Node: api.V0043Node{},
			},
			want: ObjectTypeV0043Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Node{
				V0043Node: tt.fields.V0043Node,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Node.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Node_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0043Node api.V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0043Node: api.V0043Node{},
			},
			want: &V0043Node{},
		},
		{
			name: "id",
			fields: fields{
				V0043Node: api.V0043Node{Name: ptr.To("node-0")},
			},
			want: &V0043Node{api.V0043Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Node{
				V0043Node: tt.fields.V0043Node,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Node.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Node_DeepCopy(t *testing.T) {
	type fields struct {
		V0043Node api.V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0043Node
	}{
		{
			name: "empty",
			fields: fields{
				V0043Node: api.V0043Node{},
			},
			want: &V0043Node{},
		},
		{
			name: "id",
			fields: fields{
				V0043Node: api.V0043Node{Name: ptr.To("node-0")},
			},
			want: &V0043Node{api.V0043Node{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Node{
				V0043Node: tt.fields.V0043Node,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Node.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Node_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0043Node api.V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0043NodeState]
	}{
		{
			name: "empty",
			fields: fields{
				V0043Node: api.V0043Node{},
			},
			want: set.New[api.V0043NodeState](),
		},
		{
			name: "single",
			fields: fields{
				V0043Node: api.V0043Node{
					State: ptr.To([]api.V0043NodeState{api.V0043NodeStateIDLE}),
				},
			},
			want: set.New(api.V0043NodeStateIDLE),
		},
		{
			name: "multiple",
			fields: fields{
				V0043Node: api.V0043Node{
					State: ptr.To([]api.V0043NodeState{api.V0043NodeStateIDLE, api.V0043NodeStateDRAIN}),
				},
			},
			want: set.New(api.V0043NodeStateIDLE, api.V0043NodeStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Node{
				V0043Node: tt.fields.V0043Node,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0043Node.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043NodeList_GetType(t *testing.T) {
	type fields struct {
		Items []V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0043Node{},
			},
			want: ObjectTypeV0043Node,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043NodeList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043NodeList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043Node{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0043Node{
					{V0043Node: api.V0043Node{Name: ptr.To("node-0")}},
					{V0043Node: api.V0043Node{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0043Node{api.V0043Node{Name: ptr.To("node-0")}},
				&V0043Node{api.V0043Node{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043NodeList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043NodeList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043NodeList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0043Node
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
				Items: []V0043Node{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0043Node{},
			},
			args: args{
				object: &V0043Node{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043Node{
					{V0043Node: api.V0043Node{Name: ptr.To("node-0")}},
					{V0043Node: api.V0043Node{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0043Node{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043NodeList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0043NodeList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0043NodeList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0043Node
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043Node{},
			},
			want: &V0043NodeList{
				Items: []V0043Node{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043Node{
					{V0043Node: api.V0043Node{Name: ptr.To("node-0")}},
					{V0043Node: api.V0043Node{Name: ptr.To("node-1")}},
				},
			},
			want: &V0043NodeList{
				Items: []V0043Node{
					{api.V0043Node{Name: ptr.To("node-0")}},
					{api.V0043Node{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043NodeList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043NodeList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
