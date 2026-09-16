// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

// Package hostlist expands Slurm's bracketed numeric host lists.
package hostlist

import (
	"fmt"
	"strconv"
	"strings"
)

// Slurm limits each numeric range to 64K entries. Also bound the total
// materialized result, since multiple bracket groups multiply its size.
const (
	maxRange = 64 * 1024
	maxHosts = 1024 * 1024
	maxBytes = 64 * 1024 * 1024
)

// Expand returns the hosts in expression, preserving order and duplicates.
// Hosts are separated by commas, spaces, tabs, or newlines. Bracket groups
// contain decimal numbers or inclusive ranges, for example node[0-47] or
// rack[0-1]_node[00-03]. Bare groups such as [0-7] are also supported.
//
// Padding comes from the lower endpoint: node[0-10] starts with node0, whereas
// node[00-10] starts with node00. With multiple groups, the last group varies
// fastest, followed by the preceding groups from left to right, as in Slurm.
//
// The reference is Slurm 26.05's scontrol show hostnames (ordinary decimal host
// lists, not architecture-specific multidimensional box expressions). Unlike
// Slurm's permissive C number conversion, endpoints must contain only digits
// and fit in uint64. Malformed brackets, ranges exceeding 65,536 entries, and
// results exceeding 1,048,576 hosts or a conservative 64 MiB estimate of name
// bytes are rejected, with no partial result.
//
// Reference: https://github.com/SchedMD/slurm/blob/slurm-26-05-4-1/src/common/hostlist.c
func Expand(expression string) ([]string, error) {
	hosts := []string{}
	remainingBytes := maxBytes
	start, inGroup := 0, false
	for i := 0; i <= len(expression); i++ {
		if i < len(expression) {
			switch expression[i] {
			case '[':
				if inGroup {
					return nil, fmt.Errorf("invalid hostlist %q: nested bracket", expression)
				}
				inGroup = true
			case ']':
				if !inGroup {
					return nil, fmt.Errorf("invalid hostlist %q: unmatched bracket", expression)
				}
				inGroup = false
			case 0:
				return nil, fmt.Errorf("invalid hostlist %q: NUL byte", expression)
			}
			if inGroup || !strings.ContainsRune(", \t\n", rune(expression[i])) {
				continue
			}
		}
		if inGroup {
			return nil, fmt.Errorf("invalid hostlist %q: unclosed bracket", expression)
		}
		if start < i {
			expanded, err := expandToken(expression[start:i], maxHosts-len(hosts), remainingBytes)
			if err != nil {
				return nil, fmt.Errorf("invalid hostlist %q: %w", expression, err)
			}
			hosts = append(hosts, expanded...)
			for _, host := range expanded {
				remainingBytes -= len(host)
			}
		}
		start = i + 1
	}
	return hosts, nil
}

type numberRange struct {
	first uint64
	count uint64
	width int
}

func expandToken(token string, limit, byteLimit int) ([]string, error) {
	var literals []string
	var groups [][]numberRange
	count := 1
	nameBytes := 0
	for {
		open := strings.IndexByte(token, '[')
		if open < 0 {
			literals = append(literals, token)
			nameBytes += len(token)
			break
		}
		close := strings.IndexByte(token[open:], ']') + open
		group, size, err := parseGroup(token[open+1 : close])
		if err != nil {
			return nil, err
		}
		if size > limit/count {
			return nil, fmt.Errorf("expansion exceeds %d hosts", maxHosts)
		}
		count *= size
		width := 0
		for _, r := range group {
			last := strconv.FormatUint(r.first+(r.count-1), 10)
			width = max(width, r.width, len(last))
		}
		nameBytes += open + width
		literals = append(literals, token[:open])
		groups = append(groups, group)
		token = token[close+1:]
	}
	if count > limit {
		return nil, fmt.Errorf("expansion exceeds %d hosts", maxHosts)
	}
	if nameBytes > byteLimit/count {
		return nil, fmt.Errorf("expansion exceeds %d bytes of host names", maxBytes)
	}
	if len(groups) == 0 {
		return literals, nil
	}

	// Slurm enumerates each prefix group before the prefixes to its left,
	// then enumerates the final group within each fully expanded prefix.
	prefixes := []string{""}
	for i, group := range groups {
		next := []string{}
		if i == len(groups)-1 {
			next = make([]string, 0, count)
			for _, prefix := range prefixes {
				for _, r := range group {
					for offset := range r.count {
						number := fmt.Sprintf("%0*d", r.width, r.first+offset)
						next = append(next, prefix+literals[i]+number+literals[i+1])
					}
				}
			}
		} else {
			for _, r := range group {
				for offset := range r.count {
					number := fmt.Sprintf("%0*d", r.width, r.first+offset)
					for _, prefix := range prefixes {
						next = append(next, prefix+literals[i]+number)
					}
				}
			}
		}
		prefixes = next
	}
	return prefixes, nil
}

func parseGroup(group string) ([]numberRange, int, error) {
	var ranges []numberRange
	size := 0
	for part := range strings.SplitSeq(group, ",") {
		lower, upper, isRange := strings.Cut(part, "-")
		first, err := parseNumber(lower)
		if err != nil {
			return nil, 0, err
		}
		last := first
		if isRange {
			last, err = parseNumber(upper)
			if err != nil {
				return nil, 0, err
			}
		}
		if last < first {
			return nil, 0, fmt.Errorf("invalid or oversized range %q", part)
		}
		distance := last - first
		if distance >= maxRange {
			return nil, 0, fmt.Errorf("invalid or oversized range %q", part)
		}
		count := int(distance) + 1
		if count > maxHosts-size {
			return nil, 0, fmt.Errorf("expansion exceeds %d hosts", maxHosts)
		}
		size += count
		ranges = append(ranges, numberRange{first: first, count: distance + 1, width: len(lower)})
	}
	return ranges, size, nil
}

func parseNumber(value string) (uint64, error) {
	for i := range len(value) {
		if value[i] < '0' || value[i] > '9' {
			return 0, fmt.Errorf("invalid number %q", value)
		}
	}
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", value, err)
	}
	return number, nil
}
