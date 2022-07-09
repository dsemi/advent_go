package main

import (
	"strings"
	"utils"
)

func parse(input string) []int64 {
	ns := make([]int64, 0)
	for _, line := range strings.Split(input, "\n") {
		ns = append(ns, utils.Int64(line))
	}
	return ns
}

func findFirstInvalid(ns []int64) int64 {
	n := 25
outer:
	for {
		for i := n - 25; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if ns[i]+ns[j] == ns[n] {
					n++
					continue outer
				}
			}
		}
		return ns[n]
	}
}

func Part1(input string) interface{} {
	return findFirstInvalid(parse(input))
}

func Part2(input string) interface{} {
	ns := parse(input)
	n := findFirstInvalid(ns)
	var lo, hi, acc int64
	for acc != n {
		if acc < n {
			acc += ns[hi]
			hi++
		} else {
			acc -= ns[lo]
			lo++
		}
	}
	arr := ns[lo:hi]
	low, high := utils.Extrema(arr)
	return low + high
}
