package main

import (
	"strings"
	"utils"
)

type board struct {
	grid [][]int
	done bool
}

func (b *board) isWinner() bool {
outer:
	for i := 0; i < len(b.grid); i++ {
		var r, c bool
		for j := 0; j < len(b.grid); j++ {
			r = r || b.grid[i][j] != -1
			c = c || b.grid[j][i] != -1
			if r && c {
				continue outer
			}
		}
		return true
	}
	return false
}

func winnerScores(input string) chan int {
	c := make(chan int)
	go func() {
		v := strings.Split(input, "\n\n")
		nums := v[0]
		var boards []*board
		for _, brd := range v[1:] {
			b := &board{}
			b.grid = make([][]int, 0)
			for r, line := range strings.Split(brd, "\n") {
				b.grid = append(b.grid, make([]int, 0))
				for _, v := range strings.Fields(line) {
					b.grid[r] = append(b.grid[r], utils.Int(v))
				}
			}
			boards = append(boards, b)
		}
		for _, num := range strings.Split(nums, ",") {
			n := utils.Int(num)
			for _, b := range boards {
				if b.done {
					continue
				}
				for i, row := range b.grid {
					for j, v := range row {
						if v == n {
							b.grid[i][j] = -1
						}
					}
				}
				if b.isWinner() {
					b.done = true
					var sum int
					for _, row := range b.grid {
						for _, v := range row {
							if v != -1 {
								sum += v
							}
						}
					}
					sum *= n
					c <- sum
				}
			}
		}
		close(c)
	}()
	return c
}

func Part1(input string) interface{} {
	return <-winnerScores(input)
}

func Part2(input string) interface{} {
	return utils.Last(winnerScores(input))
}
