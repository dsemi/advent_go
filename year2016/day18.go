package main

func safeOrTrap(a, b, c byte) byte {
	if a == '^' && b == '^' && c == '.' ||
		a == '.' && b == '^' && c == '^' ||
		a == '^' && b == '.' && c == '.' ||
		a == '.' && b == '.' && c == '^' {
		return '^'
	}
	return '.'
}

func numSafe(n int, input string) int {
	var total int
	state := []byte(input)
	state = append(state, '.')
	for j := 0; j < n; j++ {
		for _, x := range state[:len(state)-1] {
			if x == '.' {
				total++
			}
		}
		prev := byte('.')
		for i, v := range state[:len(state)-1] {
			state[i] = safeOrTrap(prev, v, state[i+1])
			prev = v
		}
	}
	return total
}

func Part1(input string) interface{} {
	return numSafe(40, input)
}

func Part2(input string) interface{} {
	return numSafe(400_000, input)
}
