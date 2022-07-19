package main

import "utils"

func Part1(input string) interface{} {
	step := utils.Int16(input)
	list := make([]int16, 1, 2018)
	var idx int16
	for v := int16(1); v <= 2017; v++ {
		idx = (idx+step)%v + 1
		list = append(list[:idx], list[idx-1:]...)
		list[idx] = v
	}
	return list[idx+1]
}

func Part2(input string) interface{} {
	step := utils.Int(input)
	var pos, n, valAft0 int
	for n < 50_000_000 {
		if pos == 1 {
			valAft0 = n
		}
		skip := (n-pos)/step + 1
		n += skip
		pos = (pos+skip*(step+1)-1)%n + 1
	}
	return valAft0
}
