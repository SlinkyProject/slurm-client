// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0045Reconfigure_GetKey(t *testing.T) {
	tests := []struct {
		name string
		want object.ObjectKey
	}{
		{
			name: "key",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Reconfigure{}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Reconfigure.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Reconfigure_GetType(t *testing.T) {
	tests := []struct {
		name string
		want object.ObjectType
	}{
		{
			name: "type",
			want: ObjectTypeV0045Reconfigure,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Reconfigure{}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Reconfigure.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Reconfigure_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name string
		want object.Object
	}{
		{
			name: "empty",
			want: &V0045Reconfigure{},
		},
		{
			name: "id",
			want: &V0045Reconfigure{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Reconfigure{}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Reconfigure.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045Reconfigure_DeepCopy(t *testing.T) {
	tests := []struct {
		name string
		want *V0045Reconfigure
	}{
		{
			name: "empty",
			want: &V0045Reconfigure{},
		},
		{
			name: "id",
			want: &V0045Reconfigure{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045Reconfigure{}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045Reconfigure.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ReconfigureList_GetType(t *testing.T) {
	type fields struct {
		Items []V0045Reconfigure
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectType
	}{
		{
			name: "type",
			fields: fields{
				Items: []V0045Reconfigure{},
			},
			want: ObjectTypeV0045Reconfigure,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReconfigureList{
				Items: tt.fields.Items,
			}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ReconfigureList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ReconfigureList_GetItems(t *testing.T) {
	type fields struct {
		Items []V0045Reconfigure
	}
	tests := []struct {
		name   string
		fields fields
		want   []object.Object
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Reconfigure{},
			},
			want: []object.Object{},
		},
		{
			name: "items",
			fields: fields{
				Items: []V0045Reconfigure{
					{},
				},
			},
			want: []object.Object{
				&V0045Reconfigure{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReconfigureList{
				Items: tt.fields.Items,
			}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ReconfigureList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0045ReconfigureList_AppendItem(t *testing.T) {
	type fields struct {
		Items []V0045Reconfigure
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
				Items: []V0045Reconfigure{},
			},
			args: args{
				object: nil,
			},
			wantAppend: false,
		},
		{
			name: "empty",
			fields: fields{
				Items: []V0045Reconfigure{},
			},
			args: args{
				object: &V0045Reconfigure{},
			},
			wantAppend: true,
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Reconfigure{
					{},
				},
			},
			args: args{
				object: &V0045Reconfigure{},
			},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReconfigureList{
				Items: tt.fields.Items,
			}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.args.object)
			got := len(o.GetItems())
			if want != got {
				t.Errorf("V0045ReconfigureList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0045ReconfigureList_DeepCopyObjectList(t *testing.T) {
	type fields struct {
		Items []V0045Reconfigure
	}
	tests := []struct {
		name   string
		fields fields
		want   object.ObjectList
	}{
		{
			name: "empty",
			fields: fields{
				Items: []V0045Reconfigure{},
			},
			want: &V0045ReconfigureList{
				Items: []V0045Reconfigure{},
			},
		},
		{
			name: "existing",
			fields: fields{
				Items: []V0045Reconfigure{
					{},
				},
			},
			want: &V0045ReconfigureList{
				Items: []V0045Reconfigure{
					{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0045ReconfigureList{
				Items: tt.fields.Items,
			}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0045ReconfigureList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
