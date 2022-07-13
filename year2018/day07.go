package main

import (
	"fmt"
	"strings"
	"utils"
)

type Item struct {
	t int
	k rune
}

func run(input string, workers int) (string, int) {
	m := make(map[rune]int)
	cs := make(map[rune][]func())
	for _, line := range strings.Split(input, "\n") {
		var a, b rune
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin", &a, &b)
		if _, ok := m[a]; !ok {
			m[a] = 0
		}
		m[b]++
		cs[a] = append(cs[a], func() { m[b]-- })
	}
	var keys []rune
	for k := range m {
		keys = append(keys, k)
	}
	utils.Sort(keys)
	var ans strings.Builder
	var workQ []Item
	var time int
	for {
		for i := 0; i < len(keys) && workers != len(workQ); i++ {
			k := keys[i]
			if m[k] == 0 {
				m[k] = -1
				workQ = append(workQ, Item{t: time + int(k-4), k: k})
			}
		}
		if len(workQ) == 0 {
			break
		}
		for i := 1; i < len(workQ); i++ {
			if workQ[i-1].t < workQ[i].t {
				workQ[i-1], workQ[i] = workQ[i], workQ[i-1]
			}
		}
		min := workQ[len(workQ)-1]
		workQ = workQ[:len(workQ)-1]
		ans.WriteRune(min.k)
		time = min.t
		for _, f := range cs[min.k] {
			f()
		}
	}

	return ans.String(), time
}

func Part1(input string) interface{} {
	ans, _ := run(input, 1)
	return ans
}

func Part2(input string) interface{} {
	_, ans := run(input, 5)
	return ans
}
