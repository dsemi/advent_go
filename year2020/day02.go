package year2020

import (
	"advent/types"
	"strconv"
	"strings"
)

type Day02 struct{}

func solve(f func(int, int, byte, string) bool, input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		ns := strings.Split(parts[0], "-")
		a, _ := strconv.Atoi(ns[0])
		b, _ := strconv.Atoi(ns[1])
		if f(a, b, parts[1][0], parts[2]) {
			total++
		}
	}
	return total
}

func (Day02) Part1(input string) interface{} {
	return solve(func(a, b int, c byte, str string) bool {
		cnt := strings.Count(str, string(c))
		return a <= cnt && cnt <= b
	}, input)
}

func (Day02) Part2(input string) interface{} {
	return solve(func(a, b int, c byte, str string) bool {
		return (str[a-1] == c) != (str[b-1] == c)
	}, input)
}

func init() {
	types.Register(Probs, Day02{})
}
