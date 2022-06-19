package year2021

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day02 struct{}

func (d *Day02) run(input string) (int, int, int) {
	horz, depth, aim := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		v := strings.Fields(line)
		cmd := v[0]
		n := utils.Int(v[1])
		if cmd == "forward" {
			horz += n
			depth += aim * n
		} else if cmd == "down" {
			aim += n
		} else if cmd == "up" {
			aim -= n
		}
	}
	return horz, depth, aim
}

func (d *Day02) Part1(input string) interface{} {
	horz, _, depth := d.run(input)
	return horz * depth
}

func (d *Day02) Part2(input string) interface{} {
	horz, depth, _ := d.run(input)
	return horz * depth
}

func init() {
	problems.Register(&Day02{})
}
