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

func TestV0043JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0043JobInfo api.V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfo{
				V0043JobInfo: tt.fields.V0043JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0043JobInfo api.V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{},
			},
			want: ObjectTypeV0043JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfo{
				V0043JobInfo: tt.fields.V0043JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0043JobInfo api.V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{},
			},
			want: &V0043JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0043JobInfo{api.V0043JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfo{
				V0043JobInfo: tt.fields.V0043JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0043JobInfo api.V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0043JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{},
			},
			want: &V0043JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0043JobInfo{api.V0043JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfo{
				V0043JobInfo: tt.fields.V0043JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0043JobInfo api.V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0043JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{},
			},
			want: set.New[api.V0043JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{
					JobState: ptr.To([]api.V0043JobInfoJobState{api.V0043JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0043JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0043JobInfo: api.V0043JobInfo{
					JobState: ptr.To([]api.V0043JobInfoJobState{api.V0043JobInfoJobStateRUNNING, api.V0043JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0043JobInfoJobStateRUNNING, api.V0043JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfo{
				V0043JobInfo: tt.fields.V0043JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0043JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0043JobInfo{},
			},
			want: ObjectTypeV0043JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0043JobInfo{
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)}},
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0043JobInfo{api.V0043JobInfo{JobId: ptr.To[int32](1)}},
				&V0043JobInfo{api.V0043JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0043JobInfo
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
				Items: []V0043JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0043JobInfo{},
			},
			args: args{
				object: &V0043JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043JobInfo{
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)}},
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0043JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0043JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0043JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0043JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0043JobInfo{},
			},
			want: &V0043JobInfoList{
				Items: []V0043JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0043JobInfo{
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](1)}},
					{V0043JobInfo: api.V0043JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0043JobInfoList{
				Items: []V0043JobInfo{
					{api.V0043JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0043JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
