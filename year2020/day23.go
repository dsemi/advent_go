package main

import (
	"fmt"
	"strings"
)

func read(input string, arr []int) {
	ns := make([]int, 0)
	for _, c := range input {
		ns = append(ns, int(c-'0'))
	}
	for i := 10; i < len(arr); i++ {
		ns = append(ns, i)
	}
	arr[0] = ns[0]
	for i := 0; i < len(ns); i++ {
		arr[ns[i]] = ns[(i+1)%len(ns)]
	}
}

func run(steps int, d []int) {
	m := len(d) - 1
	curr := d[0]
	for i := 0; i < steps; i++ {
		a := d[curr]
		b := d[a]
		c := d[b]
		n := curr
		for n == curr || n == a || n == b || n == c {
			n--
			if n == 0 {
				n = m
			}
		}
		d[curr] = d[c]
		d[c] = d[n]
		d[n] = a
		curr = d[curr]
	}
}

func Part1(input string) interface{} {
	arr := make([]int, 10)
	read(input, arr)
	run(100, arr)
	var res strings.Builder
	x := arr[1]
	for x != 1 {
		fmt.Fprint(&res, x)
		x = arr[x]
	}
	return res.String()
}

func Part2(input string) interface{} {
	arr := make([]int, 1_000_001)
	read(input, arr)
	run(10_000_000, arr)
	return uint64(arr[1]) * uint64(arr[arr[1]])
}
