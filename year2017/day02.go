package year2017

import (
	"advent/types"
	"advent/utils"
	"math"
	"strings"
)

type Day02 struct{}

func (Day02) Part1(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		min := math.MaxInt
		max := math.MinInt
		for _, x := range strings.Fields(line) {
			n := utils.Int(x)
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
		sum += max - min
	}
	return sum
}

func (Day02) Part2(input string) interface{} {
	var sum int
OUTER:
	for _, line := range strings.Split(input, "\n") {
		var ns []int
		for _, x := range strings.Fields(line) {
			ns = append(ns, utils.Int(x))
		}
		for i := 0; i < len(ns); i++ {
			for j := 0; j < len(ns); j++ {
				if ns[i] != ns[j] && ns[i]%ns[j] == 0 {
					sum += ns[i] / ns[j]
					continue OUTER
				}
			}
		}
	}
	return sum
}

func init() {
	types.Register(Probs, Day02{})
}
