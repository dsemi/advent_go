package main

import (
	"strings"
	"utils"
)

func parse(input string) []int {
	var ns []int
	for _, line := range strings.Fields(input) {
		ns = append(ns, utils.Int(line))
	}
	return ns
}

func Part1(input string) interface{} {
	ns := parse(input)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			if ns[i]+ns[j] == 2020 {
				return ns[i] * ns[j]
			}
		}
	}
	return -1
}

func Part2(input string) interface{} {
	ns := parse(input)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			for k := j + 1; k < len(ns); k++ {
				if ns[i]+ns[j]+ns[k] == 2020 {
					return ns[i] * ns[j] * ns[k]
				}
			}
		}
	}
	panic("unreachable")
}
