package year2015

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day23 struct{}

func (Day23) run(r map[string]int, input string) int {
	s := strings.ReplaceAll(input, ",", "")
	var instrs [][]string
	for _, line := range strings.Split(s, "\n") {
		instrs = append(instrs, strings.Fields(line))
	}
	for i := 0; i >= 0 && i < len(instrs); i++ {
		instr := instrs[i]
		switch instr[0] {
		case "hlf":
			r[instr[1]] /= 2
		case "tpl":
			r[instr[1]] *= 3
		case "inc":
			r[instr[1]]++
		case "jmp":
			i += utils.Int(instr[1]) - 1
		case "jie":
			if r[instr[1]]%2 == 0 {
				i += utils.Int(instr[2]) - 1
			}
		case "jio":
			if r[instr[1]] == 1 {
				i += utils.Int(instr[2]) - 1
			}
		}
	}
	return r["b"]
}

func (d Day23) Part1(input string) interface{} {
	return d.run(map[string]int{"a": 0, "b": 0}, input)
}

func (d Day23) Part2(input string) interface{} {
	return d.run(map[string]int{"a": 1, "b": 0}, input)
}

func init() {
	problems.Register(Day23{})
}
