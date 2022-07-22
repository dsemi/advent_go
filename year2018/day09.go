package main

import (
	"fmt"
	"utils"
)

func parse(input string) (int, int) {
	var a, b int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &a, &b)
	return a, b
}

type Zipper struct {
	left, right []int
}

func NewZipper() *Zipper {
	return &Zipper{make([]int, 0), make([]int, 0)}
}

func (z *Zipper) Push(v int) {
	z.left = append(z.left, v)
}

func (z *Zipper) Pop() int {
	if len(z.left) == 0 {
		for len(z.right) > 0 {
			z.left = append(z.left, z.right[len(z.right)-1])
			z.right = z.right[:len(z.right)-1]
		}
	}
	v := z.left[len(z.left)-1]
	z.left = z.left[:len(z.left)-1]
	return v
}

func (z *Zipper) Left(n int) {
	for i := 0; i < n; i++ {
		if len(z.left) == 0 {
			for len(z.right) > 0 {
				z.left = append(z.left, z.right[len(z.right)-1])
				z.right = z.right[:len(z.right)-1]
			}
		}
		z.right = append(z.right, z.left[len(z.left)-1])
		z.left = z.left[:len(z.left)-1]
	}
}

func (z *Zipper) Right(n int) {
	for i := 0; i < n; i++ {
		if len(z.right) == 0 {
			for len(z.left) > 0 {
				z.right = append(z.right, z.left[len(z.left)-1])
				z.left = z.left[:len(z.left)-1]
			}
		}
		z.left = append(z.left, z.right[len(z.right)-1])
		z.right = z.right[:len(z.right)-1]
	}
}

func play(n, s int) int {
	m := make([]int, n)
	arr := NewZipper()
	arr.Push(0)
	for p := 1; p <= s; p++ {
		if p%23 != 0 {
			arr.Right(1)
			arr.Push(p)
			continue
		}
		arr.Left(7)
		v := arr.Pop()
		arr.Right(1)
		m[p%n] += p + v
	}
	return utils.Maximum(m)
}

func Part1(input string) interface{} {
	a, b := parse(input)
	return play(a, b)
}

func Part2(input string) interface{} {
	a, b := parse(input)
	return play(a, b*100)
}
