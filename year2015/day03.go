package year2015

import (
	"advent/types"
	"strings"
)

type Day03 struct{}

type Coord struct {
	x int
	y int
}

func locations(input string) map[Coord]bool {
	m := make(map[Coord]bool)
	pos := Coord{x: 0, y: 0}
	m[pos] = true
	for _, c := range input {
		switch c {
		case '<':
			pos.x -= 1
		case '>':
			pos.x += 1
		case 'v':
			pos.y -= 1
		case '^':
			pos.y += 1
		}
		m[pos] = true
	}
	return m
}

func (Day03) Part1(input string) interface{} {
	return len(locations(input))
}

func (Day03) Part2(input string) interface{} {
	var b1, b2 strings.Builder
	for i, c := range input {
		if i%2 == 0 {
			b1.WriteRune(c)
		} else {
			b2.WriteRune(c)
		}
	}
	m1 := locations(b1.String())
	m2 := locations(b2.String())
	for k, _ := range m2 {
		m1[k] = true
	}
	return len(m1)
}

func init() {
	types.Register(Probs, Day03{})
}
