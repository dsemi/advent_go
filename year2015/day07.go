package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func lazy(f func() uint16) func() uint16 {
	var ans uint16
	var once sync.Once
	return func() uint16 {
		once.Do(func() {
			ans = f()
			f = nil
		})
		return ans
	}
}

func Part1(input string) interface{} {
	m := make(map[string]func() uint16)
	val := func(x string) uint16 {
		if i, err := strconv.Atoi(x); err == nil {
			return uint16(i)
		}
		return m[x]()
	}
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Fields(line)
		if len(pts) == 3 {
			m[pts[2]] = lazy(func() uint16 {
				return val(pts[0])
			})
		} else if len(pts) == 4 {
			m[pts[3]] = lazy(func() uint16 {
				return ^val(pts[1])
			})
		} else if pts[1] == "AND" {
			m[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) & val(pts[2])
			})
		} else if pts[1] == "OR" {
			m[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) | val(pts[2])
			})
		} else if pts[1] == "LSHIFT" {
			m[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) << val(pts[2])
			})
		} else if pts[1] == "RSHIFT" {
			m[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) >> val(pts[2])
			})
		}
	}
	return val("a")
}

func Part2(input string) interface{} {
	return Part1(input + fmt.Sprintf("\n%d -> b", Part1(input)))
}
