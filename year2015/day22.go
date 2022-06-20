package main

import (
	"math"
	"strings"
	"utils"
)

type Game struct {
	playerHealth, playerMana, playerArmor   int
	bossHealth, bossDamage                  int
	shieldTurns, poisonTurns, rechargeTurns int
	spentMana                               int
}

type Spell struct {
	cost   int
	effect func(*Game)
	active func(*Game) bool
}

var spells = []Spell{
	Spell{ // Magic Missile
		cost:   53,
		effect: func(state *Game) { state.bossHealth -= 4 },
		active: func(_ *Game) bool { return false },
	},
	Spell{ // Drain
		cost: 73,
		effect: func(state *Game) {
			state.playerHealth += 2
			state.bossHealth -= 2
		},
		active: func(_ *Game) bool { return false },
	},
	Spell{ // Shield
		cost: 113,
		effect: func(state *Game) {
			state.playerArmor += 7
			state.shieldTurns = 6
		},
		active: func(state *Game) bool { return state.shieldTurns > 0 },
	},
	Spell{ // Poison
		cost:   173,
		effect: func(state *Game) { state.poisonTurns = 6 },
		active: func(state *Game) bool { return state.poisonTurns > 0 },
	},
	Spell{ // Recharge
		cost:   229,
		effect: func(state *Game) { state.rechargeTurns = 5 },
		active: func(state *Game) bool { return state.rechargeTurns > 0 },
	},
}

func (state *Game) applyEffects() {
	if state.shieldTurns > 0 {
		if state.shieldTurns == 1 {
			state.playerArmor -= 7
		}
		state.shieldTurns--
	}
	if state.poisonTurns > 0 {
		state.bossHealth -= 3
		state.poisonTurns--
	}
	if state.rechargeTurns > 0 {
		state.playerMana += 101
		state.rechargeTurns--
	}
}

func parseBoss(input string) *Game {
	var v []int
	for _, line := range strings.Split(input, "\n") {
		v = append(v, utils.Int(strings.Split(line, ": ")[1]))
	}
	return &Game{
		playerHealth: 50,
		playerMana:   500,
		bossHealth:   v[0],
		bossDamage:   v[1],
	}
}

func minCostToWin(s *Game, hard bool) int {
	states := []*Game{s}
	result := math.MaxInt
	for len(states) > 0 {
		state := states[len(states)-1]
		states = states[:len(states)-1]
		if hard {
			state.playerHealth--
			if state.playerHealth <= 0 {
				continue
			}
		}
		state.applyEffects()
		if state.bossHealth <= 0 {
			result = utils.Min(result, state.spentMana)
			continue
		}
		for _, spell := range spells {
			if state.playerMana >= spell.cost && state.spentMana+spell.cost < result && !spell.active(state) {
				newState := new(Game)
				*newState = *state
				newState.playerMana -= spell.cost
				newState.spentMana += spell.cost
				spell.effect(newState)
				newState.applyEffects()
				if newState.bossHealth <= 0 {
					result = newState.spentMana
					continue
				}
				newState.playerHealth -= utils.Max(1, newState.bossDamage-newState.playerArmor)
				if newState.playerHealth > 0 {
					states = append(states, newState)
				}
			}
		}
	}
	return result
}

func Part1(input string) interface{} {
	return minCostToWin(parseBoss(input), false)
}

func Part2(input string) interface{} {
	return minCostToWin(parseBoss(input), true)
}
