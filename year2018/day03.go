package main

import (
	"fmt"
	"strings"
	"utils"
)

type Claim struct {
	num, x0, y0, x1, y1 int
}

func parse(input string) []Claim {
	claims := make([]Claim, 0)
	for _, line := range strings.Split(input, "\n") {
		var n, x, y, w, h int
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		claims = append(claims, Claim{n, x, y, x + w, y + h})
	}
	return claims
}

func coordFreqs(claims []Claim) [][]int {
	var maxX, maxY int
	for _, c := range claims {
		maxX = utils.Max(maxX, c.x1)
		maxY = utils.Max(maxY, c.y1)
	}
	result := make([][]int, maxY)
	for i := range result {
		result[i] = make([]int, maxX)
	}
	for _, claim := range claims {
		for y := claim.y0; y < claim.y1; y++ {
			for x := claim.x0; x < claim.x1; x++ {
				result[y][x]++
			}
		}
	}
	return result
}

func Part1(input string) interface{} {
	var sum int
	for _, row := range coordFreqs(parse(input)) {
		for _, v := range row {
			if v > 1 {
				sum++
			}
		}
	}
	return sum
}

func Part2(input string) interface{} {
	claims := parse(input)
	grid := coordFreqs(claims)
outer:
	for _, claim := range claims {
		for y := claim.y0; y < claim.y1; y++ {
			for x := claim.x0; x < claim.x1; x++ {
				if grid[y][x] != 1 {
					continue outer
				}
			}
		}
		return claim.num
	}
	panic("unreachable")
}
