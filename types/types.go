package types

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type Day interface {
	Part1(string) interface{}
	Part2(string) interface{}
}

func Register(probs map[int]Day, day Day) {
	_, b, _, _ := runtime.Caller(1)
	var d int
	_, err := fmt.Sscanf(filepath.Base(b), "day%d.go", &d)
	if err != nil {
		log.Fatalf("Error parsing file: %v", b)
	}
	probs[d] = day
}
