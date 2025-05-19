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
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
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
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
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
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
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
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0043PartitionInfo api.V0043PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0043PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{},
			},
			want: set.New[api.V0043PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{
					Partition: &struct {
						State *[]api.V0043PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0043PartitionInfoPartitionState{api.V0043PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0043PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0043PartitionInfo: api.V0043PartitionInfo{
					Partition: &struct {
						State *[]api.V0043PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0043PartitionInfoPartitionState{api.V0043PartitionInfoPartitionStateUP, api.V0043PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0043PartitionInfoPartitionStateUP, api.V0043PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043PartitionInfo{
				V0043PartitionInfo: tt.fields.V0043PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0043PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
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
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
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
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
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
			if want != got {
				t.Errorf("V0043PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0043PartitionInfoList_DeepCopyObjectList(t *testing.T) {
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
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
