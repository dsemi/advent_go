package main

import (
	"encoding/hex"
	"strings"
	"utils"
)

func reverse(v []int, lo, hi int) {
	l := len(v)
	for lo < hi {
		v[lo%l], v[hi%l] = v[hi%l], v[lo%l]
		lo++
		hi--
	}
}

func hash(n int, lens []int) []int {
	result := make([]int, 256)
	for i := range result {
		result[i] = i
	}
	var pos, skipSize int
	for i := 0; i < n; i++ {
		for _, l := range lens {
			reverse(result, pos, pos+l-1)
			pos += l + skipSize
			skipSize++
		}
	}
	return result
}

func Part1(input string) interface{} {
	lens := make([]int, 0)
	for _, x := range strings.Split(input, ",") {
		lens = append(lens, utils.Int(x))
	}
	res := hash(1, lens)
	return res[0] * res[1]
}

func Part2(input string) interface{} {
	lens := make([]int, 0)
	for _, x := range input {
		lens = append(lens, int(x))
	}
	lens = append(lens, 17, 31, 73, 47, 23)
	res := hash(64, lens)
	bytes := make([]byte, 0)
	for i := 0; i < len(res); i += len(res) / 16 {
		var b byte
		for _, a := range res[i : i+len(res)/16] {
			b ^= byte(a)
		}
		bytes = append(bytes, b)
	}
	return hex.EncodeToString(bytes)
}
