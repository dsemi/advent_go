package main

import (
	"strings"
	"utils"
)

type Ip struct {
	supernets, hypernets []string
}

func ips(input string) []Ip {
	result := make([]Ip, 0)
	for _, line := range strings.Split(input, "\n") {
		var ip Ip
		for i, part := range strings.FieldsFunc(line, func(x rune) bool { return x == '[' || x == ']' }) {
			if i%2 == 0 {
				ip.supernets = append(ip.supernets, part)
			} else {
				ip.hypernets = append(ip.hypernets, part)
			}
		}
		result = append(result, ip)
	}
	return result
}

func hasAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func Part1(input string) interface{} {
	var cnt int
	for _, ip := range ips(input) {
		if utils.Any(hasAbba, ip.supernets) && !utils.Any(hasAbba, ip.hypernets) {
			cnt++
		}
	}
	return cnt
}

func abas(s string) [][2]byte {
	result := make([][2]byte, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			result = append(result, [2]byte{s[i], s[i+1]})
		}
	}
	return result
}

func (ip *Ip) hasBab(ab [2]byte) bool {
	bab := string([]byte{ab[1], ab[0], ab[1]})
	return utils.Any(func(h string) bool { return strings.Contains(h, bab) }, ip.hypernets)
}

func Part2(input string) interface{} {
	var cnt int
	for _, ip := range ips(input) {
		if utils.Any(func(s string) bool { return utils.Any(ip.hasBab, abas(s)) }, ip.supernets) {
			cnt++
		}
	}
	return cnt
}
