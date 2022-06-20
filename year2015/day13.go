package main

import (
	"regexp"
	"strings"
	"utils"
)

var reg = regexp.MustCompile("(\\w+) would (gain|lose) (\\d+) happiness units by sitting next to (\\w+)\\.")

func parseHappiness(input string) [][]int {
	d := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		m := reg.FindStringSubmatch(line)
		n := utils.Int(m[3])
		if m[2] == "lose" {
			n *= -1
		}
		if _, ok := d[m[1]]; !ok {
			d[m[1]] = make(map[string]int)
		}
		if _, ok := d[m[4]]; !ok {
			d[m[4]] = make(map[string]int)
		}
		d[m[1]][m[4]] += n
		d[m[4]][m[1]] += n
	}
	var keys []string
	for k := range d {
		keys = append(keys, k)
	}
	arr := make([][]int, len(keys))
	for i := range keys {
		arr[i] = make([]int, len(keys))
	}
	for i, k1 := range keys {
		for j, k2 := range keys {
			arr[i][j] = d[k1][k2]
		}
	}
	return arr
}

func maxHappiness(d [][]int, p2 bool) int {
	ids := make([]int, len(d))
	for i := range ids {
		ids[i] = i
	}
	var max int
	for perm := range utils.Permutations(ids) {
		var sum int
		if !p2 {
			sum += d[perm[0]][perm[len(perm)-1]]
		}
		for i := range perm[1:] {
			sum += d[perm[i]][perm[i+1]]
		}
		max = utils.Max(max, sum)
	}
	return max
}

func Part1(input string) interface{} {
	return maxHappiness(parseHappiness(input), false)
}

func Part2(input string) interface{} {
	return maxHappiness(parseHappiness(input), true)
}
