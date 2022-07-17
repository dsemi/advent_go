package main

import (
	"strings"
	"utils"
)

func hash(ns []int) string {
	var b strings.Builder
	for _, n := range ns {
		b.WriteByte(byte(n))
	}
	return b.String()
}

func redistributeUntilCycle(input string) (int, int) {
	ns := make([]int, 0)
	for _, n := range strings.Fields(input) {
		ns = append(ns, utils.Int(n))
	}
	m := make(map[string]int)
	for c := 0; ; c++ {
		if v, ok := m[hash(ns)]; ok {
			return c, c - v
		}
		m[hash(ns)] = c
		var j, val int
		for i, x := range ns {
			if x > val {
				j, val = i, x
			}
		}
		ns[j] = 0
		for k := j + 1; k <= j+val; k++ {
			ns[k%len(ns)]++
		}
	}
}

func Part1(input string) interface{} {
	ans, _ := redistributeUntilCycle(input)
	return ans
}

func Part2(input string) interface{} {
	_, ans := redistributeUntilCycle(input)
	return ans
}
