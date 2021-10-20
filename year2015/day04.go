package year2015

import (
	"advent/problems"
	"crypto/md5"
	"strconv"
)

type Day04 struct{}

func (Day04) sum(s string, i int) [16]byte {
	return md5.Sum([]byte(s + strconv.Itoa(i)))
}

func (d Day04) Part1(input string) interface{} {
	i := 0
	for hash := d.sum(input, i); hash[0] != 0 || hash[1] != 0 || hash[2] > 15; hash = d.sum(input, i) {
		i++
	}
	return i
}

func (d Day04) Part2(input string) interface{} {
	i := 0
	for hash := d.sum(input, i); hash[0] != 0 || hash[1] != 0 || hash[2] != 0; hash = d.sum(input, i) {
		i++
	}
	return i
}

func init() {
	problems.Register(Day04{})
}
