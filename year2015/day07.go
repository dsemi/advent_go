package main

import (
	"fmt"
	"strconv"
	"strings"
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
		if len(pts) < 5 {
			m[pts[len(pts)-1]] = func() uint16 {
				ans := ops[pts[4-len(pts)]](val(pts[len(pts)-3]), 0)
				m[pts[len(pts)-1]] = func() uint16 { return ans }
				return ans
			}
		} else {
			m[pts[len(pts)-1]] = func() uint16 {
				ans := ops[pts[1]](val(pts[0]), val(pts[2]))
				m[pts[len(pts)-1]] = func() uint16 { return ans }
				return ans
			}
		}
	}
	return val("a")
}

func Part2(input string) interface{} {
	return Part1(input + fmt.Sprintf("\n%d -> b", Part1(input)))
}
