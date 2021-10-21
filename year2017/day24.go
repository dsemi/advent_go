package year2017

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day24 struct{}

type Pipe struct {
	a, b int
	used bool
}

func (*Day24) parsePipes(input string) []*Pipe {
	var pipes []*Pipe
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, "/")
		pipes = append(pipes, &Pipe{
			a: utils.Int(pts[0]),
			b: utils.Int(pts[1]),
		})
	}
	return pipes
}

func (d *Day24) solve(input string) (int, int) {
	pipes := d.parsePipes(input)
	var max int
	for _, pipe := range pipes {
		max = utils.Max(max, utils.Max(pipe.a, pipe.b))
	}
	arrs := [][][]*Pipe{make([][]*Pipe, max+1), make([][]*Pipe, max+1)}
	for _, pipe := range pipes {
		arrs[0][pipe.a] = append(arrs[0][pipe.a], pipe)
		if pipe.a != pipe.b {
			arrs[1][pipe.b] = append(arrs[1][pipe.b], pipe)
		}
	}
	var part1, part2, length int
	var build func(int, int, int)
	build = func(port int, strength, len int) {
		if strength > part1 {
			part1 = strength
		}
		if len > length {
			length = len
			part2 = strength
		} else if len == length && strength > part2 {
			part2 = strength
		}

		for _, arr := range arrs {
			for _, pipe := range arr[port] {
				if !pipe.used {
					pipe.used = true
					build(pipe.a+pipe.b-port, strength+pipe.a+pipe.b, len+1)
					pipe.used = false
				}
			}
		}
	}
	build(0, 0, 0)
	return part1, part2
}

func (d *Day24) Part1(input string) interface{} {
	x, _ := d.solve(input)
	return x
}

func (d *Day24) Part2(input string) interface{} {
	_, x := d.solve(input)
	return x
}

func init() {
	problems.Register(&Day24{})
}
