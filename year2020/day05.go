package main

import (
	"sort"
	"strings"
	"utils"
)

func seatIds(s string) []int {
	var ids []int
	for _, line := range strings.Split(s, "\n") {
		var n int
		for _, c := range line {
			n <<= 1
			if c == 'B' || c == 'R' {
				n++
			}
		}
		ids = append(ids, n)
	}
	return ids
}

func Part1(input string) interface{} {
	var max int
	for _, id := range seatIds(input) {
		max = utils.Max(max, id)
	}
	return max
}

func Part2(input string) interface{} {
	ids := seatIds(input)
	sort.Ints(ids)
	for i := range ids {
		if ids[i]+2 == ids[i+1] {
			return ids[i] + 1
		}
	}
	return nil
}
