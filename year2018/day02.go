package main

import (
	"strings"
)

func Part1(input string) interface{} {
	var twos, threes int
	for _, line := range strings.Split(input, "\n") {
		var n2, n3 int
		var cnts [26]int
		for _, c := range line {
			if v := cnts[c-'a']; v == 1 {
				n2++
			} else if v == 2 {
				n2--
				n3++
			} else if v == 3 {
				n3--
			}
			cnts[c-'a']++
		}
		if n2 > 0 {
			twos++
		}
		if n3 > 0 {
			threes++
		}
	}
	return twos * threes
}

func Part2(input string) interface{} {
	ids := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		ids = append(ids, line)
	}
	common := make([]byte, 0, len(ids[0]))
	for i, b1 := range ids {
		for j := i + 1; j < len(ids); j++ {
			b2 := ids[j]
			for k := range b1 {
				if b1[k] == b2[k] {
					common = append(common, b1[k])
				}
			}
			if len(common)+1 == len(b1) {
				return string(common)
			}
			common = common[:0]
		}
	}
	panic("unreachable")
}
