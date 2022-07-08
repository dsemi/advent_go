package main

import (
	"fmt"
	"utils"
)

func parse(input string) (uint64, uint64) {
	var p1, p2 uint64
	fmt.Sscanf(input, "Player 1 starting position: %d\nPlayer 2 starting position: %d", &p1, &p2)
	return p1 - 1, p2 - 1
}

func Part1(input string) interface{} {
	p1, p2 := parse(input)
	var n, p1Score, p2Score uint64
	for ; p2Score < 1000; n += 3 {
		p1 = (p1 + 3*n + 6) % 10
		p1Score += p1 + 1
		p2, p1 = p1, p2
		p2Score, p1Score = p1Score, p2Score
	}
	return utils.Min(p1Score, p2Score) * n
}

type pair struct {
	a, b uint64
}

var probs = [...]pair{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

type key struct {
	a, b, c, d uint64
}

func solve(cache map[key]pair, p1, p2, s1, s2 uint64) (uint64, uint64) {
	if s1 >= 21 {
		return 1, 0
	}
	if s2 >= 21 {
		return 0, 1
	}
	if v, ok := cache[key{p1, p2, s1, s2}]; ok {
		return v.a, v.b
	}
	var ans pair
	for _, p := range probs {
		d, n := p.a, p.b
		newP1 := (p1 + d) % 10
		x1, y1 := solve(cache, p2, newP1, s2, s1+newP1+1)
		ans.a += n * y1
		ans.b += n * x1
	}
	cache[key{p1, p2, s1, s2}] = ans
	return ans.a, ans.b
}

func Part2(input string) interface{} {
	p1, p2 := parse(input)
	x, y := solve(make(map[key]pair), p1, p2, 0, 0)
	return utils.Max(x, y)
}
