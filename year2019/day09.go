package main

import (
	"year2019/intcode"
)

func Part1(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 1
	return <-p.Output
}

func Part2(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 2
	return <-p.Output
}
