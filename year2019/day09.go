package year2019

import (
	"advent/problems"
	"advent/year2019/intcode"
)

type Day09 struct{}

func (Day09) Part1(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 1
	return <-p.Output
}

func (Day09) Part2(input string) interface{} {
	p := intcode.New(input)
	go p.Run()
	p.Input <- 2
	return <-p.Output
}

func init() {
	problems.Register(Day09{})
}
