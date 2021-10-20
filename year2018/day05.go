package year2018

import (
	"advent/problems"
	"math"
	"strings"
	"unicode"
)

type Day05 struct{}

func (Day05) react(input string) int {
	var chs []rune
	for _, c := range input {
		if len(chs) > 0 && chs[len(chs)-1] != c && unicode.ToLower(chs[len(chs)-1]) == unicode.ToLower(c) {
			chs = chs[:len(chs)-1]
		} else {
			chs = append(chs, c)
		}
	}
	return len(chs)
}

func (d Day05) Part1(input string) interface{} {
	return d.react(input)
}

func (d Day05) Part2(input string) interface{} {
	min := math.MaxInt
	for c := 'a'; c <= 'z'; c++ {
		v := d.react(strings.ReplaceAll(strings.ReplaceAll(input, string(c), ""), string(unicode.ToUpper(c)), ""))
		if v < min {
			min = v
		}
	}
	return min
}

func init() {
	problems.Register(Day05{})
}
