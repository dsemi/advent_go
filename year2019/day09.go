package year2019

import "advent/year2019/intcode"

func Day09Part1(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 1
	return <-p.Output
}

func Day09Part2(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 2
	return <-p.Output
}
