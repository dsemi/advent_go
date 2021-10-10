package year2018

import (
	"math"
	"strings"
	"unicode"
)

func react(input string) int {
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

func Day05Part1(input string) interface{} {
	return react(input)
}

func Day05Part2(input string) interface{} {
	min := math.MaxInt
	for c := 'a'; c <= 'z'; c++ {
		v := react(strings.ReplaceAll(strings.ReplaceAll(input, string(c), ""), string(unicode.ToUpper(c)), ""))
		if v < min {
			min = v
		}
	}
	return min
}
