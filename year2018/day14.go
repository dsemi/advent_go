package main

import (
	"strings"
	"utils"
)

type step struct {
	elf1, elf2 uint64
	idx        int
	rs         []byte
}

func newStep() *step {
	return &step{rs: []byte{3, 7}, elf2: 1}
}

func (s *step) next() byte {
	if s.idx >= len(s.rs) {
		elf1Score, elf2Score := s.rs[s.elf1], s.rs[s.elf2]
		tot := elf1Score + elf2Score
		if tot >= 10 {
			s.rs = append(s.rs, 1, tot%10)
		} else {
			s.rs = append(s.rs, tot)
		}
		s.elf1 = (uint64(elf1Score) + s.elf1 + 1) % uint64(len(s.rs))
		s.elf2 = (uint64(elf2Score) + s.elf2 + 1) % uint64(len(s.rs))
	}
	ans := s.rs[s.idx] + '0'
	s.idx++
	return ans
}

func Part1(input string) interface{} {
	n := utils.Int(input)
	steps := newStep()
	for i := 0; i < n; i++ {
		steps.next()
	}
	var b strings.Builder
	for i := 0; i < 10; i++ {
		b.WriteByte(steps.next())
	}
	return b.String()
}

func Part2(input string) interface{} {
	rs := make([]byte, 0)
	steps := newStep()
outer:
	for i := 1; ; i++ {
		rs = append(rs, steps.next())
		if len(rs) >= len(input) {
			for j := len(input); j > 0; j-- {
				if input[len(input)-j] != rs[len(rs)-j] {
					continue outer
				}
			}
			return i - len(input)
		}
	}
}
