package main

import (
	"strings"
	"utils"
)

type ValType uint8

const (
	Reg ValType = iota
	Lit
)

type Value struct {
	t ValType
	v int64
}

type InstrType uint8

const (
	Inp InstrType = iota
	Add
	Mul
	Div
	Mod
	Eql
)

type Instr struct {
	t InstrType
	a int64
	b Value
}

type Prog struct {
	regs [4]int64
	pc   int
}

func value(x string) Value {
	if 'w' <= x[0] && x[0] <= 'z' {
		return Value{t: Reg, v: int64(x[0] - 'w')}
	}
	return Value{t: Lit, v: utils.Int64(x)}
}

func parse(input string) []Instr {
	instrs := make([]Instr, 0)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Fields(line)
		var instr Instr
		switch pts[0] {
		case "inp":
			instr = Instr{t: Inp, a: int64(pts[1][0] - 'w')}
		case "add":
			instr = Instr{t: Add, a: int64(pts[1][0] - 'w'), b: value(pts[2])}
		case "mul":
			instr = Instr{t: Mul, a: int64(pts[1][0] - 'w'), b: value(pts[2])}
		case "div":
			instr = Instr{t: Div, a: int64(pts[1][0] - 'w'), b: value(pts[2])}
		case "mod":
			instr = Instr{t: Mod, a: int64(pts[1][0] - 'w'), b: value(pts[2])}
		case "eql":
			instr = Instr{t: Eql, a: int64(pts[1][0] - 'w'), b: value(pts[2])}
		default:
			panic("Invalid input")
		}
		instrs = append(instrs, instr)
	}
	return instrs
}

func (p *Prog) val(v Value) int64 {
	switch v.t {
	case Reg:
		return p.regs[v.v]
	case Lit:
		return v.v
	}
	panic("impossible")
}

func (p *Prog) runNext(instrs []Instr, inp int64) bool {
	var a int64
	for {
		instr := instrs[p.pc]
		switch instr.t {
		case Inp:
			p.regs[instr.a] = inp
		case Add:
			p.regs[instr.a] += p.val(instr.b)
		case Mul:
			p.regs[instr.a] *= p.val(instr.b)
		case Div:
			if instr.a == 3 && instr.b.t == Lit {
				a = instr.b.v
			}
			p.regs[instr.a] /= p.val(instr.b)
		case Mod:
			p.regs[instr.a] %= p.val(instr.b)
		case Eql:
			p.regs[instr.a] = int64(utils.IntBool(p.regs[instr.a] == p.val(instr.b)))
		}
		p.pc++
		if p.pc >= len(instrs) || instrs[p.pc].t == Inp {
			break
		}
	}
	if a == 0 {
		panic("bad assumption")
	}
	return a != 26 || p.regs[1] == 0
}

type key struct {
	a, b int64
}

func dfs(vis map[key]bool, instrs []Instr, prog Prog, n, d int64, p2 bool) (int64, bool) {
	if d == 0 {
		return n, prog.regs[3] == 0
	}
	if ok := vis[key{prog.regs[3], d}]; ok {
		return 0, false
	}
	var i, diff, end int64 = 9, -1, 0
	if p2 {
		i, diff, end = 1, 1, 10
	}
	for ; i != end; i += diff {
		p := prog
		if !p.runNext(instrs, i) {
			continue
		}
		if v, ok := dfs(vis, instrs, p, n*10+i, d-1, p2); ok {
			return v, true
		}
	}
	vis[key{prog.regs[3], d}] = true
	return 0, false
}

func Part1(input string) interface{} {
	instrs := parse(input)
	ans, _ := dfs(make(map[key]bool), instrs, Prog{}, 0, 14, false)
	return ans
}

func Part2(input string) interface{} {
	instrs := parse(input)
	ans, _ := dfs(make(map[key]bool), instrs, Prog{}, 0, 14, true)
	return ans
}
