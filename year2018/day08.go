package main

import (
	"strings"
	"utils"
)

type Tree struct {
	val      []int
	children []*Tree
}

func (t *Tree) Sum() int {
	sum := utils.Sum(t.val)
	for _, c := range t.children {
		sum += c.Sum()
	}
	return sum
}

func parseTree(q *utils.Queue[int]) *Tree {
	n, m := q.Pop(), q.Pop()
	nodes := make([]*Tree, 0)
	for i := 0; i < n; i++ {
		nodes = append(nodes, parseTree(q))
	}
	vals := make([]int, 0)
	for i := 0; i < m; i++ {
		vals = append(vals, q.Pop())
	}
	return &Tree{val: vals, children: nodes}
}

func parse(input string) *Tree {
	q := utils.NewQueue[int]()
	for _, pt := range strings.Fields(input) {
		q.Push(utils.Int(pt))
	}
	return parseTree(q)
}

func Part1(input string) interface{} {
	return parse(input).Sum()
}

func Part2(input string) interface{} {
	tree := parse(input)
	stack := utils.NewStack[*Tree]()
	stack.Push(tree)
	var result int
	for stack.Len() > 0 {
		node := stack.Pop()
		if len(node.children) == 0 {
			result += utils.Sum(node.val)
			continue
		}
		for _, i := range node.val {
			if i-1 < len(node.children) {
				stack.Push(node.children[i-1])
			}
		}
	}
	return result
}
