package year2020

import (
	"advent/problems"
	"strconv"
	"strings"
)

type Day01 struct{}

func (*Day01) parse(input string) []int {
	var ns []int
	for _, line := range strings.Fields(input) {
		n, _ := strconv.Atoi(line)
		ns = append(ns, n)
	}
	return ns
}

func (d *Day01) Part1(input string) interface{} {
	ns := d.parse(input)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			if ns[i]+ns[j] == 2020 {
				return ns[i] * ns[j]
			}
		}
	}
	return -1
}

func (d *Day01) Part2(input string) interface{} {
	ns := d.parse(input)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			for k := j + 1; k < len(ns); k++ {
				if ns[i]+ns[j]+ns[k] == 2020 {
					return ns[i] * ns[j] * ns[k]
				}
			}
		}
	}
	return -1
}

func init() {
	problems.Register(&Day01{})
}
