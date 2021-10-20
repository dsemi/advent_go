package year2015

import (
	"advent/problems"
	"advent/utils"
	"math"
	"strings"
)

type Day21 struct{}

type Equip struct {
	cost, damage, armor int
}

type Person struct {
	Equip
	hp int
}

var shop1 = []Equip{
	Equip{cost: 8, damage: 4, armor: 0},  // Dagger
	Equip{cost: 10, damage: 5, armor: 0}, // Shortsword
	Equip{cost: 25, damage: 6, armor: 0}, // Warhammer
	Equip{cost: 40, damage: 7, armor: 0}, // Longsword
	Equip{cost: 74, damage: 8, armor: 0}, // Greataxe
}

var shop2 = []Equip{
	Equip{cost: 13, damage: 0, armor: 1},  // Leather
	Equip{cost: 31, damage: 0, armor: 2},  // Chainmail
	Equip{cost: 53, damage: 0, armor: 3},  // Splintmail
	Equip{cost: 75, damage: 0, armor: 4},  // Bandedmail
	Equip{cost: 102, damage: 0, armor: 5}, // Platemail
}

var shop3 = []Equip{
	Equip{cost: 25, damage: 1, armor: 0},  // Damage +1
	Equip{cost: 50, damage: 2, armor: 0},  // Damage +2
	Equip{cost: 100, damage: 3, armor: 0}, // Damage +3
	Equip{cost: 20, damage: 0, armor: 1},  // Defense +1
	Equip{cost: 40, damage: 0, armor: 2},  // Defense +2
	Equip{cost: 80, damage: 0, armor: 3},  // Defense +3
	Equip{cost: 0, damage: 0, armor: 0},   // None
}

func (Day21) person(equips ...Equip) Person {
	p := Person{hp: 100}
	for _, equip := range equips {
		p.cost += equip.cost
		p.damage += equip.damage
		p.armor += equip.armor
	}
	return p
}

func (d Day21) allEquipCombos() []Person {
	var v []Person
	for _, weapon := range shop1 {
		for _, armor := range shop2 {
			for i, ring1 := range shop3 {
				for _, ring2 := range shop3[i+1:] {
					v = append(v, d.person(weapon, armor, ring1), d.person(weapon, armor, ring1, ring2))
				}
			}
		}
	}
	return v
}

func (Day21) parseBoss(input string) Person {
	var v []int
	for _, line := range strings.Split(input, "\n") {
		pts := strings.Fields(line)
		v = append(v, utils.Int(pts[len(pts)-1]))
	}
	return Person{
		hp: v[0],
		Equip: Equip{
			damage: v[1],
			armor:  v[2],
		},
	}
}

func (Day21) isWinning(boss, player Person) bool {
	ttd := func(p1, p2 Person) int {
		q := p1.hp / utils.Max(1, p2.damage-p1.armor)
		if p1.hp%utils.Max(1, p2.damage-p1.armor) != 0 {
			q++
		}
		return q
	}
	return ttd(player, boss) >= ttd(boss, player)
}

func (d Day21) Part1(input string) interface{} {
	boss := d.parseBoss(input)
	m := math.MaxInt
	for _, p := range d.allEquipCombos() {
		if d.isWinning(boss, p) {
			m = utils.Min(m, p.cost)
		}
	}
	return m
}

func (d Day21) Part2(input string) interface{} {
	boss := d.parseBoss(input)
	m := 0
	for _, p := range d.allEquipCombos() {
		if !d.isWinning(boss, p) {
			m = utils.Max(m, p.cost)
		}
	}
	return m
}

func init() {
	problems.Register(Day21{})
}
