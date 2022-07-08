package main

import (
	"fmt"
	"math"
	"strings"
	"utils"
)

type key struct {
	a, b, c, d, e, f int64
}

func solve(input string, lo, hi int64) int64 {
	cubes := make(map[key]int64)
	for _, line := range strings.Split(input, "\n") {
		var w string
		var nx0, nx1, ny0, ny1, nz0, nz1 int64
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &w, &nx0, &nx1, &ny0, &ny1, &nz0, &nz1)
		update := make(map[key]int64)
		for k, es := range cubes {
			ex0, ex1, ey0, ey1, ez0, ez1 := k.a, k.b, k.c, k.d, k.e, k.f
			x0, x1 := utils.Max(nx0, ex0), utils.Min(nx1, ex1)
			y0, y1 := utils.Max(ny0, ey0), utils.Min(ny1, ey1)
			z0, z1 := utils.Max(nz0, ez0), utils.Min(nz1, ez1)
			if x0 <= x1 && y0 <= y1 && z0 <= z1 {
				update[key{x0, x1, y0, y1, z0, z1}] -= es
			}
		}
		if w == "on" {
			update[key{nx0, nx1, ny0, ny1, nz0, nz1}]++
		}
		for k, v := range update {
			cubes[k] += v
		}
	}
	var sum int64
	for k, s := range cubes {
		x0, x1, y0, y1, z0, z1 := k.a, k.b, k.c, k.d, k.e, k.f
		x0, x1 = utils.Max(lo, x0), utils.Min(hi, x1)
		y0, y1 = utils.Max(lo, y0), utils.Min(hi, y1)
		z0, z1 = utils.Max(lo, z0), utils.Min(hi, z1)
		sum += utils.Max(0, x1-x0+1) * utils.Max(0, y1-y0+1) * utils.Max(0, z1-z0+1) * s
	}
	return sum
}

func Part1(input string) interface{} {
	return solve(input, -50, 50)
}

func Part2(input string) interface{} {
	return solve(input, math.MinInt64, math.MaxInt64)
}
