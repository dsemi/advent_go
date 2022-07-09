package main

import (
	"strings"
	"utils"
)

func path(input string) chan utils.Coord {
	c := make(chan utils.Coord)
	go func() {
		defer close(c)
		dir := utils.Coord{0, 1}
		pos := utils.Coord{0, 0}
		for _, cmd := range strings.Split(input, ", ") {
			if cmd[0] == 'R' {
				dir = utils.Coord{dir.Y, -dir.X}
			} else {
				dir = utils.Coord{-dir.Y, dir.X}
			}
			n := utils.Int(cmd[1:])
			for i := 0; i < n; i++ {
				pos.X += dir.X
				pos.Y += dir.Y
				c <- pos
			}
		}
	}()
	return c
}

func Part1(input string) interface{} {
	var p utils.Coord
	for p = range path(input) {
	}
	return utils.Abs(p.X) + utils.Abs(p.Y)
}

func Part2(input string) interface{} {
	m := make(map[utils.Coord]bool)
	for p := range path(input) {
		if m[p] {
			return utils.Abs(p.X) + utils.Abs(p.Y)
		}
		m[p] = true
	}
	return nil
}
