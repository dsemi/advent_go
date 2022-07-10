package main

import (
	"strings"
	"utils"
)

type RuleType uint8

const (
	Single RuleType = iota
	Multi
)

type Rule struct {
	t   RuleType
	c   byte
	rss [][]int
}

func parse(input string) ([]Rule, []string) {
	parts := strings.Split(input, "\n\n")
	rules := make([]Rule, len(strings.Split(parts[0], "\n")))
	for _, line := range strings.Split(parts[0], "\n") {
		pts := strings.Split(line, ": ")
		idx := utils.Int(pts[0])
		rule := Rule{}
		if strings.HasPrefix(pts[1], "\"") {
			rule.t = Single
			rule.c = pts[1][1]
		} else {
			rule.t = Multi
			for _, p := range strings.Split(pts[1], " | ") {
				r := make([]int, 0)
				for _, x := range strings.Fields(p) {
					r = append(r, utils.Int(x))
				}
				rule.rss = append(rule.rss, r)
			}
		}
		rules[idx] = rule
	}
	return rules, strings.Split(parts[1], "\n")
}

func check(rules []Rule, s string, seq []int) bool {
	if len(s) == 0 || len(seq) == 0 {
		return len(s) == 0 && len(seq) == 0
	}
	rule := rules[seq[0]]
	switch rule.t {
	case Single:
		return s[0] == rule.c && check(rules, s[1:], seq[1:])
	case Multi:
		for _, rs := range rule.rss {
			if check(rules, s, append(rs, seq[1:]...)) {
				return true
			}
		}
	}
	return false
}

func countMatches(rules []Rule, messages []string) int {
	var cnt int
	for _, message := range messages {
		if check(rules, message, []int{0}) {
			cnt++
		}
	}
	return cnt
}

func Part1(input string) interface{} {
	rules, messages := parse(input)
	return countMatches(rules, messages)
}

func Part2(input string) interface{} {
	rules, messages := parse(input)
	rules[8] = Rule{t: Multi, rss: [][]int{{42}, {42, 8}}}
	rules[11] = Rule{t: Multi, rss: [][]int{{42, 31}, {42, 11, 31}}}
	return countMatches(rules, messages)
}
