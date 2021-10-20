package year2020

import (
	"advent/problems"
	"fmt"
	"regexp"
	"strings"
)

type Day04 struct{}

func (Day04) Part1(input string) interface{} {
	res := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var cnt int
OUTER:
	for _, line := range strings.Split(input, "\n\n") {
		for _, re := range res {
			if ok, _ := regexp.MatchString(fmt.Sprintf("(^|\\s)%s:", re), line); !ok {
				continue OUTER
			}
		}
		cnt++
	}
	return cnt
}

func (Day04) Part2(input string) interface{} {
	res := []string{
		"(^|\\s)byr:(19[2-9][0-9]|200[0-2])(\\s|$)",
		"(^|\\s)iyr:(201[0-9]|2020)(\\s|$)",
		"(^|\\s)eyr:(202[0-9]|2030)(\\s|$)",
		"(^|\\s)hgt:(1[5-8][0-9]|19[0-3])cm|hgt:(59|6[0-9]|7[0-6])in(\\s|$)",
		"(^|\\s)hcl:#[0-9a-f]{6}(\\s|$)",
		"(^|\\s)ecl:(amb|blu|brn|gry|grn|hzl|oth)(\\s|$)",
		"(^|\\s)pid:[0-9]{9}(\\s|$)",
	}
	var cnt int
OUTER:
	for _, line := range strings.Split(input, "\n\n") {
		for _, re := range res {
			if ok, _ := regexp.MatchString(re, line); !ok {
				continue OUTER
			}
		}
		cnt++
	}
	return cnt
}

func init() {
	problems.Register(Day04{})
}
