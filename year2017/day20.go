package main

import (
	"fmt"
	"strings"
	"utils"
)

type Particle struct {
	pos, vel, acc utils.Coord3
}

func parse(input string) []*Particle {
	particles := make([]*Particle, 0)
	for _, line := range strings.Split(input, "\n") {
		p := &Particle{}
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.pos.X, &p.pos.Y, &p.pos.Z, &p.vel.X, &p.vel.Y, &p.vel.Z, &p.acc.X, &p.acc.Y, &p.acc.Z)
		particles = append(particles, p)
	}
	return particles
}

func Part1(input string) interface{} {
	return utils.ArgMinBy(parse(input), func(p *Particle) int {
		return p.acc.Dist(utils.Coord3{0, 0, 0})
	})
}

func Part2(input string) interface{} {
	ps := parse(input)
	for i := 0; i < 100; i++ {
		tbl := make(map[utils.Coord3]int)
		for _, p := range ps {
			p.vel = p.vel.Add(p.acc)
			p.pos = p.pos.Add(p.vel)
			tbl[p.pos]++
		}
		j := 0
		for _, p := range ps {
			if tbl[p.pos] <= 1 {
				ps[j] = p
				j++
			}
		}
		ps = ps[:j]
	}
	return len(ps)
}
