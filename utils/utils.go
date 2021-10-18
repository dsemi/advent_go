package utils

import (
	"log"
	"sort"
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

func Min64(a, b int64) int64 {
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

func Sum(ns []int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}
	return sum
}

func Sum64(ns []int64) int64 {
	var sum int64
	for _, n := range ns {
		sum += n
	}
	return sum
}

func Product64(ns []int64) int64 {
	prod := int64(1)
	for _, n := range ns {
		prod *= n
	}
	return prod
}

func Last(c chan int64) int64 {
	var n int64
	for n = range c {
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

func Permutations64(xs []int64, callback func([]int64)) {
	ns := make([]int64, len(xs))
	copy(ns, xs)
	var f func(i int)
	f = func(i int) {
		if i > len(ns) {
			callback(ns)
		} else {
			f(i + 1)
			for j := i + 1; j < len(ns); j++ {
				ns[i], ns[j] = ns[j], ns[i]
				f(i + 1)
				ns[i], ns[j] = ns[j], ns[i]
			}
		}
	}
	f(0)
}

func Combinations(xs []int, n int, callback func([]int)) {
	arr := make([]int, n)
	var f func([]int, int)
	f = func(xs []int, n int) {
		if n == 0 {
			callback(arr)
		} else {
			for i, x := range xs {
				arr[n-1] = x
				f(xs[i+1:], n-1)
			}
		}
	}
	f(xs, n)
}

func Combinations64(xs []int64, n int, callback func([]int64)) {
	arr := make([]int64, n)
	var f func([]int64, int)
	f = func(xs []int64, n int) {
		if n == 0 {
			callback(arr)
		} else {
			for i, x := range xs {
				arr[n-1] = x
				f(xs[i+1:], n-1)
			}
		}
	}
	f(xs, n)
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

func Reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

type Coord struct {
	X, Y int
}

func (a Coord) Add(b Coord) Coord {
	return Coord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

type Counter struct {
	cnts map[rune]int
	keys []rune
}

func NewCounter(s string) *Counter {
	c := &Counter{
		cnts: make(map[rune]int),
	}
	for _, r := range s {
		c.Add(r)
	}
	return c
}

func (c *Counter) Add(r rune) {
	c.cnts[r]++
}

func (c *Counter) Runes() []rune {
	c.keys = []rune{}
	for k := range c.cnts {
		c.keys = append(c.keys, k)
	}
	sort.Sort(c)
	return c.keys
}

func (c *Counter) Less(i, j int) bool {
	a, b := c.cnts[c.keys[i]], c.cnts[c.keys[j]]
	if a == b {
		return c.keys[i] < c.keys[j]
	}
	return a > b
}

func (c *Counter) Swap(i, j int) {
	c.keys[i], c.keys[j] = c.keys[j], c.keys[i]
}

func (c *Counter) Len() int {
	return len(c.keys)
}
