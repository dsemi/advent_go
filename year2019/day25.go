package main

import (
	"regexp"
	"year2019/intcode"
)

const instrs = `north
east
take astrolabe
south
take space law space brochure
north
west
north
north
north
north
take weather machine
north
take antenna
west
south
`

func Part1(input string) interface{} {
	prog := intcode.New(input)
	go prog.Run()
	go func() {
		for _, c := range instrs {
			prog.Input <- int64(c)
		}
	}()
	var s string
	for c := range prog.Output {
		s += string(c)
	}
	reg := regexp.MustCompile("\\d+")
	var res string
	for _, res = range reg.FindAllString(s, -1) {
	}
	return res
}

func Part2(input string) interface{} {
	return ""
}
