package utils

import (
	"container/heap"
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
	n, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatalf("Failed to convert %v to integer: %v", input, err)
	}
	return n
}

func Uint32(input string) uint32 {
	n, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		log.Fatalf("Failed to convert %v to integer: %v", input, err)
	}
	return uint32(n)
}

func Uint64(input string) uint64 {
	n, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		log.Fatalf("Failed to convert %v to integer: %v", input, err)
	}
	return n
}

func IntBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Uint64Bool(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

type Integer interface {
	int | int32 | int64 | uint64
}

type Number interface {
	Integer | float32 | float64
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

func MapSum[K comparable, V Number](ns map[K]V) V {
	var sum V
	for _, v := range ns {
		sum += v
	}
	return sum
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

func CountOnes[T Integer](n T) T {
	var cnt T
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

func (a Coord) Mul(b Coord) Coord {
	return Coord{
		X: a.X*b.X - a.Y*b.Y,
		Y: a.X*b.Y + a.Y*b.X,
	}
}

func (a Coord) Sgn() Coord {
	return Coord{
		X: Sgn(a.X),
		Y: Sgn(a.Y),
	}
}

func (a Coord) Scale(n int) Coord {
	return Coord{
		X: a.X * n,
		Y: a.Y * n,
	}
}

func (a Coord) Pow(n int) Coord {
	if n == 0 {
		return Coord{X: 1, Y: 0}
	}
	if n&1 != 0 {
		return a.Mul(a.Pow(n - 1))
	}
	return a.Mul(a).Pow(n / 2)
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

const deBruijn64 = 0x03f79d71b4ca8b09

var deBruijn64Lookup = []byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

func countTrailingZeros(x uint64) int {
	return int(deBruijn64Lookup[((x&-x)*(deBruijn64))>>58])
}

type Bits struct {
	N uint64
	v int
}

func (b *Bits) Next() bool {
	if b.N == 0 {
		return false
	}
	b.v = countTrailingZeros(b.N)
	b.N &= b.N - 1
	return true
}

func (b *Bits) Get() int {
	return b.v
}

type pq struct {
	items []any
	less  func(a, b any) bool
}

func (q *pq) Len() int {
	return len(q.items)
}

func (q *pq) Less(i, j int) bool {
	return q.less(q.items[i], q.items[j])
}

func (q *pq) Swap(i, j int) {
	q.items[j], q.items[i] = q.items[i], q.items[j]
}

func (q *pq) Push(x any) {
	q.items = append(q.items, x)
}

func (q *pq) Pop() any {
	v := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return v
}

type PQueue[T any] struct {
	q pq
}

func NewPQueue[T any](less func(a, b T) bool) *PQueue[T] {
	return &PQueue[T]{q: pq{less: func(a, b any) bool {
		return less(a.(T), b.(T))
	}}}
}

func (q *PQueue[T]) Len() int {
	return q.q.Len()
}

func (q *PQueue[T]) Push(x T) {
	heap.Push(&q.q, x)
}

func (q *PQueue[T]) Pop() T {
	return heap.Pop(&q.q).(T)
}

// Fix, FixAll (Init), Remove

func Mod(a, b int64) int64 {
	m := a % b
	if m < 0 {
		if b < 0 {
			m -= b
		} else {
			m += b
		}
	}
	return m
}

func DivMod(a, b int64) (int64, int64) {
	d, m := a/b, a%b
	if m < 0 {
		if b < 0 {
			d++
			m -= b
		} else {
			d--
			m += b
		}
	}
	return d, m
}

func mulInv(a, b0 int64) int64 {
	b := b0
	var x0, x1 int64 = 0, 1
	if b == 1 {
		return 1
	}
	for a > 1 {
		q, r := DivMod(a, b)
		a, b = b, r
		x0, x1 = x1-q*x0, x0
	}
	if x1 < 0 {
		x1 += b0
	}
	return x1
}

func ChineseRemainder(as, ns []int64) int64 {
	var sum int64
	prod := Product(ns)
	for i := range as {
		p := prod / ns[i]
		sum += as[i] * mulInv(p, ns[i]) * p
	}
	return Mod(sum, prod)
}
