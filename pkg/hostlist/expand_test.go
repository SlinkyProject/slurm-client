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

func TestExpandSlurmFixtures(t *testing.T) {
	data, err := os.ReadFile("testdata/expand.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Cases []struct {
			Expression string
			Hosts      []string
			Invalid    bool
		}
	}
	if err := json.Unmarshal(data, &fixture); err != nil {
		t.Fatal(err)
	}
	for _, tt := range fixture.Cases {
		t.Run(tt.Expression, func(t *testing.T) {
			got, err := hostlist.Expand(tt.Expression)
			if (err != nil) != tt.Invalid {
				t.Fatalf("Expand() error = %v, want invalid = %v", err, tt.Invalid)
			}
			if !slices.Equal(got, tt.Hosts) {
				t.Errorf("Expand() = %q, want %q", got, tt.Hosts)
			}
		})
	}
}

func TestExpandRejectsInvalidInput(t *testing.T) {
	for _, input := range []string{
		"node1]", "node[[1-2]]", "node[+1-2]", "node[ 1-2]", "node[1-]",
		"node[18446744073709551616]", "node[0-65536]", "node[0-18446744073709551615]",
		"rack[0-65535]node[0-65535]", "node[0-65535,0-65535,0-65535]gpu[0-9]",
		"node[" + strings.Repeat("0-65535,", 16) + "0]", "node\x00suffix",
		strings.Repeat("prefix", 256) + "[0-65535]",
		"valid,node[2-1]", "valid,node[1-2", "valid,node[1--2]",
	} {
		t.Run(input, func(t *testing.T) {
			got, err := hostlist.Expand(input)
			if err == nil || got != nil {
				t.Fatalf("Expand() returned %d hosts and error %v; want nil hosts and an error", len(got), err)
			}
		})
	}
}

func TestExpandRangeBoundary(t *testing.T) {
	hosts, err := hostlist.Expand("node[0-65535]")
	if err != nil {
		t.Fatal(err)
	}
	if len(hosts) != 65536 || hosts[0] != "node0" || hosts[len(hosts)-1] != "node65535" {
		t.Fatalf("unexpected expansion of maximum range: %d hosts", len(hosts))
	}
	// Slurm 26.05.3/4's scontrol crashes expanding a range ending at the
	// uint64 maximum. Check this boundary independently of the C reference.
	hosts, err = hostlist.Expand("node[18446744073709551614-18446744073709551615]")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(hosts, []string{"node18446744073709551614", "node18446744073709551615"}) {
		t.Fatalf("unexpected expansion at uint64 boundary: %q", hosts)
	}
}

func FuzzExpand(f *testing.F) {
	for _, seed := range []string{"", "dgx8-[0-47]", "node[00-10]", "rack[0-1]node[0-1]gpu[0-1]", "node[1,,2]"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 128 {
			t.Skip()
		}
		hosts, err := hostlist.Expand(input)
		if err != nil {
			if hosts != nil {
				t.Fatal("error returned with partial results")
			}
			return
		}
		// Expansion produces literal names. Expanding them again must preserve
		// every identity, including padding, order, and duplicate entries.
		again, err := hostlist.Expand(strings.Join(hosts, ","))
		if err != nil || !slices.Equal(hosts, again) {
			t.Fatalf("expansion was not idempotent: %v", err)
		}
	})
}

func ExampleExpand() {
	hosts, _ := hostlist.Expand("node[0-2],node[00-02]")
	for _, host := range hosts {
		fmt.Println(host)
	}
	// Output:
	// node0
	// node1
	// node2
	// node00
	// node01
	// node02
}
