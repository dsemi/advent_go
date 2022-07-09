package main

import (
	"fmt"
	"strings"
	"utils"
)

type bit struct {
	i int
	c rune
}

type Cmd struct {
	mask []bit
	r, v uint64
}

func parse(input string) []Cmd {
	var mask []bit
	res := make([]Cmd, 0)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = make([]bit, 0)
			for i, x := range strings.Fields(line)[2] {
				mask = append(mask, bit{35 - i, x})
			}
		} else {
			var r, v uint64
			fmt.Sscanf(line, "mem[%d] = %d", &r, &v)
			res = append(res, Cmd{mask, r, v})
		}
	}
	return res
}

func Part1(input string) interface{} {
	cmds := parse(input)
	m := make(map[uint64]uint64)
	for _, cmd := range cmds {
		v := cmd.v
		for _, b := range cmd.mask {
			switch b.c {
			case '1':
				v |= 1 << b.i
			case '0':
				v &= ^(1 << b.i)
			}
		}
		m[cmd.r] = v
	}
	return utils.MapSum(m)
}

func setVals(m map[uint64]uint64, xs []bit, r, v uint64) {
	if len(xs) == 0 {
		m[r] = v
		return
	}
	i, c := xs[0].i, xs[0].c
	switch c {
	case '1':
		setVals(m, xs[1:], r|(1<<i), v)
	case '0':
		setVals(m, xs[1:], r, v)
	case 'X':
		setVals(m, xs[1:], r|(1<<i), v)
		setVals(m, xs[1:], r & ^(1<<i), v)
	default:
		panic("Invalid bit")
	}
}

func Part2(input string) interface{} {
	m := make(map[uint64]uint64)
	for _, cmd := range parse(input) {
		setVals(m, cmd.mask, cmd.r, cmd.v)
	}
	return utils.MapSum(m)
}
