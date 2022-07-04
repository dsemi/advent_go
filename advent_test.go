package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"problems"
	"strconv"
	"testing"
)

var expected map[string]map[string][]string

func getExpectedSolutions(t *testing.T, year int, day int) (*string, *string) {
	if expected == nil {
		testFile := filepath.Join(problems.Basepath, "expectedAnswers.json")
		buf, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatal("Error reading test file")
		}
		if err := json.Unmarshal(buf, &expected); err != nil {
			t.Fatalf("Error unmarshaling json: %v", err)
		}
	}
	if yr, ok := expected[strconv.Itoa(year)]; ok {
		if parts, ok := yr[strconv.Itoa(day)]; ok {
			return &parts[0], &parts[1]
		}
	}
	return nil, nil
}

func TestProblems(t *testing.T) {
	for year := 2015; year <= 2021; year++ {
		for day := 1; day <= 25; day++ {
			p1, p2 := problems.GetProb(year, day)
			if p1 == nil || p2 == nil {
				continue
			}
			p1e, p2e := getExpectedSolutions(t, year, day)
			if p1e != nil || p2e != nil {
				t.Run(fmt.Sprintf("Year%dDay%02d", year, day), func(t *testing.T) {
					input := problems.GetInput(year, day, false)
					p1a := fmt.Sprint(p1(input))
					if *p1e != p1a {
						t.Errorf("Expected: %v\nObserved: %v", *p1e, p1a)
					}
					p2a := fmt.Sprint(p2(input))
					if *p2e != p2a {
						t.Errorf("Expected: %v\nObserved: %v", *p2e, p2a)
					}
				})
			}
		}
	}
}
