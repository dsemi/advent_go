package main

import (
	"strings"
	"utils"
)

type pair struct {
	a, b int
}

func parse(input string) []pair {
	pairs := make([]pair, 0)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, ": ")
		pairs = append(pairs, pair{utils.Int(pts[0]), 2*utils.Int(pts[1]) - 2})
	}
	return pairs
}

func Part1(input string) interface{} {
	var sum int
	for _, p := range parse(input) {
		if p.a%p.b == 0 {
			sum += p.a * (p.b + 2) / 2
		}
	}
	return sum
}

func Part2(input string) interface{} {
	scrs := parse(input)
outer:
	for i := 0; ; i++ {
		for _, p := range scrs {
			if (p.a+i)%p.b == 0 {
				continue outer
			}
		}
		return i
	}
}
