// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"
)

type ObjectA struct {
	Str    string
	StrPtr *string
	Int    int32
	IntPtr *int32
}

type ObjectB struct {
	ObjectA
}

func sampleA() ObjectA {
	return ObjectA{
		Str:    "foo",
		StrPtr: new("bar"),
		Int:    1,
		IntPtr: new(int32(2)),
	}
}

func TestClone(t *testing.T) {
	in := sampleA()
	got := Clone(&in)
	if got == &in || got.StrPtr == in.StrPtr || got.IntPtr == in.IntPtr {
		t.Fatalf("Clone() did not deep copy pointers")
	}
	if got.Str != in.Str || *got.StrPtr != *in.StrPtr || got.Int != in.Int || *got.IntPtr != *in.IntPtr {
		t.Fatalf("Clone() = %#v, want %#v", got, in)
	}
	if Clone[ObjectA](nil) != nil {
		t.Fatalf("Clone(nil) should be nil")
	}
}

type cloneAll struct {
	NilPtr   *int
	Ptr      *int
	NilSlice []string
	Slice    []string
	NilMap   map[string]int
	Map      map[string]int
	Arr      [2]int
	Nested   ObjectA
	NilAny   any
	Any      any
	priv     int
}

func TestClone_kinds(t *testing.T) {
	n := 7
	in := &cloneAll{
		Ptr:    &n,
		Slice:  []string{"a"},
		Map:    map[string]int{"k": 1},
		Arr:    [2]int{1, 2},
		Nested: sampleA(),
		Any:    sampleA(),
		priv:   9,
	}
	got := Clone(in)
	if got == in || got.Ptr == in.Ptr || got.Slice[0] != "a" || got.Map["k"] != 1 || got.Arr != in.Arr {
		t.Fatalf("Clone() = %#v", got)
	}
	if got.NilPtr != nil || got.NilSlice != nil || got.NilMap != nil || got.NilAny != nil {
		t.Fatalf("Clone() did not preserve nils: %#v", got)
	}
	if got.Nested.StrPtr == in.Nested.StrPtr || got.priv != 9 {
		t.Fatalf("Clone() nested/unexported = %#v", got)
	}
}
