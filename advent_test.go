package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"problems"
	"strconv"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

var (
	year = flag.Int("year", 0, "Year of the problem to test.")
	day  = flag.Int("day", 0, "Day of the problem to test.")

	expected map[string]map[string][]string
)

func getExpectedSolutions(t *testing.T, year int, day int) (*string, *string) {
	if expected == nil {
		filepath, err := bazel.Runfile("expectedAnswers.json")
		if err != nil {
			t.Fatalf("Error determining runfile for expected answers: %v", err)
		}
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Fatalf("Error reading expected answers file: %v", err)
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

func TestProblem(t *testing.T) {
	p1, p2 := problems.GetProb(*year, *day)
	if p1 == nil || p2 == nil {
		t.Fatalf("Could not find problem year %d day %d", *year, *day)
	}
	p1e, p2e := getExpectedSolutions(t, *year, *day)
	if p1e != nil || p2e != nil {
		input := problems.GetInput(*year, *day, false)
		p1a := fmt.Sprint(p1(input))
		if *p1e != p1a {
			t.Errorf("Part 1 error. Expected: %v\nObserved: %v", *p1e, p1a)
		}
		p2a := fmt.Sprint(p2(input))
		if *p2e != p2a {
			t.Errorf("Part 2 error. Expected: %v\nObserved: %v", *p2e, p2a)
		}
	}
}
