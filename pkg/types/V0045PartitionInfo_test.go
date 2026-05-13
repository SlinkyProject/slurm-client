// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"
	"k8s.io/utils/set"

	api "github.com/SlinkyProject/slurm-client/api/v0045"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045PartitionInfo_GetKey(t *testing.T) {
	type fields struct {
		V0045PartitionInfo api.V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")},
			},
			want: "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfo{
				V0045PartitionInfo: tt.fields.V0045PartitionInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfo_GetType(t *testing.T) {
	type fields struct {
		V0045PartitionInfo api.V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{},
			},
			want: ObjectTypeV0045PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfo{
				V0045PartitionInfo: tt.fields.V0045PartitionInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045PartitionInfo api.V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{},
			},
			want: &V0045PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")},
			},
			want: &V0045PartitionInfo{api.V0045PartitionInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfo{
				V0045PartitionInfo: tt.fields.V0045PartitionInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0045PartitionInfo api.V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045PartitionInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{},
			},
			want: &V0045PartitionInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")},
			},
			want: &V0045PartitionInfo{api.V0045PartitionInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfo{
				V0045PartitionInfo: tt.fields.V0045PartitionInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0045PartitionInfo api.V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0045PartitionInfoPartitionState]
	}{
		{
			name: "empty",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{},
			},
			want: set.New[api.V0045PartitionInfoPartitionState](),
		},
		{
			name: "single",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{
					Partition: &struct {
						Exclusive     *string                                 "json:\"exclusive,omitempty\""
						Oversubscribe *string                                 "json:\"oversubscribe,omitempty\""
						State         *[]api.V0045PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0045PartitionInfoPartitionState{api.V0045PartitionInfoPartitionStateUP}),
					},
				},
			},
			want: set.New(api.V0045PartitionInfoPartitionStateUP),
		},
		{
			name: "multiple",
			fields: fields{
				V0045PartitionInfo: api.V0045PartitionInfo{
					Partition: &struct {
						Exclusive     *string                                 "json:\"exclusive,omitempty\""
						Oversubscribe *string                                 "json:\"oversubscribe,omitempty\""
						State         *[]api.V0045PartitionInfoPartitionState "json:\"state,omitempty\""
					}{
						State: ptr.To([]api.V0045PartitionInfoPartitionState{api.V0045PartitionInfoPartitionStateUP, api.V0045PartitionInfoPartitionStateDRAIN}),
					},
				},
			},
			want: set.New(api.V0045PartitionInfoPartitionStateUP, api.V0045PartitionInfoPartitionStateDRAIN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfo{
				V0045PartitionInfo: tt.fields.V0045PartitionInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0045PartitionInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045PartitionInfo{},
			},
			want: ObjectTypeV0045PartitionInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045PartitionInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045PartitionInfo{
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")}},
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0045PartitionInfo{api.V0045PartitionInfo{Name: ptr.To("test_0")}},
				&V0045PartitionInfo{api.V0045PartitionInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045PartitionInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045PartitionInfo
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
				Items: []V0045PartitionInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045PartitionInfo{},
			},
			args: args{
				object: &V0045PartitionInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045PartitionInfo{
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")}},
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0045PartitionInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045PartitionInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045PartitionInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045PartitionInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045PartitionInfo{},
			},
			want: &V0045PartitionInfoList{
				Items: []V0045PartitionInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045PartitionInfo{
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("test_0")}},
					{V0045PartitionInfo: api.V0045PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0045PartitionInfoList{
				Items: []V0045PartitionInfo{
					{api.V0045PartitionInfo{Name: ptr.To("test_0")}},
					{api.V0045PartitionInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045PartitionInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045PartitionInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
