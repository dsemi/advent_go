package year2018

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day01 struct{}

func (*Day01) Part1(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sum += utils.Int(line)
	}
	return sum
}

func (*Day01) Part2(input string) interface{} {
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

func init() {
	problems.Register(&Day01{})
}
