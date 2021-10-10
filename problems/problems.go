package problems

import (
	"advent/year2015"
	"advent/year2016"
	"advent/year2017"
	"advent/year2018"
	"advent/year2019"
	"advent/year2020"
)

type Day struct {
	Part1 func(string) interface{}
	Part2 func(string) interface{}
}

var Probs = map[int]map[int]Day{
	2015: map[int]Day{
		1: Day{year2015.Day01Part1, year2015.Day01Part2},
		2: Day{year2015.Day02Part1, year2015.Day02Part2},
		3: Day{year2015.Day03Part1, year2015.Day03Part2},
		4: Day{year2015.Day04Part1, year2015.Day04Part2},
		5: Day{year2015.Day05Part1, year2015.Day05Part2},
		6: Day{year2015.Day06Part1, year2015.Day06Part2},
		7: Day{year2015.Day07Part1, year2015.Day07Part2},
		8: Day{year2015.Day08Part1, year2015.Day08Part2},
		12: Day{year2015.Day12Part1, year2015.Day12Part2},
	},
	2016: map[int]Day{
		1: Day{year2016.Day01Part1, year2016.Day01Part2},
		3: Day{year2016.Day03Part1, year2016.Day03Part2},
		10: Day{year2016.Day10Part1, year2016.Day10Part2},
		19: Day{year2016.Day19Part1, year2016.Day19Part2},
	},
	2017: map[int]Day{
		1: Day{year2017.Day01Part1, year2017.Day01Part2},
		2: Day{year2017.Day02Part1, year2017.Day02Part2},
		4: Day{year2017.Day04Part1, year2017.Day04Part2},
		18: Day{year2017.Day18Part1, year2017.Day18Part2},
		24: Day{year2017.Day24Part1, year2017.Day24Part2},
	},
	2018: map[int]Day{
		1: Day{year2018.Day01Part1, year2018.Day01Part2},
		5: Day{year2018.Day05Part1, year2018.Day05Part2},
		7: Day{year2018.Day07Part1, year2018.Day07Part2},
		20: Day{year2018.Day20Part1, year2018.Day20Part2},
	},
	2019: map[int]Day{
		1: Day{year2019.Day01Part1, year2019.Day01Part2},
		2: Day{year2019.Day02Part1, year2019.Day02Part2},
		4: Day{year2019.Day04Part1, year2019.Day04Part2},
		5: Day{year2019.Day05Part1, year2019.Day05Part2},
		7: Day{year2019.Day07Part1, year2019.Day07Part2},
		9: Day{year2019.Day09Part1, year2019.Day09Part2},
		25: Day{year2019.Day25Part1, year2019.Day25Part2},
	},
	2020: map[int]Day{
		1: Day{year2020.Day01Part1, year2020.Day01Part2},
		2: Day{year2020.Day02Part1, year2020.Day02Part2},
		3: Day{year2020.Day03Part1, year2020.Day03Part2},
		4: Day{year2020.Day04Part1, year2020.Day04Part2},
		5: Day{year2020.Day05Part1, year2020.Day05Part2},
		6: Day{year2020.Day06Part1, year2020.Day06Part2},
		10: Day{year2020.Day10Part1, year2020.Day10Part2},
	},
}
