package year2017

import (
	"advent/utils"
	"sort"
	"strings"
)

func Day04Part1(input string) interface{} {
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
    sort.Sort(utils.SortRunes(r))
    return string(r)
}

func Day04Part2(input string) interface{} {
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
