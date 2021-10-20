package year2015

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day17 struct{}

func (Day17) allCombos(input string) chan [][]int {
	var xs []int
	for _, line := range strings.Split(input, "\n") {
		xs = append(xs, utils.Int(line))
	}
	c := make(chan [][]int)
	go func() {
		defer close(c)
		for n := 1; n <= len(xs); n++ {
			var combos [][]int
			utils.Combinations(xs, n, func(combo []int) {
				if utils.Sum(combo) == 150 {
					combos = append(combos, combo)
				}
			})
			c <- combos
		}
	}()
	return c
}

func (d Day17) Part1(input string) interface{} {
	var sum int
	for v := range d.allCombos(input) {
		sum += len(v)
	}
	return sum
}

func (d Day17) Part2(input string) interface{} {
	for v := range d.allCombos(input) {
		if len(v) > 0 {
			return len(v)
		}
	}
	return -1
}

func init() {
	problems.Register(Day17{})
}
