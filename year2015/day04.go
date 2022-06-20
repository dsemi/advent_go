package main

import (
	"crypto/md5"
	"strconv"
)

func sum(s string, i int) [16]byte {
	return md5.Sum([]byte(s + strconv.Itoa(i)))
}

func Part1(input string) interface{} {
	i := 0
	for hash := sum(input, i); hash[0] != 0 || hash[1] != 0 || hash[2] > 15; hash = sum(input, i) {
		i++
	}
	return i
}

func Part2(input string) interface{} {
	i := 0
	for hash := sum(input, i); hash[0] != 0 || hash[1] != 0 || hash[2] != 0; hash = sum(input, i) {
		i++
	}
	return i
}
