package main

import (
	"strings"
	"utils"
)

func stabilize(input string, p2 bool) int {
	grid := make([][]byte, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]byte, 0)
		for _, c := range line {
			row = append(row, byte(c))
		}
		grid = append(grid, row)
	}
	seats := make([]utils.Coord, 0)
	neighbs := make([][]utils.Coord, 0)
	for r, row := range grid {
		for c, v := range row {
			if v != 'L' {
				continue
			}
			stCoord := utils.Coord{r, c}
			vec := make([]utils.Coord, 0)
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					drc := utils.Coord{dr, dc}
					start := stCoord.Add(drc)
					for coord := start; coord.X >= 0 && coord.X < len(grid) &&
						coord.Y >= 0 && coord.Y < len(grid[0]); coord = coord.Add(drc) {
						if grid[coord.X][coord.Y] == 'L' {
							if x := stCoord.Sub(coord); p2 || (utils.Abs(x.X) <= 1 && utils.Abs(x.Y) <= 1) {
								vec = append(vec, coord)
							}
							break
						}
					}
				}
			}
			seats = append(seats, stCoord)
			neighbs = append(neighbs, vec)
		}
	}
	for changed := true; changed; {
		changed = false
		grid2 := make([][]byte, len(grid))
		for i, row := range grid {
			grid2[i] = make([]byte, len(row))
			copy(grid2[i], row)
		}
		for i := range seats {
			coord, adjs := seats[i], neighbs[i]
			var adjsOcc int
			for _, adj := range adjs {
				if grid[adj.X][adj.Y] == '#' {
					adjsOcc++
				}
			}
			r, c := coord.X, coord.Y
			if grid[r][c] == 'L' && adjsOcc == 0 {
				grid2[r][c] = '#'
				changed = true
			} else if grid[r][c] == '#' && adjsOcc >= 4+utils.ToInt[int](p2) {
				grid2[r][c] = 'L'
				changed = true
			}
		}
		grid = grid2
	}
	var sum int
	for _, row := range grid {
		for _, v := range row {
			if v == '#' {
				sum++
			}
		}
	}
	return sum
}

func Part1(input string) interface{} {
	return stabilize(input, false)
}

func Part2(input string) interface{} {
	return stabilize(input, true)
}
