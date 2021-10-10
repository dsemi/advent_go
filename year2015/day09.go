package year2015

import (
	"advent/utils"
	"math"
	"regexp"
	"strings"
)

func allDists(input string) chan int {
	dists := make(map[string]map[string]int)
	re := regexp.MustCompile("(\\w+) to (\\w+) = (\\d+)")
	for _, line := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(line)
		if _, ok := dists[m[1]]; !ok {
			dists[m[1]] = make(map[string]int)
		}
		if _, ok := dists[m[2]]; !ok {
			dists[m[2]] = make(map[string]int)
		}
		n := utils.Int(m[3])
		dists[m[1]][m[2]] = n
		dists[m[2]][m[1]] = n
	}
	var keys []string
	for k := range dists {
		keys = append(keys, k)
	}
	c := make(chan int)
	go func() {
		for perm := range utils.PermutationsString(keys) {
			var dist int
			for i := range perm[1:] {
				dist += dists[perm[i]][perm[i+1]]
			}
			c <- dist
		}
		close(c)
	}()
	return c
}

func Day09Part1(input string) interface{} {
	min := math.MaxInt
	for d := range allDists(input) {
		min = utils.Min(min, d)
	}
	return min
}

func Day09Part2(input string) interface{} {
	var max int
	for d := range allDists(input) {
		max = utils.Max(max, d)
	}
	return max
}
