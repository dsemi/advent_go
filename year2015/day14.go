package year2015

import (
	"advent/problems"
	"advent/utils"
	"fmt"
	"log"
	"strings"
)

type Day14 struct{}

type Reindeer struct {
	speed, flyTime, restTime int
	timeLeft                 int
	dist                     int
	resting                  bool
	score                    int
}

func (r *Reindeer) tick() {
	if !r.resting {
		r.dist += r.speed
	}
	r.timeLeft--
	if r.timeLeft == 0 {
		r.resting = !r.resting
		if r.resting {
			r.timeLeft = r.restTime
		} else {
			r.timeLeft = r.flyTime
		}
	}
}

func (Day14) parseReindeer(input string) []*Reindeer {
	var reindeer []*Reindeer
	for _, line := range strings.Split(input, "\n") {
		var name string
		var speed, flyTime, restTime int
		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds",
			&name, &speed, &flyTime, &restTime)
		if err != nil {
			log.Fatalf("Error parsing input: %v", err)
		}
		reindeer = append(reindeer, &Reindeer{
			speed:    speed,
			flyTime:  flyTime,
			restTime: restTime,
			timeLeft: flyTime,
		})
	}
	return reindeer
}

func (d Day14) Part1(input string) interface{} {
	rs := d.parseReindeer(input)
	for i := 0; i < 2503; i++ {
		for _, r := range rs {
			r.tick()
		}
	}
	var max int
	for _, r := range rs {
		max = utils.Max(max, r.dist)
	}
	return max
}

func (d Day14) Part2(input string) interface{} {
	rs := d.parseReindeer(input)
	for i := 0; i < 2503; i++ {
		var max int
		for _, r := range rs {
			r.tick()
			max = utils.Max(max, r.dist)
		}
		for _, r := range rs {
			if r.dist == max {
				r.score++
			}
		}
	}
	var max int
	for _, r := range rs {
		max = utils.Max(max, r.score)
	}
	return max
}

func init() {
	problems.Register(Day14{})
}
