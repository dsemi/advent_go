package year2019

import (
	"advent/types"
	"strconv"
	"strings"
)

type Day04 struct{}

func solve(n int, f func(int) bool) bool {
	prev := n % 10
	c := 1
	b := false
	n /= 10
	for n != 0 {
		m := n % 10
		if m == prev {
			c++
		} else if m > prev {
			return false
		} else {
			b = b || f(c)
			c = 1
			prev = m
		}
		n /= 10
	}
	return b || f(c)
}

func numValid(input string, f func(int) bool) int {
	var cnt int
	pts := strings.Split(input, "-")
	var lo int
	var hi int
	var err error
	if lo, err = strconv.Atoi(pts[0]); err != nil {
		panic("Bad input")
	}
	if hi, err = strconv.Atoi(pts[1]); err != nil {
		panic("Bad input")
	}
	for i := lo; i <= hi; i++ {
		if f(i) {
			cnt++
		}
	}
	return cnt
}

func (Day04) Part1(input string) interface{} {
	return numValid(input, func(n int) bool {
		return solve(n, func(x int) bool {
			return x >= 2
		})
	})
}

func (Day04) Part2(input string) interface{} {
	return numValid(input, func(n int) bool {
		return solve(n, func(x int) bool {
			return x == 2
		})
	})
}

func init() {
	types.Register(Probs, Day04{})
}
