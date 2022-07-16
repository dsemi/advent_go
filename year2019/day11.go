package main

import (
	"math"
	"strings"
	"utils"
	"year2019/intcode"
)

func runRobot(prog *intcode.Program, t map[utils.Coord]int64) {
	pos, dir := utils.Coord{0, 0}, utils.Coord{0, -1}
	go func() {
		for {
			prog.Input <- t[pos]
			col := <-prog.Output
			d := <-prog.Output
			t[pos] = col
			if d == 1 {
				dir = dir.Mul(utils.Coord{0, 1})
			} else {
				dir = dir.Mul(utils.Coord{0, -1})
			}
			pos = pos.Add(dir)
		}
	}()
	prog.Run()
}

func Part1(input string) interface{} {
	p := intcode.New(input)
	m := make(map[utils.Coord]int64)
	runRobot(&p, m)
	return len(m)
}

func draw(points map[utils.Coord]int64) string {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for pt := range points {
		minX = utils.Min(minX, pt.X)
		minY = utils.Min(minY, pt.Y)
		maxX = utils.Max(maxX, pt.X)
		maxY = utils.Max(maxY, pt.Y)
	}
	var chrs strings.Builder
	for y := minY; y <= maxY; y++ {
		chrs.WriteRune('\n')
		for x := minX; x <= maxX; x++ {
			if points[utils.Coord{x, y}] == 0 {
				chrs.WriteRune(' ')
			} else {
				chrs.WriteRune('#')
			}
		}
	}
	return chrs.String()
}

func Part2(input string) interface{} {
	p := intcode.New(input)
	m := make(map[utils.Coord]int64)
	m[utils.Coord{0, 0}] = 1
	runRobot(&p, m)
	return draw(m)
}
