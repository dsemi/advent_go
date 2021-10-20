package year2017

import (
	"advent/problems"
	"advent/utils"
	"sort"
	"strings"
)

type Day04 struct{}

func (Day04) Part1(input string) interface{} {
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

func (Day04) SortString(s string) string {
	r := []rune(s)
	sort.Sort(utils.SortRunes(r))
	return string(r)
}

func (d Day04) Part2(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		ps := strings.Fields(line)
		for i, p := range ps {
			ps[i] = d.SortString(p)
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

func init() {
	problems.Register(Day04{})
}
