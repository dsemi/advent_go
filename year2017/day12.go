package main

import (
	"strings"
	"utils"
)

type Pipe struct {
	connects []*Pipe
	visited  bool
}

func parse(input string) []*Pipe {
	pipes := make([]*Pipe, 0)
	for _ = range strings.Split(input, "\n") {
		pipes = append(pipes, &Pipe{})
	}
	for i, line := range strings.Split(input, "\n") {
		for _, n := range strings.Split(strings.Split(line, " <-> ")[1], ", ") {
			pipes[i].connects = append(pipes[i].connects, pipes[utils.Int(n)])
		}
	}
	return pipes
}

func traverse(p *Pipe) int {
	if p.visited {
		return 0
	}
	p.visited = true
	cnt := 1
	for _, c := range p.connects {
		cnt += traverse(c)
	}
	return cnt
}

func Part1(input string) interface{} {
	return traverse(parse(input)[0])
}

func Part2(input string) interface{} {
	pipes := parse(input)
	var cnt int
	for _, p := range pipes {
		if traverse(p) > 0 {
			cnt++
		}
	}
	return cnt
}
