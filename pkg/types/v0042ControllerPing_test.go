// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042ControllerPing_GetKey(t *testing.T) {
	type fields struct {
		V0042ControllerPing api.V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: "controller-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPing{
				V0042ControllerPing: tt.fields.V0042ControllerPing,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPing.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPing_GetType(t *testing.T) {
	type fields struct {
		V0042ControllerPing api.V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{},
			},
			want: ObjectTypeV0042ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPing{
				V0042ControllerPing: tt.fields.V0042ControllerPing,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPing.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPing_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0042ControllerPing api.V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{},
			},
			want: &V0042ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0042ControllerPing{api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPing{
				V0042ControllerPing: tt.fields.V0042ControllerPing,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPing.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPing_DeepCopy(t *testing.T) {
	type fields struct {
		V0042ControllerPing api.V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0042ControllerPing
	}{
		{
			name: "empty",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{},
			},
			want: &V0042ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0042ControllerPing{api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPing{
				V0042ControllerPing: tt.fields.V0042ControllerPing,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPing.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPingList_GetType(t *testing.T) {
	type fields struct {
		Items []V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0042ControllerPing{},
			},
			want: ObjectTypeV0042ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPingList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPingList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042ControllerPing{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0042ControllerPing{
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: []object.Object{
				&V0042ControllerPing{api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
				&V0042ControllerPing{api.V0042ControllerPing{Hostname: ptr.To("controller-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPingList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ControllerPingList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0042ControllerPing
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
				Items: []V0042ControllerPing{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0042ControllerPing{},
			},
			args: args{
				object: &V0042ControllerPing{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042ControllerPing{
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			args: args{
				object: &V0042ControllerPing{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPingList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0042ControllerPingList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042ControllerPingList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0042ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042ControllerPing{},
			},
			want: &V0042ControllerPingList{
				Items: []V0042ControllerPing{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042ControllerPing{
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0042ControllerPing: api.V0042ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: &V0042ControllerPingList{
				Items: []V0042ControllerPing{
					{api.V0042ControllerPing{Hostname: ptr.To("controller-0")}},
					{api.V0042ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ControllerPingList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ControllerPingList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
