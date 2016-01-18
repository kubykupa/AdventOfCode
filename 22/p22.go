package main

import (
	"fmt"
)

type GameState struct {
	Hard                                 bool
	Actor                                bool
	WizHP, BossHP, Mana                  int
	ShieldTime, PoisonTime, RechargeTime int
	//Hystory                              string
}

var StateCache = map[GameState]int{}
var BossDMG int

const (
	MAGIC_MISSILE_COST = 53
	DRAIN_COST         = 73
	SHIELD_COST        = 113
	POISON_COST        = 173
	RECHARGE_COST      = 229

	NONE = -8
)

func calcDMG(dmg, def int) int {
	if dmg > def {
		return dmg - def
	}
	return 1
}

func Battle(state GameState) int {
	//func Battle(state GameState) (int, string) {
	if minMana, ok := StateCache[state]; ok {
		return minMana
		//return minMana, state.Hystory
	}
	minMana := Round(state)
	StateCache[state] = minMana
	//return minMana, state.Hystory
	return minMana
}

func Best(minMana, cost int, state GameState) int {
	//res, _ := Battle(state)
	res := Battle(state)
	if res == NONE {
		return minMana
	}
	res += cost
	if minMana == NONE {
		return res
	}
	if res < minMana {
		return res
	}
	return minMana
}

func Round(state GameState) int {
	if state.Hard && state.Actor {
		state.WizHP--
		if state.WizHP < 1 {
			return NONE
		}
	}
	if state.WizHP < 1 {
		return NONE
	}

	def := 0
	if state.ShieldTime > 0 {
		state.ShieldTime--
		if state.ShieldTime > 0 {
			def = 7
		}
	}

	if state.PoisonTime > 0 {
		state.BossHP -= 3
		state.PoisonTime--
	}
	if state.RechargeTime > 0 {
		state.Mana += 101
		state.RechargeTime--
	}

	if state.BossHP < 1 {
		fmt.Println("WIN:", state)
		return 0
	}

	if !state.Actor {
		res := Battle(GameState{state.Hard, true, state.WizHP - calcDMG(BossDMG, def), state.BossHP, state.Mana, state.ShieldTime, state.PoisonTime, state.RechargeTime /*, state.Hystory*/})
		return res
	}

	res := NONE
	if state.Mana >= MAGIC_MISSILE_COST {
		res = Best(res, MAGIC_MISSILE_COST, GameState{state.Hard, false, state.WizHP, state.BossHP - 4, state.Mana - MAGIC_MISSILE_COST, state.ShieldTime, state.PoisonTime, state.RechargeTime /*, state.Hystory + "M "*/})
	}

	if state.Mana >= DRAIN_COST {
		res = Best(res, DRAIN_COST, GameState{state.Hard, false, state.WizHP + 2, state.BossHP - 2, state.Mana - DRAIN_COST, state.ShieldTime, state.PoisonTime, state.RechargeTime /*, state.Hystory + "D "*/})
	}

	if state.Mana >= SHIELD_COST {
		res = Best(res, SHIELD_COST, GameState{state.Hard, false, state.WizHP, state.BossHP, state.Mana - SHIELD_COST, 6, state.PoisonTime, state.RechargeTime /*, state.Hystory + "S "*/})
	}

	if state.Mana >= POISON_COST {
		res = Best(res, POISON_COST, GameState{state.Hard, false, state.WizHP, state.BossHP, state.Mana - POISON_COST, state.ShieldTime, 6, state.RechargeTime /*, state.Hystory + "P "*/})
	}

	if state.Mana >= RECHARGE_COST {
		res = Best(res, RECHARGE_COST, GameState{state.Hard, false, state.WizHP, state.BossHP, state.Mana - RECHARGE_COST, state.ShieldTime, state.PoisonTime, 5 /*,state.Hystory + "R "*/})
	}
	return res
}

func Fight(label string, hard bool, wizHP, bossHP, mana, dmg int) {
	BossDMG = dmg
	mana = Battle(GameState{hard, true, wizHP, bossHP, mana, 0, 0, 0 /*, ""*/})
	fmt.Println(label, mana /*, log*/)
}

func main() {
	//Fight("Test 1:", 10, 13, 250, 8)
	//Fight("Test 2:", 10, 14, 250, 8)

	Fight("Part 1:", false, 50, 71, 500, 10)
	Fight("Part 2:", true, 50, 71, 500, 10)
}
