package main

import (
	"strings"
	"utils"
)

func quantumEntanglement(n int64, s string) int64 {
	var wts []int64
	for _, line := range strings.Split(s, "\n") {
		wts = append(wts, utils.Int64(line))
	}
	groupSize := utils.Sum(wts) / n
	i := 1
	for {
		var m *int64
		utils.Combinations(wts, i, func(combo []int64) {
			if utils.Sum(combo) == groupSize {
				p := utils.Product(combo)
				if m == nil {
					m = &p
				} else {
					*m = utils.Min(*m, p)
				}
			}
		})
		if m != nil {
			return *m
		}
		i++
	}
}

func Part1(input string) interface{} {
	return quantumEntanglement(3, input)
}

func Part2(input string) interface{} {
	return quantumEntanglement(4, input)
}
