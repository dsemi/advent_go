package main

import (
	"strings"
	"utils"
)

func run(r map[string]int, input string) int {
	var instrs [][]string
	for _, line := range strings.Split(strings.ReplaceAll(input, ",", ""), "\n") {
		instrs = append(instrs, strings.Fields(line))
	}
	for i := 0; i >= 0 && i < len(instrs); i++ {
		instr := instrs[i]
		if instr[0] == "hlf" {
			r[instr[1]] /= 2
		} else if instr[0] == "tpl" {
			r[instr[1]] *= 3
		} else if instr[0] == "inc" {
			r[instr[1]]++
		} else if instr[0] == "jmp" {
			i += utils.Int(instr[1]) - 1
		} else if instr[0] == "jie" && r[instr[1]]%2 == 0 {
			i += utils.Int(instr[2]) - 1
		} else if instr[0] == "jio" && r[instr[1]] == 1 {
			i += utils.Int(instr[2]) - 1
		}
	}
	return r["b"]
}

func Part1(input string) interface{} {
	return run(map[string]int{"a": 0, "b": 0}, input)
}

func Part2(input string) interface{} {
	return run(map[string]int{"a": 1, "b": 0}, input)
}
