package main

import (
	"advent/problems"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:generate ./problems/build_problems.sh

func getInput(year, day int) string {
	// TODO network
	_, b, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(filepath.Dir(b), fmt.Sprintf("inputs/%d/input%d.txt", year, day))
	buf, _ := ioutil.ReadFile(inputFile)
	return strings.TrimRightFunc(string(buf), unicode.IsSpace)
}

func colorizeTime(t float64) string {
	var color string
	if t < 0.5 {
		color = "\x1b[32m"
	} else if t < 1.0 {
		color = "\x1b[33m"
	} else {
		color = "\x1b[31m"
	}
	return fmt.Sprintf("%v%.3f\x1b[0m", color, t)
}

func runProblem(year, day int, input string) float64 {
	prob := func() problems.Day {
		if y := problems.Probs[year]; y != nil {
			return y[day]
		}
		return problems.Day{}
	}()
	if prob.Part1 == nil {
		fmt.Println("Day", day, "not implemented")
		return 0
	}
	fmt.Println("Day", day)
	start := time.Now()
	ans := prob.Part1(input)
	t1 := time.Since(start).Seconds()
	fmt.Printf("Part 1: %32v  Elapsed time %v seconds\n", ans, colorizeTime(t1))
	start = time.Now()
	ans = prob.Part2(input)
	t2 := time.Since(start).Seconds()
	fmt.Printf("Part 2: %32v  Elapsed time %v seconds\n", ans, colorizeTime(t2))
	fmt.Println()
	return t1 + t2
}

func main() {
	year, _ := strconv.Atoi(os.Args[1])
	days := os.Args[2:]
	var total float64
	for _, day := range days {
		day, _ := strconv.Atoi(day)
		input := getInput(year, day)
		total += runProblem(year, day, input)
	}
	fmt.Printf("Total: %53.3f seconds\n", total)
}
