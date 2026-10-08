// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"
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

func TestGetStateAsSet(t *testing.T) {
	if got := GetStateAsSet[string](nil); !set.New[string]().Equal(got) {
		t.Fatalf("GetStateAsSet(nil) = %v", got)
	}
	states := ptr.To([]string{"IDLE", "DRAIN"})
	if got := GetStateAsSet(states); !set.New("IDLE", "DRAIN").Equal(got) {
		t.Fatalf("GetStateAsSet() = %v", got)
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

func TestParseNodeName(t *testing.T) {
	tests := []struct {
		name     string
		nodeConf string
		want     string
		wantErr  bool
	}{
		{
			name:     "Valid node configuration",
			nodeConf: "NodeName=node-0 CPUs=4 State=EXTERNAL",
			want:     "node-0",
			wantErr:  false,
		},
		{
			name:     "Missing NodeName",
			nodeConf: "CPUs=4 State=EXTERNAL",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Empty string",
			nodeConf: "",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Lowercase nodename",
			nodeConf: "nodename=node-0 CPUs=4 State=EXTERNAL",
			want:     "node-0",
			wantErr:  false,
		},
		{
			name:     "Uppercase NODENAME",
			nodeConf: "NODENAME=compute-01 CPUs=8 RealMemory=16384",
			want:     "compute-01",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNodeName(tt.nodeConf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNodeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
