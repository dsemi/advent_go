package main

import (
	"strings"
)

func parse(input string) [][]uint32 {
	var ns [][]uint32
	for _, line := range strings.Split(input, "\n") {
		var n []uint32
		for _, c := range line {
			n = append(n, uint32(c-'0'))
		}
		ns = append(ns, n)
	}
	return ns
}

func Part1(input string) interface{} {
	var gamma uint32
	ns := parse(input)
	for i := 0; i < len(ns[0]); i++ {
		var ones int
		for _, n := range ns {
			if n[i] == 1 {
				ones++
			}
		}
		gamma <<= 1
		if ones >= (len(ns)+1)/2 {
			gamma |= 1
		}
	}
	return gamma * ((1<<len(ns[0]) - 1) ^ gamma)
}

func mostMatched(ns [][]uint32, pred func(int, int) bool) uint32 {
	for i := 0; i < len(ns[0]); i++ {
		var ones, zeros [][]uint32
		for _, n := range ns {
			if n[i] == 1 {
				ones = append(ones, n)
			} else {
				zeros = append(zeros, n)
			}
		}
		if pred(len(ones), len(zeros)) {
			ns = ones
		} else {
			ns = zeros
		}
	}
	var x uint32
	for _, n := range ns[0] {
		x = x<<1 | n
	}
	return x
}

func Part2(input string) interface{} {
	ns := parse(input)
	return mostMatched(ns, func(a, b int) bool { return a >= b }) *
		mostMatched(ns, func(a, b int) bool { return a < b && a != 0 || b == 0 })
}
