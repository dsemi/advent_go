#!/usr/bin/python3

import os
import os.path as op
import sys

TEMPLATE = """package year{year}

import "advent/problems"

type Day{day} struct{{}}

func (*Day{day}) Part1(input string) interface{{}} {{
	return input
}}

func (*Day{day}) Part2(input string) interface{{}} {{
	return ""
}}

func init() {{
	problems.Register(&Day{day}{{}})
}}
"""

def main():
  args = sys.argv[1:]
  assert len(args) == 2
  [year, day] = args
  day = "{:02}".format(int(day))
  filename = f"year{year}/day{day}.go"
  if op.exists(filename):
    print("File already exists")
    exit(1)
  if not op.exists(f"year{year}"):
    os.mkdir(f"year{year}")
  with open(filename, 'w') as f:
    f.write(TEMPLATE.format(year=year, day=day))

if __name__ == '__main__':
  main()
