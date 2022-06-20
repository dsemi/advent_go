package main

import (
	"strconv"
	"strings"
)

func Part1(input string) interface{} {
	var total int
	for _, line := range strings.Split(input, "\n") {
		var i int
		var err error
		if i, err = strconv.Atoi(line); err != nil {
			panic("Bad input")
		}
		total += i/3 - 2
	}
	return total
}

func Part2(input string) interface{} {
	var total int
	for _, line := range strings.Split(input, "\n") {
		var i int
		var err error
		if i, err = strconv.Atoi(line); err != nil {
			panic("Bad input")
		}
		for fuel := i/3 - 2; fuel > 0; {
			total += fuel
			fuel = fuel/3 - 2
		}
	}
	return total
}
