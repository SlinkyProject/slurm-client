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

func TestV0045JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0045JobInfo api.V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfo{
				V0045JobInfo: tt.fields.V0045JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0045JobInfo api.V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{},
			},
			want: ObjectTypeV0045JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfo{
				V0045JobInfo: tt.fields.V0045JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0045JobInfo api.V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{},
			},
			want: &V0045JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0045JobInfo{api.V0045JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfo{
				V0045JobInfo: tt.fields.V0045JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0045JobInfo api.V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0045JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{},
			},
			want: &V0045JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0045JobInfo{api.V0045JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfo{
				V0045JobInfo: tt.fields.V0045JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0045JobInfo api.V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0045JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{},
			},
			want: set.New[api.V0045JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{
					JobState: ptr.To([]api.V0045JobInfoJobState{api.V0045JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0045JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0045JobInfo: api.V0045JobInfo{
					JobState: ptr.To([]api.V0045JobInfoJobState{api.V0045JobInfoJobStateRUNNING, api.V0045JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0045JobInfoJobStateRUNNING, api.V0045JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfo{
				V0045JobInfo: tt.fields.V0045JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0045JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045JobInfo{},
			},
			want: ObjectTypeV0045JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045JobInfo{
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)}},
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0045JobInfo{api.V0045JobInfo{JobId: ptr.To[int32](1)}},
				&V0045JobInfo{api.V0045JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045JobInfo
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
				Items: []V0045JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045JobInfo{},
			},
			args: args{
				object: &V0045JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045JobInfo{
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)}},
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0045JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045JobInfo{},
			},
			want: &V0045JobInfoList{
				Items: []V0045JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045JobInfo{
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](1)}},
					{V0045JobInfo: api.V0045JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0045JobInfoList{
				Items: []V0045JobInfo{
					{api.V0045JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0045JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
