package main

import (
	"strings"
	"utils"
)

type Snailfish struct {
	val         uint64
	left, right *Snailfish
}

func parse(i string) (*Snailfish, string) {
	if i[0] >= '0' && i[0] <= '9' {
		return &Snailfish{val: uint64(i[0] - '0')}, i[1:]
	}
	if i[0] != '[' {
		panic("Bad start parse")
	}
	fish := &Snailfish{}
	fish.left, i = parse(i[1:])
	if i[0] != ',' {
		panic("Bad mid parse")
	}
	fish.right, i = parse(i[1:])
	if i[0] != ']' {
		panic("Bad end parse")
	}
	return fish, i[1:]
}

func NewSnailfish(input string) *Snailfish {
	fish, rest := parse(input)
	if rest != "" {
		panic("leftovers")
	}
	return fish
}

func (s *Snailfish) Copy() *Snailfish {
	if s == nil {
		return nil
	}
	return &Snailfish{val: s.val, left: s.left.Copy(), right: s.right.Copy()}
}

func (s *Snailfish) exp(prev, next **uint64, depth int) bool {
	if s.left == nil {
		if *next != nil {
			s.val += **next
			return true
		}
		*prev = &s.val
		return false
	} else if *next != nil || depth != 4 {
		return s.left.exp(prev, next, depth+1) || s.right.exp(prev, next, depth+1)
	}
	if *prev != nil {
		**prev += s.left.val
	}
	*next = &s.right.val
	s.left = nil
	s.right = nil
	s.val = 0
	return false
}

func (s *Snailfish) Explode() bool {
	var prev, next *uint64
	return s.exp(&prev, &next, 0) || next != nil
}

func (s *Snailfish) Split() bool {
	if s.val > 9 {
		s.left = &Snailfish{val: s.val / 2}
		s.right = &Snailfish{val: (s.val + 1) / 2}
		s.val = 0
		return true
	}
	if s.left != nil {
		return s.left.Split() || s.right.Split()
	}
	return false
}

func (s *Snailfish) Magnitude() uint64 {
	if s.left == nil {
		return s.val
	}
	return 3*s.left.Magnitude() + 2*s.right.Magnitude()
}

func add(a, b *Snailfish) *Snailfish {
	x := &Snailfish{left: a, right: b}
	for x.Explode() || x.Split() {
	}
	return x
}

func Part1(input string) interface{} {
	lines := strings.Split(input, "\n")
	fish := NewSnailfish(lines[0])
	for _, line := range lines[1:] {
		fish = add(fish, NewSnailfish(line))
	}
	return fish.Magnitude()
}

func Part2(input string) interface{} {
	fish := make([]*Snailfish, 0)
	for _, line := range strings.Split(input, "\n") {
		fish = append(fish, NewSnailfish(line))
	}
	var max uint64
	for _, a := range fish {
		for _, b := range fish {
			max = utils.Max(max, add(a.Copy(), b.Copy()).Magnitude())
		}
	}
	return max
}
