package year2015

import (
	"advent/problems"
	"strings"
)

type Day08 struct{}

func (*Day08) encodedLen(s string) int {
	var length int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			i++
			if s[i] == 'x' {
				i += 2
			}
			length++
		case '"':
		default:
			length++
		}
	}
	return length
}

func (d *Day08) Part1(input string) interface{} {
	var diff int
	for _, line := range strings.Split(input, "\n") {
		diff += len(line) - d.encodedLen(line)
	}
	return diff
}

func (*Day08) Part2(input string) interface{} {
	var diff int
	for _, line := range strings.Split(input, "\n") {
		diff += 2
		for _, c := range line {
			if c == '\\' || c == '"' {
				diff++
			}
		}
	}
	return diff
}

func init() {
	problems.Register(&Day08{})
}
