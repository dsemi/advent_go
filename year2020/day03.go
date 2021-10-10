package year2020

import "strings"

func run(right int, down int, input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(lines); i += down {
		if lines[i][i/down*right%len(lines[i])] == '#' {
			total++
		}
	}
	return total
}

func Day03Part1(input string) interface{} {
	return run(3, 1, input)
}

func Day03Part2(input string) interface{} {
	return run(1, 1, input) *
		run(3, 1, input) *
		run(5, 1, input) *
		run(7, 1, input) *
		run(1, 2, input)
}
