// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/SlinkyProject/slurm-client/pkg/hostlist"
)

var _ = Describe("Hostlist compression", Label("hostlist"), func() {
	It("matches scontrol and preserves every host through both expanders", func() {
		data, err := os.ReadFile("../hostlist/testdata/compress.json")
		Expect(err).NotTo(HaveOccurred())
		var fixture struct {
			Cases []struct {
				Hosts      []string
				Expression string
			}
		}
		Expect(json.Unmarshal(data, &fixture)).To(Succeed())
		var cases [][]string
		for _, tt := range fixture.Cases {
			cases = append(cases, tt.Hosts)
		}
		cases = append(cases, generatedHosts()...)
		inputs := make([]string, len(cases))
		for i, hosts := range cases {
			inputs[i] = strings.Join(hosts, ",")
		}
		outputs := scontrolHostlists("hostlist", inputs)
		expressions := make([]string, len(cases))
		for i, hosts := range cases {
			original := slices.Clone(hosts)
			expression, err := hostlist.Compress(hosts)
			Expect(err).NotTo(HaveOccurred(), "hosts: %q", original)
			Expect(hosts).To(Equal(original))
			Expect(expression).To(Equal(strings.TrimSuffix(outputs[i], "\n")), "hosts: %q", original)
			if i < len(fixture.Cases) {
				Expect(expression).To(Equal(fixture.Cases[i].Expression))
			}
			expanded, err := hostlist.Expand(expression)
			Expect(err).NotTo(HaveOccurred())
			Expect(hostnamesOutput(expanded)).To(Equal(hostnamesOutput(original)), "expression: %q", expression)
			expressions[i] = expression
		}
		expanded := scontrolHostlists("hostnames", expressions)
		for i, hosts := range cases {
			Expect(expanded[i]).To(Equal(hostnamesOutput(hosts)), "expression: %q", expressions[i])
		}
	})
})

func generatedHosts() [][]string {
	var cases [][]string
	for first := 0; first < 16; first++ {
		for count := 0; count <= 8; count++ {
			for width := 0; width < 4; width++ {
				hosts := []string{}
				for n := first; n < first+count; n++ {
					hosts = append(hosts, fmt.Sprintf("node%0*d", width, n))
				}
				cases = append(cases, hosts)
			}
		}
	}
	// Stop below uint64's maximum: Slurm 26.05.3/4 cannot reliably round-trip
	// that endpoint. TestCompressRangeBoundary covers it independently.
	for _, first := range []uint64{8, 9, 98, 99, 998, 999, 1<<31 - 1, 1<<32 - 1, 1<<64 - 8} {
		for _, count := range []uint64{1, 2, 4} {
			for _, width := range []int{0, 2, 3, 20} {
				var hosts []string
				for n := first; n < first+count; n++ {
					hosts = append(hosts, fmt.Sprintf("node%0*d", width, n))
				}
				cases = append(cases, hosts)
			}
		}
	}
	rng := rand.New(rand.NewPCG(73, 29)) //nolint:gosec // A fixed seed makes differential failures reproducible.
	for range 512 {
		hosts := []string{}
		for range rng.IntN(65) {
			prefix := []string{"node", "rack0node", "other", ""}[rng.IntN(4)]
			host := fmt.Sprintf("%s%0*d", prefix, rng.IntN(6), rng.IntN(1024))
			if rng.IntN(5) == 0 {
				host += ".example"
			}
			hosts = append(hosts, host)
			if rng.IntN(4) == 0 {
				hosts = append(hosts, host)
			}
		}
		rng.Shuffle(len(hosts), func(i, j int) { hosts[i], hosts[j] = hosts[j], hosts[i] })
		cases = append(cases, hosts)
	}
	return cases
}
