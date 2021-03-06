package main

import (
	"fmt"
	"strings"
	"utils"
)

type pair struct {
	a utils.Coord
	b utils.Coord
}

func solve(input string, p2 bool) int {
	var lines []pair
	maxX, maxY := 0, 0
	for _, line := range strings.Split(input, "\n") {
		var x0, y0, x1, y1 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		maxX = utils.Max(maxX, utils.Max(x0, x1))
		maxY = utils.Max(maxY, utils.Max(y0, y1))
		lines = append(lines, pair{a: utils.Coord{x0, y0}, b: utils.Coord{x1, y1}})
	}
	grid := make([][]int, maxX+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, maxY+1)
	}
	for _, p := range lines {
		if !p2 && p.a.X != p.b.X && p.a.Y != p.b.Y {
			continue
		}
		d := p.b.Sub(p.a).Sgn()
		curr := p.a
		end := p.b.Add(d)
		for curr != end {
			grid[curr.X][curr.Y]++
			curr = curr.Add(d)
		}
	}
	var cnt int
	for _, row := range grid {
		for _, v := range row {
			if v > 1 {
				cnt++
			}
		}
	}
	return cnt
}

func Part1(input string) interface{} {
	return solve(input, false)
}

func Part2(input string) interface{} {
	return solve(input, true)
}
