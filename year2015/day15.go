package year2015

import (
	"advent/problems"
	"advent/utils"
	"regexp"
	"strings"
)

type Day15 struct{}

func (*Day15) parseIngredients(input string) [][]int {
	reg := regexp.MustCompile("-?\\d+")
	var ings [][]int
	for _, line := range strings.Split(input, "\n") {
		var ns []int
		for _, x := range reg.FindAllString(line, -1) {
			ns = append(ns, utils.Int(x))
		}
		ings = append(ings, ns)
	}
	return ings
}

func (*Day15) maxScore(total int, calPred func(int) bool, ings [][]int) int {
	var max int
	utils.Partitions(len(ings), total, func(ms []int) {
		v := make([]int, 5)
		for i := 0; i < len(ms); i++ {
			for j := 0; j < len(v); j++ {
				v[j] += ms[i] * ings[i][j]
			}
		}
		if calPred(v[4]) {
			prod := 1
			for _, x := range v[:4] {
				prod *= utils.Max(0, x)
			}
			max = utils.Max(max, prod)
		}
	})
	return max
}

func (d *Day15) Part1(input string) interface{} {
	return d.maxScore(100, func(x int) bool { return true }, d.parseIngredients(input))
}

func (d *Day15) Part2(input string) interface{} {
	return d.maxScore(100, func(x int) bool { return x == 500 }, d.parseIngredients(input))
}

func init() {
	problems.Register(&Day15{})
}
