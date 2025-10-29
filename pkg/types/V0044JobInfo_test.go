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

func TestV0044JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0044JobInfo api.V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfo{
				V0044JobInfo: tt.fields.V0044JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0044JobInfo api.V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{},
			},
			want: ObjectTypeV0044JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfo{
				V0044JobInfo: tt.fields.V0044JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044JobInfo api.V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{},
			},
			want: &V0044JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0044JobInfo{api.V0044JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfo{
				V0044JobInfo: tt.fields.V0044JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0044JobInfo api.V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{},
			},
			want: &V0044JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0044JobInfo{api.V0044JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfo{
				V0044JobInfo: tt.fields.V0044JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0044JobInfo api.V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0044JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{},
			},
			want: set.New[api.V0044JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{
					JobState: ptr.To([]api.V0044JobInfoJobState{api.V0044JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0044JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0044JobInfo: api.V0044JobInfo{
					JobState: ptr.To([]api.V0044JobInfoJobState{api.V0044JobInfoJobStateRUNNING, api.V0044JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0044JobInfoJobStateRUNNING, api.V0044JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfo{
				V0044JobInfo: tt.fields.V0044JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0044JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044JobInfo{},
			},
			want: ObjectTypeV0044JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044JobInfo{
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)}},
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0044JobInfo{api.V0044JobInfo{JobId: ptr.To[int32](1)}},
				&V0044JobInfo{api.V0044JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044JobInfo
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
				Items: []V0044JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044JobInfo{},
			},
			args: args{
				object: &V0044JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044JobInfo{
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)}},
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0044JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044JobInfo{},
			},
			want: &V0044JobInfoList{
				Items: []V0044JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044JobInfo{
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](1)}},
					{V0044JobInfo: api.V0044JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0044JobInfoList{
				Items: []V0044JobInfo{
					{api.V0044JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0044JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
