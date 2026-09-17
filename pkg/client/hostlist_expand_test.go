// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/SlinkyProject/slurm-client/pkg/hostlist"
)

var _ = Describe("Hostlist expansion", Label("hostlist"), func() {
	It("matches scontrol show hostnames for the reference cases", func() {
		data, err := os.ReadFile("../hostlist/testdata/expand.json")
		Expect(err).NotTo(HaveOccurred())
		var fixture struct {
			Cases []struct {
				Expression string
				Hosts      []string
				Invalid    bool
			}
		}
		Expect(json.Unmarshal(data, &fixture)).To(Succeed())
		for _, tt := range fixture.Cases {
			By(fmt.Sprintf("expanding %q", tt.Expression))
			stdout, stderr := slurmHostnames(tt.Expression)
			hosts, err := hostlist.Expand(tt.Expression)
			if tt.Invalid {
				// scontrol can exit successfully while reporting invalid input
				// on stderr, so its exit status alone is insufficient.
				Expect(stderr).To(ContainSubstring("Invalid hostlist"))
				Expect(err).To(HaveOccurred())
				Expect(hosts).To(BeNil())
				continue
			}
			Expect(stderr).To(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
			Expect(stdout).To(Equal(hostnamesOutput(tt.Hosts)))
			Expect(stdout).To(Equal(hostnamesOutput(hosts)))
		}
	})

	It("matches scontrol across padding widths and multiple bracket groups", func() {
		var expressions []string
		for _, first := range []int{0, 1, 8, 9, 98, 99} {
			for lowerWidth := 1; lowerWidth <= 3; lowerWidth++ {
				for upperWidth := 1; upperWidth <= 3; upperWidth++ {
					expressions = append(expressions,
						fmt.Sprintf("node[%0*d-%0*d]", lowerWidth, first, upperWidth, first+3))
				}
			}
		}
		expressions = append(expressions,
			"r[1,0]n[02,0-1]g[1,0]", "a[0-1]b[0-1]c[0-1]d[0-1]",
			"node[0,00,1,01]", "[0-15]", "r[0-1]n[0-1].example,other[00-02]",
		)
		for _, expression := range expressions {
			By(fmt.Sprintf("expanding %q", expression))
			stdout, stderr := slurmHostnames(expression)
			Expect(stderr).To(BeEmpty())
			hosts, err := hostlist.Expand(expression)
			Expect(err).NotTo(HaveOccurred())
			Expect(hostnamesOutput(hosts)).To(Equal(stdout))
		}
	})

	It("matches scontrol for generated numeric ranges and mixed host lists", func() {
		expressions := generatedHostlists()
		outputs := scontrolHostlists("hostnames", expressions)
		for i, expression := range expressions {
			hosts, err := hostlist.Expand(expression)
			Expect(err).NotTo(HaveOccurred(), "expression: %q", expression)
			Expect(hostnamesOutput(hosts)).To(Equal(outputs[i]), "expression: %q", expression)
		}
	})
})

// Enumerate small ranges and digit boundaries, then add repeatable combinations
// of multiple groups, mixed padding, duplicate terms, and separators. Compare
// the complete ordered output, rather than just the count or the set of names.
func generatedHostlists() []string {
	var expressions []string
	for first := 0; first < 16; first++ {
		for span := 0; span < 4; span++ {
			for lowerWidth := 0; lowerWidth < 4; lowerWidth++ {
				for upperWidth := 0; upperWidth < 4; upperWidth++ {
					expressions = append(expressions,
						fmt.Sprintf("node[%0*d-%0*d]", lowerWidth, first, upperWidth, first+span))
				}
			}
		}
	}
	// Slurm 26.05.3/4 crashes for ranges ending at uint64's maximum.
	// Keep live comparisons below that endpoint; TestExpandRangeBoundary
	// independently checks the Go implementation at the maximum itself.
	for _, first := range []uint64{8, 9, 10, 98, 99, 100, 998, 999, 1000, 1<<31 - 1, 1<<32 - 1, 1<<64 - 8} {
		for _, span := range []uint64{0, 1, 3} {
			for _, lowerWidth := range []int{0, 2, 3, 20} {
				for _, upperWidth := range []int{0, 2, 3, 20} {
					expressions = append(expressions,
						fmt.Sprintf("node[%0*d-%0*d]", lowerWidth, first, upperWidth, first+span))
				}
			}
		}
	}
	rng := rand.New(rand.NewPCG(42, 17)) //nolint:gosec // A fixed seed makes differential failures reproducible.
	for range 512 {
		var terms []string
		for range 1 + rng.IntN(3) {
			var term strings.Builder
			for group := range 1 + rng.IntN(4) {
				fmt.Fprintf(&term, "r%d[", group)
				for part := range 1 + rng.IntN(2) {
					if part != 0 {
						term.WriteByte(',')
					}
					first := rng.IntN(1000)
					fmt.Fprintf(&term, "%0*d", rng.IntN(5), first)
					if rng.IntN(2) == 0 {
						fmt.Fprintf(&term, "-%0*d", rng.IntN(5), first+rng.IntN(2))
					}
				}
				term.WriteByte(']')
			}
			if rng.IntN(3) == 0 {
				term.WriteString(".example")
			}
			terms = append(terms, term.String())
		}
		if rng.IntN(3) == 0 {
			terms = append(terms, terms[0])
		}
		separator := []string{",", " ", "\t", "\n"}[rng.IntN(4)]
		expressions = append(expressions, strings.Join(terms, separator))
	}
	return expressions
}

// Batch container execs to keep thousands of comparisons practical in CI.
// Expressions are positional arguments to a fixed script, never shell code.
// NUL-delimited results preserve individual outputs, including empty lists.
func scontrolHostlists(command string, expressions []string) []string {
	GinkgoHelper()
	container, err := compose.ServiceContainer(ctx, "control-plane")
	Expect(err).NotTo(HaveOccurred())
	var outputs []string
	for start := 0; start < len(expressions); start += 64 {
		end := min(start+64, len(expressions))
		By(fmt.Sprintf("comparing scontrol %s cases %d-%d of %d", command, start+1, end, len(expressions)))
		const script = `ulimit -c 0
command=$1; shift
for expression do
    scontrol show "$command" "$expression" || {
        result=$?
        printf 'scontrol %s failed for %s\n' "$command" "$expression" >&2
        exit "$result"
    }
    printf '\000'
done`
		args := []string{"sh", "-c", script, "hostlist-test", command}
		args = append(args, expressions[start:end]...)
		rc, reader, err := container.Exec(ctx, args)
		Expect(err).NotTo(HaveOccurred())
		stdout, stderr := demultiplexReader(reader)
		Expect(rc).To(Equal(0), stderr)
		Expect(stderr).To(BeEmpty(), "expressions: %q", expressions[start:end])
		batch := strings.Split(stdout, "\x00")
		Expect(batch).To(HaveLen(end - start + 1))
		Expect(batch[len(batch)-1]).To(BeEmpty())
		outputs = append(outputs, batch[:len(batch)-1]...)
	}
	return outputs
}

func slurmHostnames(expression string) (string, string) {
	GinkgoHelper()
	container, err := compose.ServiceContainer(ctx, "control-plane")
	Expect(err).NotTo(HaveOccurred())
	rc, reader, err := container.Exec(ctx, []string{"scontrol", "show", "hostnames", expression})
	Expect(err).NotTo(HaveOccurred())
	stdout, stderr := demultiplexReader(reader)
	Expect(rc).To(Equal(0), stderr)
	return stdout, stderr
}

func hostnamesOutput(hosts []string) string {
	if len(hosts) == 0 {
		return ""
	}
	return strings.Join(hosts, "\n") + "\n"
}
