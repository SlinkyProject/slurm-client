// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044ReservationInfo_GetKey(t *testing.T) {
	type fields struct {
		V0044ReservationInfo api.V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectKey
	}{
		{
			name: "empty",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{},
			},
			want: "",
		},
		{
			name: "key",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")},
			},
			want: "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfo{
				V0044ReservationInfo: tt.fields.V0044ReservationInfo,
			}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfo_GetType(t *testing.T) {
	type fields struct {
		V0044ReservationInfo api.V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{},
			},
			want: ObjectTypeV0044ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfo{
				V0044ReservationInfo: tt.fields.V0044ReservationInfo,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfo_DeepCopyObject(t *testing.T) {
	type fields struct {
		V0044ReservationInfo api.V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.Object
	}{
		{
			name: "empty",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{},
			},
			want: &V0044ReservationInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")},
			},
			want: &V0044ReservationInfo{api.V0044ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfo{
				V0044ReservationInfo: tt.fields.V0044ReservationInfo,
			}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfo_DeepCopy(t *testing.T) {
	type fields struct {
		V0044ReservationInfo api.V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   *V0044ReservationInfo
	}{
		{
			name: "empty",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{},
			},
			want: &V0044ReservationInfo{},
		},
		{
			name: "id",
			fields: fields{
				V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")},
			},
			want: &V0044ReservationInfo{api.V0044ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfo{
				V0044ReservationInfo: tt.fields.V0044ReservationInfo,
			}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfoList_GetType(t *testing.T) {
	type fields struct {
		Items []V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0044ReservationInfo{},
			},
			want: ObjectTypeV0044ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfoList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044ReservationInfo{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0044ReservationInfo{
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")}},
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			want: []object.Object{
				&V0044ReservationInfo{api.V0044ReservationInfo{Name: ptr.To("test_0")}},
				&V0044ReservationInfo{api.V0044ReservationInfo{Name: ptr.To("node-1")}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfoList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044ReservationInfoList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0044ReservationInfo
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
				Items: []V0044ReservationInfo{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0044ReservationInfo{},
			},
			args: args{
				object: &V0044ReservationInfo{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044ReservationInfo{
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")}},
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			args: args{
				object: &V0044ReservationInfo{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfoList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0044ReservationInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044ReservationInfoList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0044ReservationInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0044ReservationInfo{},
			},
			want: &V0044ReservationInfoList{
				Items: []V0044ReservationInfo{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0044ReservationInfo{
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("test_0")}},
					{V0044ReservationInfo: api.V0044ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
			want: &V0044ReservationInfoList{
				Items: []V0044ReservationInfo{
					{api.V0044ReservationInfo{Name: ptr.To("test_0")}},
					{api.V0044ReservationInfo{Name: ptr.To("node-1")}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044ReservationInfoList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044ReservationInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
