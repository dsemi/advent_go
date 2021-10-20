package year2020

import (
	"advent/problems"
	"strings"
)

type Day03 struct{}

func (Day03) run(right int, down int, input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(lines); i += down {
		if lines[i][i/down*right%len(lines[i])] == '#' {
			total++
		}
	}
	return total
}

func (d Day03) Part1(input string) interface{} {
	return d.run(3, 1, input)
}

func (d Day03) Part2(input string) interface{} {
	return d.run(1, 1, input) *
		d.run(3, 1, input) *
		d.run(5, 1, input) *
		d.run(7, 1, input) *
		d.run(1, 2, input)
}

func init() {
	problems.Register(Day03{})
}
