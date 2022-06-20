package main

import (
	"math"
	"utils"
)

func Part1(input string) interface{} {
	n := utils.Int(input)
	return 1 + 2*(n-int(math.Pow(2, math.Floor(math.Log2(float64(n))))))
}

func Part2(input string) interface{} {
	n := utils.Int(input)
	p3 := int(math.Pow(3, math.Floor(math.Log(float64(n))/math.Log(3))))
	ans := n - p3
	ans2 := ans + utils.Max(0, ans-p3)
	if ans2 == 0 {
		return p3
	} else {
		return ans
	}
}
