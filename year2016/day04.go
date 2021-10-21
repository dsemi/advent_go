package year2016

import (
	"advent/problems"
	"advent/utils"
	"strings"
)

type Day04 struct{}

type Room struct {
	name     string
	sectorId int
	checksum string
}

func (r Room) isReal() bool {
	c := utils.NewCounter(strings.ReplaceAll(r.name, "-", ""))
	return r.checksum == string(c.Runes()[:5])
}

func (*Day04) parseRooms(input string) chan Room {
	c := make(chan Room)
	go func() {
		defer close(c)
		for _, line := range strings.Split(input, "\n") {
			idx := strings.LastIndex(line, "-")
			name, rest := line[:idx], line[idx+1:]
			pts := strings.Split(rest, "[")
			c <- Room{
				name:     name,
				sectorId: utils.Int(pts[0]),
				checksum: pts[1][:len(pts[1])-1],
			}
		}
	}()
	return c
}

func (d *Day04) Part1(input string) interface{} {
	var sum int
	for r := range d.parseRooms(input) {
		if r.isReal() {
			sum += r.sectorId
		}
	}
	return sum
}

func (*Day04) rotate(n int, s string) string {
	b := []rune(s)
	for i := range b {
		if b[i] == ' ' {
			b[i] = '-'
		} else {
			b[i] = ((b[i]-rune(n)-97)%26+26)%26 + 97
		}
	}
	return string(b)
}

func (d *Day04) Part2(input string) interface{} {
	for r := range d.parseRooms(input) {
		if strings.Contains(r.name, d.rotate(r.sectorId, "northpole")) {
			return r.sectorId
		}
	}
	return -1
}

func init() {
	problems.Register(&Day04{})
}