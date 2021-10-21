package year2015

import "advent/problems"

type Day01 struct{}

func (*Day01) Part1(input string) interface{} {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func (*Day01) Part2(input string) interface{} {
	floor := 0
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}

func init() {
	problems.Register(&Day01{})
}
