package main

import (
	"advent/problems"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:generate ./problems/build_problems.sh

const rateLimit = 5 * time.Second
var last = new(time.Time)

func getInput(year, day int, download bool) string {
	_, b, _, _ := runtime.Caller(0)
	inputFile := filepath.Join(filepath.Dir(b), fmt.Sprintf("inputs/%d/input%d.txt", year, day))
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && download {
			fmt.Printf("Downloading input for Year %d Day %d\n", year, day)
			if last != nil {
				time.Sleep(time.Until(last.Add(rateLimit)))
			}
			*last = time.Now()
			url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Cookie", os.Getenv("AOC_SESSION"))
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalf("Problem input fetch failed: %v", err)
			}
			if resp.StatusCode < 200 || resp.StatusCode > 299 {
				log.Fatalf("Bad HTTP response: %v", resp)
			}
			b := new(bytes.Buffer)
			if 	_, err = b.ReadFrom(resp.Body); err != nil {
				log.Fatalf("Error reading HTTP response body: %v", err)
			}
			if err = resp.Body.Close(); err != nil {
				log.Fatalf("Error closing HTTP response body: %v", err)
			}
			buf = b.Bytes()
			if err = ioutil.WriteFile(inputFile, buf, 0644); err != nil {
				log.Fatalf("Unable to write to output file: %v", err)
			}
		} else {
			log.Fatalf("Error reading problem input file: %v", err)
		}
	}
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
		input := getInput(year, day, true)
		total += runProblem(year, day, input)
	}
	fmt.Printf("Total: %53.3f seconds\n", total)
}
