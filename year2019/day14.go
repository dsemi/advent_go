package main

import (
	"strings"
	"utils"
)

type Chemical struct {
	n      int64
	inputs map[string]int64
}

func parse(input string) map[string]*Chemical {
	result := make(map[string]*Chemical)
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Split(line, " => ")
		outp := strings.Fields(pts[1])
		chem := &Chemical{n: utils.Int64(outp[0]), inputs: make(map[string]int64)}
		for _, inp := range strings.Split(pts[0], ", ") {
			pts2 := strings.Fields(inp)
			chem.inputs[pts2[1]] = utils.Int64(pts2[0])
		}
		result[outp[1]] = chem
	}
	return result
}

func numOre(reactions map[string]*Chemical, surplus map[string]int64, k string, c int64) int64 {
	var ore int64
	if chemical, ok := reactions[k]; ok {
		n, chems := chemical.n, chemical.inputs
		q, r := utils.DivMod(c, n)
		for chem, a := range chems {
			amt := a * (q + utils.ToInt[int64](r != 0))
			val := surplus[chem]
			surplus[chem] = utils.Max(0, val-amt)
			if amt > val {
				ore += numOre(reactions, surplus, chem, amt-val)
			}
		}
		if r != 0 {
			surplus[k] += n - r
		}
	} else {
		ore += c
	}
	return ore
}

func Part1(input string) interface{} {
	return numOre(parse(input), make(map[string]int64), "FUEL", 1)
}

const trillion int64 = 1_000_000_000_000

func Part2(input string) interface{} {
	reactions := parse(input)
	var a, b int64 = 0, trillion
	for a < b {
		mid := (a + b) / 2
		if numOre(reactions, make(map[string]int64), "FUEL", mid) > trillion {
			b = mid - 1
		} else {
			a = mid + 1
		}
	}
	if numOre(reactions, make(map[string]int64), "FUEL", a) > trillion {
		return a - 1
	}
	return a
}
