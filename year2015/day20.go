package year2015

import (
	"advent/types"
	"advent/utils"
)

type Day20 struct{}

var primes = []int{2, 3, 5, 7, 11, 13}

func solve(goal, primeIndex int) int {
	if primeIndex < 0 {
		return goal
	}
	p := primes[primeIndex]
	pPower, pSum := 1, 1
	best := solve(goal, primeIndex-1)
	for pSum < goal {
		pPower *= p
		pSum += pPower
		subgoal := (goal + pSum - 1) / pSum
		best = utils.Min(best, pPower*solve(subgoal, primeIndex-1))
	}
	return best
}

func (Day20) Part1(input string) interface{} {
	n := utils.Int(input)
	return solve(n/10, len(primes)-1)
}

func (Day20) Part2(input string) interface{} {
	n := utils.Int(input)
	vec := make([]int, 1000000)
	for i := 1; i < len(vec); i++ {
		for c, j := 0, i; c < 50 && j < len(vec); c, j = c+1, j+i {
			vec[j] += 11 * i
		}
	}
	for i, v := range vec {
		if v >= n {
			return i
		}
	}
	return -1
}

func init() {
	types.Register(Probs, Day20{})
}
