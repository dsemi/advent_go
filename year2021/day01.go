package year2021

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day01 struct{}

func (d *Day01) solve(input string, offset int) int {
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

func (d *Day01) Part1(input string) interface{} {
	return d.solve(input, 1)
}

func (d *Day01) Part2(input string) interface{} {
	return d.solve(input, 3)
}

func init() {
	problems.Register(&Day01{})
}
