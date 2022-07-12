package main

import (
	"math"
	"strings"
	"utils"
)

func flipTiles(input string) map[utils.Coord]bool {
	tiles := make(map[utils.Coord]int)
	for _, line := range strings.Split(input, "\n") {
		var tile utils.Coord
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'e':
				tile = tile.Add(utils.Coord{1, -1})
			case 's':
				i++
				if line[i] == 'e' {
					tile = tile.Add(utils.Coord{0, -1})
				} else {
					tile = tile.Add(utils.Coord{-1, 0})
				}
			case 'w':
				tile = tile.Add(utils.Coord{-1, 1})
			case 'n':
				i++
				if line[i] == 'w' {
					tile = tile.Add(utils.Coord{0, 1})
				} else {
					tile = tile.Add(utils.Coord{1, 0})
				}
			default:
				panic("Invalid dir")
			}
		}
		tiles[tile]++
	}
	set := make(map[utils.Coord]bool)
	for t, v := range tiles {
		if v%2 == 1 {
			set[t] = true
		}
	}
	return set
}

func Part1(input string) interface{} {
	return len(flipTiles(input))
}

const steps = 100

func clone(grid [][]bool) [][]bool {
	grid2 := make([][]bool, len(grid))
	for i, row := range grid {
		grid2[i] = make([]bool, len(row))
		copy(grid2[i], row)
	}
	return grid2
}

func Part2(input string) interface{} {
	tiles := flipTiles(input)
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for coord := range tiles {
		minX = utils.Min(minX, coord.X)
		minY = utils.Min(minY, coord.Y)
		maxX = utils.Max(maxX, coord.X)
		maxY = utils.Max(maxY, coord.Y)
	}
	xOffset := -minX + steps + 1
	yOffset := -minY + steps + 1
	minX += xOffset
	minY += yOffset
	maxX += xOffset
	maxY += yOffset
	grid := make([][]bool, maxY+steps+2)
	for i := range grid {
		grid[i] = make([]bool, maxX+steps+2)
	}
	for tile := range tiles {
		grid[tile.Y+yOffset][tile.X+xOffset] = true
	}
	grid2 := clone(grid)
	for i := 0; i < steps; i++ {
		minX--
		minY--
		maxX++
		maxY++
		for r := minY; r <= maxY; r++ {
			for c := minX; c <= maxX; c++ {
				var adj int
				for _, d := range []utils.Coord{{1, -1}, {0, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 0}} {
					if grid[r+d.Y][c+d.X] {
						adj++
					}
				}
				if grid[r][c] {
					grid2[r][c] = adj != 0 && adj <= 2
				} else {
					grid2[r][c] = adj == 2
				}
			}
		}
		grid2, grid = grid, grid2
	}
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
