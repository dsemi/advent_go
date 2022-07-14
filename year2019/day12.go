package main

import (
	"fmt"
	"strings"
	"utils"
)

type Moon struct {
	pos, vel []int64
}

func (m *Moon) Clone() *Moon {
	m2 := &Moon{
		pos: make([]int64, len(m.pos)),
		vel: make([]int64, len(m.vel)),
	}
	copy(m2.pos, m.pos)
	copy(m2.vel, m.vel)
	return m2
}

func (m *Moon) Equals(m2 *Moon) bool {
	if len(m.pos) != len(m2.pos) || len(m.vel) != len(m2.vel) {
		return false
	}
	for i := range m.pos {
		if m.pos[i] != m2.pos[i] {
			return false
		}
	}
	for i := range m.vel {
		if m.vel[i] != m2.vel[i] {
			return false
		}
	}
	return true
}

func parse(input string) []*Moon {
	moons := make([]*Moon, 0)
	for _, line := range strings.Split(input, "\n") {
		var x, y, z int64
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, &Moon{
			pos: []int64{x, y, z},
			vel: []int64{0, 0, 0},
		})
	}
	return moons
}

func step(moons []*Moon) {
	for i := range moons {
		for j := range moons {
			for x := range moons[i].pos {
				moons[i].vel[x] += int64(utils.Compare(moons[j].pos[x], moons[i].pos[x]))
			}
		}
	}
	for _, moon := range moons {
		for x := range moon.pos {
			moon.pos[x] += moon.vel[x]
		}
	}
}

func findCycle(moons []*Moon) uint64 {
	ms := make([]*Moon, len(moons))
	for i := range ms {
		ms[i] = moons[i].Clone()
	}
outer:
	for i := uint64(1); ; i++ {
		step(ms)
		for j := range ms {
			if !ms[j].Equals(moons[j]) {
				continue outer
			}
		}
		return i
	}
}

func Part1(input string) interface{} {
	m := parse(input)
	for i := 0; i < 1000; i++ {
		step(m)
	}
	var sum int64
	for _, v := range m {
		var posSum, velSum int64
		for x := range v.pos {
			posSum += utils.Abs(v.pos[x])
			velSum += utils.Abs(v.vel[x])
		}
		sum += posSum * velSum
	}
	return sum
}

func Part2(input string) interface{} {
	moons := parse(input)
	var ans uint64 = 1
	for i := 0; i < 3; i++ {
		ms := make([]*Moon, len(moons))
		for j := range ms {
			ms[j] = &Moon{
				pos: []int64{moons[j].pos[i]},
				vel: []int64{moons[j].vel[i]},
			}
		}
		ans = utils.Lcm(ans, findCycle(ms))
	}
	return ans
}
