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

type Number interface {
	int | int32 | int64 | uint64 | float32 | float64
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Min[T Number](a, b T) T {
	if b < a {
		return b
	}
	return a
}

func Max[T Number](a, b T) T {
	if b > a {
		return b
	}
	return a
}

func Sgn[T Number](n T) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func Minimum[T Number](ns []T) T {
	n := ns[0]
	for i := 1; i < len(ns); i++ {
		n = Min(n, ns[i])
	}
	return n
}

func Maximum[T Number](ns []T) T {
	n := ns[0]
	for i := 1; i < len(ns); i++ {
		n = Max(n, ns[i])
	}
	return n
}

func Extrema[T Number](ns []T) (T, T) {
	min := ns[0]
	max := ns[0]
	for _, v := range ns[1:] {
		min = Min(min, v)
		max = Max(max, v)
	}
	return min, max
}

func anyVal[K comparable, V any](ns map[K]V) V {
	for _, v := range ns {
		return v
	}
	panic("Empty collection")
}

func MapExtrema[K comparable, V Number](ns map[K]V) (V, V) {
	min := anyVal(ns)
	max := min
	for _, v := range ns {
		min = Min(min, v)
		max = Max(max, v)
	}
	return min, max
}

func Sum[T Number](ns []T) T {
	var sum T
	for _, n := range ns {
		sum += n
	}
	return sum
}

func Product[T Number](ns []T) T {
	var prod T = 1
	for _, n := range ns {
		prod *= n
	}
	return prod
}

func Last[T any](c chan T) T {
	var n T
	for n = range c {
	}
	return n
}

func Permutations[T any](ns []T) chan []T {
	c := make(chan []T)
	var f func(i int)
	f = func(i int) {
		if i > len(ns) {
			nns := make([]T, len(ns))
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

func Combinations[T any](xs []T, n int, callback func([]T)) {
	arr := make([]T, n)
	var f func([]T, int)
	f = func(xs []T, n int) {
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

type Sortable[T rune | int | int64] []T

func (s Sortable[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sortable[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sortable[T]) Len() int {
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

func (a Coord) Sub(b Coord) Coord {
	return Coord{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Coord) Sgn() Coord {
	return Coord{
		X: Sgn(a.X),
		Y: Sgn(a.Y),
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

func (c *Counter) Get(k rune) int {
	v, ok := c.cnts[k]
	if !ok {
		return 0
	}
	return v
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
