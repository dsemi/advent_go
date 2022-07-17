package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

func guardSleepFreqs(input string) map[int][]int {
	result := make(map[int][]int)
	lines := strings.Split(input, "\n")
	sort.Strings(lines)
	var guard, lastM, st int
	for _, line := range lines {
		var y, m, d, h, min, guardNo int
		if n, _ := fmt.Sscanf(line, "[%d-%d-%d %d:%d] Guard #%d begins shift", &y, &m, &d, &h, &min, &guardNo); n == 6 {
			if lastM > 0 {
				for i := lastM; i < 60; i++ {
					result[guard][i] += st
				}
			}
			guard = guardNo
			st = 0
			lastM = 0
			if _, ok := result[guard]; !ok {
				result[guard] = make([]int, 60)
			}
		} else {
			for i := lastM; i < min; i++ {
				result[guard][i] += st
			}
			st ^= 1
			lastM = min
		}
	}
	for i := lastM; i < 60; i++ {
		result[guard][i] += st
	}
	return result
}

func Part1(input string) interface{} {
	sleepFreqs := guardSleepFreqs(input)
	var max, maxSum int
	for k, v := range sleepFreqs {
		if sum := utils.Sum(v); sum > maxSum {
			max, maxSum = k, sum
		}
	}
	return max * utils.ArgMax(sleepFreqs[max])
}

func Part2(input string) interface{} {
	var maxV, maxI, max int
	for k, v := range guardSleepFreqs(input) {
		for i, x := range v {
			if x > maxV {
				maxV, maxI, max = x, i, k
			}
		}
	}
	return maxI * max
}
