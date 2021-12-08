package year2017

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day05 struct{}

func (*Day05) calcSteps(input string, f func(int) int) int {
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

func (d *Day05) Part1(input string) interface{} {
	return d.calcSteps(input, func(x int) int { return x + 1 })
}

func (d *Day05) Part2(input string) interface{} {
	return d.calcSteps(input, func(x int) int {
		if x >= 3 {
			return x - 1
		} else {
			return x + 1
		}
	})
}

func init() {
	problems.Register(&Day05{})
}
