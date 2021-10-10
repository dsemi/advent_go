package year2015

func lookAndSay(n int, input string) int {
	one := []rune(input)
	two := make([]rune, 0, len(one))
	inp := &one
	out := &two
	for i := 0; i < n; i++ {
		curr := (*inp)[0]
		count := 1
		for _, c := range (*inp)[1:] {
			if curr == c {
				count++
				continue
			}
			*out = append(*out, rune(count)+'0', curr)
			curr = c
			count = 1
		}
		*out = append(*out, rune(count)+'0', curr)
		inp, out = out, inp
		*out = (*out)[:0]
	}
	return len(*inp)
}

func Day10Part1(input string) interface{} {
	return lookAndSay(40, input)
}

func Day10Part2(input string) interface{} {
	return lookAndSay(50, input)
}
