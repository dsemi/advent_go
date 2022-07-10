package main

import "strings"

type coord struct {
	w, x, y, z int
}

func solve(s string, p2 bool) int {
	onCubes := make(map[coord]bool)
	for y, line := range strings.Split(s, "\n") {
		for x, v := range line {
			if v != '#' {
				continue
			}
			onCubes[coord{x: x, y: y}] = true
		}
	}
	for i := 0; i < 6; i++ {
		m := make(map[coord]int)
		for pos := range onCubes {
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					for z := -1; z <= 1; z++ {
						for w := -1; w <= 1; w++ {
							if !p2 && w != 0 {
								continue
							}
							pos2 := coord{x: pos.x + x, y: pos.y + y, z: pos.z + z, w: pos.w + w}
							if pos != pos2 {
								m[pos2]++
							}
						}
					}
				}
			}
		}
		nextOn := make(map[coord]bool)
		for pos := range onCubes {
			if v := m[pos]; v == 2 || v == 3 {
				nextOn[pos] = true
			}
		}
		for pos, v := range m {
			if !onCubes[pos] && v == 3 {
				nextOn[pos] = true
			}
		}
		onCubes = nextOn
	}
	return len(onCubes)
}

func Part1(input string) interface{} {
	return solve(input, false)
}

func Part2(input string) interface{} {
	return solve(input, true)
}
