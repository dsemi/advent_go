package year2020

import (
	"advent/utils"
	"strings"
)

func sumCounts(input string, f func(int, int) int) int {
	var total int
	toInt := func(s string) int {
		var n int
		for _, c := range s {
			n |= 1 << int(c-'a')
		}
		return n
	}
	for _, line := range strings.Split(input, "\n\n") {
		fs := strings.Fields(line)
		num := toInt(fs[0])
		for _, x := range fs[1:] {
			num = f(num, toInt(x))
		}
		total += utils.CountOnes(num)
	}
	return total
}

func Day06Part1(input string) interface{} {
	return sumCounts(input, func(a, b int) int {
		return a | b
	})
}

func Day06Part2(input string) interface{} {
	return sumCounts(input, func(a, b int) int {
		return a & b
	})
}
