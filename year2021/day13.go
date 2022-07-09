package main

import (
	"fmt"
	"strings"
	"utils"
)

func parse(input string) (map[utils.Coord]bool, string) {
	pts := strings.Split(input, "\n\n")
	dots := make(map[utils.Coord]bool)
	for _, line := range strings.Split(pts[0], "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		dots[utils.Coord{x, y}] = true
	}
	return dots, pts[1]
}

func fold(paper map[utils.Coord]bool, instr string) map[utils.Coord]bool {
	var d rune
	var n int
	fmt.Sscanf(instr, "fold along %c=%d", &d, &n)
	newPaper := make(map[utils.Coord]bool)
	if d == 'x' {
		for k := range paper {
			newPaper[utils.Coord{utils.Min(k.X, 2*n-k.X), k.Y}] = true
		}
	} else {
		for k := range paper {
			newPaper[utils.Coord{k.X, utils.Min(k.Y, 2*n-k.Y)}] = true
		}
	}
	return newPaper
}

func Part1(input string) interface{} {
	paper, instrs := parse(input)
	return len(fold(paper, strings.Split(instrs, "\n")[0]))
}

func Part2(input string) interface{} {
	paper, instrs := parse(input)
	for _, instr := range strings.Split(instrs, "\n") {
		paper = fold(paper, instr)
	}
	var mx, my int
	for k := range paper {
		mx = utils.Max(mx, k.X)
		my = utils.Max(my, k.Y)
	}
	var display strings.Builder
	for y := 0; y <= my; y++ {
		display.WriteRune('\n')
		for x := 0; x <= mx; x++ {
			if paper[utils.Coord{x, y}] {
				display.WriteRune('#')
			} else {
				display.WriteRune(' ')
			}
		}
	}
	return display.String()
}
