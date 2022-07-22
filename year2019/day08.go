package main

import (
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	var cnts *utils.Counter[rune]
	for i := 0; i < len(input); i += 150 {
		t := utils.NewCounter([]rune(input[i : i+150]))
		if cnts == nil || t.Get('0') < cnts.Get('0') {
			cnts = t
		}
	}
	return cnts.Get('1') * cnts.Get('2')
}

func Part2(input string) interface{} {
	pts := make([]rune, 150)
	for i := range pts {
		pts[i] = '2'
	}
	for i := 0; i < len(input); i += 150 {
		for i, c := range input[i : i+150] {
			if pts[i] == '2' {
				pts[i] = c
			}
		}
	}
	var b strings.Builder
	for i := 0; i < len(pts); i += 25 {
		b.WriteRune('\n')
		for _, x := range pts[i : i+25] {
			if x == '0' {
				b.WriteRune(' ')
			} else {
				b.WriteRune('#')
			}
		}
	}
	return b.String()
}
