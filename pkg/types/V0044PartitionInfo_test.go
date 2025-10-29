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

func TestV0044PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0044PartitionInfo api.V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfo{
				V0044PartitionInfo: tt.fields.V0044PartitionInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0044PartitionInfo api.V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{},
			},
			want: ObjectTypeV0044PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfo{
				V0044PartitionInfo: tt.fields.V0044PartitionInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044PartitionInfo api.V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{},
			},
			want: &V0044PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0044PartitionInfo{api.V0044PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfo{
				V0044PartitionInfo: tt.fields.V0044PartitionInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0044PartitionInfo api.V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{},
			},
			want: &V0044PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0044PartitionInfo{api.V0044PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfo{
				V0044PartitionInfo: tt.fields.V0044PartitionInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0044PartitionInfo api.V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0044PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{},
			},
			want: set.New[api.V0044PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{
					Partition: &struct {
						State *[]api.V0044PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0044PartitionInfoPartitionState{api.V0044PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0044PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0044PartitionInfo: api.V0044PartitionInfo{
					Partition: &struct {
						State *[]api.V0044PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0044PartitionInfoPartitionState{api.V0044PartitionInfoPartitionStateUP, api.V0044PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0044PartitionInfoPartitionStateUP, api.V0044PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfo{
				V0044PartitionInfo: tt.fields.V0044PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0044PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044PartitionInfo{},
			},
			want: ObjectTypeV0044PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044PartitionInfo{
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")}},
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0044PartitionInfo{api.V0044PartitionInfo{Name: ptr.To("node-0")}},
				&V0044PartitionInfo{api.V0044PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044PartitionInfo
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
				Items: []V0044PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044PartitionInfo{},
			},
			args: args{
				object: &V0044PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044PartitionInfo{
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")}},
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0044PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044PartitionInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044PartitionInfo{},
			},
			want: &V0044PartitionInfoList{
				Items: []V0044PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044PartitionInfo{
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-0")}},
					{V0044PartitionInfo: api.V0044PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0044PartitionInfoList{
				Items: []V0044PartitionInfo{
					{api.V0044PartitionInfo{Name: ptr.To("node-0")}},
					{api.V0044PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
