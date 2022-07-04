package main

import (
	"strings"
	"utils"
)

func parse(input string) [][]uint8 {
	grid := make([][]uint8, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]uint8, 0)
		for _, c := range line {
			row = append(row, uint8(c-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}

func Part1(input string) interface{} {
	grid := parse(input)
	var sum int
	for i, row := range grid {
		for j, v := range row {
			if (i == 0 || grid[i-1][j] > v) &&
				(j == 0 || grid[i][j-1] > v) &&
				(i+1 == len(grid) || grid[i+1][j] > v) &&
				(j+1 == len(grid[i]) || grid[i][j+1] > v) {
				sum += int(v) + 1
			}
		}
	}
	return sum
}

func dfs(grid [][]uint8, vis [][]bool, i, j int) int {
	if vis[i][j] || grid[i][j] == 9 {
		return 0
	}
	vis[i][j] = true
	sz := 1
	if i > 0 {
		sz += dfs(grid, vis, i-1, j)
	}
	if j > 0 {
		sz += dfs(grid, vis, i, j-1)
	}
	if i+1 < len(grid) {
		sz += dfs(grid, vis, i+1, j)
	}
	if j+1 < len(grid[i]) {
		sz += dfs(grid, vis, i, j+1)
	}
	return sz
}

func Part2(input string) interface{} {
	grid := parse(input)
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}
	basins := make([]int, 3)
	for i := range grid {
		for j := range grid[i] {
			if size := dfs(grid, vis, i, j); size > 0 {
				for x := range basins {
					if size > basins[x] {
						size, basins[x] = basins[x], size
					}
				}
			}
		}
	}
	return utils.Product(basins)
}
