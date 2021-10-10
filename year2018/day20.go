package year2018

type Coord struct {
	x int
	y int
}

func parseEdges(input string) map[Coord]int {
	var stack []Coord
	pos := Coord{x: 0, y: 0}
	result := make(map[Coord]int)
	for i := 1; i < len(input)-1; i++ {
		switch input[i] {
		case '(':
			stack = append(stack, pos)
		case ')':
			pos = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case '|':
			pos = stack[len(stack)-1]
		default:
			v := result[pos] + 1
			switch input[i] {
			case 'N':
				pos.y -= 1
			case 'E':
				pos.x += 1
			case 'S':
				pos.y += 1
			case 'W':
				pos.x -= 1
			}
			if d, ok := result[pos]; !ok || v < d {
				result[pos] = v
			}
		}
	}
	return result
}

func Day20Part1(input string) interface{} {
	var max int
	for _, v := range parseEdges(input) {
		if v > max {
			max = v
		}
	}
	return max
}

func Day20Part2(input string) interface{} {
	var cnt int
	for _, v := range parseEdges(input) {
		if v >= 1000 {
			cnt++
		}
	}
	return cnt
}
