package main

import (
	"strings"
	"utils"
)

func process(input string, f func(int, int, int) int) int {
	var total int
	for _, line := range strings.Split(input, "\n") {
		v := strings.Split(line, "x")
		total += f(utils.Int(v[0]), utils.Int(v[1]), utils.Int(v[2]))
	}
	return total
}

func Part1(input string) interface{} {
	return process(input, func(l int, w int, h int) int {
		return 2*l*w + 2*l*h + 2*w*h + utils.Minimum([]int{l * w, l * h, w * h})
	})
}

func Part2(input string) interface{} {
	return process(input, func(l int, w int, h int) int {
		return l*w*h + 2*utils.Minimum([]int{l + w, l + h, w + h})
	})
}
