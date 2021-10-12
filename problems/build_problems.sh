#!/bin/bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

function output_problems {
  echo "package problems"
  echo ""
  echo "import ("
  for d in *; do
    if [[ $d == year* ]]; then
      echo -e "\t\"advent/$d\""
    fi
  done
  echo ")"
  echo ""
  echo "type Day struct {"
  echo -e "\tPart1 func(string) interface{}"
  echo -e "\tPart2 func(string) interface{}"
  echo "}"
  echo ""
  echo "var Probs = map[int]map[int]Day{"
  for d in *; do
    if [[ $d == year* ]]; then
      echo -e "\t${d:4}: map[int]Day{"
      for i in {01..25}; do
        if [[ -f "$d/day$i.go" ]]; then
          n=$(echo $i | sed 's/^0*//')
          echo -e "\t\t$n: Day{$d.Day${i}Part1, $d.Day${i}Part2},"
        fi
      done
      echo -e "\t},"
    fi
  done
  echo "}"
}

function output_tests {
  echo "package problems"
  echo ""
  echo "import ("
  echo -e "\t\"encoding/json\""
  echo -e "\t\"fmt\""
  echo -e "\t\"io/ioutil\""
  echo -e "\t\"path/filepath\""
  echo -e "\t\"strconv\""
  echo -e "\t\"testing\""
  echo ")"
  echo ""
  echo "var expected map[string]map[string][]string"
  echo ""
  echo "func getExpectedSolutions(t *testing.T, year int, day int) (string, string) {"
  echo -e "\tif expected == nil {"
  echo -e "\t\ttestFile := filepath.Join(basepath, \"test/expectedAnswers.json\")"
  echo -e "\t\tbuf, err := ioutil.ReadFile(testFile)"
  echo -e "\t\tif err != nil {"
  echo -e "\t\t\tt.Fatal(\"Error reading test file\")"
  echo -e "\t\t}"
  echo -e "\t\tjson.Unmarshal(buf, &expected)"
  echo -e "\t}"
  echo -e "\tparts := expected[strconv.Itoa(year)][strconv.Itoa(day)]"
  echo -e "\treturn parts[0], parts[1]"
  echo "}"
  echo ""
  echo "func runTest(t *testing.T, year, day int) {"
  echo -e "\tp := Probs[year][day]"
  echo -e "\tp1e, p2e := getExpectedSolutions(t, year, day)"
  echo -e "\tinput := GetInput(year, day, false)"
  echo -e "\tp1a := fmt.Sprint(p.Part1(input))"
  echo -e "\tif p1e != p1a {"
  echo -e "\t\tt.Fatalf(\"Expected: %v\\\nObserved: %v\", p1e, p1a)"
  echo -e "\t}"
  echo -e "\tp2a := fmt.Sprint(p.Part2(input))"
  echo -e "\tif p2e != p2a {"
  echo -e "\t\tt.Fatalf(\"Expected: %v\\\nObserved: %v\", p2e, p2a)"
  echo -e "\t}"
  echo "}"

  for d in *; do
    if [[ $d == year* ]]; then
      for i in {01..25}; do
        if [[ -f "$d/day$i.go" ]]; then
          y="${d:4}"
          n=$(echo $i | sed 's/^0*//')
          echo ""
          echo "func TestYear${y}Day$i(t *testing.T) {"
          echo -e "\trunTest(t, $y, $n)"
          echo "}"
        fi
      done
    fi
  done
}

output_problems > "$SCRIPT_DIR/problems.go"

output_tests > "$SCRIPT_DIR/problems_test.go"
