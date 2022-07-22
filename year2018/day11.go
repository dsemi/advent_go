package main

import (
	"fmt"
	"utils"
)

func makeSat(input string) [][]int64 {
	serialNum := utils.Int64(input)
	grid := make([][]int64, 301)
	for y := 1; y < len(grid); y++ {
		row := make([]int64, 301)
		for x := 1; x < len(row); x++ {
			rackId := int64(x) + 10
			powerLevel := rackId * int64(y)
			powerLevel += serialNum
			powerLevel *= rackId
			powerLevel = (powerLevel / 100) % 10
			powerLevel -= 5
			row[x] = powerLevel
		}
		grid[y] = row
	}
	for y := len(grid) - 1; y > 0; y-- {
		for x := len(grid[y]) - 1; x > 0; x-- {
			if y+1 < len(grid) {
				grid[y][x] += grid[y+1][x]
			}
			if x+1 < len(grid[y]) {
				grid[y][x] += grid[y][x+1]
			}
			if y+1 < len(grid) && x+1 < len(grid[y]) {
				grid[y][x] -= grid[y+1][x+1]
			}
		}
	}
	return grid
}

func maxCell(size int, sat [][]int64) (int, int, int64) {
	var a, b int
	var c int64
	for y := 1; y < len(sat)-size; y++ {
		for x := 1; x < len(sat[y])-size; x++ {
			sum := sat[y][x] - sat[y+size][x] - sat[y][x+size] + sat[y+size][x+size]
			if sum > c {
				a, b, c = x, y, sum
			}
		}
	}
	return a, b, c
}

func Part1(input string) interface{} {
	sat := makeSat(input)
	x, y, _ := maxCell(3, sat)
	return fmt.Sprintf("%d,%d", x, y)
}

func Part2(input string) interface{} {
	sat := makeSat(input)
	var x, y, j int
	var max int64
	for i := 1; i < 300; i++ {
		a, b, c := maxCell(i, sat)
		if c > max {
			x, y, j, max = a, b, i, c
		}
	}
	return fmt.Sprintf("%d,%d,%d", x, y, j)
}
