package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint uint = ^uint(0)
const MaxInt int = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var Lens = make(map[[2]string]int)
var Town []string
var Names = make(map[string]int)

func parse(line string) (string, string, int) {
	sub := strings.Split(line, " ")
	from, to := sub[0], sub[2]
	dist, _ := strconv.Atoi(sub[4])

	if _, ok := Names[from]; !ok {
		Names[from] = len(Names)
		Town = append(Town, from)
	}
	if _, ok := Names[to]; !ok {
		Names[to] = len(Names)
		Town = append(Town, to)
	}
	Lens[[2]string{from, to}] = dist
	Lens[[2]string{to, from}] = dist

	return from, to, dist
}

func calc(way []int, show bool) int {
	var res int
	if show {
		fmt.Print(way, way[0], ":", Town[way[0]], "->")
	}
	for i := 1; i < len(way); i++ {
		dist, _ := Lens[[2]string{Town[way[i-1]], Town[way[i]]}]
		res += dist
		if show {
			fmt.Print("(", dist, ")", way[i], ":", Town[way[i]], "->")
		}
	}
	if show {
		fmt.Println("dist: ", res)
	}
	return res
}

type Result struct {
	Length int
	Way    []int
}

func getMask(m uint) string {
	return fmt.Sprintf("%08b", m)
}

func mutate(way []int, mask uint, res *Result) {
	N := len(Town)

	//fmt.Println("way:", way, " N: ", N, "used:", getMask(mask))
	if N == len(way) {
		d := calc(way, true)
		if d > res.Length {
			res.Length = d
			res.Way = way
		}
		fmt.Println()
		return
	}

	for i := 0; i < N; i++ {
		bit := mask & (1 << uint(i))
		if bit == 0 {
			w := append(way, i)
			m := mask | (1 << uint(i))
			fmt.Println("way:", w, i, getMask(m))
			mutate(w, mask|(1<<uint(i)), res) // fix level head element i
		}
	}
}

func main() {
	//f, err := os.OpenFile("easy_09", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_09.txt", os.O_RDONLY, 0666)
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
		from, to, dist := parse(line)
		fmt.Println(from, to, dist)
	}
	var res Result
	res.Length = MinInt

	mutate([]int{}, uint(0), &res)
	fmt.Println("RESULT: ", res.Length, res.Way)
}
