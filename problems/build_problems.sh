#!/bin/bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

echo "package problems" > "$SCRIPT_DIR/problems.go"
echo "" >> "$SCRIPT_DIR/problems.go"
echo "import (" >> "$SCRIPT_DIR/problems.go"
for d in *; do
  if [[ $d == year* ]]; then
    echo -e "\t\"advent/$d\"" >> "$SCRIPT_DIR/problems.go"
  fi
done
echo ")" >> "$SCRIPT_DIR/problems.go"
echo "" >> "$SCRIPT_DIR/problems.go"
echo "type Day struct {" >> "$SCRIPT_DIR/problems.go"
echo -e "\tPart1 func(string) interface{}" >> "$SCRIPT_DIR/problems.go"
echo -e "\tPart2 func(string) interface{}" >> "$SCRIPT_DIR/problems.go"
echo "}" >> "$SCRIPT_DIR/problems.go"
echo "" >> "$SCRIPT_DIR/problems.go"
echo "var Probs = map[int]map[int]Day{" >> "$SCRIPT_DIR/problems.go"
for d in *; do
  if [[ $d == year* ]]; then
    echo -e "\t${d:4}: map[int]Day{" >> "$SCRIPT_DIR/problems.go"
    for i in {01..25}; do
      if [[ -f "$d/day$i.go" ]]; then
        n=$(echo $i | sed 's/^0*//')
        echo -e "\t\t$n: Day{$d.Day${i}Part1, $d.Day${i}Part2}," >> "$SCRIPT_DIR/problems.go"
      fi
    done
    echo -e "\t}," >> "$SCRIPT_DIR/problems.go"
  fi
done
echo "}" >> "$SCRIPT_DIR/problems.go"
