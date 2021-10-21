package year2017

import "advent/problems"

type Day01 struct{}

func (*Day01) Part1(input string) interface{} {
	var sum int
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return sum
}

func (*Day01) Part2(input string) interface{} {
	var sum int
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+len(input)/2)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return sum
}

func init() {
	problems.Register(&Day01{})
}
