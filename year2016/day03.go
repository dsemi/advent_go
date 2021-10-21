package year2016

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day03 struct{}

func (*Day03) parse(input string) [][3]int {
	var ts [][3]int
	for _, line := range strings.Split(input, "\n") {
		ts = append(ts, [3]int{})
		for i, field := range strings.Fields(line) {
			ts[len(ts)-1][i] = utils.Int(field)
		}
	}
	return ts
}

func (*Day03) valid(sides [3]int) bool {
	return sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0]
}

func (d *Day03) Part1(input string) interface{} {
	var cnt int
	for _, t := range d.parse(input) {
		if d.valid(t) {
			cnt++
		}
	}
	return cnt
}

func (d *Day03) Part2(input string) interface{} {
	ts := d.parse(input)
	var cnt int
	for i := 0; i+2 < len(ts); i += 3 {
		for j := 0; j < 3; j++ {
			if d.valid([3]int{ts[i][j], ts[i+1][j], ts[i+2][j]}) {
				cnt++
			}
		}
	}
	return cnt
}

func init() {
	problems.Register(&Day03{})
}
