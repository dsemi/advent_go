package main

import (
	"sort"
	"strings"
	"utils"
)

func parse(input string) []utils.Coord {
	coords := make([]utils.Coord, 0)
	for y, line := range strings.Split(input, "\n") {
		for x, v := range line {
			if v == '#' {
				coords = append(coords, utils.Coord{x, y})
			}
		}
	}
	return coords
}

type Angle struct {
	x, y int
}

func newAngle(a, b utils.Coord) Angle {
	x, y := b.X-a.X, b.Y-a.Y
	gcd := utils.Gcd(utils.Abs(x), utils.Abs(y))
	return Angle{x: x / gcd, y: y / gcd}
}

func visibilities(pt utils.Coord, pts []utils.Coord) [][]utils.Coord {
	m := utils.NewOrderedMap[Angle, []utils.Coord]()
	for _, p := range pts {
		if p != pt {
			angle := newAngle(pt, p)
			dist := pt.Dist(p)
			elems := m.Get(angle)
			idx := sort.Search(len(elems), func(i int) bool {
				return dist < pt.Dist(elems[i])
			})
			if idx == len(elems) {
				elems = append(elems, p)
			} else {
				elems = append(elems[:idx+1], elems[idx:]...)
				elems[idx] = p
			}
			m.Put(angle, elems)
		}
	}
	vals := make([][]utils.Coord, 0)
	less := func(a, b Angle) bool {
		if a.x >= 0 && b.x < 0 {
			return true
		} else if a.x < 0 && b.x >= 0 {
			return false
		} else if a.x == 0 && b.x == 0 {
			return a.y < b.y
		} else {
			det := a.x*(-b.y) - (-a.y)*b.x
			return det < 0
		}
	}
	for _, k := range m.Keys(less) {
		vals = append(vals, m.Get(k))
	}
	return vals
}

func maxDetected(asts []utils.Coord) [][]utils.Coord {
	var max [][]utils.Coord
	for _, ast := range asts {
		x := visibilities(ast, asts)
		if len(x) > len(max) {
			max = x
		}
	}
	return max
}

func Part1(input string) interface{} {
	return len(maxDetected(parse(input)))
}

func Part2(input string) interface{} {
	pts := maxDetected(parse(input))
	for {
		var cnt int
		for i := range pts {
			if len(pts[i]) > 0 {
				cnt++
				if cnt == 200 {
					c := pts[i][0]
					return 100*c.X + c.Y
				}
				pts[i] = pts[i][1:]
			}
		}
	}
}
