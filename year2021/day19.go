package main

import (
	"fmt"
	"strings"
	"utils"
)

type Pt = [3]int

func mins(a, b *Pt) Pt {
	return [3]int{utils.Min(a[0], b[0]), utils.Min(a[1], b[1]), utils.Min(a[2], b[2])}
}

func hash(p *Pt) uint64 {
	return (uint64(p[0]) << 42) ^ (uint64(p[1]) << 21) ^ uint64(p[2])
}

type Scanner struct {
	ps     []Pt
	offset Pt
	min    Pt
}

func (s *Scanner) add(p Pt) {
	s.min = mins(&s.min, &p)
	s.ps = append(s.ps, p)
}

func parseScanners(input string) []*Scanner {
	scanners := make([]*Scanner, 0)
	for _, sc := range strings.Split(input, "\n\n") {
		scanner := &Scanner{}
		for _, line := range strings.Split(sc, "\n")[1:] {
			var x, y, z int
			fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			scanner.add([3]int{x, y, z})
		}
		scanners = append(scanners, scanner)
	}
	return scanners
}

func canAlign(aa int, a, b *Scanner) (bool, int, int, bool) {
	var collision [4096 * 6]uint8
	for _, pa := range a.ps {
		for _, pb := range b.ps {
			var base int
			for _, n := range []int{2048 + (pb[0] - b.min[0]) - (pa[aa] - a.min[aa]),
				(pb[0] - b.min[0]) + (pa[aa] - a.min[aa]),
				2048 + (pb[1] - b.min[1]) - (pa[aa] - a.min[aa]),
				(pb[1] - b.min[1]) + (pa[aa] - a.min[aa]),
				2048 + (pb[2] - b.min[2]) - (pa[aa] - a.min[aa]),
				(pb[2] - b.min[2]) + (pa[aa] - a.min[aa])} {
				idx := base + n
				collision[idx]++
				if collision[idx] == 12 {
					ori := idx / 4096
					axis := ori / 2
					negate := ori%2 == 1
					n += b.min[axis]
					if negate {
						n += a.min[aa]
					} else {
						n -= a.min[aa] + 2048
					}
					return true, n, axis, negate
				}
				base += 4096
			}
		}
	}
	return false, 0, 0, false
}

func align(aa int, b *Scanner, n int, axis int, negate bool) {
	if negate {
		b.offset[aa] = -n
	} else {
		b.offset[aa] = n
	}
	if axis != aa {
		b.min[axis], b.min[aa] = b.min[aa], b.min[axis]
		for i := range b.ps {
			b.ps[i][axis], b.ps[i][aa] = b.ps[i][aa], b.ps[i][axis]
		}
	}
	if negate {
		b.min[aa] = n - b.min[aa] - 2047
		for i := range b.ps {
			b.ps[i][aa] = n - b.ps[i][aa]
		}
	} else {
		b.min[aa] -= n
		for i := range b.ps {
			b.ps[i][aa] -= n
		}
	}
}

func combine(scanners []*Scanner) {
	need := (uint64(1) << len(scanners)) - 2
	todo := []int{0}
	for len(todo) > 0 {
		i := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		bs := utils.Bits{N: need}
		for bs.Next() {
			j := bs.Get()
			if can, n, axis, negate := canAlign(0, scanners[i], scanners[j]); can {
				align(0, scanners[j], n, axis, negate)
				if can, n, axis, negate := canAlign(1, scanners[i], scanners[j]); can {
					align(1, scanners[j], n, axis, negate)
				}
				if can, n, axis, negate := canAlign(2, scanners[i], scanners[j]); can {
					align(2, scanners[j], n, axis, negate)
				}
				need ^= 1 << j
				todo = append(todo, j)
			}
		}
	}
}

func Part1(input string) interface{} {
	scanners := parseScanners(input)
	combine(scanners)
	set := make(map[uint64]bool)
	for _, s := range scanners {
		for _, p := range s.ps {
			set[hash(&p)] = true
		}
	}
	return len(set)
}

func Part2(input string) interface{} {
	scanners := parseScanners(input)
	combine(scanners)
	var result int
	for _, a := range scanners {
		for _, b := range scanners {
			dist := utils.Abs(a.offset[0]-b.offset[0]) +
				utils.Abs(a.offset[1]-b.offset[1]) +
				utils.Abs(a.offset[2]-b.offset[2])
			result = utils.Max(result, dist)
		}
	}
	return result
}
