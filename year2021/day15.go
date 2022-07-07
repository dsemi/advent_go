package main

import (
	"container/heap"
	"math"
	"strings"
)

type pos struct {
	visited bool
	risk    int
	neighbs []*pos
}

func parse(input string) [][]int {
	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		grid = append(grid, row)
	}
	return grid
}

type PQueue struct {
	items []any
	less  func(a, b any) bool
}

func (q *PQueue) Len() int {
	return len(q.items)
}

func (q *PQueue) Less(i, j int) bool {
	return q.less(q.items[i], q.items[j])
}

func (q *PQueue) Swap(i, j int) {
	q.items[j], q.items[i] = q.items[i], q.items[j]
}

func (q *PQueue) Push(x any) {
	q.items = append(q.items, x)
}

func (q *PQueue) Pop() any {
	v := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return v
}

func dijkstra(ds [][]int) int {
	grid := make([][]*pos, len(ds))
	for i, row := range ds {
		grid[i] = make([]*pos, len(row))
		for j, n := range row {
			grid[i][j] = &pos{risk: n}
		}
	}
	for i, row := range grid {
		for j := range row {
			if i > 0 {
				grid[i][j].neighbs = append(grid[i][j].neighbs, grid[i-1][j])
			}
			if j > 0 {
				grid[i][j].neighbs = append(grid[i][j].neighbs, grid[i][j-1])
			}
			if i < len(grid)-1 {
				grid[i][j].neighbs = append(grid[i][j].neighbs, grid[i+1][j])
			}
			if j < len(grid[i])-1 {
				grid[i][j].neighbs = append(grid[i][j].neighbs, grid[i][j+1])
			}
		}
	}
	dist := make(map[*pos]int)
	for _, row := range grid {
		for _, p := range row {
			dist[p] = math.MaxInt
		}
	}
	dist[grid[0][0]] = 0
	q := &PQueue{less: func(a, b any) bool {
		return dist[a.(*pos)] < dist[b.(*pos)]
	}}
	heap.Push(q, grid[0][0])

	for q.Len() > 0 {
		u := heap.Pop(q).(*pos)
		u.visited = true
		if u == grid[len(grid)-1][len(grid[0])-1] {
			return dist[u]
		}

		for _, v := range u.neighbs {
			if v.visited {
				continue
			}
			alt := dist[u] + v.risk
			if alt < dist[v] && dist[u] != math.MaxInt {
				dist[v] = alt
				heap.Push(q, v)
			}
		}
	}
	return 0
}

func Part1(input string) interface{} {
	return dijkstra(parse(input))
}

func Part2(input string) interface{} {
	smallGrid := parse(input)
	grid := make([][]int, len(smallGrid)*5)
	for r := range grid {
		grid[r] = make([]int, len(smallGrid[0])*5)
		for c := range grid[r] {
			rd, rm := r/len(smallGrid), r%len(smallGrid)
			cd, cm := c/len(smallGrid[0]), c%len(smallGrid[0])
			grid[r][c] = (smallGrid[rm][cm]-1+rd+cd)%9 + 1
		}
	}
	return dijkstra(grid)
}
