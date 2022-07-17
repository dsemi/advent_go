package main

import (
	"strings"
	"utils"
)

func runCmd(reg map[string]int64, line string) int64 {
	pts := strings.Fields(line)
	r, op, n, r2, cond, n2 := pts[0], pts[1], pts[2], pts[4], pts[5], pts[6]
	var cmpFn func(int64, int64) bool
	if cond == "==" {
		cmpFn = func(a, b int64) bool { return a == b }
	} else if cond == "!=" {
		cmpFn = func(a, b int64) bool { return a != b }
	} else if cond == ">" {
		cmpFn = func(a, b int64) bool { return a > b }
	} else if cond == ">=" {
		cmpFn = func(a, b int64) bool { return a >= b }
	} else if cond == "<" {
		cmpFn = func(a, b int64) bool { return a < b }
	} else if cond == "<=" {
		cmpFn = func(a, b int64) bool { return a <= b }
	} else {
		panic("Parse cond error")
	}
	if cmpFn(reg[r2], utils.Int64(n2)) {
		if op == "inc" {
			reg[r] += utils.Int64(n)
		} else {
			reg[r] -= utils.Int64(n)
		}
	}
	return reg[r]
}

func Part1(input string) interface{} {
	tbl := make(map[string]int64)
	for _, line := range strings.Split(input, "\n") {
		runCmd(tbl, line)
	}
	var max int64
	for _, v := range tbl {
		if v > max {
			max = v
		}
	}
	return max
}

func Part2(input string) interface{} {
	tbl := make(map[string]int64)
	var max int64
	for _, line := range strings.Split(input, "\n") {
		if v := runCmd(tbl, line); v > max {
			max = v
		}
	}
	return max
}
