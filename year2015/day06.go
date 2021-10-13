package year2015

import (
	"advent/types"
	"advent/utils"
	"regexp"
	"strings"
)

type Day06 struct{}

func runCommands(input string, turnOff func(int) int, turnOn func(int) int, toggle func(int) int) int {
	reg := regexp.MustCompile("(turn off|turn on|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)")
	var arr [1000][1000]int
	for _, line := range strings.Split(input, "\n") {
		m := reg.FindStringSubmatch(line)
		var f func(int) int
		if m[1] == "turn off" {
			f = turnOff
		} else if m[1] == "turn on" {
			f = turnOn
		} else {
			f = toggle
		}
		x0, y0, x1, y1 := utils.Int(m[2]), utils.Int(m[3]), utils.Int(m[4]), utils.Int(m[5])
		for y := y0; y <= y1; y++ {
			for x := x0; x <= x1; x++ {
				arr[y][x] = f(arr[y][x])
			}
		}
	}
	var total int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			total += arr[i][j]
		}
	}
	return total
}

func (Day06) Part1(input string) interface{} {
	return runCommands(input, func(x int) int { return 0 }, func(x int) int { return 1 }, func(x int) int { return x ^ 1 })
}

func (Day06) Part2(input string) interface{} {
	return runCommands(input, func(x int) int {
		if x == 0 {
			return 0
		} else {
			return x - 1
		}
	}, func(x int) int { return x + 1 }, func(x int) int { return x + 2 })
}

func init() {
	types.Register(Probs, Day06{})
}
