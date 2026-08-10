// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043ControllerPing_GetKey(t *testing.T) {
	type fields struct {
		V0043ControllerPing api.V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: "controller-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPing{
				V0043ControllerPing: tt.fields.V0043ControllerPing,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPing_GetType(t *testing.T) {
	type fields struct {
		V0043ControllerPing api.V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{},
			},
			want: ObjectTypeV0043ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPing{
				V0043ControllerPing: tt.fields.V0043ControllerPing,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPing_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0043ControllerPing api.V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{},
			},
			want: &V0043ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0043ControllerPing{api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPing{
				V0043ControllerPing: tt.fields.V0043ControllerPing,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPing_DeepCopy(t *testing.T) {
	type fields struct {
		V0043ControllerPing api.V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0043ControllerPing
	}{
		{
			name: "empty",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{},
			},
			want: &V0043ControllerPing{},
		},
		{
			name: "id",
			fields: fields{
				V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")},
			},
			want: &V0043ControllerPing{api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPing{
				V0043ControllerPing: tt.fields.V0043ControllerPing,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPingList_GetType(t *testing.T) {
	type fields struct {
		Items []V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0043ControllerPing{},
			},
			want: ObjectTypeV0043ControllerPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPingList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPingList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043ControllerPing{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0043ControllerPing{
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: []object.Object{
				&V0043ControllerPing{api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
				&V0043ControllerPing{api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPingList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043ControllerPingList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0043ControllerPing
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
				Items: []V0043ControllerPing{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0043ControllerPing{},
			},
			args: args{
				object: &V0043ControllerPing{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043ControllerPing{
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			args: args{
				object: &V0043ControllerPing{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPingList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			require.Equal(t, want, got)
		})
	}
}

func TestV0043ControllerPingList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0043ControllerPing
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043ControllerPing{},
			},
			want: &V0043ControllerPingList{
				Items: []V0043ControllerPing{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043ControllerPing{
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
					{V0043ControllerPing: api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
			want: &V0043ControllerPingList{
				Items: []V0043ControllerPing{
					{api.V0043ControllerPing{Hostname: ptr.To("controller-0")}},
					{api.V0043ControllerPing{Hostname: ptr.To("controller-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ControllerPingList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
