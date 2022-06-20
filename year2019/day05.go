package main

import (
	"year2019/intcode"
)

func Part1(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 1
	var x int64
	for x = range prog.Output {
	}
	return x
}

func Part2(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	prog.Input <- 5
	var x int64
	for x = range prog.Output {
	}
	return x
}
