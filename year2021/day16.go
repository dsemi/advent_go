package main

import "utils"

type bitstream struct {
	vsum uint64
	bits []uint8
}

func newBs(input string) *bitstream {
	bs := &bitstream{}
	for _, c := range input {
		var b rune
		if c >= 'A' {
			b = c - 'A' + 10
		} else {
			b = c - '0'
		}
		for i := 3; i >= 0; i-- {
			bs.bits = append(bs.bits, uint8((b>>i)&1))
		}
	}
	return bs
}

func (b *bitstream) next(n int) uint64 {
	var v uint64
	for i := 0; i < n; i++ {
		v = v<<1 | uint64(b.bits[i])
	}
	b.bits = b.bits[n:]
	return v
}

func (b *bitstream) packet() uint64 {
	b.vsum += b.next(3)
	typeId := b.next(3)
	if typeId == 4 {
		var n uint64
		for b.next(1) == 1 {
			n = n<<4 | b.next(4)
		}
		return n<<4 | b.next(4)
	}
	ns := make([]uint64, 0)
	if b.next(1) == 0 {
		i := b.next(15)
		r := &bitstream{bits: b.bits[:i]}
		b.bits = b.bits[i:]
		for len(r.bits) > 0 {
			ns = append(ns, r.packet())
		}
		b.vsum += r.vsum
	} else {
		for i := b.next(11); i > 0; i-- {
			ns = append(ns, b.packet())
		}
	}
	switch typeId {
	case 0:
		return utils.Sum(ns)
	case 1:
		return utils.Product(ns)
	case 2:
		return utils.Minimum(ns)
	case 3:
		return utils.Maximum(ns)
	case 5:
		return utils.IntBool(ns[0] > ns[1])
	case 6:
		return utils.IntBool(ns[0] < ns[1])
	case 7:
		return utils.IntBool(ns[0] == ns[1])
	default:
		panic("Bad type id")
	}
}

func Part1(input string) interface{} {
	b := newBs(input)
	b.packet()
	return b.vsum
}

func Part2(input string) interface{} {
	return newBs(input).packet()
}
