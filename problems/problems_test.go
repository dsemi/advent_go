package problems

import (
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
		testFile := filepath.Join(basepath, "test/expectedAnswers.json")
		buf, err := ioutil.ReadFile(testFile)
		if err != nil {
			t.Fatal("Error reading test file")
		}
		json.Unmarshal(buf, &expected)
	}
	parts := expected[strconv.Itoa(year)][strconv.Itoa(day)]
	return parts[0], parts[1]
}

func runTest(t *testing.T, year, day int) {
	p := Probs[year][day]
	p1e, p2e := getExpectedSolutions(t, year, day)
	input := GetInput(year, day, false)
	p1a := fmt.Sprint(p.Part1(input))
	if p1e != p1a {
		t.Fatalf("Expected: %v\nObserved: %v", p1e, p1a)
	}
	p2a := fmt.Sprint(p.Part2(input))
	if p2e != p2a {
		t.Fatalf("Expected: %v\nObserved: %v", p2e, p2a)
	}
}

func TestYear2015Day01(t *testing.T) {
	runTest(t, 2015, 1)
}

func TestYear2015Day02(t *testing.T) {
	runTest(t, 2015, 2)
}

func TestYear2015Day03(t *testing.T) {
	runTest(t, 2015, 3)
}

func TestYear2015Day04(t *testing.T) {
	runTest(t, 2015, 4)
}

func TestYear2015Day05(t *testing.T) {
	runTest(t, 2015, 5)
}

func TestYear2015Day06(t *testing.T) {
	runTest(t, 2015, 6)
}

func TestYear2015Day07(t *testing.T) {
	runTest(t, 2015, 7)
}

func TestYear2015Day08(t *testing.T) {
	runTest(t, 2015, 8)
}

func TestYear2015Day09(t *testing.T) {
	runTest(t, 2015, 9)
}

func TestYear2015Day10(t *testing.T) {
	runTest(t, 2015, 10)
}

func TestYear2015Day11(t *testing.T) {
	runTest(t, 2015, 11)
}

func TestYear2015Day12(t *testing.T) {
	runTest(t, 2015, 12)
}

func TestYear2016Day01(t *testing.T) {
	runTest(t, 2016, 1)
}

func TestYear2016Day03(t *testing.T) {
	runTest(t, 2016, 3)
}

func TestYear2016Day10(t *testing.T) {
	runTest(t, 2016, 10)
}

func TestYear2016Day19(t *testing.T) {
	runTest(t, 2016, 19)
}

func TestYear2017Day01(t *testing.T) {
	runTest(t, 2017, 1)
}

func TestYear2017Day02(t *testing.T) {
	runTest(t, 2017, 2)
}

func TestYear2017Day04(t *testing.T) {
	runTest(t, 2017, 4)
}

func TestYear2017Day18(t *testing.T) {
	runTest(t, 2017, 18)
}

func TestYear2017Day24(t *testing.T) {
	runTest(t, 2017, 24)
}

func TestYear2018Day01(t *testing.T) {
	runTest(t, 2018, 1)
}

func TestYear2018Day05(t *testing.T) {
	runTest(t, 2018, 5)
}

func TestYear2018Day07(t *testing.T) {
	runTest(t, 2018, 7)
}

func TestYear2018Day20(t *testing.T) {
	runTest(t, 2018, 20)
}

func TestYear2019Day01(t *testing.T) {
	runTest(t, 2019, 1)
}

func TestYear2019Day02(t *testing.T) {
	runTest(t, 2019, 2)
}

func TestYear2019Day04(t *testing.T) {
	runTest(t, 2019, 4)
}

func TestYear2019Day05(t *testing.T) {
	runTest(t, 2019, 5)
}

func TestYear2019Day07(t *testing.T) {
	runTest(t, 2019, 7)
}

func TestYear2019Day09(t *testing.T) {
	runTest(t, 2019, 9)
}

func TestYear2019Day25(t *testing.T) {
	runTest(t, 2019, 25)
}

func TestYear2020Day01(t *testing.T) {
	runTest(t, 2020, 1)
}

func TestYear2020Day02(t *testing.T) {
	runTest(t, 2020, 2)
}

func TestYear2020Day03(t *testing.T) {
	runTest(t, 2020, 3)
}

func TestYear2020Day04(t *testing.T) {
	runTest(t, 2020, 4)
}

func TestYear2020Day05(t *testing.T) {
	runTest(t, 2020, 5)
}

func TestYear2020Day06(t *testing.T) {
	runTest(t, 2020, 6)
}

func TestYear2020Day10(t *testing.T) {
	runTest(t, 2020, 10)
}
