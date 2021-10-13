package year2015

import (
	"advent/types"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Day07 struct{}

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

func (Day07) Part1(input string) interface{} {
	d := make(map[string]func() uint16)
	val := func(x string) uint16 {
		if i, err := strconv.Atoi(x); err == nil {
			return uint16(i)
		}
		return d[x]()
	}
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Fields(line)
		if len(pts) == 3 {
			d[pts[2]] = lazy(func() uint16 {
				return val(pts[0])
			})
		} else if len(pts) == 4 {
			d[pts[3]] = lazy(func() uint16 {
				return ^val(pts[1])
			})
		} else if pts[1] == "AND" {
			d[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) & val(pts[2])
			})
		} else if pts[1] == "OR" {
			d[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) | val(pts[2])
			})
		} else if pts[1] == "LSHIFT" {
			d[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) << val(pts[2])
			})
		} else if pts[1] == "RSHIFT" {
			d[pts[4]] = lazy(func() uint16 {
				return val(pts[0]) >> val(pts[2])
			})
		}
	}
	return val("a")
}

func (d Day07) Part2(input string) interface{} {
	return d.Part1(input + fmt.Sprintf("\n%d -> b", d.Part1(input)))
}

func init() {
	types.Register(Probs, Day07{})
}
