package main

import "utils"

const roomXor uint32 = 0xffaa5500

var (
	costArr   = [4]uint16{1, 10, 100, 1000}
	safeSkips = [...]uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x3, 0x7, 0x7, 0x7, 0x7, 0, 0, 0, 0x3,
		0x707, 0xf0f, 0xf0f, 0xf0f, 0, 0, 0, 0x3, 0x707, 0xf0f0f, 0x1f1f1f, 0x1f1f1f, 0}
)

func readInput(input string) uint32 {
	var room uint32
	for i := 0; i < 4; i++ {
		room |= uint32(input[31]-'A') << (8 * i)
		room |= uint32(input[45]-'A') << (8*i + 2)
		input = input[2:]
	}
	return room ^ (roomXor & 0x0f0f0f0f)
}

func insertPart2(room uint32) uint32 {
	room = (room & 0x03030303) | ((room << 4) & 0xc0c0c0c0)
	return room ^ 0x1c2c0c3c
}

func moveCost(dist int) int {
	return 2 * utils.Max(1, utils.Abs(dist))
}

func baseCost(room uint32) (int, int) {
	room ^= roomXor
	var base, secondRow int
	for i := 0; i < 4; i++ {
		glyph0 := int(room & 3)
		glyph1 := int((room >> 2) & 3)
		cost0 := int(costArr[glyph0])
		cost1 := int(costArr[glyph1])
		base += cost0 * (moveCost(i-glyph0) + 1)
		base += cost1 * moveCost(i-glyph1)
		secondRow += cost1
		room >>= 8
	}
	cost1 := base + secondRow*2 + 3333
	cost2 := base + secondRow*4 + 29115
	return cost1, cost2
}

type Room struct {
	room uint32
}

func (self *Room) empty(r int) bool {
	return (self.room>>(8*r))&0xff == 0
}

func (self *Room) get(r int) int {
	return r ^ int((self.room>>(8*r))&3)
}

func (self *Room) pop(r int) {
	mask1 := uint32(0xff) << (8 * r)
	mask2 := uint32(0x3f) << (8 * r)
	self.room = ((self.room >> 2) & mask2) | (self.room & ^mask1)
}

type Hall struct {
	hall uint32
}

func (self *Hall) empty(h int) bool {
	return self.hall&(4<<(4*h)) == 0
}

func (self *Hall) clear(h int) {
	self.hall &= ^(0xf << (4 * h))
}

func (self *Hall) set(h, g int) {
	self.hall |= (4 | uint32(g)) << (4 * h)
}

func (self *Hall) get(h int) int {
	return int((self.hall >> (4 * h)) & 3)
}

func (self *Hall) mask() uint32 {
	return self.hall & 0x4444444
}

type State struct {
	room Room
	hall Hall
}

func NewState(hash uint64) State {
	return State{
		room: Room{uint32(hash)},
		hall: Hall{uint32(hash >> 32)},
	}
}

func (s *State) hash() uint64 {
	return (uint64(s.hall.hall) << 32) | uint64(s.room.room)
}

func (s *State) solved() bool {
	return s.room.room|s.hall.hall == 0
}

func roomL(r int) int {
	return r + 1
}

func roomR(r int) int {
	return r + 2
}

func (s *State) obstructed(r, h int) bool {
	var lo, hi int
	if h <= roomL(r) {
		lo, hi = h+1, roomL(r)
	} else {
		lo, hi = roomR(r), h-1
	}
	mask := (uint32(16) << (4 * hi)) - (1 << (4 * lo))
	return s.hall.hall&mask != 0
}

func (s *State) forceOne() bool {
	bits := utils.Bits{N: uint64(s.hall.mask())}
	for bits.Next() {
		b := bits.Get()
		h := int(b / 4)
		r := s.hall.get(h)
		if s.room.empty(r) && !s.obstructed(int(r), h) {
			s.hall.clear(h)
			return true
		}
	}
	for r := 0; r < 4; r++ {
		if s.room.empty(r) {
			continue
		}
		g := s.room.get(r)
		if g == r || !s.room.empty(g) {
			continue
		}
		var x int
		if r < g {
			x = roomR(g)
		} else {
			x = roomL(g)
		}
		if !s.obstructed(r, x) {
			s.room.pop(r)
			return true
		}
	}
	return false
}

func (s *State) deadlocked() bool {
	h43 := s.hall.hall & 0x0077000
	if h43 == 0x0047000 || h43 == 0x0057000 {
		return true
	}
	h42 := s.hall.hall & 0x0070700
	if h42 == 0x0040700 {
		return true
	}
	h32 := s.hall.hall & 0x0007700
	if h32 == 0x0004600 || h32 == 0x0004700 {
		return true
	}
	return false
}

