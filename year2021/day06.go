package main

import (
	"strings"
	"utils"
)

func solve(s string, n int) uint64 {
	fish := make([]uint64, 9)
	for _, c := range strings.Split(s, ",") {
		fish[utils.Int(c)]++
	}
	for i := 0; i < n; i++ {
		fish[(i+7)%9] += fish[i%9]
	}
	return utils.Sum(fish)
}

func Part1(input string) interface{} {
	return solve(input, 80)
}

func Part2(input string) interface{} {
	return solve(input, 256)
}
