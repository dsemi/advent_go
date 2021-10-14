package utils

import (
	"log"
	"strconv"
)

func Int(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Failed to convert %v to integer: %v", input, err)
	}
	return n
}

func Int64(input string) int64 {
	return int64(Int(input))
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func Minimum(ns []int) int {
	n := ns[0]
	for i := 1; i < len(ns); i++ {
		n = Min(n, ns[i])
	}
	return n
}

func Maximum(ns []int) int {
	n := ns[0]
	for i := 1; i < len(ns); i++ {
		n = Max(n, ns[i])
	}
	return n
}

func PermutationsString(ns []string) chan []string {
	c := make(chan []string)
	var f func(i int)
	f = func(i int) {
		if i > len(ns) {
			nns := make([]string, len(ns))
			copy(nns, ns)
			c <- nns
			return
		}
		f(i + 1)
		for j := i + 1; j < len(ns); j++ {
			ns[i], ns[j] = ns[j], ns[i]
			f(i + 1)
			ns[i], ns[j] = ns[j], ns[i]
		}
	}
	go func() {
		defer close(c)
		f(0)
	}()
	return c
}

func Permutations(ns []int) chan []int {
	c := make(chan []int)
	var f func(i int)
	f = func(i int) {
		if i > len(ns) {
			nns := make([]int, len(ns))
			copy(nns, ns)
			c <- nns
			return
		}
		f(i + 1)
		for j := i + 1; j < len(ns); j++ {
			ns[i], ns[j] = ns[j], ns[i]
			f(i + 1)
			ns[i], ns[j] = ns[j], ns[i]
		}
	}
	go func() {
		defer close(c)
		f(0)
	}()
	return c
}

func Permutations64(ns []int64) chan []int64 {
	c := make(chan []int64)
	var f func(i int)
	f = func(i int) {
		if i > len(ns) {
			nns := make([]int64, len(ns))
			copy(nns, ns)
			c <- nns
			return
		}
		f(i + 1)
		for j := i + 1; j < len(ns); j++ {
			ns[i], ns[j] = ns[j], ns[i]
			f(i + 1)
			ns[i], ns[j] = ns[j], ns[i]
		}
	}
	go func() {
		defer close(c)
		f(0)
	}()
	return c
}

type SortRunes []rune

func (s SortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortRunes) Len() int {
	return len(s)
}

type SortInt64s []int64

func (s SortInt64s) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortInt64s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortInt64s) Len() int {
	return len(s)
}

func CountOnes(n int) int {
	var cnt int
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

func Partitions(n, t int, f func([]int)) {
	ns := make([]int, n)
	var recur func(int, int)
	recur = func(n, t int) {
		if n == 0 {
			ns[n] = t
			f(ns)
		} else {
			for x := 0; x <= t; x++ {
				ns[n] = x
				recur(n-1, t-x)
			}
		}
	}
	recur(n-1, t)
}
