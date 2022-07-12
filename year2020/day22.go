package main

import (
	"strings"
	"utils"
)

func parse(input string) ([]int, []int) {
	parts := strings.Split(input, "\n\n")
	ns := make([][]int, 2)
	for i := range ns {
		ns[i] = make([]int, 0)
		for _, n := range strings.Split(parts[i], "\n")[1:] {
			ns[i] = append(ns[i], utils.Int(n))
		}
	}
	return ns[0], ns[1]
}

func hash(xs, ys []int) string {
	var b strings.Builder
	for _, x := range xs {
		b.WriteByte(byte(x))
	}
	b.WriteByte(255)
	for _, y := range ys {
		b.WriteByte(byte(y))
	}
	return b.String()
}

func play(as, bs []int, p2, sub bool) (int, bool) {
	s := make(map[string]bool)
	if sub && utils.Maximum(as) > utils.Maximum(bs) {
		return 0, true
	}
	for len(as) > 0 && len(bs) > 0 {
		if p2 {
			key := hash(as, bs)
			if s[key] {
				return 0, true
			}
			s[key] = true
		}
		a := as[0]
		as = as[1:]
		b := bs[0]
		bs = bs[1:]
		p1Win := a > b
		if p2 && a <= len(as) && b <= len(bs) {
			var as2, bs2 []int
			_, p1Win = play(append(as2, as[:a]...), append(bs2, bs[:b]...), p2, true)
		}
		if p1Win {
			as = append(as, a, b)
		} else {
			bs = append(bs, b, a)
		}
	}
	xs := bs
	win := len(bs) == 0
	if win {
		xs = as
	}

	var sum int
	for i, x := range xs {
		sum += (len(xs) - i) * x
	}
	return sum, win
}

func Part1(input string) interface{} {
	a, b := parse(input)
	ans, _ := play(a, b, false, false)
	return ans
}

func Part2(input string) interface{} {
	a, b := parse(input)
	ans, _ := play(a, b, true, false)
	return ans
}
