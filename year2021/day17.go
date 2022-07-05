package main

import (
	"fmt"
	"math"
	"utils"
)

func parse(input string) (int64, int64, int64, int64) {
	var x0, x1, y0, y1 int64
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x0, &x1, &y0, &y1)
	return x0, x1, y0, y1
}

func Part1(input string) interface{} {
	_, _, y0, _ := parse(input)
	return y0 * (y0 + 1) / 2
}

func hitsTarget(x0, x1, y0, y1, vx, vy int64) bool {
	var x, y int64
	for x <= x1 && y >= y0 {
		x, y = x+vx, y+vy
		vx = utils.Max(0, vx-1)
		vy--
		if x0 <= x && x <= x1 && y0 <= y && y <= y1 {
			return true
		}
	}
	return false
}

func Part2(input string) interface{} {
	x0, x1, y0, y1 := parse(input)
	// First triangular number > x0 is lower bound.
	// n^2 + n - 2x0 = 0
	mx := int64(math.Ceil(math.Sqrt(float64(1+8*x0))/2 - 0.5))
	var cnt int
	for x := mx; x <= x1; x++ {
		for y := y0; y <= -y0; y++ {
			if hitsTarget(x0, x1, y0, y1, x, y) {
				cnt++
			}
		}
	}
	return cnt
}
