package main

import (
	"strings"
	"utils"
)

func ap(nums *utils.Stack[int64], ops *utils.Stack[byte]) {
	a := nums.Pop()
	b := nums.Pop()
	switch ops.Pop() {
	case '+':
		nums.Push(a + b)
	case '*':
		nums.Push(a * b)
	default:
		panic("Invalid op")
	}
}

func calc(s *strings.Reader, prec func(byte) uint8) int64 {
	nums := utils.NewStack[int64]()
	ops := utils.NewStack[byte]()
	for s.Len() > 0 {
		c, _ := s.ReadByte()
		if '0' <= c && c <= '9' {
			nums.Push(int64(c - '0'))
		} else if c == '(' {
			nums.Push(calc(s, prec))
		} else if c == ')' {
			break
		} else if c == '+' || c == '*' {
			if ops.Len() > 0 && prec(c) <= prec(ops.Peek()) {
				ap(nums, ops)
			}
			ops.Push(c)
		}
	}
	for ops.Len() > 0 {
		ap(nums, ops)
	}
	return nums.Pop()
}

func Part1(input string) interface{} {
	prec := func(byte) uint8 { return 1 }
	var sum int64
	for _, line := range strings.Split(input, "\n") {
		sum += calc(strings.NewReader(line), prec)
	}
	return sum
}

func Part2(input string) interface{} {
	prec := func(x byte) uint8 {
		if x == '+' {
			return 2
		}
		return 1
	}
	var sum int64
	for _, line := range strings.Split(input, "\n") {
		sum += calc(strings.NewReader(line), prec)
	}
	return sum
}
