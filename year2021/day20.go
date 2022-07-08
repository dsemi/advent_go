package main

import "strings"

type Grid struct {
	arr [][]byte
	iea []byte
}

func (g *Grid) pad(c byte) {
	row := make([]byte, len(g.arr[0]))
	for r := range row {
		row[r] = c
	}
	g.arr = append([][]byte{row}, g.arr...)
	row2 := make([]byte, len(g.arr[0]))
	copy(row2, row)
	g.arr = append(g.arr, row2)
	for r := range g.arr {
		g.arr[r] = append([]byte{c}, g.arr[r]...)
		g.arr[r] = append(g.arr[r], c)
	}
}

func (g *Grid) enhance() {
	g.pad(g.arr[0][0])
	grid := make([][]byte, len(g.arr))
	for r := range g.arr {
		grid[r] = make([]byte, len(g.arr[r]))
		for c := range g.arr[r] {
			if r == 0 || r == len(g.arr)-1 || c == 0 || c == len(g.arr[r])-1 {
				grid[r][c] = g.arr[r][c]
				continue
			}
			var idx int
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					idx <<= 1
					if g.arr[r+dr][c+dc] == '#' {
						idx |= 1
					}
				}
			}
			grid[r][c] = g.iea[idx]
		}
	}
	g.arr = grid
	var idx int
	if g.arr[0][0] == '#' {
		idx = 0b111111111
	}
	ch := g.iea[idx]
	last := len(g.arr) - 1
	for i, row := range g.arr {
		row[0] = ch
		row[last] = ch
		if i == 0 || i == last {
			for j := range row {
				row[j] = ch
			}
		}
	}
}

func run(input string, times int) int {
	pts := strings.Split(input, "\n\n")
	arr := make([][]byte, 0)
	for _, line := range strings.Split(pts[1], "\n") {
		row := make([]byte, 0)
		for _, c := range line {
			row = append(row, byte(c))
		}
		arr = append(arr, row)
	}
	grid := &Grid{
		iea: []byte(pts[0]),
		arr: arr,
	}
	grid.pad('.')
	for i := 0; i < times; i++ {
		grid.enhance()
	}
	var sum int
	for _, row := range grid.arr {
		for _, v := range row {
			if v == '#' {
				sum++
			}
		}
	}
	return sum
}

func Part1(input string) interface{} {
	return run(input, 2)
}

func Part2(input string) interface{} {
	return run(input, 50)
}
