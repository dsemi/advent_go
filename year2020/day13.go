package main

import (
	"math"
	"strings"
	"utils"
)

func parse(input string) (int64, []int64, []int64) {
	pts := strings.Split(input, "\n")
	t := utils.Int64(pts[0])
	var as, ns []int64
	for i, x := range strings.Split(pts[1], ",") {
		if x == "x" {
			continue
		}
		as = append(as, -int64(i))
		ns = append(ns, utils.Int64(x))
	}
	return t, as, ns
}

func Part1(input string) interface{} {
	t, _, buses := parse(input)
	var ans, min int64 = 0, math.MaxInt64
	for _, b := range buses {
		x := b - t%b
		if x < min {
			min = x
			ans = b * x
		}
	}
	return ans
}

func Part2(input string) interface{} {
	_, as, ns := parse(input)
	return utils.ChineseRemainder(as, ns)
}
