package year2015

import "advent/problems"

type Day11 struct{}

func (Day11) increment(b []rune) {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 'z' {
			b[i] = 'a'
		} else {
			b[i]++
			if b[i] == 'i' || b[i] == 'o' || b[i] == 'l' {
				b[i]++
				for j := i + 1; j < len(b); j++ {
					b[j] = 'a'
				}
			}
			break
		}
	}
}

func (Day11) isValid(b []rune) bool {
	for i := 0; i < len(b)-2; i++ {
		if b[i]+2 == b[i+1]+1 && b[i+1]+1 == b[i+2] {
			goto NEXT
		}
	}
	return false
NEXT:
	var cnt int
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			cnt++
			i++
		}
	}
	return cnt >= 2
}

func (d Day11) nextValidPw(s string) string {
	b := []rune(s)
	d.increment(b)
	for !d.isValid(b) {
		d.increment(b)
	}
	return string(b)
}

func (d Day11) Part1(input string) interface{} {
	return d.nextValidPw(input)
}

func (d Day11) Part2(input string) interface{} {
	return d.nextValidPw(d.nextValidPw(input))
}

func init() {
	problems.Register(Day11{})
}
