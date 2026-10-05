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

func TestV0043PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0043PartitionInfo api.V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfo{
				V0043PartitionInfo: tt.fields.V0043PartitionInfo,
			}
			got := o.GetKey()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0043PartitionInfo api.V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{},
			},
			want: ObjectTypeV0043PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfo{
				V0043PartitionInfo: tt.fields.V0043PartitionInfo,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0043PartitionInfo api.V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{},
			},
			want: &V0043PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0043PartitionInfo{api.V0043PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfo{
				V0043PartitionInfo: tt.fields.V0043PartitionInfo,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0043PartitionInfo api.V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0043PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{},
			},
			want: &V0043PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0043PartitionInfo{api.V0043PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfo{
				V0043PartitionInfo: tt.fields.V0043PartitionInfo,
			}
			got := o.DeepCopy()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0043PartitionInfo{},
			},
			want: ObjectTypeV0043PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfoList{
				Items: tt.fields.Items,
			}
			got := o.GetType()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0043PartitionInfo{
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")}},
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0043PartitionInfo{api.V0043PartitionInfo{Name: ptr.To("node-0")}},
				&V0043PartitionInfo{api.V0043PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfoList{
				Items: tt.fields.Items,
			}
			got := o.GetItems()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestV0043PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0043PartitionInfo
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
				Items: []V0043PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0043PartitionInfo{},
			},
			args: args{
				object: &V0043PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043PartitionInfo{
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")}},
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0043PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfoList{
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

func TestV0043PartitionInfoList_DeepCopyObject(t *testing.T) {
	type fields struct {
		Items []V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043PartitionInfo{},
			},
			want: &V0043PartitionInfoList{
				Items: []V0043PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043PartitionInfo{
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-0")}},
					{V0043PartitionInfo: api.V0043PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0043PartitionInfoList{
				Items: []V0043PartitionInfo{
					{api.V0043PartitionInfo{Name: ptr.To("node-0")}},
					{api.V0043PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfoList{
				Items: tt.fields.Items,
			}
			got := o.DeepCopyObject()
			require.Equal(t, tt.want, got)
		})
	}
}
