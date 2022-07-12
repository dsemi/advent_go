package main

import (
	"fmt"
	"math"
	"strings"
	"utils"
)

type Grid = [][]bool

type Tile struct {
	num  uint64
	grid Grid
}

func parse(input string) []Tile {
	tiles := make([]Tile, 0)
	for _, x := range strings.Split(input, "\n\n") {
		grid := make(Grid, 0)
		gen := strings.Split(x, "\n")
		var n uint64
		fmt.Sscanf(gen[0], "Tile %d:", &n)
		for _, row := range gen[1:] {
			r := make([]bool, 0)
			for _, y := range row {
				r = append(r, y == '#')
			}
			grid = append(grid, r)
		}
		tiles = append(tiles, Tile{num: n, grid: grid})
	}
	return tiles
}

func transpose(grid Grid) {
	for i := range grid {
		for j := i + 1; j < len(grid); j++ {
			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}
	}
}

func reverse(grid Grid) {
	last := len(grid) - 1
	for i := 0; i < len(grid)/2; i++ {
		grid[i], grid[last-i] = grid[last-i], grid[i]
	}
}

func clone(grid Grid) Grid {
	grid2 := make(Grid, len(grid))
	for i, row := range grid {
		grid2[i] = make([]bool, len(row))
		copy(grid2[i], row)
	}
	return grid2
}

func orientations(tile Grid) []Grid {
	t := clone(tile)
	v := make([]Grid, 8)
	for i := 0; i < 4; i++ {
		v[2*i] = clone(t)
		transpose(t)
		v[2*i+1] = clone(t)
		reverse(t)
	}
	return v
}

func firstCol(grid Grid) uint64 {
	var hash uint64
	for _, row := range grid {
		hash <<= 1
		if row[0] {
			hash |= 1
		}
	}
	return hash
}

func firstRow(grid Grid) uint64 {
	var hash uint64
	for _, v := range grid[0] {
		hash <<= 1
		if v {
			hash |= 1
		}
	}
	return hash
}

func lastCol(grid Grid) uint64 {
	var hash uint64
	for _, row := range grid {
		hash <<= 1
		if row[len(row)-1] {
			hash |= 1
		}
	}
	return hash
}

func lastRow(grid Grid) uint64 {
	var hash uint64
	for _, v := range grid[len(grid)-1] {
		hash <<= 1
		if v {
			hash |= 1
		}
	}
	return hash
}

func cornersAndHashes(tiles []Tile) ([]uint64, map[uint64][]Tile) {
	m := make(map[uint64][]Tile)
	for _, tile := range tiles {
		for _, t := range orientations(tile.grid) {
			hash := firstCol(t)
			m[hash] = append(m[hash], Tile{num: tile.num, grid: t})
		}
	}
	m2 := make(map[uint64]int)
	for _, v := range m {
		if len(v) == 1 {
			m2[v[0].num]++
		}
	}
	corners := make([]uint64, 0)
	for k, v := range m2 {
		if v == 4 {
			corners = append(corners, k)
		}
	}
	return corners, m
}

func Part1(input string) interface{} {
	tiles := parse(input)
	corners, _ := cornersAndHashes(tiles)
	return utils.Product(corners)
}

func placeTiles(tiles []Tile) ([]Tile, int) {
	size := int(math.Sqrt(float64(len(tiles))))
	grid := make([]Tile, 0)
	corners, m := cornersAndHashes(tiles)
	var start Tile
outer:
	for _, tile := range tiles {
		for _, c := range corners {
			if tile.num == c {
				start = tile
				break outer
			}
		}
	}
	for len(m[lastCol(start.grid)]) < 2 || len(m[lastRow(start.grid)]) < 2 {
		transpose(start.grid)
		reverse(start.grid)
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			var g Tile
			if r == 0 && c == 0 {
				g.num = start.num
				g.grid = clone(start.grid)
			} else if c == 0 {
				prev := grid[(r-1)*size+c]
				for _, t := range m[lastRow(prev.grid)] {
					if t.num != prev.num {
						g.num = t.num
						g.grid = clone(t.grid)
						transpose(g.grid)
						break
					}
				}
			} else {
				prev := grid[r*size+c-1]
				for _, t := range m[lastCol(prev.grid)] {
					if t.num != prev.num {
						g.num = t.num
						g.grid = clone(t.grid)
						break
					}
				}
			}
			grid = append(grid, g)
		}
	}
	return grid, size
}

func findSeaMonsters(pic [][]bool) int {
	size := 20
	mons := []int{
		0b00000000000000000010,
		0b10000110000110000111,
		0b01001001001001001000,
	}
	var cnt int
	for _, m := range mons {
		cnt += utils.CountOnes(m)
	}
	var sum int
	for i := 0; i < len(pic)-2; i++ {
		wind := pic[i : i+3]
		rs := clone(wind)
		var tot int
		for len(rs[0]) >= size {
			matched := true
			for r := range rs {
				if utils.BitsToInt[int](rs[r][:size])&mons[r] != mons[r] {
					matched = false
				}
			}
			if matched {
				tot += cnt
			}
			for j := range rs {
				rs[j] = rs[j][1:]
			}
		}
		sum += tot
	}
	return sum
}

func Part2(input string) interface{} {
	grid, size := placeTiles(parse(input))
	var innerSize int
	for i := range grid {
		grid[i].grid = grid[i].grid[1 : len(grid[i].grid)-1]
		for j := range grid[i].grid {
			grid[i].grid[j] = grid[i].grid[j][1 : len(grid[i].grid[j])-1]
		}
		innerSize = len(grid[i].grid)
	}
	pic := make([][]bool, 0)
	for i := 0; i < len(grid); i += size {
		chunk := grid[i : i+size]
		for row := 0; row < innerSize; row++ {
			r := make([]bool, 0)
			for _, t := range chunk {
				r = append(r, t.grid[row]...)
			}
			pic = append(pic, r)
		}
	}

	for _, p := range orientations(pic) {
		ms := findSeaMonsters(p)
		if ms != 0 {
			var tot int
			for _, row := range p {
				for _, v := range row {
					if v {
						tot++
					}
				}
			}
			return tot - ms
		}
	}
	panic("unreachable")
}
