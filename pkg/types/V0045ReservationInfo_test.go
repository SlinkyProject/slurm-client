// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"k8s.io/utils/ptr"

	"github.com/stretchr/testify/require"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045ReservationInfo_GetKey(t *testing.T) {
	type fields struct {
		V0045ReservationInfo api.V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")},
			},
			want: "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfo{
				V0045ReservationInfo: tt.fields.V0045ReservationInfo,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfo_GetType(t *testing.T) {
	type fields struct {
		V0045ReservationInfo api.V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{},
			},
			want: ObjectTypeV0045ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfo{
				V0045ReservationInfo: tt.fields.V0045ReservationInfo,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045ReservationInfo api.V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{},
			},
			want: &V0045ReservationInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")},
			},
			want: &V0045ReservationInfo{api.V0045ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfo{
				V0045ReservationInfo: tt.fields.V0045ReservationInfo,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0045ReservationInfo api.V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045ReservationInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{},
			},
			want: &V0045ReservationInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")},
			},
			want: &V0045ReservationInfo{api.V0045ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfo{
				V0045ReservationInfo: tt.fields.V0045ReservationInfo,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045ReservationInfo{},
			},
			want: ObjectTypeV0045ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfoList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045ReservationInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045ReservationInfo{
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")}},
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0045ReservationInfo{api.V0045ReservationInfo{Name: ptr.To("test_0")}},
				&V0045ReservationInfo{api.V0045ReservationInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfoList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0045ReservationInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045ReservationInfo
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
				Items: []V0045ReservationInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045ReservationInfo{},
			},
			args: args{
				object: &V0045ReservationInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045ReservationInfo{
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")}},
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0045ReservationInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfoList{
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

func TestV0045ReservationInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045ReservationInfo{},
			},
			want: &V0045ReservationInfoList{
				Items: []V0045ReservationInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045ReservationInfo{
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("test_0")}},
					{V0045ReservationInfo: api.V0045ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0045ReservationInfoList{
				Items: []V0045ReservationInfo{
					{api.V0045ReservationInfo{Name: ptr.To("test_0")}},
					{api.V0045ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReservationInfoList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObjectList()
			require.Equal(t, tt.want, got)
		})
	}
}
