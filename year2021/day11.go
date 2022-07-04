package main

import (
	"strings"
	"utils"
)

type grid struct {
	arr [10][10]int
}

func newGrid(input string) grid {
	gr := grid{}
	for r, line := range strings.Split(input, "\n") {
		for c, v := range line {
			gr.arr[r][c] = int(v - '0')
		}
	}
	return gr
}

func (g *grid) flash(r, c int) int {
	g.arr[r][c] = -1
	flashes := 1
	for _, r1 := range []int{r - 1, r, r + 1} {
		for _, c1 := range []int{c - 1, c, c + 1} {
			if r1 >= 0 && r1 < len(g.arr) && c1 >= 0 && c1 < len(g.arr[r1]) && g.arr[r1][c1] != -1 {
				g.arr[r1][c1]++
				if g.arr[r1][c1] > 9 {
					flashes += g.flash(r1, c1)
				}
			}
		}
	}
	return flashes
}

func (g *grid) step() int {
	for r := range g.arr {
		for c := range g.arr[r] {
			g.arr[r][c] = utils.Max(g.arr[r][c], 0) + 1
		}
	}
	var sum int
	for r := range g.arr {
		for c := range g.arr[r] {
			if g.arr[r][c] > 9 {
				sum += g.flash(r, c)
			}
		}
	}
	return sum
}

func Part1(input string) interface{} {
	g := newGrid(input)
	var sum int
	for i := 0; i < 100; i++ {
		sum += g.step()
	}
	return sum
}

func Part2(input string) interface{} {
	g := newGrid(input)
	for i := 1; ; i++ {
		if g.step() == 100 {
			return i
		}
	}
}
