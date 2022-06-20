package main

import (
	"utils"
)

func midpt(x, y int) int {
	return (x + y) / 2
}

func corners() chan int {
	c := make(chan int)
	go func() {
		v := 1
		c <- v
		for i := 1; ; i++ {
			v += i
			c <- v
			v += i
			c <- v
		}
	}()
	return c
}

func Part1(input string) interface{} {
	n := utils.Int(input)
	var ns []int
	for x := range corners() {
		ns = append(ns, x)
		if x >= n {
			break
		}
	}
	a := ns[len(ns)-1]
	b := ns[len(ns)-2]
	c := ns[len(ns)-3]
	return b - midpt(b, c) + utils.Abs(n-midpt(a, b))
}

func spiralPath() chan int {
	dirs := []utils.Coord{
		utils.Coord{1, 0},
		utils.Coord{0, 1},
		utils.Coord{-1, 0},
		utils.Coord{0, -1},
	}
	idx := 0
	pos := utils.Coord{0, 0}
	tbl := make(map[utils.Coord]int)
	tbl[utils.Coord{0, 0}] = 1
	c := make(chan int)
	go func() {
		for i := 1; ; i++ {
			for k := 0; k < 2; k++ {
				for j := 0; j < i; j++ {
					pos = pos.Add(dirs[idx])
					var val int
					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if x != 0 || y != 0 {
								val += tbl[pos.Add(utils.Coord{x, y})]
							}
						}
					}
					tbl[pos] = val
					c <- val
				}
				idx = (idx + 1) % 4
			}
		}
	}()
	return c
}

func Part2(input string) interface{} {
	n := utils.Int(input)
	for x := range spiralPath() {
		if x > n {
			return x
		}
	}
	panic("No solution")
}
