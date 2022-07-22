package main

import (
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	grid = utils.Transpose(grid)
	var b strings.Builder
	for _, row := range grid {
		mostCommon := utils.NewCounter(row)
		b.WriteRune(mostCommon.Keys()[0])
	}
	return b.String()
}

func Part2(input string) interface{} {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	grid = utils.Transpose(grid)
	var b strings.Builder
	for _, row := range grid {
		mostCommon := utils.NewCounter(row)
		keys := mostCommon.Keys()
		b.WriteRune(keys[len(keys)-1])
	}
	return b.String()
}
