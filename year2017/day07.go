package main

import (
	"sort"
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	inp := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		inp = append(inp, strings.Split(line, " -> "))
	}
	inp = utils.Transpose(inp)
	c := make(map[string]bool)
	for _, x := range inp[1] {
		for _, y := range strings.Split(x, ", ") {
			c[y] = true
		}
	}
	s := make(map[string]bool)
	for _, x := range inp[0] {
		s[strings.Fields(x)[0]] = true
	}
	for x := range s {
		if !c[x] {
			return x
		}
	}
	panic("unreachable")
}

type Node struct {
	weight   int64
	children []string
}

func findImbalance(m map[string]*Node, curr string) (int64, bool) {
	node := m[curr]
	if len(node.children) == 0 {
		return node.weight, false
	}
	wts := make([]int64, 0)
	for _, x := range node.children {
		w, b := findImbalance(m, x)
		if b {
			return w, b
		}
		wts = append(wts, w)
	}
	count := make(map[int64]int)
	for _, w := range wts {
		count[w]++
	}
	if len(count) == 1 {
		return node.weight + utils.Sum(wts), false
	}
	cts := make([]int64, 0)
	for c := range count {
		cts = append(cts, c)
	}
	sort.Slice(cts, func(i, j int) bool { return count[cts[i]] > count[cts[j]] })
	anomaly := cts[len(cts)-1]
	expected := cts[0]
	for i, v := range wts {
		if v == anomaly {
			ans := expected - anomaly + m[node.children[i]].weight
			return ans, true
		}
	}
	panic("Could not find imbalance")
}

func Part2(input string) interface{} {
	m := make(map[string]*Node)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		pts := strings.SplitN(parts[0], " (", 2)
		n, w := pts[0], pts[1]
		node := &Node{weight: utils.Int64(w[:len(w)-1])}
		if len(parts) > 1 {
			node.children = strings.Split(parts[1], ", ")
		}
		m[n] = node
	}
	root := Part1(input).(string)
	ans, _ := findImbalance(m, root)
	return ans
}
