// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package hostlist_test

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/SlinkyProject/slurm-client/pkg/hostlist"
)

func TestCompressSlurmFixtures(t *testing.T) {
	data, err := os.ReadFile("testdata/compress.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Cases []struct {
			Hosts      []string
			Expression string
		}
	}
	if err := json.Unmarshal(data, &fixture); err != nil {
		t.Fatal(err)
	}
	for _, tt := range fixture.Cases {
		t.Run(tt.Expression, func(t *testing.T) {
			original := slices.Clone(tt.Hosts)
			expression, err := hostlist.Compress(tt.Hosts)
			if err != nil || expression != tt.Expression {
				t.Fatalf("Compress() = %q, %v; want %q", expression, err, tt.Expression)
			}
			if !slices.Equal(tt.Hosts, original) {
				t.Fatal("Compress modified its input")
			}
			hosts, err := hostlist.Expand(expression)
			if err != nil || !slices.Equal(hosts, original) {
				t.Fatalf("round trip = %q, %v; want %q", hosts, err, original)
			}
		})
	}
}

func TestCompressRejectsInvalidHosts(t *testing.T) {
	for _, host := range []string{"", "node[0-1]", "node]", "node1,node2", "node1 node2", "node\t", "node\n", "node\x00", "node18446744073709551616"} {
		t.Run(host, func(t *testing.T) {
			expression, err := hostlist.Compress([]string{"valid1", host})
			if err == nil || expression != "" {
				t.Fatalf("Compress() = %q, %v; want empty expression and an error", expression, err)
			}
		})
	}
}

func TestCompressLimits(t *testing.T) {
	for name, hosts := range map[string][]string{
		"host count":    make([]string, 1024*1024+1),
		"literal bytes": slices.Repeat([]string{strings.Repeat("x", 1024*1024)}, 65),
		"numeric bytes": slices.Repeat([]string{strings.Repeat("x", 1024*1024) + "1"}, 65),
		"expansion estimate": append([]string{"node" + strings.Repeat("0", 1024) + "1"},
			slices.Repeat([]string{"node1"}, 65535)...),
	} {
		t.Run(name, func(t *testing.T) {
			expression, err := hostlist.Compress(hosts)
			if err == nil || expression != "" {
				t.Fatalf("Compress() returned %d bytes and %v; want empty expression and an error", len(expression), err)
			}
		})
	}
}

func TestCompressRangeBoundary(t *testing.T) {
	hosts := make([]string, 65537)
	for i := range hosts {
		hosts[i] = fmt.Sprintf("node%d", i)
	}
	expression, err := hostlist.Compress(hosts)
	if err != nil || expression != "node[0-65535,65536]" {
		t.Fatalf("Compress() = %q, %v", expression, err)
	}
	expanded, err := hostlist.Expand(expression)
	if err != nil || !slices.Equal(expanded, hosts) {
		t.Fatalf("round trip across range limit failed: %v", err)
	}
	hosts = []string{"node18446744073709551614", "node18446744073709551615"}
	expression, err = hostlist.Compress(hosts)
	if err != nil || expression != "node[18446744073709551614-18446744073709551615]" {
		t.Fatalf("Compress() = %q, %v", expression, err)
	}
	expanded, err = hostlist.Expand(expression)
	if err != nil || !slices.Equal(expanded, hosts) {
		t.Fatalf("round trip at uint64 maximum failed: %v", err)
	}
	// Slurm 26.05.4 incorrectly joins these across unsigned wraparound and
	// drops node0. Preserve both names; do not reproduce that overflow bug.
	hosts = []string{"node18446744073709551615", "node0"}
	expression, err = hostlist.Compress(hosts)
	if err != nil || expression != "node[18446744073709551615,0]" {
		t.Fatalf("Compress() = %q, %v", expression, err)
	}
	expanded, err = hostlist.Expand(expression)
	if err != nil || !slices.Equal(expanded, hosts) {
		t.Fatalf("round trip across uint64 boundary failed: %v", err)
	}
}

func FuzzCompressRoundTrip(f *testing.F) {
	for _, input := range []string{"node0,node1", "node0,node00,node1,node01", "node2,node1,node2", "rack0node0,rack1node1", "", "a,b", "7,8,9,10"} {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		hosts := strings.Split(input, ",")
		original := slices.Clone(hosts)
		expression, err := hostlist.Compress(hosts)
		if !slices.Equal(hosts, original) {
			t.Fatal("Compress modified its input")
		}
		if err != nil {
			if expression != "" {
				t.Fatal("error returned with partial expression")
			}
			return
		}
		expanded, err := hostlist.Expand(expression)
		if err != nil || !slices.Equal(expanded, original) {
			t.Fatalf("round trip through %q = %q, %v; want %q", expression, expanded, err, original)
		}
	})
}

func ExampleCompress() {
	expression, _ := hostlist.Compress([]string{"node2", "node1", "node2"})
	fmt.Println(expression)
	// Output: node[2,1-2]
}
