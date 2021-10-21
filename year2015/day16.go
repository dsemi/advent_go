package year2015

import (
	"advent/problems"
	"advent/utils"
	"regexp"
	"strings"
)

type Day16 struct{}

var sue = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func (*Day16) findSue(input string, f func(string, int) bool) int {
	reg := regexp.MustCompile("(\\w+): (\\d+)")
OUTER:
	for i, line := range strings.Split(input, "\n") {
		for _, m := range reg.FindAllStringSubmatch(line, -1) {
			if !f(m[1], utils.Int(m[2])) {
				continue OUTER
			}
		}
		return i + 1
	}
	return -1
}

func (d *Day16) Part1(input string) interface{} {
	return d.findSue(input, func(k string, v int) bool {
		return sue[k] == v
	})
}

func (d *Day16) Part2(input string) interface{} {
	return d.findSue(input, func(k string, v int) bool {
		if k == "cats" {
			return v > 7
		} else if k == "pomeranians" {
			return v < 3
		} else if k == "goldfish" {
			return v < 5
		} else if k == "trees" {
			return v > 3
		}
		return sue[k] == v
	})
}

func init() {
	problems.Register(&Day16{})
}
