// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0041PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0041PartitionInfo api.V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfo{
				V0041PartitionInfo: tt.fields.V0041PartitionInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0041PartitionInfo api.V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{},
			},
			want: ObjectTypeV0041PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfo{
				V0041PartitionInfo: tt.fields.V0041PartitionInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0041PartitionInfo api.V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{},
			},
			want: &V0041PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0041PartitionInfo{api.V0041PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfo{
				V0041PartitionInfo: tt.fields.V0041PartitionInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0041PartitionInfo api.V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0041PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{},
			},
			want: &V0041PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0041PartitionInfo{api.V0041PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfo{
				V0041PartitionInfo: tt.fields.V0041PartitionInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0041PartitionInfo api.V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0041PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{},
			},
			want: set.New[api.V0041PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{
					Partition: &struct {
						State *[]api.V0041PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0041PartitionInfoPartitionState{api.V0041PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0041PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0041PartitionInfo: api.V0041PartitionInfo{
					Partition: &struct {
						State *[]api.V0041PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0041PartitionInfoPartitionState{api.V0041PartitionInfoPartitionStateUP, api.V0041PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0041PartitionInfoPartitionStateUP, api.V0041PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfo{
				V0041PartitionInfo: tt.fields.V0041PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0041PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0041PartitionInfo{},
			},
			want: ObjectTypeV0041PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0041PartitionInfo{
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")}},
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0041PartitionInfo{api.V0041PartitionInfo{Name: ptr.To("node-0")}},
				&V0041PartitionInfo{api.V0041PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0041PartitionInfo
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
				Items: []V0041PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0041PartitionInfo{},
			},
			args: args{
				object: &V0041PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041PartitionInfo{
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")}},
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0041PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0041PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041PartitionInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0041PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041PartitionInfo{},
			},
			want: &V0041PartitionInfoList{
				Items: []V0041PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041PartitionInfo{
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-0")}},
					{V0041PartitionInfo: api.V0041PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0041PartitionInfoList{
				Items: []V0041PartitionInfo{
					{api.V0041PartitionInfo{Name: ptr.To("node-0")}},
					{api.V0041PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
