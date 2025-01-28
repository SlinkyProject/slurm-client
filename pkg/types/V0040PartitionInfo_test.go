// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0040"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0040PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0040PartitionInfo api.V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfo{
				V0040PartitionInfo: tt.fields.V0040PartitionInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0040PartitionInfo api.V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{},
			},
			want: ObjectTypeV0040PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfo{
				V0040PartitionInfo: tt.fields.V0040PartitionInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0040PartitionInfo api.V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{},
			},
			want: &V0040PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0040PartitionInfo{api.V0040PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfo{
				V0040PartitionInfo: tt.fields.V0040PartitionInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0040PartitionInfo api.V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0040PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{},
			},
			want: &V0040PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0040PartitionInfo{api.V0040PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfo{
				V0040PartitionInfo: tt.fields.V0040PartitionInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0040PartitionInfo api.V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0040PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{},
			},
			want: set.New[api.V0040PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{
					Partition: &struct {
						State *[]api.V0040PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0040PartitionInfoPartitionState{api.V0040PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0040PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0040PartitionInfo: api.V0040PartitionInfo{
					Partition: &struct {
						State *[]api.V0040PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0040PartitionInfoPartitionState{api.V0040PartitionInfoPartitionStateUP, api.V0040PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0040PartitionInfoPartitionStateUP, api.V0040PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfo{
				V0040PartitionInfo: tt.fields.V0040PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0040PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0040PartitionInfo{},
			},
			want: ObjectTypeV0040PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0040PartitionInfo{
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")}},
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0040PartitionInfo{api.V0040PartitionInfo{Name: ptr.To("node-0")}},
				&V0040PartitionInfo{api.V0040PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0040PartitionInfo
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
				Items: []V0040PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0040PartitionInfo{},
			},
			args: args{
				object: &V0040PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040PartitionInfo{
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")}},
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0040PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0040PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0040PartitionInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0040PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040PartitionInfo{},
			},
			want: &V0040PartitionInfoList{
				Items: []V0040PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040PartitionInfo{
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-0")}},
					{V0040PartitionInfo: api.V0040PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0040PartitionInfoList{
				Items: []V0040PartitionInfo{
					{api.V0040PartitionInfo{Name: ptr.To("node-0")}},
					{api.V0040PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
