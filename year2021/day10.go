package main

import (
	"strings"
	"utils"
)

var (
	pair = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	score = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

type block struct {
	corrupted bool
	score     int
}

func process(line string) block {
	stack := make([]rune, 0)
	for _, c := range line {
		if strings.ContainsRune("([{<", c) {
			stack = append(stack, pair[c])
			continue
		}
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v != c {
			return block{corrupted: true, score: score[c]}
		}
	}
	var sc int
	for len(stack) > 0 {
		sc *= 5
		sc += strings.IndexRune(")]}>", stack[len(stack)-1]) + 1
		stack = stack[:len(stack)-1]
	}
	return block{score: sc}
}

func Part1(input string) interface{} {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		if bl := process(line); bl.corrupted {
			sum += bl.score
		}
	}
	return sum
}

func Part2(input string) interface{} {
	scores := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if bl := process(line); !bl.corrupted {
			scores = append(scores, bl.score)
		}
	}
	utils.Sort(scores)
	return scores[len(scores)/2]
}
