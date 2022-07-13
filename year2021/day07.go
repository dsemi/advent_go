package main

import (
	"math"
	"strings"
	"utils"
)

func parse(input string) []int {
	var ns []int
	for _, n := range strings.Split(input, ",") {
		ns = append(ns, utils.Int(n))
	}
	return ns
}

func Part1(input string) interface{} {
	ns := parse(input)
	utils.Sort(ns)
	var med int
	if len(ns)%2 == 0 {
		med = (ns[len(ns)/2-1] + ns[len(ns)/2]) / 2
	} else {
		med = ns[len(ns)/2]
	}
	var sum int
	for _, n := range ns {
		sum += utils.Abs(n - med)
	}
	return sum
}

func g(n int) int {
	return n * (n + 1) / 2
}

func Part2(input string) interface{} {
	ns := parse(input)
	mean := float64(utils.Sum(ns)) / float64(len(ns))
	var sum1, sum2 int
	for _, n := range ns {
		sum1 += g(utils.Abs(int(float64(n) - math.Floor(mean))))
		sum2 += g(utils.Abs(int(float64(n) - math.Ceil(mean))))
	}
	return utils.Min(sum1, sum2)
}
