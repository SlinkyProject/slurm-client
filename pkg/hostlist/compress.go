// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package hostlist

import (
	"fmt"
	"strings"
)

// Compress returns a compact Slurm host list without modifying hosts.
// It preserves order, duplicates, and padding, like scontrol show hostlist.
// Consecutive hosts with the same prefix are grouped by their trailing decimal
// numbers; other names remain literal. For example, node0,node1 becomes
// node[0-1], while rack0node0,rack1node1 remains unchanged.
//
// Each input must be one nonempty, unbracketed host name without host-list
// separators or NUL bytes. Numeric suffixes must fit in uint64. The host-count
// and name-byte limits are the same as Expand. Ranges are split at 65,536
// entries so Expand can read the result. Empty input produces an empty string.
// On error, no partial expression is returned.
func Compress(hosts []string) (string, error) {
	if len(hosts) > maxHosts {
		return "", fmt.Errorf("host list exceeds %d hosts", maxHosts)
	}
	var expressions []string
	remainingBytes := maxBytes
	for i := 0; i < len(hosts); {
		prefix, number, err := splitHost(hosts[i])
		if err != nil {
			return "", err
		}
		if number.text == "" {
			if len(prefix) > remainingBytes {
				return "", fmt.Errorf("host list exceeds %d bytes of host names", maxBytes)
			}
			expressions = append(expressions, prefix)
			remainingBytes -= len(prefix)
			i++
			continue
		}

		var numbers []hostNumber
		widest, bytes := 0, 0
		j := i
		for ; j < len(hosts); j++ {
			nextPrefix, next, err := splitHost(hosts[j])
			if err != nil {
				return "", err
			}
			if nextPrefix != prefix || next.text == "" {
				break
			}
			if len(hosts[j]) > remainingBytes-bytes {
				return "", fmt.Errorf("host list exceeds %d bytes of host names", maxBytes)
			}
			bytes += len(hosts[j])
			widest = max(widest, len(hosts[j]))
			numbers = append(numbers, next)
		}
		// Match Expand's conservative bound for a bracket group's name bytes.
		if widest > remainingBytes/len(numbers) {
			return "", fmt.Errorf("host list exceeds %d bytes of host names", maxBytes)
		}
		remainingBytes -= bytes
		if len(numbers) == 1 {
			expressions = append(expressions, hosts[i])
		} else {
			expressions = append(expressions, prefix+"["+compressNumbers(numbers)+"]")
		}
		i = j
	}
	return strings.Join(expressions, ","), nil
}

type hostNumber struct {
	text  string
	value uint64
}

func splitHost(host string) (string, hostNumber, error) {
	if host == "" || strings.ContainsAny(host, ", \t\n[]\x00") {
		return "", hostNumber{}, fmt.Errorf("invalid literal host name %q", host)
	}
	i := len(host)
	for i > 0 && host[i-1] >= '0' && host[i-1] <= '9' {
		i--
	}
	if i == len(host) {
		return host, hostNumber{}, nil
	}
	number, err := parseNumber(host[i:])
	if err != nil {
		return "", hostNumber{}, fmt.Errorf("invalid host name %q: %w", host, err)
	}
	return host[:i], hostNumber{text: host[i:], value: number}, nil
}

func compressNumbers(numbers []hostNumber) string {
	var ranges []string
	for i := 0; i < len(numbers); {
		j := i + 1
		for j < len(numbers) && j-i < maxRange {
			previous, next := numbers[j-1], numbers[j]
			if next.value <= previous.value || next.value-previous.value != 1 ||
				fmt.Sprintf("%0*d", len(numbers[i].text), next.value) != next.text {
				break
			}
			j++
		}
		part := numbers[i].text
		if j > i+1 {
			part += "-" + numbers[j-1].text
		}
		ranges = append(ranges, part)
		i = j
	}
	return strings.Join(ranges, ",")
}
