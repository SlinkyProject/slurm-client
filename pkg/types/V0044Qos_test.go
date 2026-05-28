// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Qos_GetKey(t *testing.T) {
	type fields struct {
		V0044Qos api.V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044Qos: api.V0044Qos{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0044Qos: api.V0044Qos{Name: ptr.To("test_0")},
			},
			want: "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Qos{
				V0044Qos: tt.fields.V0044Qos,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Qos.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Qos_GetType(t *testing.T) {
	type fields struct {
		V0044Qos api.V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044Qos: api.V0044Qos{},
			},
			want: ObjectTypeV0044Qos,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Qos{
				V0044Qos: tt.fields.V0044Qos,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Qos.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Qos_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044Qos api.V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044Qos: api.V0044Qos{},
			},
			want: &V0044Qos{},
		},
		{
			name: "id",
			fields: fields{
				V0044Qos: api.V0044Qos{Name: ptr.To("test_0")},
			},
			want: &V0044Qos{api.V0044Qos{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Qos{
				V0044Qos: tt.fields.V0044Qos,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Qos.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Qos_DeepCopy(t *testing.T) {
	type fields struct {
		V0044Qos api.V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044Qos
	}{
		{
			name: "empty",
			fields: fields{
				V0044Qos: api.V0044Qos{},
			},
			want: &V0044Qos{},
		},
		{
			name: "id",
			fields: fields{
				V0044Qos: api.V0044Qos{Name: ptr.To("test_0")},
			},
			want: &V0044Qos{api.V0044Qos{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Qos{
				V0044Qos: tt.fields.V0044Qos,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Qos.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044QosList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044Qos{},
			},
			want: ObjectTypeV0044Qos,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044QosList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044QosList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044QosList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Qos{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044Qos{
					{V0044Qos: api.V0044Qos{Name: ptr.To("test_0")}},
					{V0044Qos: api.V0044Qos{Name: ptr.To("normal")}},
				},
			},
			want: []object.Object{
				&V0044Qos{api.V0044Qos{Name: ptr.To("test_0")}},
				&V0044Qos{api.V0044Qos{Name: ptr.To("normal")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044QosList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044QosList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044QosList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044Qos
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
				Items: []V0044Qos{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044Qos{},
			},
			args: args{
				object: &V0044Qos{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Qos{
					{V0044Qos: api.V0044Qos{Name: ptr.To("test_0")}},
					{V0044Qos: api.V0044Qos{Name: ptr.To("normal")}},
				},
			},
			args: args{
				object: &V0044Qos{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044QosList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044QosList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044QosList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044Qos
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044Qos{},
			},
			want: &V0044QosList{
				Items: []V0044Qos{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044Qos{
					{V0044Qos: api.V0044Qos{Name: ptr.To("test_0")}},
					{V0044Qos: api.V0044Qos{Name: ptr.To("normal")}},
				},
			},
			want: &V0044QosList{
				Items: []V0044Qos{
					{api.V0044Qos{Name: ptr.To("test_0")}},
					{api.V0044Qos{Name: ptr.To("normal")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044QosList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044QosList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
