package main

import (
	"fmt"
)

const (
	SKIP         = 0
	MAGIC_MISSLE = 1
	DRAIN        = 2
	SHIELD       = 3
	POISON       = 4
	RECHARGE     = 5
)

type Spell struct {
	Type     int
	Name     string
	ManaCost int
	Time     int
}

func (spell *Spell) String() string {
	return fmt.Sprintf("{%s %d}", spell.Name, spell.Time)
}

type Pers struct {
	HP, Mana int
	DMG, DEF int
	Bufs     map[int]*Spell
}

func calcDMG(dmg, def int) int {
	if dmg > def {
		return dmg - def
	}
	return 1
}

func Clone(pers *Pers) (res Pers) {
	res = Pers{pers.HP, pers.Mana, pers.DMG, pers.DEF, map[int]*Spell{}}
	for _, ptr := range pers.Bufs {
		s := *ptr
		res.Bufs[s.Type] = &s
	}
	return
}

func Begin(spell *Spell, caster, receiver *Pers) {
	switch spell.Type {
	case SHIELD:
		caster.DEF += 7
		break
	}
}

func End(spell *Spell, caster, receiver *Pers) {
	switch spell.Type {
	case SHIELD:
		caster.DEF -= 7
		break
	}
}

func Action(spell *Spell, caster, receiver *Pers) {
	switch spell.Type {
	case MAGIC_MISSLE:
		receiver.HP -= 4
		break
	case DRAIN:
		receiver.HP -= 2
		caster.HP += 2
		break
	case POISON:
		receiver.HP -= 3
		break
	case RECHARGE:
		caster.Mana += 101
		break
	}
}

func Cast(spell Spell, caster, receiver *Pers) bool {
	//fmt.Println("Try cast:", spell)
	if caster.Mana >= spell.ManaCost {
		if spell.Time > 0 {
			if _, ok := caster.Bufs[spell.Type]; ok {
				return false
			} else {
				caster.Bufs[spell.Type] = &spell
				Begin(&spell, caster, receiver)
			}
		} else {
			Action(&spell, caster, receiver)
		}
		caster.Mana -= spell.ManaCost
		//fmt.Println("Cast:", spell.Name, "->", *caster, *receiver)
		return true
	}
	return false
}

func Update(spell *Spell, caster, receiver *Pers) bool {
	if spell.Time > 0 {
		spell.Time--
		Action(spell, caster, receiver)
		return true
	} else {
		End(spell, caster, receiver)
		delete(caster.Bufs, spell.Type)
		//fmt.Println("End buf:", spell, "hero:", *caster)
	}
	return false
}

var SPELLS = []Spell{
	{MAGIC_MISSLE, "MagicMissle", 53, 0},
	{DRAIN, "Drain", 73, 0},
	{SHIELD, "Shield", 113, 6},
	{POISON, "Poison", 173, 6},
	{RECHARGE, "Recharge", 229, 5},
}

func BestMana(usedMana int, minMana *int) {
	if *minMana > usedMana {
		*minMana = usedMana
	}
}

func battle(hero, boss *Pers, usedMana int, minMana *int, log string, depth int) {
	if hero.Mana < SPELLS[0].ManaCost {
		fmt.Println("To low mana:", log)
		return
	}
	//fmt.Println("IN", depth, "log:", log)

	for _, spell := range SPELLS {
		h := Clone(hero)
		b := Clone(boss)

		//fmt.Println("-- Player turn --")
		//fmt.Println("- Player has", h.HP, "hit points,", h.DEF, "armor,", h.Mana, "mana bufs:", h.Bufs)
		//fmt.Println("- Boss has", b.HP, "hit points")
		if Cast(spell, &h, &b) {
			m := usedMana + spell.ManaCost
			l := log + fmt.Sprint("[+", spell.Name, "] ")
			for _, buf := range h.Bufs {
				Update(buf, &h, &b)
			}

			if b.HP < 1 {
				BestMana(m, minMana)
				fmt.Println("WIN: ", *minMana, "hero:", h, " log:", l)
				continue
			}
			//fmt.Println("-- Boss turn --")
			//fmt.Println("- Player has", h.HP, "hit points,", h.DEF, "armor,", h.Mana, "mana")
			//fmt.Println("- Boss has", b.HP, "hit points")
			h.HP -= calcDMG(b.DMG, h.DEF)
			if h.HP < 1 {
				fmt.Println("DIE", depth, "log:", log)
				continue
			}

			for _, buf := range h.Bufs {
				Update(buf, &h, &b)
			}
			if b.HP < 1 {
				BestMana(m, minMana)
				fmt.Println("WIN: ", *minMana, "hero:", h, " log:", l)
				continue
			}
			battle(&h, &b, m, minMana, l, depth+1)
		} else {
			log += fmt.Sprint("[-", spell.Name, "] ")
		}
	}

	//fmt.Println("OUT", depth, "log:", log)
	if depth == 0 {
		fmt.Println("Min Mana:", *minMana)
	}
}

func run(wizHP, wizMana, bossHP, bossDMG int) int {
	minMana := 100000
	def := map[int]*Spell{}

	battle(&(Pers{wizHP, wizMana, 0, 0, def}),
		&(Pers{bossHP, 0, bossDMG, 0, def}),
		0, &minMana, "", 0)
	return minMana
}

func main() {
	//fmt.Println("\ntest 1:", run(10, 250, 13, 8))
	//fmt.Println("\ntest 2:", run(10, 250, 14, 8))

	fmt.Println("result:", run(50, 500, 71, 10))
}
