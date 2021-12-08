package year2019

import (
	"advent/problems"
	"advent/year2019/intcode"
)

type Day23 struct{}

type Network struct {
	computers []intcode.Program
	x, y      int
}

// func newNetwork(input string) *Network {
// 	p := intcode.New(input)
// 	n := &Network{}
// 	for i := 0; i < 50; i++ {
// 		prog := p.Copy()
// 		prog.Input <- int64(i)
// 		n.computers = append(n.computers, prog)
// 	}
// 	return n
// }

// func (n *Network) Run() {
// 	for {
// 		empty := true
// 		packets := make([][3]int64, len(n.computers))
// 		for j, comp := range n.computers {
// 			for i := 0; i < len(packets); i++ {
// 				packets[j][i] = <-comp.Output
// 				empty = false
// 			}
// 		}
// 		if
// 	}
// }

func (*Day23) Part1(input string) interface{} {
	return input
}

func (*Day23) Part2(input string) interface{} {
	return ""
}

func init() {
	problems.Register(&Day23{})
}
