package main

import (
	"sort"
	"strings"
	"utils"
)

type IpRanges struct {
	ips [][2]uint64
}

func parse(input string) *IpRanges {
	ips := make([][2]uint64, 0)
	for _, line := range strings.Split(input, "\n") {
		ns := strings.Split(line, "-")
		ips = append(ips, [2]uint64{utils.Uint64(ns[0]), utils.Uint64(ns[1])})
	}
	sort.Slice(ips, func(a, b int) bool {
		if ips[a][0] == ips[b][0] {
			return ips[a][1] < ips[b][1]
		}
		return ips[a][0] < ips[b][0]
	})
	return &IpRanges{ips}
}

func (r *IpRanges) Next() ([2]uint64, bool) {
	if len(r.ips) == 0 {
		return [2]uint64{}, false
	}
	curr := r.ips[0]
	r.ips = r.ips[1:]
	for len(r.ips) > 0 && r.ips[0][0] <= curr[1]+1 {
		curr[1] = utils.Max(curr[1], r.ips[0][1])
		r.ips = r.ips[1:]
	}
	return curr, true
}

func Part1(input string) interface{} {
	span, _ := parse(input).Next()
	if span[0] > 0 {
		return 0
	}
	return span[1] + 1
}

func Part2(input string) interface{} {
	sum := utils.Pow[uint64](2, 32)
	ranges := parse(input)
	for span, ok := ranges.Next(); ok; span, ok = ranges.Next() {
		sum -= span[1] - span[0] + 1
	}
	return sum
}
