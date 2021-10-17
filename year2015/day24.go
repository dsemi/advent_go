package year2015

import (
	"advent/types"
	"advent/utils"
	"strings"
)

type Day24 struct{}

func quantumEntanglement(n int64, s string) int64 {
	var wts []int64
	for _, line := range strings.Split(s, "\n") {
		wts = append(wts, utils.Int64(line))
	}
	groupSize := utils.Sum64(wts) / n
	i := 1
	for {
		var m *int64
		utils.Combinations64(wts, i, func(combo []int64) {
			if utils.Sum64(combo) == groupSize {
				p := utils.Product64(combo)
				if m == nil {
					m = &p
				} else {
					*m = utils.Min64(*m, p)
				}
			}
		})
		if m != nil {
			return *m
		}
		i++
	}
}

func (Day24) Part1(input string) interface{} {
	return quantumEntanglement(3, input)
}

func (Day24) Part2(input string) interface{} {
	return quantumEntanglement(4, input)
}

func init() {
	types.Register(Probs, Day24{})
}
