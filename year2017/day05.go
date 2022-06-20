package main

import (
	"strings"
	"utils"
)

func calcSteps(input string, f func(int) int) int {
	var ns []int
	for _, line := range strings.Split(input, "\n") {
		ns = append(ns, utils.Int(line))
	}
	var i, res int
	for i >= 0 && i < len(ns) {
		val := ns[i]
		ns[i] = f(val)
		i += val
		res += 1
	}
	return res
}

func Part1(input string) interface{} {
	return calcSteps(input, func(x int) int { return x + 1 })
}

func Part2(input string) interface{} {
	return calcSteps(input, func(x int) int {
		if x >= 3 {
			return x - 1
		} else {
			return x + 1
		}
	})
}
