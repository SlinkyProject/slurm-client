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

func TestV0040JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0040JobInfo api.V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfo{
				V0040JobInfo: tt.fields.V0040JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0040JobInfo api.V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{},
			},
			want: ObjectTypeV0040JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfo{
				V0040JobInfo: tt.fields.V0040JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0040JobInfo api.V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{},
			},
			want: &V0040JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0040JobInfo{api.V0040JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfo{
				V0040JobInfo: tt.fields.V0040JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0040JobInfo api.V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0040JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{},
			},
			want: &V0040JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0040JobInfo{api.V0040JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfo{
				V0040JobInfo: tt.fields.V0040JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0040JobInfo api.V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0040JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{},
			},
			want: set.New[api.V0040JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{
					JobState: ptr.To([]api.V0040JobInfoJobState{api.V0040JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0040JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0040JobInfo: api.V0040JobInfo{
					JobState: ptr.To([]api.V0040JobInfoJobState{api.V0040JobInfoJobStateRUNNING, api.V0040JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0040JobInfoJobStateRUNNING, api.V0040JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfo{
				V0040JobInfo: tt.fields.V0040JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0040JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0040JobInfo{},
			},
			want: ObjectTypeV0040JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0040JobInfo{
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)}},
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0040JobInfo{api.V0040JobInfo{JobId: ptr.To[int32](1)}},
				&V0040JobInfo{api.V0040JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0040JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0040JobInfo
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
				Items: []V0040JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0040JobInfo{},
			},
			args: args{
				object: &V0040JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040JobInfo{
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)}},
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0040JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0040JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0040JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0040JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0040JobInfo{},
			},
			want: &V0040JobInfoList{
				Items: []V0040JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0040JobInfo{
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](1)}},
					{V0040JobInfo: api.V0040JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0040JobInfoList{
				Items: []V0040JobInfo{
					{api.V0040JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0040JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0040JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0040JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