func (s *State) crowded() bool {
	h0 := 0
	h := s.hall.hall>>2 | 0x10000000
	var satisfied bool
	for i := 0; i < 8; i++ {
		if h&1 != 0 {
			if h0 < i {
				r0 := utils.Max(0, h0-2)
				r1 := utils.Max(3, i-2)
				space := i - h0
				mask := uint32(3) << (2 * space)
				for r := r0; r <= r1; r++ {
					rr := (s.room.room >> (8 * r)) & 0xff
					if rr&mask == 0 {
						satisfied = true
					}
				}
			}
			h0 = i + 1
		}
		h >>= 4
	}
	return !satisfied
}

type neighb struct {
	a int
	b State
	c uint32
}

func (s *State) neighbors(skip uint32) []neighb {
	ns := make([]neighb, 0)
	skipRooms := 0
	for i := 0; i < 3; i++ {
		h := i + 2
		if !s.hall.empty(h) {
			mask := 0b1110 << i
			if i < s.hall.get(h) {
				skipRooms |= ^mask
			} else {
				skipRooms |= mask
			}
		}
	}
	for r := 0; r < 4; r++ {
		if skipRooms&(1<<r) != 0 || s.room.empty(r) {
			continue
		}
		g := s.room.get(r)
		var lo, hi int
		if r < g {
			lo, hi = roomR(r), roomL(g)
		} else if r > g {
			lo, hi = roomR(g), roomL(r)
		} else {
			lo, hi = roomL(r), roomR(r)
		}
		for h := 0; h < 7; h++ {
			if r != g && h >= lo && h <= hi {
				continue
			}
			skipIdx := 8*r + h
			if (skip>>skipIdx)&1 != 0 {
				continue
			}
			if !s.hall.empty(h) || s.obstructed(r, h) {
				continue
			}
			var cost int
			if h < lo {
				cost = lo - h
			} else if hi < h {
				cost = h - hi
			}
			cost *= 2
			cost -= utils.ToInt[int]((utils.ToInt[int](cost == 0)|utils.ToInt[int](r == g)) == 0) +
				(utils.ToInt[int](h == 0) | utils.ToInt[int](h == 6))
			cost *= 2
			n := *s
			n.room.pop(r)
			n.hall.set(h, g)
			if n.deadlocked() {
				continue
			}
			skips := safeSkips[skipIdx]
			for n.forceOne() {
				skips = 0
			}
			if n.crowded() {
				continue
			}
			ns = append(ns, neighb{cost * int(costArr[g]), n, skips})
		}
	}
	return ns
}

type val struct {
	a uint64
	b uint16
	c uint32
}

const SIZE int = 14983

type Hash struct {
	table [SIZE]val
}

func (h *Hash) find(key uint64) int {
	idx := int(key % uint64(SIZE))
	for h.table[idx].a != 0 && h.table[idx].a != ^key {
		idx++
		idx &= -utils.ToInt[int](idx < SIZE)
	}
	return idx
}

func (h *Hash) insert(key uint64, v1 uint16, v2 uint32) {
	idx := h.find(key)
	h.table[idx] = val{^key, v1, v2}
}

func (h *Hash) get(idx int) (uint16, uint32) {
	return h.table[idx].b, h.table[idx].c
}

func (h *Hash) exists(idx int) bool {
	return h.table[idx].a != 0
}

type pair struct {
	a int
	b uint64
}

func less(a, b pair) bool {
	if a.a == b.a {
		return a.b < b.b
	}
	return a.a < b.a
}

func solve(start State) int {
	cost := &Hash{}
	cost.insert(start.hash(), 0, 0)
	q := utils.NewPQueue(less)
	q.Push(pair{0, start.hash()})
	for q.Len() > 0 {
		elem := q.Pop()
		queueCost, curHash := elem.a, elem.b
		curCost, curSkips := cost.get(cost.find(curHash))
		if queueCost != int(curCost) {
			continue
		}
		cur := NewState(curHash)
		if cur.solved() {
			break
		}
		for _, nb := range cur.neighbors(curSkips) {
			delta, state, skips := nb.a, nb.b, nb.c
			hash := state.hash()
			newCost := int(curCost) + delta
			newIdx := cost.find(hash)
			if !cost.exists(newIdx) {
				cost.insert(hash, uint16(newCost), skips)
				q.Push(pair{newCost, hash})
			} else {
				prevCost, _ := cost.get(newIdx)
				if newCost == int(prevCost) {
					cost.table[newIdx].c &= skips
				} else if newCost < int(prevCost) {
					cost.table[newIdx].b = uint16(newCost)
					cost.table[newIdx].c = skips
					q.Push(pair{newCost, hash})
				}
			}

		}
	}
	ans, _ := cost.get(cost.find(0))
	return int(ans)
}

func Part1(input string) interface{} {
	room := readInput(input)
	base, _ := baseCost(room)
	return base + solve(NewState(uint64(room)))
}

func Part2(input string) interface{} {
	room := readInput(input)
	_, base := baseCost(room)
	return base + solve(NewState(uint64(insertPart2(room))))
}
