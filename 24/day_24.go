package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	GROUPS = 4
	BAD    = 666666
)

type Val [2]int

var PACKS []int
var GROUP_SUM int

var BEST = Val{BAD, BAD}
var FIRST = map[uint]bool{}
var SECOND = map[uint]uint8{}

func getMask(mask uint) string {
	p := ""
	s := 0
	for i := 0; i < len(PACKS); i++ {
		bit := mask & (1 << uint(i))
		if bit > 0 {
			p += fmt.Sprintf("%d ", PACKS[i])
			s += PACKS[i]
		}
	}
	return fmt.Sprintf("%sval: %d", p, mask)
}

func best(a, b Val) (Val, bool) {
	if a[0] < b[0] {
		return a, true
	} else if a[0] == b[0] {
		if a[1] < b[1] {
			return a, true
		}
	}
	return b, false
}

func second(mask uint, group uint8, sum int) uint8 {
	if sum > GROUP_SUM {
		return 0
	}
	if v, ok := SECOND[mask]; ok {
		return v
	}
	if sum == GROUP_SUM {
		if group == GROUPS {
			return 1
		} else {
			return second(mask, group+1, 0)
		}
	}

	res := uint8(0)
	for i := 0; i < len(PACKS); i++ {
		bit := mask & (1 << uint(i))
		if bit == 0 {
			res = second(mask|(1<<uint(i)), group, sum+PACKS[i])
			if res == 1 {
				break
			}
		}
	}
	SECOND[mask] = res
	return res
}

func first(mask uint, sum, size, qe int) {
	//fmt.Println("mask:", getMask(mask), "sum:", sum, "size:", size, "qe:", qe)
	if sum > GROUP_SUM || size > BEST[0] {
		return
	}
	if _, ok := FIRST[mask]; ok {
		return
	}

	if sum == GROUP_SUM {
		//fmt.Println("Step:", state.Step, "sum:", stepSum, "for:", show(&state))
		v := Val{size, qe}
		if _, ok := best(v, BEST); ok {
			if second(mask, 2, 0) == 1 {
				fmt.Println("\nNew best:", BEST, "to:", v, " mask:", getMask(mask))
				BEST = v
			}
		}
		return
	}

	FIRST[mask] = true

	for i := 0; i < len(PACKS); i++ {
		bit := mask & (1 << uint(i))
		if bit == 0 {
			first(mask|(1<<uint(i)), sum+PACKS[i], size+1, qe*PACKS[i])
		}
	}
}

func main() {
	f, err := os.OpenFile("input_24.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_24", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		p, _ := strconv.Atoi(line)
		PACKS = append(PACKS, p)
		GROUP_SUM += p
	}
	GROUP_SUM /= GROUPS

	fmt.Println("Size:", len(PACKS), "group sum:", GROUP_SUM)
	sort.Sort(sort.Reverse(sort.IntSlice(PACKS)))
	fmt.Println("Packs:", PACKS)

	first(0, 0, 0, 1)
	fmt.Println("BEST:", BEST)
}
