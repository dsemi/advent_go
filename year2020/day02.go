package main

import (
	"fmt"
	"strings"
)

func solve(f func(int, int, byte, string) bool, input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		var str string
		var c byte
		var a, b int
		fmt.Sscanf(line, "%d-%d %c: %s", &a, &b, &c, &str)
		if f(a, b, c, str) {
			total++
		}
	}
	return total
}

func Part1(input string) interface{} {
	return solve(func(a, b int, c byte, str string) bool {
		cnt := strings.Count(str, string(c))
		return a <= cnt && cnt <= b
	}, input)
}

func Part2(input string) interface{} {
	return solve(func(a, b int, c byte, str string) bool {
		return (str[a-1] == c) != (str[b-1] == c)
	}, input)
}
