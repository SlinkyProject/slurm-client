// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0042PartitionInfo api.V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")},
			},
			want: "node-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfo{
				V0042PartitionInfo: tt.fields.V0042PartitionInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0042PartitionInfo api.V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{},
			},
			want: ObjectTypeV0042PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfo{
				V0042PartitionInfo: tt.fields.V0042PartitionInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0042PartitionInfo api.V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{},
			},
			want: &V0042PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0042PartitionInfo{api.V0042PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfo{
				V0042PartitionInfo: tt.fields.V0042PartitionInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0042PartitionInfo api.V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0042PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{},
			},
			want: &V0042PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")},
			},
			want: &V0042PartitionInfo{api.V0042PartitionInfo{Name: ptr.To("node-0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfo{
				V0042PartitionInfo: tt.fields.V0042PartitionInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0042PartitionInfo api.V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0042PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{},
			},
			want: set.New[api.V0042PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{
					Partition: &struct {
						State *[]api.V0042PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0042PartitionInfoPartitionState{api.V0042PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0042PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0042PartitionInfo: api.V0042PartitionInfo{
					Partition: &struct {
						State *[]api.V0042PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0042PartitionInfoPartitionState{api.V0042PartitionInfoPartitionStateUP, api.V0042PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0042PartitionInfoPartitionStateUP, api.V0042PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfo{
				V0042PartitionInfo: tt.fields.V0042PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0042PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0042PartitionInfo{},
			},
			want: ObjectTypeV0042PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0042PartitionInfo{
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")}},
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0042PartitionInfo{api.V0042PartitionInfo{Name: ptr.To("node-0")}},
				&V0042PartitionInfo{api.V0042PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0042PartitionInfo
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
				Items: []V0042PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0042PartitionInfo{},
			},
			args: args{
				object: &V0042PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042PartitionInfo{
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")}},
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0042PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0042PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042PartitionInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0042PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042PartitionInfo{},
			},
			want: &V0042PartitionInfoList{
				Items: []V0042PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042PartitionInfo{
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-0")}},
					{V0042PartitionInfo: api.V0042PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0042PartitionInfoList{
				Items: []V0042PartitionInfo{
					{api.V0042PartitionInfo{Name: ptr.To("node-0")}},
					{api.V0042PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
