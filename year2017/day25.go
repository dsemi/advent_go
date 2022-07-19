package main

import (
	"fmt"
	"io"
	"strings"
	"utils"
)

type Dir int

const (
	left Dir = iota
	right

	tmpl0 = `Begin in state %c.
Perform a diagnostic checksum after %d steps.`

	tmpl = `  If the current value is %d:
    - Write the value %d.
    - Move one slot to the %s
    - Continue with state %c.`
)

type Rule struct {
	write int
	dir   Dir
	state int
}

func parse(input string) (int, int, [][2]Rule) {
	r := strings.NewReader(input)
	var start rune
	var steps int
	fmt.Fscanf(r, tmpl0, &start, &steps)
	r.Seek(1, io.SeekCurrent)
	rules := make([][2]Rule, 0)
	for i := 0; i < 6; i++ {
		var st rune
		fmt.Fscanf(r, "\nIn state %c:", &st)
		r.Seek(1, io.SeekCurrent)
		var branches [2]Rule
		for j := 0; j < 2; j++ {
			var dir string
			var x, write int
			var state rune
			fmt.Fscanf(r, tmpl, &x, &write, &dir, &state)
			r.Seek(1, io.SeekCurrent)
			branches[j] = Rule{
				write: write,
				state: int(state - 'A'),
			}
			if dir == "right." {
				branches[j].dir = right
			}
		}
		rules = append(rules, branches)
	}
	return int(start - 'A'), steps, rules
}

func Part1(input string) interface{} {
	state, steps, rules := parse(input)
	tape := make([]int, 1)
	var i int
	for j := 0; j < steps; j++ {
		rule := rules[state][tape[i]]
		tape[i] = rule.write
		if rule.dir == left {
			if i == 0 {
				tape = append([]int{0}, tape...)
			} else {
				i--
			}
		} else {
			i++
			if i >= len(tape) {
				tape = append(tape, 0)
			}
		}
		state = rule.state
	}
	return utils.Sum(tape)
}

func Part2(input string) interface{} {
	return ""
}
