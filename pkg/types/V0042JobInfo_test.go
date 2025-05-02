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

func TestV0042JobInfo_GetKey(t *testing.T) {
	type fields struct {
		V0042JobInfo api.V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{},
			},
			want: "0",
		},
		{
			name: "key",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfo{
				V0042JobInfo: tt.fields.V0042JobInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfo_GetType(t *testing.T) {
	type fields struct {
		V0042JobInfo api.V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{},
			},
			want: ObjectTypeV0042JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfo{
				V0042JobInfo: tt.fields.V0042JobInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0042JobInfo api.V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{},
			},
			want: &V0042JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0042JobInfo{api.V0042JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfo{
				V0042JobInfo: tt.fields.V0042JobInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0042JobInfo api.V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0042JobInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{},
			},
			want: &V0042JobInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)},
			},
			want: &V0042JobInfo{api.V0042JobInfo{JobId: ptr.To[int32](1)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfo{
				V0042JobInfo: tt.fields.V0042JobInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfo_GetStateAsSet(t *testing.T) {
	type fields struct {
		V0042JobInfo api.V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   set.Set[api.V0042JobInfoJobState]
	}{
		{
			name: "empty",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{},
			},
			want: set.New[api.V0042JobInfoJobState](),
		},
		{
			name: "single",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{
					JobState: ptr.To([]api.V0042JobInfoJobState{api.V0042JobInfoJobStatePENDING}),
				},
			},
			want: set.New(api.V0042JobInfoJobStatePENDING),
		},
		{
			name: "multiple",
			fields: fields{
				V0042JobInfo: api.V0042JobInfo{
					JobState: ptr.To([]api.V0042JobInfoJobState{api.V0042JobInfoJobStateRUNNING, api.V0042JobInfoJobStateREQUEUEFED}),
				},
			},
			want: set.New(api.V0042JobInfoJobStateRUNNING, api.V0042JobInfoJobStateREQUEUEFED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfo{
				V0042JobInfo: tt.fields.V0042JobInfo,
			}
			if got := o.GetStateAsSet(); !tt.want.Equal(got) {
				t.Errorf("V0042JobInfo.GetStateAsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0042JobInfo{},
			},
			want: ObjectTypeV0042JobInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042JobInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0042JobInfo{
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)}},
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: []object.Object{
				&V0042JobInfo{api.V0042JobInfo{JobId: ptr.To[int32](1)}},
				&V0042JobInfo{api.V0042JobInfo{JobId: ptr.To[int32](2)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042JobInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0042JobInfo
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
				Items: []V0042JobInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0042JobInfo{},
			},
			args: args{
				object: &V0042JobInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042JobInfo{
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)}},
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			args: args{
				object: &V0042JobInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0042JobInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042JobInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0042JobInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0042JobInfo{},
			},
			want: &V0042JobInfoList{
				Items: []V0042JobInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0042JobInfo{
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](1)}},
					{V0042JobInfo: api.V0042JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
			want: &V0042JobInfoList{
				Items: []V0042JobInfo{
					{api.V0042JobInfo{JobId: ptr.To[int32](1)}},
					{api.V0042JobInfo{JobId: ptr.To[int32](2)}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042JobInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042JobInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
