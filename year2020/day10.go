package main

import (
	"strings"
	"utils"
)

func parseNums(input string) []int64 {
	var ns []int64
	for _, x := range strings.Split(input, "\n") {
		ns = append(ns, utils.Int64(x))
	}
	utils.Sort(ns)
	ns = append([]int64{0}, ns...)
	return append(ns, ns[len(ns)-1]+3)
}

func Part1(input string) interface{} {
	ns := parseNums(input)
	cnt := make(map[int64]int)
	for i := 1; i < len(ns); i++ {
		cnt[ns[i]-ns[i-1]]++
	}
	return cnt[1] * cnt[3]
}

func Part2(input string) interface{} {
	ns := parseNums(input)
	dp := make([]int64, ns[len(ns)-1]+1)
	dp[0] = 1
	for _, n := range ns[1:] {
		for i := n - 3; i < n; i++ {
			if i >= 0 {
				dp[n] += dp[i]
			}
		}
	}
	return dp[ns[len(ns)-1]]
}
