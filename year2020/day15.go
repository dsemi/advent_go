package main

import (
	"strings"
	"utils"
)

func solve(n uint32, input string) uint32 {
	m := make([]uint32, n)
	j := uint32(1)
	for _, v := range strings.Split(input, ",") {
		m[utils.Uint32(v)] = j
		j++
	}
	var result uint32
	for i := j; i < n; i++ {
		m[result], result = i, i-m[result]
		if result == i {
			result = 0
		}
	}
	return result
}

func Part1(input string) interface{} {
	return solve(2020, input)
}

func Part2(input string) interface{} {
	return solve(30_000_000, input)
}
