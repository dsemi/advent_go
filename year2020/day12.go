package main

import (
	"strings"
	"utils"
)

func travel(input string, x, y int, moveWay bool) int {
	st := []utils.Coord{{0, 0}, {x, y}}
	idx := utils.IntBool(moveWay)
	for _, line := range strings.Split(input, "\n") {
		n := utils.Int(line[1:])
		switch line[0] {
		case 'N':
			st[idx].Y += n
		case 'S':
			st[idx].Y -= n
		case 'E':
			st[idx].X += n
		case 'W':
			st[idx].X -= n
		case 'R':
			st[1] = st[1].Mul(utils.Coord{0, -1}.Pow(n / 90))
		case 'L':
			st[1] = st[1].Mul(utils.Coord{0, 1}.Pow(n / 90))
		case 'F':
			st[0] = st[0].Add(st[1].Scale(n))
		default:
			panic("Invalid cmd")
		}
	}
	return utils.Abs(st[0].X) + utils.Abs(st[0].Y)
}

func Part1(input string) interface{} {
	return travel(input, 1, 0, false)
}

func Part2(input string) interface{} {
	return travel(input, 10, 1, true)
}
