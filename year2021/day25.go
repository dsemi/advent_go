package main

import "strings"

const height = 137

type Cucumbers struct {
	c [height][4]uint64
}

func (a Cucumbers) or(b Cucumbers) Cucumbers {
	var result Cucumbers
	for r := 0; r < height; r++ {
		for i := 0; i < 4; i++ {
			result.c[r][i] = a.c[r][i] | b.c[r][i]
		}
	}
	return result
}

func (a Cucumbers) and(b Cucumbers) Cucumbers {
	var result Cucumbers
	for r := 0; r < height; r++ {
		for i := 0; i < 4; i++ {
			result.c[r][i] = a.c[r][i] & b.c[r][i]
		}
	}
	return result
}

func (a Cucumbers) xor(b Cucumbers) Cucumbers {
	var result Cucumbers
	for r := 0; r < height; r++ {
		for i := 0; i < 4; i++ {
			result.c[r][i] = a.c[r][i] ^ b.c[r][i]
		}
	}
	return result
}

func (a *Cucumbers) shiftUp() Cucumbers {
	var result Cucumbers
	result.c[height-1] = a.c[0]
	for r := 1; r < height; r++ {
		result.c[r-1] = a.c[r]
	}
	return result
}

func (a *Cucumbers) shiftDown() Cucumbers {
	var result Cucumbers
	result.c[0] = a.c[height-1]
	for r := 1; r < height; r++ {
		result.c[r] = a.c[r-1]
	}
	return result
}

func (a *Cucumbers) shiftLeft() Cucumbers {
	var result Cucumbers
	for r := 0; r < height; r++ {
		result.c[r][0] = (a.c[r][0] >> 1) | (a.c[r][1] << 63)
		result.c[r][1] = (a.c[r][1] >> 1) | (a.c[r][2] << 63)
		result.c[r][2] = (a.c[r][2] >> 1) | (a.c[r][0] << 10)
		result.c[r][2] &= 0x7ff
		result.c[r][3] = 0

	}
	return result
}

func (a *Cucumbers) shiftRight() Cucumbers {
	var result Cucumbers
	for r := 0; r < height; r++ {
		result.c[r][0] = (a.c[r][0] << 1) | (a.c[r][2] >> 10)
		result.c[r][1] = (a.c[r][1] << 1) | (a.c[r][0] >> 63)
		result.c[r][2] = (a.c[r][2] << 1) | (a.c[r][1] >> 63)
		result.c[r][2] &= 0x7ff
		result.c[r][3] = 0
	}
	return result
}

func advanceRight(d, r *Cucumbers) Cucumbers {
	result := r.shiftRight()
	blocked := result.and(d.or(*r))
	return result.xor(blocked).or(blocked.shiftLeft())
}

func advanceDown(d, r *Cucumbers) Cucumbers {
	result := d.shiftDown()
	blocked := result.and(d.or(*r))
	return result.xor(blocked).or(blocked.shiftUp())
}

func toMask(count int, inp string) (uint64, uint64) {
	var dMask, rMask uint64
	for i, c := range inp {
		if i == count {
			break
		}
		if c == 'v' {
			dMask |= 1 << i
		} else if c == '>' {
			rMask |= 1 << i
		}
	}
	return dMask, rMask
}

func Part1(input string) interface{} {
	var d, r Cucumbers
	for row, line := range strings.Split(input, "\n") {
		d.c[row][0], r.c[row][0] = toMask(64, line)
		d.c[row][1], r.c[row][1] = toMask(64, line[64:])
		d.c[row][2], r.c[row][2] = toMask(11, line[128:])
		d.c[row][3], r.c[row][3] = 0, 0
	}
	for cnt := 1; ; cnt++ {
		nextR := advanceRight(&d, &r)
		done := nextR == r
		r = nextR
		nextD := advanceDown(&d, &r)
		if done && nextD == d {
			return cnt
		}
		d = nextD
	}
}

func Part2(input string) interface{} {
	return ""
}
