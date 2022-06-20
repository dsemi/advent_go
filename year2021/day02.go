package main

import (
	"strings"
	"utils"
)

func run(input string) (int, int, int) {
	horz, depth, aim := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		v := strings.Fields(line)
		cmd := v[0]
		n := utils.Int(v[1])
		if cmd == "forward" {
			horz += n
			depth += aim * n
		} else if cmd == "down" {
			aim += n
		} else if cmd == "up" {
			aim -= n
		}
	}
	return horz, depth, aim
}

func Part1(input string) interface{} {
	horz, _, depth := run(input)
	return horz * depth
}

func Part2(input string) interface{} {
	horz, depth, _ := run(input)
	return horz * depth
}
