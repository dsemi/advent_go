package year2016

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day01 struct{}

func (Day01) path(input string) chan utils.Coord {
	c := make(chan utils.Coord)
	go func() {
		defer close(c)
		dir := utils.Coord{X: 0, Y: 1}
		pos := utils.Coord{X: 0, Y: 0}
		for _, cmd := range strings.Split(input, ", ") {
			if cmd[0] == 'R' {
				dir = utils.Coord{X: dir.Y, Y: -dir.X}
			} else {
				dir = utils.Coord{X: -dir.Y, Y: dir.X}
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

func (d Day01) Part1(input string) interface{} {
	var p utils.Coord
	for p = range d.path(input) {
	}
	return utils.Abs(p.X) + utils.Abs(p.Y)
}

func (d Day01) Part2(input string) interface{} {
	m := make(map[utils.Coord]bool)
	for p := range d.path(input) {
		if m[p] {
			return utils.Abs(p.X) + utils.Abs(p.Y)
		}
		m[p] = true
	}
	return nil
}

func init() {
	problems.Register(Day01{})
}
