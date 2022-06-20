package year2021

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day06 struct{}

func (d *Day06) solve(s string, n int) uint64 {
	fish := make([]uint64, 9)
	for _, c := range strings.Split(s, ",") {
		fish[utils.Int(c)]++
	}
	for i := 0; i < n; i++ {
		fish[(i+7)%9] += fish[i%9]
	}
	return utils.Sum(fish)
}

func (d *Day06) Part1(input string) interface{} {
	return d.solve(input, 80)
}

func (d *Day06) Part2(input string) interface{} {
	return d.solve(input, 256)
}

func init() {
	problems.Register(&Day06{})
}
