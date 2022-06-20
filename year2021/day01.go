package main

import (
	"strings"
	"utils"
)

func solve(input string, offset int) int {
	ns := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		ns = append(ns, utils.Int(line))
	}
	var count int
	for i := 0; i+offset < len(ns); i++ {
		if ns[i] < ns[i+offset] {
			count++
		}
	}
	return count
}

func Part1(input string) interface{} {
	return solve(input, 1)
}

func Part2(input string) interface{} {
	return solve(input, 3)
}
