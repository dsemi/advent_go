package main

import (
	"advent/problems"
	"advent/types"
	"advent/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	prob := func() types.Day {
		if y := problems.Probs[year]; y != nil {
			return y[day]
		}
		return nil
	}()
	if prob == nil {
		fmt.Println(year, "Day", day, "not implemented")
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

func parseDays(dayStr string) (int, int) {
	if strings.ContainsRune(dayStr, '-') {
		pts := strings.Split(dayStr, "-")
		if len(pts) != 2 {
			log.Fatalf("Failed to parse days: %v", dayStr)
		}
		a, b := utils.Int(pts[0]), utils.Int(pts[1])
		if a > b {
			log.Fatalf("Invalid range: %v", dayStr)
		}
		return a, b
	} else {
		n := utils.Int(dayStr)
		return n, n
	}
}

func main() {
	year, _ := strconv.Atoi(os.Args[1])
	days := os.Args[2:]
	var total float64
	for _, dayStr := range days {
		start, end := parseDays(dayStr)
		for day := start; day <= end; day++ {
			input := problems.GetInput(year, day, true)
			total += runProblem(year, day, input)
		}
	}
	fmt.Printf("Total: %53.3f seconds\n", total)
}
