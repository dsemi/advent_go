package main

func process(input string) (int, int) {
	var score, depth, garbageCount int
	var inGarbage, ignoreNext bool
	for _, x := range input {
		if ignoreNext {
			ignoreNext = false
		} else if inGarbage {
			if x == '!' {
				ignoreNext = true
			} else if x == '>' {
				inGarbage = false
			} else {
				garbageCount++
			}
		} else if x == '}' {
			score += depth
			depth--
		} else if x == '{' {
			depth++
		} else if x == '<' {
			inGarbage = true
		}
	}
	return score, garbageCount
}

func Part1(input string) interface{} {
	ans, _ := process(input)
	return ans
}

func Part2(input string) interface{} {
	_, ans := process(input)
	return ans
}
