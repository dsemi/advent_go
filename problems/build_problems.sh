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

output_problems > "$SCRIPT_DIR/problems.go"
