package main

import (
	"fmt"
	"strings"
)

type pair struct {
	n    int
	name string
}

func parseBags(input string) map[string][]pair {
	result := make(map[string][]pair)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.SplitN(line, " bags contain ", 2)
		outerBag, innerBags := pts[0], pts[1]
		bags := make([]pair, 0)
		for _, bag := range strings.Split(innerBags, ", ") {
			if bag == "no other bags." {
				continue
			}
			var n int
			var name1, name2 string
			fmt.Sscanf(bag, "%d %s %s", &n, &name1, &name2)
			bags = append(bags, pair{n, name1 + " " + name2})
		}
		result[outerBag] = bags
	}
	return result
}

func Part1(input string) interface{} {
	m := parseBags(input)
	rev := make(map[string][]string)
	for k, v := range m {
		for _, es := range v {
			rev[es.name] = append(rev[es.name], k)
		}
	}
	stack := make([]string, len(rev["shiny gold"]))
	copy(stack, rev["shiny gold"])
	visited := make(map[string]bool)
	var ans int
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[v] {
			visited[v] = true
			ans++
			if len(rev[v]) > 0 {
				stack = append(stack, rev[v]...)
			}
		}
	}
	return ans
}

func countBags(m map[string][]pair, k string) int {
	var sum int
	for _, p := range m[k] {
		sum += p.n + p.n*countBags(m, p.name)
	}
	return sum
}

func Part2(input string) interface{} {
	return countBags(parseBags(input), "shiny gold")
}
