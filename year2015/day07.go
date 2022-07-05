package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	m := make(map[string]func() uint16)
	val := func(x string) uint16 {
		if i, err := strconv.Atoi(x); err == nil {
			return uint16(i)
		}
		return m[x]()
	}
	ops := map[string]func(a, b uint16) uint16{
		"->":     func(a, b uint16) uint16 { return a },
		"NOT":    func(a, b uint16) uint16 { return ^a },
		"AND":    func(a, b uint16) uint16 { return a & b },
		"OR":     func(a, b uint16) uint16 { return a | b },
		"LSHIFT": func(a, b uint16) uint16 { return a << b },
		"RSHIFT": func(a, b uint16) uint16 { return a >> b },
	}
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Fields(line)
		m[pts[len(pts)-1]] = func() uint16 {
			ans := ops[pts[utils.Abs(4-len(pts))]](val(pts[1-utils.Abs(4-len(pts))]), val(pts[len(pts)-3]))
			m[pts[len(pts)-1]] = func() uint16 { return ans }
			return ans
		}
	}
	return val("a")
}

func Part2(input string) interface{} {
	return Part1(input + fmt.Sprintf("\n%d -> b", Part1(input)))
}
