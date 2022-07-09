package main

import (
	"strings"
	"utils"
)

type InstrType uint8

const (
	Acc InstrType = iota
	Jmp
	Nop
)

type Instr struct {
	t InstrType
	v int64
}

func parse(input string) []Instr {
	instrs := make([]Instr, 0)
	for _, line := range strings.Split(input, "\n") {
		w := strings.Fields(line)
		var instr Instr
		switch w[0] {
		case "acc":
			instr = Instr{t: Acc, v: utils.Int64(w[1])}
		case "jmp":
			instr = Instr{t: Jmp, v: utils.Int64(w[1])}
		case "nop":
			instr = Instr{t: Nop, v: utils.Int64(w[1])}
		default:
			panic("Invalid instruction")
		}
		instrs = append(instrs, instr)
	}
	return instrs
}

func runProg(prog []Instr) (int64, bool) {
	visited := make(map[int]bool)
	var acc int64
	for i := 0; 0 <= i && i < len(prog); i++ {
		if visited[i] {
			return acc, false
		}
		visited[i] = true
		n := prog[i].v
		switch prog[i].t {
		case Acc:
			acc += n
		case Jmp:
			i += int(n) - 1
		}
	}
	return acc, true
}

func Part1(input string) interface{} {
	ans, _ := runProg(parse(input))
	return ans
}

func flip(prog []Instr, i int) {
	switch prog[i].t {
	case Jmp:
		prog[i].t = Nop
	case Nop:
		prog[i].t = Jmp
	}
}

func Part2(input string) interface{} {
	prog := parse(input)
	for i := range prog {
		flip(prog, i)
		if ans, fin := runProg(prog); fin {
			return ans
		}
		flip(prog, i)
	}
	panic("unreachable")
}
