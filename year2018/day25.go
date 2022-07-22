package main

import (
	"strings"
	"utils"
)

type node struct {
	pt     utils.Coord4
	parent *node
	rank   int
}

func newNode(p utils.Coord4) *node {
	n := &node{pt: p}
	n.parent = n
	return n
}

func find(n *node) *node {
	for n.parent != n {
		n = n.parent
	}
	return n
}

func union(x, y *node) {
	xRoot, yRoot := find(x), find(y)
	if xRoot == yRoot {
		return
	}
	if xRoot.rank < yRoot.rank {
		xRoot.parent = yRoot
	} else if xRoot.rank > yRoot.rank {
		yRoot.parent = xRoot
	} else {
		yRoot.parent = xRoot
		xRoot.rank++
	}
}

func parse(input string) []*node {
	nodes := make([]*node, 0)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, ",")
		nodes = append(nodes, newNode(utils.Coord4{
			utils.Int(pts[0]), utils.Int(pts[1]), utils.Int(pts[2]), utils.Int(pts[3])}))
	}
	return nodes
}

func constellations(pts []*node) int {
	for i, a := range pts {
		for _, b := range pts[i+1:] {
			if a.pt.Dist(b.pt) <= 3 {
				union(a, b)
			}
		}
	}
	m := make(map[*node]bool)
	for _, p := range pts {
		m[find(p)] = true
	}
	return len(m)
}

func Part1(input string) interface{} {
	return constellations(parse(input))
}

func Part2(input string) interface{} {
	return ""
}
