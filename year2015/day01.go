package year2015

func Day01Part1(input string) interface{} {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func Day01Part2(input string) interface{} {
	floor := 0
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
