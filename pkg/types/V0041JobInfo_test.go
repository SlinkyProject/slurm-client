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

func TestV0041JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0041JobInfo api.V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfo{
				V0041JobInfo: tt.fields.V0041JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0041JobInfo api.V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{},
			},
			want: ObjectTypeV0041JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfo{
				V0041JobInfo: tt.fields.V0041JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0041JobInfo api.V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{},
			},
			want: &V0041JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0041JobInfo{api.V0041JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfo{
				V0041JobInfo: tt.fields.V0041JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0041JobInfo api.V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0041JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{},
			},
			want: &V0041JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0041JobInfo{api.V0041JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfo{
				V0041JobInfo: tt.fields.V0041JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0041JobInfo api.V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0041JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{},
			},
			want: set.New[api.V0041JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{
					JobState: ptr.To([]api.V0041JobInfoJobState{api.V0041JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0041JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0041JobInfo: api.V0041JobInfo{
					JobState: ptr.To([]api.V0041JobInfoJobState{api.V0041JobInfoJobStateRUNNING, api.V0041JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0041JobInfoJobStateRUNNING, api.V0041JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfo{
				V0041JobInfo: tt.fields.V0041JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0041JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0041JobInfo{},
			},
			want: ObjectTypeV0041JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0041JobInfo{
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)}},
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0041JobInfo{api.V0041JobInfo{JobId: ptr.To[int32](1)}},
				&V0041JobInfo{api.V0041JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0041JobInfo
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
				Items: []V0041JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0041JobInfo{},
			},
			args: args{
				object: &V0041JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041JobInfo{
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)}},
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0041JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0041JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0041JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0041JobInfo{},
			},
			want: &V0041JobInfoList{
				Items: []V0041JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0041JobInfo{
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](1)}},
					{V0041JobInfo: api.V0041JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0041JobInfoList{
				Items: []V0041JobInfo{
					{api.V0041JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0041JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
