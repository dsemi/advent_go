package year2019

import (
	"advent/problems"
	"advent/utils"
	"advent/year2019/intcode"
)

type Day07 struct{}

func (*Day07) chain(p intcode.Program, phases []int64, cycle bool) chan int64 {
	c := make(chan int64)
	var progs []intcode.Program
	var prev chan int64
	for _, phase := range phases {
		prog2 := p.Copy()
		if prev != nil {
			prog2.Input = prev
		}
		prev = prog2.Output
		prog2.Input <- phase
		progs = append(progs, prog2)
	}
	progs[0].Input <- 0
	for i := range progs {
		go progs[i].Run()
	}
	go func() {
		defer close(c)
		for v := range progs[len(progs)-1].Output {
			c <- v
			if cycle {
				progs[0].Input <- v
			}
		}
	}()
	return c
}

func (d *Day07) Part1(input string) interface{} {
	p := intcode.New(input)
	var v int64
	for perm := range utils.Permutations([]int64{0, 1, 2, 3, 4}) {
		xc := d.chain(p, perm, false)
		if x := <-xc; x > v {
			v = x
		}
	}
	return v
}

func (d *Day07) Part2(input string) interface{} {
	p := intcode.New(input)
	var v int64
	for perm := range utils.Permutations([]int64{5, 6, 7, 8, 9}) {
		xc := d.chain(p, perm, true)
		if x := utils.Last(xc); x > v {
			v = x
		}
	}
	return v
}

func init() {
	problems.Register(&Day07{})
}
