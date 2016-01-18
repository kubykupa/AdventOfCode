package main

import (
	"fmt"
)

const (
	PRESENT_COUNT = 33100000
	MAX_STEPS     = 50
)

func part1() {
	number := PRESENT_COUNT

	var elfs []int
	for house := 1; house < PRESENT_COUNT; house++ {
		p := 0
		elfs = append(elfs, house)
		for e, h := range elfs {
			if h > house {
				continue
			}
			if h == house {
				//fmt.Println("elf:", e, " moved by:", e+1, " to:", h+e+1)
				elfs[e] += e + 1
				p += 10 * (e + 1)
			}
		}
		if p >= PRESENT_COUNT {
			if house < number {
				number = house
				fmt.Println("Part1: house №:", house, "presents:", p)
				return
			}
		}
	}
}

func part2() {
	gifts := 11
	stopHouse := PRESENT_COUNT
	//stopHouse := 10

	var houses = map[int]int{}
	var elfs = map[int]int{}

	for house := 1; house < stopHouse; house++ {
		elfs[house] = house
		for e, h := range elfs {
			p := houses[h] + gifts*e
			//fmt.Println("Elf", e, "at house:", h, "step:", h/e, "presents:", houses[h], "->", p)
			houses[h] = p
			if p >= PRESENT_COUNT {
				if h < stopHouse {
					stopHouse = h
					fmt.Println("New stop house №:", stopHouse, "presents:", p)
				}
			}

			if h/e == MAX_STEPS {
				//fmt.Println("Stop elf:", e, "in house:", h)
				delete(elfs, e)
			} else {
				step := h + e
				elfs[e] = step
			}
		}

		//fmt.Println(house, ":", houses[house])
		delete(houses, house)
	}
	fmt.Println("Part2:", stopHouse)
}

func main() {
	//part1()
	part2()

}
