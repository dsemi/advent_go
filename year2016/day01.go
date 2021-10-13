package year2016

import (
	"advent/types"
	"advent/utils"
	"strings"
)

type Day01 struct{}

type Pos struct {
	x int
	y int
}

func path(input string) chan Pos {
	c := make(chan Pos)
	go func() {
		defer close(c)
		dir := Pos{x: 0, y: 1}
		pos := Pos{x: 0, y: 0}
		for _, cmd := range strings.Split(input, ", ") {
			if cmd[0] == 'R' {
				dir = Pos{x: dir.y, y: -dir.x}
			} else {
				dir = Pos{x: -dir.y, y: dir.x}
			}
			n := utils.Int(cmd[1:])
			for i := 0; i < n; i++ {
				pos.x += dir.x
				pos.y += dir.y
				c <- pos
			}
		}
	}()
	return c
}

func (Day01) Part1(input string) interface{} {
	var p Pos
	for p = range path(input) {
	}
	return utils.Abs(p.x) + utils.Abs(p.y)
}

func (Day01) Part2(input string) interface{} {
	m := make(map[Pos]bool)
	for p := range path(input) {
		if m[p] {
			return utils.Abs(p.x) + utils.Abs(p.y)
		}
		m[p] = true
	}
	return nil
}

func init() {
	types.Register(Probs, Day01{})
}
