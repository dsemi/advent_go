package year2015

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day03 struct{}

func (Day03) locations(input string) map[utils.Coord]bool {
	m := make(map[utils.Coord]bool)
	pos := utils.Coord{X: 0, Y: 0}
	m[pos] = true
	for _, c := range input {
		switch c {
		case '<':
			pos.X -= 1
		case '>':
			pos.X += 1
		case 'v':
			pos.Y -= 1
		case '^':
			pos.Y += 1
		}
		m[pos] = true
	}
	return m
}

func (d Day03) Part1(input string) interface{} {
	return len(d.locations(input))
}

func (d Day03) Part2(input string) interface{} {
	var b1, b2 strings.Builder
	for i, c := range input {
		if i%2 == 0 {
			b1.WriteRune(c)
		} else {
			b2.WriteRune(c)
		}
	}
	m1 := d.locations(b1.String())
	m2 := d.locations(b2.String())
	for k, _ := range m2 {
		m1[k] = true
	}
	return len(m1)
}

func init() {
	problems.Register(Day03{})
}
