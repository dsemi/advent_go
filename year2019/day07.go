package year2019

import (
	"advent/utils"
	"advent/year2019/intcode"
)

func chain(p intcode.Program, phases []int64, cycle bool) chan int64 {
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

func Day07Part1(input string) interface{} {
	p := intcode.New(input)
	c := utils.Permutations64([]int64{0, 1, 2, 3, 4})
	var v int64
	for perm := range c {
		xc := chain(p, perm, false)
		if x := <-xc; x > v {
			v = x
		}
	}
	return v
}

func Day07Part2(input string) interface{} {
	p := intcode.New(input)
	c := utils.Permutations64([]int64{5, 6, 7, 8, 9})
	var v int64
	for perm := range c {
		xc := chain(p, perm, true)
		var x int64
		for x = range xc {
		}
		if x > v {
			v = x
		}
	}
	return v
}
