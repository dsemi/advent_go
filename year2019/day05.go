package year2019

import "advent/year2019/intcode"

func Day05Part1(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 1
	var x int64
	for x = range prog.Output {
	}
	return x
}

func Day05Part2(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 5
	var x int64
	for x = range prog.Output {
	}
	return x
}
