package year2019

import (
	"advent/problems"
	"advent/year2019/intcode"
)

type Day05 struct{}

func (Day05) Part1(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 1
	var x int64
	for x = range prog.Output {
	}
	return x
}

func (Day05) Part2(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 5
	var x int64
	for x = range prog.Output {
	}
	return x
}

func init() {
	problems.Register(Day05{})
}
