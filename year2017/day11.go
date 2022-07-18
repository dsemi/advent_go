package main

import (
	"strings"
	"utils"
)

func path(input string) []utils.Coord3 {
	pts := make([]utils.Coord3, 1)
	pts[0] = utils.Coord3{0, 0, 0}
	for _, d := range strings.Split(input, ",") {
		var x utils.Coord3
		if d == "n" {
			x = utils.Coord3{0, 1, -1}
		} else if d == "ne" {
			x = utils.Coord3{1, 0, -1}
		} else if d == "se" {
			x = utils.Coord3{1, -1, 0}
		} else if d == "s" {
			x = utils.Coord3{0, -1, 1}
		} else if d == "sw" {
			x = utils.Coord3{-1, 0, 1}
		} else if d == "nw" {
			x = utils.Coord3{-1, 1, 0}
		} else {
			panic("Parse error")
		}
		pts = append(pts, pts[len(pts)-1].Add(x))
	}
	return pts
}

func distFromOrigin(pt utils.Coord3) int {
	return utils.Maximum([]int{utils.Abs(pt.X), utils.Abs(pt.Y), utils.Abs(pt.Z)})
}

func Part1(input string) interface{} {
	pts := path(input)
	return distFromOrigin(pts[len(pts)-1])
}

func Part2(input string) interface{} {
	var ans int
	for _, p := range path(input) {
		ans = utils.Max(ans, distFromOrigin(p))
	}
	return ans
}
