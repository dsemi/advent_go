package intcode

import (
	"math"
	"strconv"
	"strings"
)

type Program struct {
	idx     int64
	relBase int64
	mem     []int64
	Input   chan int64
	Output  chan int64
}

func parseInstrs(input string) []int64 {
	var instrs []int64
	for _, x := range strings.Split(input, ",") {
		var i int
		var err error
		if i, err = strconv.Atoi(x); err != nil {
			panic("Bad input")
		}
		instrs = append(instrs, int64(i))
	}
	return instrs
}

func New(input string) Program {
	return Program{
		mem:    parseInstrs(input),
		Input:  make(chan int64, 100),
		Output: make(chan int64, 100),
	}
}

func (p *Program) Copy() Program {
	prog := Program{
		idx:     p.idx,
		relBase: p.relBase,
		mem:     make([]int64, len(p.mem)),
		Input:   make(chan int64, 100),
		Output:  make(chan int64, 100),
	}
	copy(prog.mem, p.mem)
	return prog
}

func (p *Program) get(i int64) int64 {
	if i >= 0 && i < int64(len(p.mem)) {
		return p.mem[i]
	}
	return 0
}

func (p *Program) set(i int64, v int64) {
	if l := i - int64(len(p.mem)) + 1; l > 0 {
		s := make([]int64, l)
		p.mem = append(p.mem, s...)
	}
	p.mem[i] = v
}

func (p *Program) arg(i int64) int64 {
	mode := p.get(p.idx) / int64(math.Pow(10.0, float64(i+1))) % 10
	switch mode {
	case 0:
		return p.get(p.idx + i)
	case 1:
		return p.idx + i
	case 2:
		return p.get(p.idx+i) + p.relBase
	}
	panic("Unknown mode")
}

func (p *Program) Run() {
	defer close(p.Output)
	for {
		switch op := p.get(p.idx) % 100; op {
		case 1: // Add
			p.set(p.arg(3), p.get(p.arg(1))+p.get(p.arg(2)))
			p.idx += 4
		case 2: // Mul
			p.set(p.arg(3), p.get(p.arg(1))*p.get(p.arg(2)))
			p.idx += 4
		case 3: // Sav
			p.set(p.arg(1), <-p.Input)
			p.idx += 2
		case 4: // Out
			p.Output <- p.get(p.arg(1))
			p.idx += 2
		case 5: // Jit
			if p.get(p.arg(1)) != 0 {
				p.idx = p.get(p.arg(2))
			} else {
				p.idx += 3
			}
		case 6: // Jif
			if p.get(p.arg(1)) == 0 {
				p.idx = p.get(p.arg(2))
			} else {
				p.idx += 3
			}
		case 7: // Lt
			if p.get(p.arg(1)) < p.get(p.arg(2)) {
				p.set(p.arg(3), 1)
			} else {
				p.set(p.arg(3), 0)
			}
			p.idx += 4
		case 8: // Eql
			if p.get(p.arg(1)) == p.get(p.arg(2)) {
				p.set(p.arg(3), 1)
			} else {
				p.set(p.arg(3), 0)
			}
			p.idx += 4
		case 9: // Arb
			p.relBase += p.get(p.arg(1))
			p.idx += 2
		case 99: // Hlt
			return
		default:
			panic("Invalid op code")
		}
	}
}

func (p *Program) RunNoIo(a int64, b int64) int64 {
	p.mem[1] = a
	p.mem[2] = b
	p.Run()
	return p.mem[0]
}
