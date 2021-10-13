package main

import (
	"advent/problems"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"
)

var expected map[string]map[string][]string

func getExpectedSolutions(t *testing.T, year int, day int) (string, string) {
	if expected == nil {
		testFile := filepath.Join(problems.Basepath, "test/expectedAnswers.json")
		buf, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatal("Error reading test file")
		}
		json.Unmarshal(buf, &expected)
	}
	parts := expected[strconv.Itoa(year)][strconv.Itoa(day)]
	return parts[0], parts[1]
}

func TestProblems(t *testing.T) {
	for year, days := range problems.Probs {
		for day, p := range days {
			t.Run(fmt.Sprintf("Year%dDay%02d", year, day), func(t *testing.T) {
				p1e, p2e := getExpectedSolutions(t, year, day)
				input := problems.GetInput(year, day, false)
				p1a := fmt.Sprint(p.Part1(input))
				if p1e != p1a {
					t.Fatalf("Expected: %v\nObserved: %v", p1e, p1a)
				}
				p2a := fmt.Sprint(p.Part2(input))
				if p2e != p2a {
					t.Fatalf("Expected: %v\nObserved: %v", p2e, p2a)
				}
			})
		}
	}
}
