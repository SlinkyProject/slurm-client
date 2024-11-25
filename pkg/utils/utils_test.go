// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"

	"k8s.io/utils/ptr"
)

type objectA struct {
	Str    string  `json:"str"`
	StrPtr *string `json:"str_ptr"`
	Int    int32   `json:"int"`
	IntPtr *int32  `json:"int_ptr"`
}

type objectB struct {
	objectA
}

func TestRemarshal(t *testing.T) {
	objA := objectA{
		Str:    "foo",
		StrPtr: ptr.To("bar"),
		Int:    1,
		IntPtr: ptr.To[int32](2),
	}
	type args struct {
		in  any
		out any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Empty object",
			args: args{
				in:  objectA{},
				out: &objectB{},
			},
			wantErr: false,
		},
		{
			name: "With non-nil data",
			args: args{
				in:  objA,
				out: &objectB{},
			},
			wantErr: false,
		},
		{
			name: "With pointer in struct",
			args: args{
				in:  &objA,
				out: &objectB{},
			},
			wantErr: false,
		},
		{
			name: "With non-pointer in, out struct",
			args: args{
				in:  objA,
				out: objectB{},
			},
			wantErr: true,
		},
		{
			name: "With non-pointer out struct",
			args: args{
				in:  &objA,
				out: objectB{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remarshal(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Remarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemarshalOrDie(t *testing.T) {
	objA := objectA{
		Str:    "foo",
		StrPtr: ptr.To("bar"),
		Int:    1,
		IntPtr: ptr.To[int32](2),
	}
	objB := &objectB{}
	type args struct {
		in  any
		out any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty object",
			args: args{
				in:  objectA{},
				out: &objectB{},
			},
		},
		{
			name: "With non-nil data",
			args: args{
				in:  objA,
				out: &objB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemarshalOrDie(tt.args.in, tt.args.out)
		})
	}
}
