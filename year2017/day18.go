package year2017

import (
	"strconv"
	"strings"
)

type Op int

const (
	Snd Op = iota
	Set
	Add
	Mul
	Mod
	Rcv
	Jgz
)

type ValT int

const (
	Lit ValT = iota
	Reg
)

type Val struct {
	t ValT
	v int
}

type Instr struct {
	op Op
	a  Val
	b  Val
}

type Sim struct {
	line    int
	reg     [26]int
	instrs  []Instr
	sends   int
	recvs   int
	waiting bool
}

func val(x string) Val {
	if i, err := strconv.Atoi(x); err == nil {
		return Val{t: Lit, v: i}
	}
	return Val{t: Reg, v: int(x[0] - 'a')}
}

func ParseInstrs(input string) Sim {
	var instrs []Instr
	for _, line := range strings.Split(input, "\n") {
		switch pts := strings.Fields(line); pts[0] {
		case "snd":
			instrs = append(instrs, Instr{op: Snd, a: val(pts[1])})
		case "set":
			instrs = append(instrs, Instr{op: Set, a: val(pts[1]), b: val(pts[2])})
		case "add":
			instrs = append(instrs, Instr{op: Add, a: val(pts[1]), b: val(pts[2])})
		case "mul":
			instrs = append(instrs, Instr{op: Mul, a: val(pts[1]), b: val(pts[2])})
		case "mod":
			instrs = append(instrs, Instr{op: Mod, a: val(pts[1]), b: val(pts[2])})
		case "rcv":
			instrs = append(instrs, Instr{op: Rcv, a: val(pts[1])})
		case "jgz":
			instrs = append(instrs, Instr{op: Jgz, a: val(pts[1]), b: val(pts[2])})
		}
	}
	return Sim{instrs: instrs}
}

func (s *Sim) Val(v Val) int {
	switch v.t {
	case Lit:
		return v.v
	case Reg:
		return s.reg[v.v]
	}
	panic("unreachable")
}

func (s *Sim) Run(in <-chan int, out chan<- int) {
	defer close(out)
	for ; s.line >= 0 && s.line < len(s.instrs); s.line++ {
		switch instr := s.instrs[s.line]; instr.op {
		case Snd:
			s.sends++
			out <- s.Val(instr.a)
		case Set:
			s.reg[instr.a.v] = s.Val(instr.b)
		case Add:
			s.reg[instr.a.v] += s.Val(instr.b)
		case Mul:
			s.reg[instr.a.v] *= s.Val(instr.b)
		case Mod:
			b := s.Val(instr.b)
			s.reg[instr.a.v] %= b
			if s.reg[instr.a.v] < 0 {
				s.reg[instr.a.v] += b
			}
		case Rcv:
			s.waiting = true
			v := <-in
			s.waiting = false
			s.recvs++
			s.reg[instr.a.v] = v
		case Jgz:
			if s.Val(instr.a) > 0 {
				s.line += s.Val(instr.b) - 1
			}
		}
	}
}

func Day18Part1(input string) interface{} {
	sim := ParseInstrs(input)
	in := make(chan int)
	out := make(chan int)
	go sim.Run(in, out)
	var v int
	for v = range out {
		if sim.waiting {
			if v == 0 {
				in <- 0
			} else {
				break
			}
		}
	}
	return v
}

func Day18Part2(input string) interface{} {
	sim0 := ParseInstrs(input)
	sim1 := ParseInstrs(input)
	sim1.reg[int('p'-'a')] = 1
	in := make(chan int, 100)
	out := make(chan int, 100)
	go sim0.Run(in, out)
	go sim1.Run(out, in)
	for !(sim0.waiting && sim1.waiting && sim0.sends == sim1.recvs && sim1.sends == sim0.recvs) {
	}
	return sim1.sends
}
