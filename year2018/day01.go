package year2018

import (
	"advent/utils"
	"strings"
)

func Day01Part1(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sum += utils.Int(line)
	}
	return sum
}

func Day01Part2(input string) interface{} {
	var sum int
	m := make(map[int]bool)
	var ns []int
	for _, line := range strings.Split(input, "\n") {
		n := utils.Int(line)
		ns = append(ns, n)
		sum += n
		if m[sum] {
			return sum
		}
		m[sum] = true
	}
	var i int
	for {
		sum += ns[i]
		if m[sum] {
			return sum
		}
		m[sum] = true
		i = (i + 1) % len(ns)
	}
}