package main

import (
	"strings"
)

type cave struct {
	lowercase bool
	visited   int
	start     bool
	end       bool
	neighbors []*cave
}

func makeCave(name string) *cave {
	return &cave{
		lowercase: strings.ToLower(name) == name,
		start:     name == "start",
		end:       name == "end",
		neighbors: make([]*cave, 0),
	}
}

func parse(input string) *cave {
	caves := make(map[string]*cave)
	for _, line := range strings.Split(input, "\n") {
		v := strings.Split(line, "-")
		a, ok := caves[v[0]]
		if !ok {
			a = makeCave(v[0])
			caves[v[0]] = a
		}
		b, ok := caves[v[1]]
		if !ok {
			b = makeCave(v[1])
			caves[v[1]] = b
		}
		a.neighbors = append(a.neighbors, b)
		b.neighbors = append(b.neighbors, a)
	}
	return caves["start"]
}

func dfs(c *cave, canRevisit bool) int {
	if c.end {
		return 1
	} else if c.lowercase && c.visited > 0 {
		if !canRevisit || c.start {
			return 0
		}
		canRevisit = false
	}
	c.visited++
	var sum int
	for _, neighb := range c.neighbors {
		sum += dfs(neighb, canRevisit)
	}
	c.visited--
	return sum
}

func Part1(input string) interface{} {
	return dfs(parse(input), false)
}

func Part2(input string) interface{} {
	return dfs(parse(input), true)
}
