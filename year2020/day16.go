package main

import (
	"fmt"
	"strings"
	"utils"
)

type Rule struct {
	name               string
	id                 uint64
	alo, ahi, blo, bhi uint64
}

func parse(input string) ([]Rule, []uint64, [][]uint64) {
	parts := strings.Split(input, "\n\n")
	rules := make([]Rule, 0)
	for i, line := range strings.Split(parts[0], "\n") {
		rule := Rule{id: 1 << i}
		pts := strings.Split(line, ": ")
		rule.name = pts[0]
		fmt.Sscanf(pts[1], "%d-%d or %d-%d", &rule.alo, &rule.ahi, &rule.blo, &rule.bhi)
		rules = append(rules, rule)
	}
	yours := make([]uint64, 0)
	for _, x := range strings.Split(strings.Split(parts[1], "\n")[1], ",") {
		yours = append(yours, utils.Uint64(x))
	}
	others := make([][]uint64, 0)
	for _, line := range strings.Split(parts[2], "\n")[1:] {
		other := make([]uint64, 0)
		for _, x := range strings.Split(line, ",") {
			other = append(other, utils.Uint64(x))
		}
		others = append(others, other)
	}
	return rules, yours, others
}

func invalidValues(rules []Rule, ticket []uint64) []uint64 {
	vals := make([]uint64, 0)
outer:
	for _, field := range ticket {
		for _, rule := range rules {
			if rule.alo <= field && field <= rule.ahi || rule.blo <= field && field <= rule.bhi {
				continue outer
			}
		}
		vals = append(vals, field)
	}
	return vals
}

func Part1(input string) interface{} {
	rules, _, tix := parse(input)
	var sum uint64
	for _, t := range tix {
		for _, v := range invalidValues(rules, t) {
			sum += v
		}
	}
	return sum
}

func Part2(input string) interface{} {
	rules, yours, tix := parse(input)
	ruleMap := make(map[uint64]Rule)
	for _, rule := range rules {
		ruleMap[rule.id] = rule
	}
	tix2 := make([][]uint64, 0)
	for _, t := range tix {
		if len(invalidValues(rules, t)) == 0 {
			tix2 = append(tix2, t)
		}
	}
	tix = tix2
	poss := make([][]Rule, len(yours))
	for i := range poss {
		poss[i] = make([]Rule, len(rules))
		copy(poss[i], rules)
	}
	for _, t := range tix {
		for i, field := range t {
			validRules := make([]Rule, 0)
			for _, rule := range poss[i] {
				if rule.alo <= field && field <= rule.ahi || rule.blo <= field && field <= rule.bhi {
					validRules = append(validRules, rule)
				}
			}
			poss[i] = validRules
		}
	}
	possSet := make([]uint64, len(poss))
	for i, p := range poss {
		var key uint64
		for _, x := range p {
			key |= x.id
		}
		possSet[i] = key
	}
	for {
		var ones uint64
		for _, p := range possSet {
			if utils.CountOnes(p) == 1 {
				ones |= p
			}
		}
		if uint64(len(possSet)) == utils.CountOnes(ones) {
			break
		}
		for i, p := range possSet {
			if utils.CountOnes(p) > 1 {
				possSet[i] &= ^ones
			}
		}
	}
	var prod uint64 = 1
	for i, k := range possSet {
		if strings.HasPrefix(ruleMap[k].name, "departure") {
			prod *= yours[i]
		}
	}
	return prod
}
