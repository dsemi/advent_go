package year2016

import (
	"advent/problems"
	"advent/utils"
	"regexp"
	"strings"
	"sync"
)

type Day10 struct{}

type Bot struct {
	a, b int
	c    chan int
}

type Factory struct {
	bots map[string]*Bot
}

func (f *Factory) GetBot(name string) *Bot {
	if _, ok := f.bots[name]; !ok {
		f.bots[name] = &Bot{c: make(chan int, 1)}
	}
	return f.bots[name]
}

func (*Day10) buildFactory(input string) Factory {
	botRe := regexp.MustCompile("(bot \\d+) gives low to (\\w+ \\d+) and high to (\\w+ \\d+)")
	valRe := regexp.MustCompile("value (\\d+) goes to (bot \\d+)")
	factory := Factory{bots: make(map[string]*Bot)}
	var wg sync.WaitGroup
	for _, line := range strings.Split(input, "\n") {
		if m := botRe.FindStringSubmatch(line); m != nil {
			src := factory.GetBot(m[1])
			lo, hi := factory.GetBot(m[2]), factory.GetBot(m[3])
			wg.Add(1)
			go func() {
				defer wg.Done()
				src.a = <-src.c
				src.b = <-src.c
				lo.c <- utils.Min(src.a, src.b)
				hi.c <- utils.Max(src.a, src.b)
			}()
		} else if m := valRe.FindStringSubmatch(line); m != nil {
			factory.GetBot(m[2]).c <- utils.Int(m[1])
		}
	}
	wg.Wait()
	return factory
}

func (d *Day10) Part1(input string) interface{} {
	f := d.buildFactory(input)
	for k, v := range f.bots {
		if v.a == 17 && v.b == 61 || v.a == 61 && v.b == 17 {
			return k[4:]
		}
	}
	return nil
}

func (d *Day10) Part2(input string) interface{} {
	f := d.buildFactory(input)
	return <-f.GetBot("output 0").c * <-f.GetBot("output 1").c * <-f.GetBot("output 2").c
}

func init() {
	problems.Register(&Day10{})
}
