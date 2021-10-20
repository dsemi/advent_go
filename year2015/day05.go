package year2015

import (
	"advent/problems"
	"strings"
)

type Day05 struct{}

func (Day05) Part1(input string) interface{} {
	var total int
	for _, line := range strings.Split(input, "\n") {
		var cnt int
		var double bool
		for i := 0; i < len(line) && (!double || cnt < 3); i++ {
			if strings.Contains("aeiou", string(line[i])) {
				cnt++
			}
			if i > 0 && line[i-1] == line[i] {
				double = true
			}
		}
		c3 := true
		for _, x := range []string{"ab", "cd", "pq", "xy"} {
			if strings.Contains(line, x) {
				c3 = false
				break
			}
		}
		if cnt >= 3 && double && c3 {
			total++
		}

	}
	return total
}

func (Day05) Part2(input string) interface{} {
	var total int
	for _, line := range strings.Split(input, "\n") {
		if func() bool {
			for i := 0; i < len(line)-3; i++ {
				for j := i + 2; j < len(line)-1; j++ {
					if line[i] == line[j] && line[i+1] == line[j+1] {
						return true
					}
				}
			}
			return false
		}() && func() bool {
			for i := 0; i < len(line)-2; i++ {
				if line[i] == line[i+2] {
					return true
				}
			}
			return false
		}() {
			total++
		}
	}
	return total
}

func init() {
	problems.Register(Day05{})
}
