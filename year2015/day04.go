package main

import (
	"crypto/md5"
	"strconv"
)

func solve(input string, n byte) int {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] <= n {
			return i
		}
	}
}

func Part1(input string) interface{} {
	return solve(input, 15)
}

func Part2(input string) interface{} {
	return solve(input, 0)
}
