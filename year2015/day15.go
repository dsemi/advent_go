package year2015

import (
	"advent/types"
	"advent/utils"
	"regexp"
	"strings"
)

type Day15 struct{}

func parseIngredients(input string) [][]int {
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

func maxScore(total int, calPred func(int) bool, ings [][]int) int {
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

func (Day15) Part1(input string) interface{} {
	return maxScore(100, func(x int) bool { return true }, parseIngredients(input))
}

func (Day15) Part2(input string) interface{} {
	return maxScore(100, func(x int) bool { return x == 500 }, parseIngredients(input))
}

func init() {
	types.Register(Probs, Day15{})
}
