package main

import (
	"flag"
	"fmt"
	"log"
	"problems"
	"strconv"
	"strings"
	"time"
	"utils"
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

func runProblem(year, day int) float64 {
	p1, p2 := problems.GetProb(year, day)
	if p1 == nil || p2 == nil {
		fmt.Println(year, "Day", day, "not implemented")
		return 0
	}
	input := problems.GetInput(year, day, true)
	fmt.Println("Day", day)
	start := time.Now()
	ans := p1(input)
	t1 := time.Since(start).Seconds()
	fmt.Printf("Part 1: %50v  Elapsed time %v seconds\n", ans, colorizeTime(t1))
	start = time.Now()
	ans = p2(input)
	t2 := time.Since(start).Seconds()
	fmt.Printf("Part 2: %50v  Elapsed time %v seconds\n", ans, colorizeTime(t2))
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
	flag.Parse()
	year, _ := strconv.Atoi(flag.Arg(0))
	days := flag.Args()[1:]
	var total float64
	for _, dayStr := range days {
		start, end := parseDays(dayStr)
		for day := start; day <= end; day++ {
			total += runProblem(year, day)
		}
	}
	fmt.Printf("Total: %71.3f seconds\n", total)
}
