package year2016

import (
	"advent/types"
	"advent/utils"
	"strings"
)

type Day02 struct{}

func run(input string, pad []string) string {
	d := make(map[utils.Coord]string)
	var xy utils.Coord
	for y, line := range pad {
		for x, c := range strings.Fields(line) {
			if c != "." {
				coord := utils.Coord{X: x, Y: y}
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
					return utils.Coord{X: 0, Y: -1}
				case 'D':
					return utils.Coord{X: 0, Y: 1}
				case 'L':
					return utils.Coord{X: -1, Y: 0}
				case 'R':
					return utils.Coord{X: 1, Y: 0}
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

func (Day02) Part1(input string) interface{} {
	return run(input,
		[]string{
			"1 2 3",
			"4 5 6",
			"7 8 9",
		})
}

func (Day02) Part2(input string) interface{} {
	return run(input,
		[]string{
			". . 1 . .",
			". 2 3 4 .",
			"5 6 7 8 9",
			". A B C .",
			". . D . .",
		})
}

func init() {
	types.Register(Probs, Day02{})
}
