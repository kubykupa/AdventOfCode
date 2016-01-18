package main

import (
	"fmt"
)

type Equip struct {
	Cost     int
	DMG, DEF int
}

type Pers struct {
	HP, DMG, DEF int
	Cost         int
}

var weapons = []Equip{
	{0, 0, 0},
	{8, 4, 0},  // Dagger
	{10, 5, 0}, // Shortsword
	{25, 6, 0}, // Warhammer
	{40, 7, 0}, // Longsword
	{74, 8, 0}, // Greataxe
}

var armors = []Equip{
	{0, 0, 0},
	{13, 0, 1},  // Leather
	{31, 0, 2},  // Chainmail
	{53, 0, 3},  // Splintmail
	{75, 0, 4},  // Bandedmail
	{102, 0, 5}, // Platemail
}

var rings = []Equip{
	{0, 0, 0},
	{25, 1, 0},  // Damage +1
	{50, 2, 0},  // Damage +2
	{100, 3, 0}, // Damage +3
	{20, 0, 1},  // Defense +1
	{40, 0, 2},  // Defense +2
	{80, 0, 3},  // Defense +3
}

func calcDMG(dmg, def int) int {
	if dmg > def {
		return dmg - def
	}
	return 1
}

func Battle(hero, boss Pers) bool {
	heroAttack := calcDMG(hero.DMG, boss.DEF)
	bossAttack := calcDMG(boss.DMG, hero.DEF)
	for {
		boss.HP -= heroAttack
		if boss.HP < 1 {
			return true
		}
		hero.HP -= bossAttack
		if hero.HP < 1 {
			break
		}
		//fmt.Println(hero, boss)
	}
	return false
}

func main() {
	//Battle(Pers{8, 5, 5, -1}, Pers{12, 7, 2, -1})
	//return

	boss := Pers{103, 9, 2, -1}

	minCost := 10000000
	maxCost := 0
	for _, w := range weapons {
		for _, a := range armors {
			for _, l := range rings {
				for _, r := range rings {
					hero := Pers{
						100,
						w.DMG + a.DMG + l.DMG + r.DMG,
						a.DEF + l.DEF + r.DEF,
						w.Cost + a.Cost + l.Cost + r.Cost}

					if Battle(hero, boss) {
						if hero.Cost < minCost {
							minCost = hero.Cost
						}
					} else {
						if w.Cost > 0 && r.Cost != l.Cost && hero.Cost > maxCost {
							fmt.Println(w, a, l, r, " = ", hero.Cost)
							maxCost = hero.Cost
						}
					}
				}
			}
		}
	}
	fmt.Println("Win with min cost:", minCost)
	fmt.Println("Lose with max cost:", maxCost)

}
