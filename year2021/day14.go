package main

import (
	"math"
	"strings"
	"utils"
)

func polymerize(input string, n int) uint64 {
	pts := strings.Split(input, "\n\n")
	vtmpl := pts[0]
	d := make(map[string]byte)
	for _, line := range strings.Split(pts[1], "\n") {
		pt := strings.Split(line, " -> ")
		d[pt[0]] = pt[1][0]
	}
	cnts := make(map[string]uint64)
	for k := 0; k < len(vtmpl)-1; k++ {
		cnts[vtmpl[k:k+2]]++
	}
	for i := 0; i < n; i++ {
		cnts2 := make(map[string]uint64)
		for k, v := range cnts {
			rep := d[k]
			cnts2[string([]byte{k[0], rep})] += v
			cnts2[string([]byte{rep, k[1]})] += v
		}
		cnts = cnts2
	}
	lets := make(map[byte]uint64)
	for k, v := range cnts {
		lets[k[0]] += v
	}
	lets[vtmpl[len(vtmpl)-1]]++
	var (
		min uint64 = math.MaxUint64
		max uint64 = 0
	)
	for _, v := range lets {
		min = utils.Min(min, v)
		max = utils.Max(max, v)
	}
	return max - min
}

func Part1(input string) interface{} {
	return polymerize(input, 10)
}

func Part2(input string) interface{} {
	return polymerize(input, 40)
}
