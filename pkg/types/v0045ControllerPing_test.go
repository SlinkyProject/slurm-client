// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045ControllerPing_GetKey(t *testing.T) {
	type fields struct {
		V0045ControllerPing api.V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: "controller-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPing{
				V0045ControllerPing: tt.fields.V0045ControllerPing,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPing.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPing_GetType(t *testing.T) {
	type fields struct {
		V0045ControllerPing api.V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{},
			},
			want: ObjectTypeV0045ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPing{
				V0045ControllerPing: tt.fields.V0045ControllerPing,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPing.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPing_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045ControllerPing api.V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{},
			},
			want: &V0045ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0045ControllerPing{api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPing{
				V0045ControllerPing: tt.fields.V0045ControllerPing,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPing.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPing_DeepCopy(t *testing.T) {
	type fields struct {
		V0045ControllerPing api.V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045ControllerPing
	}{
		{
			name: "empty",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{},
			},
			want: &V0045ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0045ControllerPing{api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPing{
				V0045ControllerPing: tt.fields.V0045ControllerPing,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPing.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPingList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045ControllerPing{},
			},
			want: ObjectTypeV0045ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPingList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPingList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045ControllerPing{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045ControllerPing{
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: []object.Object{
				&V0045ControllerPing{api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
				&V0045ControllerPing{api.V0045ControllerPing{Hostname: ptr.To("controller-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPingList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ControllerPingList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045ControllerPing
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
				Items: []V0045ControllerPing{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045ControllerPing{},
			},
			args: args{
				object: &V0045ControllerPing{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045ControllerPing{
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			args: args{
				object: &V0045ControllerPing{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPingList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045ControllerPingList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045ControllerPingList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045ControllerPing{},
			},
			want: &V0045ControllerPingList{
				Items: []V0045ControllerPing{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045ControllerPing{
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0045ControllerPing: api.V0045ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: &V0045ControllerPingList{
				Items: []V0045ControllerPing{
					{api.V0045ControllerPing{Hostname: ptr.To("controller-0")}},
					{api.V0045ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ControllerPingList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
