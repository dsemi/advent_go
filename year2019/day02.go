package year2019

import "advent/year2019/intcode"

func Day02Part1(input string) interface{} {
	prog := intcode.New(input)
	return prog.RunNoIo(12, 2)
}

func Day02Part2(input string) interface{} {
	prog := intcode.New(input)
	for noun := int64(0); noun < 100; noun++ {
		for verb := int64(0); verb < 100; verb++ {
			prog2 := prog.Copy()
			if prog2.RunNoIo(noun, verb) == 19690720 {
				return 100*noun + verb
			}
		}
	}
	panic("No solution")
}
