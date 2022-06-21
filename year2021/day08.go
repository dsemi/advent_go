package main

import (
	"strings"
	"utils"
)

var freqs = []int{42, 17, 34, 39, 30, 37, 41, 25, 49, 45}

func parse(input string) [][]int {
	var results [][]int
	for _, line := range strings.Split(input, "\n") {
		v := strings.Split(line, " | ")
		hist := utils.NewCounter(strings.ReplaceAll(v[0], " ", ""))
		var res []int
	outer:
		for _, n := range strings.Fields(v[1]) {
			var x int
			for _, d := range n {
				x += hist.Get(d)
			}
			for i, f := range freqs {
				if x == f {
					res = append(res, i)
					continue outer
				}
			}
			panic("freq not found")
		}
		results = append(results, res)
	}
	return results
}

func Part1(input string) interface{} {
	var cnt int
	for _, ns := range parse(input) {
		for _, n := range ns {
			if n == 1 || n == 4 || n == 7 || n == 8 {
				cnt++
			}
		}
	}
	return cnt
}

func Part2(input string) interface{} {
	var sum int
	for _, ns := range parse(input) {
		var a int
		for _, b := range ns {
			a = 10*a + b
		}
		sum += a
	}
	return sum
}
