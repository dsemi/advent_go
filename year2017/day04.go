package main

import (
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		ps := strings.Fields(line)
		m := make(map[string]bool)
		for _, p := range ps {
			m[p] = true
		}
		if len(ps) == len(m) {
			sum++
		}
	}
	return sum
}

func SortString(s string) string {
	r := []rune(s)
	utils.Sort(r)
	return string(r)
}

func Part2(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		ps := strings.Fields(line)
		for i, p := range ps {
			ps[i] = SortString(p)
		}
		m := make(map[string]bool)
		for _, p := range ps {
			m[p] = true
		}
		if len(ps) == len(m) {
			sum++
		}
	}
	return sum
}
