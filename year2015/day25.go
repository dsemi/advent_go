package main

import (
	"math/big"
	"regexp"
	"utils"
)

func Part1(input string) interface{} {
	reg := regexp.MustCompile("\\d+")
	v := reg.FindAllString(input, -1)
	r, c := utils.Int(v[0]), utils.Int(v[1])
	n := r + c - 1
	index := big.NewInt(int64(n*(n-1)/2 + c - 1))
	m := big.NewInt(33554393)
	x := big.NewInt(252533)
	return x.Exp(x, index, m).Mul(x, big.NewInt(20151125)).Mod(x, m)
}

func Part2(input string) interface{} {
	return ""
}
