package main

import (
	"math"
	"math/big"
	"strings"
	"utils"
)

func Part1(input string) interface{} {
	pts := strings.Fields(input)
	card, door := utils.Int64(pts[0]), utils.Int64(pts[1])
	var md int64 = 20201227
	m := int64(math.Ceil(math.Sqrt(float64(md))))
	tbl := make(map[int64]int64)
	for i, n := int64(0), int64(1); i < m; i, n = i+1, (n*7)%md {
		tbl[n] = i
	}
	var f big.Int
	f.Exp(big.NewInt(7), big.NewInt(md-m-1), big.NewInt(md))
	factor := f.Int64()
	n := door
	for i := int64(0); i < m; i++ {
		if v, ok := tbl[n]; ok {
			var x big.Int
			return x.Exp(big.NewInt(card), big.NewInt(i*m+v), big.NewInt(md))
		}
		n = (n * factor) % md
	}
	panic("unreachable")
}

func Part2(input string) interface{} {
	return ""
}
