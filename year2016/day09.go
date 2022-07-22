package main

import (
	"fmt"
	"strings"
)

func decompressedLen(f func(string) int64, input string) int64 {
	i := strings.IndexByte(input, '(')
	if i == -1 {
		return int64(len(input))
	}
	input = input[i:]
	var dataLen, repeat int64
	fmt.Sscanf(input, "(%dx%d)", &dataLen, &repeat)
	rest := input[strings.IndexByte(input, ')')+1:]
	return int64(i) + repeat*f(rest[:dataLen]) + decompressedLen(f, rest[dataLen:])
}

func Part1(input string) interface{} {
	return decompressedLen(func(s string) int64 { return int64(len(s)) }, input)
}

func Part2(input string) interface{} {
	return decompressedLen(func(s string) int64 { return Part2(s).(int64) }, input)
}
