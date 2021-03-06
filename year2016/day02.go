package main

import (
	"strings"
	"utils"
)

func run(input string, pad []string) string {
	d := make(map[utils.Coord]string)
	var xy utils.Coord
	for y, line := range pad {
		for x, c := range strings.Fields(line) {
			if c != "." {
				coord := utils.Coord{x, y}
				d[coord] = c
				if c == "5" {
					xy = coord
				}
			}
		}
	}
	var result string
	for _, line := range strings.Split(input, "\n") {
		for _, c := range line {
			dir := func() utils.Coord {
				switch c {
				case 'U':
					return utils.Coord{0, -1}
				case 'D':
					return utils.Coord{0, 1}
				case 'L':
					return utils.Coord{-1, 0}
				case 'R':
					return utils.Coord{1, 0}
				default:
					panic("Bad dir")
				}
			}()
			if _, ok := d[xy.Add(dir)]; ok {
				xy = xy.Add(dir)
			}
		}
		result += d[xy]
	}
	return result
}

func Part1(input string) interface{} {
	return run(input,
		[]string{
			"1 2 3",
			"4 5 6",
			"7 8 9",
		})
}

func Part2(input string) interface{} {
	return run(input,
		[]string{
			". . 1 . .",
			". 2 3 4 .",
			"5 6 7 8 9",
			". A B C .",
			". . D . .",
		})
}
