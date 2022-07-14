package main

import (
	"strings"
	"utils"
)

type pair struct {
	a, b string
}

func parse(input string) []*pair {
	pairs := make([]*pair, 0)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, ")")
		pairs = append(pairs, &pair{pts[0], pts[1]})
	}
	return pairs
}

func Part1(input string) interface{} {
	t := make(map[string][]string)
	for _, p := range parse(input) {
		t[p.a] = append(t[p.a], p.b)
	}
	keys := []string{"COM"}
	var result int
	for depth := 0; len(keys) > 0; depth++ {
		result += depth * len(keys)
		keys2 := make([]string, 0)
		for _, k := range keys {
			for _, v := range t[k] {
				keys2 = append(keys2, v)
			}
		}
		keys = keys2
	}
	return result
}

func pathFromCom(t map[string]string, key string) []string {
	result := make([]string, 0)
	for v, ok := t[key]; ok; v, ok = t[v] {
		result = append(result, v)
	}
	utils.Reverse(result)
	return result
}

func Part2(input string) interface{} {
	t := make(map[string]string)
	for _, p := range parse(input) {
		t[p.b] = p.a
	}
	xs := pathFromCom(t, "YOU")
	ys := pathFromCom(t, "SAN")
	for i := 0; i < utils.Min(len(xs), len(ys)); i++ {
		if xs[i] != ys[i] {
			return len(xs) + len(ys) - 2*i
		}
	}
	panic("unreachable")
}
