package year2015

import (
	"advent/types"
	"strings"
)

type Day18 struct{}

func parseGrid(input string) [][]bool {
	var grid [][]bool
	for _, line := range strings.Split(input, "\n") {
		row := make([]bool, len(line))
		for i, c := range line {
			row[i] = c == '#'
		}
		grid = append(grid, row)
	}
	return grid
}

func step(grid [][]bool) {
	neighbs := make([][]int, len(grid))
	for i := range grid {
		neighbs[i] = make([]int, len(grid[i]))
	}
	for i, row := range grid {
		for j, _ := range row {
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x != 0 || y != 0 {
						i2 := i + y
						j2 := j + x
						if i2 >= 0 && i2 < len(grid) && j2 >= 0 && j2 < len(grid[i2]) && grid[i2][j2] {
							neighbs[i][j]++
						}
					}
				}
			}
		}
	}
	for i, row := range grid {
		for j, v := range row {
			grid[i][j] = v && (neighbs[i][j] == 2 || neighbs[i][j] == 3) || !v && neighbs[i][j] == 3
		}
	}
}

func countSquares(grid [][]bool) int {
	var sum int
	for _, row := range grid {
		for _, v := range row {
			if v {
				sum++
			}
		}
	}
	return sum
}

func (Day18) Part1(input string) interface{} {
	grid := parseGrid(input)
	for i := 0; i < 100; i++ {
		step(grid)
	}
	return countSquares(grid)
}

func (Day18) Part2(input string) interface{} {
	grid := parseGrid(input)
	grid[0][0] = true
	grid[0][99] = true
	grid[99][0] = true
	grid[99][99] = true
	for i := 0; i < 100; i++ {
		step(grid)
		grid[0][0] = true
		grid[0][99] = true
		grid[99][0] = true
		grid[99][99] = true
	}
	return countSquares(grid)
}

func init() {
	types.Register(Probs, Day18{})
}
