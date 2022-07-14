package main

import (
	"strings"
	"utils"
)

type segment struct {
	horz bool
	a, b utils.Coord
	d    int
	r    bool
}

func parse(input string) [][]segment {
	wires := make([][]segment, 0)
	for _, line := range strings.Split(input, "\n") {
		segments := make([]segment, 0)
		pos := utils.Coord{0, 0}
		steps := 0
		for _, p := range strings.Split(line, ",") {
			segment := segment{}
			var d utils.Coord
			switch p[0] {
			case 'U':
				segment.horz = false
				d = utils.Coord{0, 1}
			case 'D':
				segment.horz = false
				d = utils.Coord{0, -1}
			case 'L':
				segment.horz = true
				d = utils.Coord{-1, 0}
			case 'R':
				segment.horz = true
				d = utils.Coord{1, 0}
			default:
				panic("Unknown direction")
			}
			n := utils.Int(p[1:])
			prev := pos
			pos = pos.Add(d.Scale(n))
			if prev.Less(pos) {
				segment.d = steps
				segment.a = prev
				segment.b = pos
				segment.r = false
			} else {
				segment.d = steps + n
				segment.a = pos
				segment.b = prev
				segment.r = true
			}
			steps += n
			segments = append(segments, segment)
		}
		wires = append(wires, segments)
	}
	return wires
}

func intersections(a, b []segment) [][2]int {
	pts := make([][2]int, 0)
	for _, w1 := range a {
		for _, w2 := range b {
			if w1.horz == w2.horz {
				continue
			}
			var hs, vs segment
			if w1.horz {
				hs, vs = w1, w2
			} else {
				hs, vs = w2, w1
			}
			if hs.a.X <= vs.a.X && vs.a.X <= hs.b.X && vs.a.Y <= hs.a.Y && hs.a.Y <= vs.b.Y {
				b := hs.d + vs.d
				if hs.r {
					b -= utils.Abs(hs.a.X - vs.a.X)
				} else {
					b += utils.Abs(hs.a.X - vs.a.X)
				}
				if vs.r {
					b -= utils.Abs(vs.a.Y - hs.a.Y)
				} else {
					b += utils.Abs(vs.a.Y - hs.a.Y)
				}
				pts = append(pts, [2]int{utils.Abs(vs.a.X) + utils.Abs(hs.a.Y), b})
			}
		}
	}
	return pts
}

func solve(input string, idx int) int {
	wires := parse(input)
	xs := make([]int, 0)
	for w1 := range wires {
		for w2 := w1 + 1; w2 < len(wires); w2++ {
			for _, p := range intersections(wires[w1], wires[w2]) {
				xs = append(xs, p[idx])
			}
		}
	}
	return utils.Minimum(xs)
}

func Part1(input string) interface{} {
	return solve(input, 0)
}

func Part2(input string) interface{} {
	return solve(input, 1)
}
