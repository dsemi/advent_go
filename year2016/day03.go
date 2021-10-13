package year2016

import (
	"advent/types"
	"advent/utils"
	"strings"
)

type Day03 struct{}

func parse(input string) [][3]int {
	var ts [][3]int
	for _, line := range strings.Split(input, "\n") {
		ts = append(ts, [3]int{})
		for i, field := range strings.Fields(line) {
			ts[len(ts)-1][i] = utils.Int(field)
		}
	}
	return ts
}

func valid(sides [3]int) bool {
	return sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0]
}

func (Day03) Part1(input string) interface{} {
	var cnt int
	for _, t := range parse(input) {
		if valid(t) {
			cnt++
		}
	}
	return cnt
}

func (Day03) Part2(input string) interface{} {
	ts := parse(input)
	var cnt int
	for i := 0; i+2 < len(ts); i += 3 {
		for j := 0; j < 3; j++ {
			if valid([3]int{ts[i][j], ts[i+1][j], ts[i+2][j]}) {
				cnt++
			}
		}
	}
	return cnt
}

func init() {
	types.Register(Probs, Day03{})
}
