package main

import (
	"strings"
	"utils"
)

func parse(input string) map[utils.Coord]rune {
	m := make(map[utils.Coord]rune)
	for r, line := range strings.Split(input, "\n") {
		for c, v := range line {
			if v != ' ' {
				m[utils.Coord{r, c}] = v
			}
		}
	}
	return m
}

var (
	left  = utils.Coord{0, 1}
	right = utils.Coord{0, -1}
)

func turn(grid map[utils.Coord]rune, dir, pos utils.Coord) utils.Coord {
	if _, ok := grid[left.Mul(dir).Add(pos)]; ok {
		return left.Mul(dir)
	}
	return right.Mul(dir)
}

func followPath(grid map[utils.Coord]rune) []rune {
	var coord utils.Coord
	for k := range grid {
		coord = k
		break
	}
	for k := range grid {
		if k.X == coord.X && k.Y < coord.Y || k.X < coord.X {
			coord = k
		}
	}
	dir := utils.Coord{1, 0}
	result := make([]rune, 0)
	for v, ok := grid[coord]; ok; v, ok = grid[coord] {
		result = append(result, v)
		if v == '+' {
			dir = turn(grid, dir, coord)
		}
		coord = coord.Add(dir)
	}
	return result
}

func Part1(input string) interface{} {
	ans := make([]rune, 0)
	for _, x := range followPath(parse(input)) {
		if x != '|' && x != '-' && x != '+' {
			ans = append(ans, x)
		}
	}
	return string(ans)
}

func Part2(input string) interface{} {
	return len(followPath(parse(input)))
}
