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

var Happy = make(map[[2]string]int)
var Names []string
var Name2Index = make(map[string]int)

func parse(line string) (string, string, int) {
	sub := strings.Split(line, " ")

	from, to := sub[0], sub[10]
	happy, _ := strconv.Atoi(sub[3])

	if _, ok := Name2Index[from]; !ok {
		Name2Index[from] = len(Names)
		Names = append(Names, from)
	}
	if _, ok := Name2Index[to]; !ok {
		Name2Index[to] = len(Names)
		Names = append(Names, to)
	}

	if sub[2] == "lose" {
		happy = -happy
	}

	Happy[[2]string{from, to}] = happy

	return from, to, happy
}

func getHappy(from, to int, show bool) int {
	var res int
	A := Names[from]
	B := Names[to]

	happy, _ := Happy[[2]string{A, B}]
	res += happy
	if show {
		fmt.Print(A, ",", B, ":", happy, " + ")
	}
	happy, _ = Happy[[2]string{B, A}]
	res += happy
	if show {
		fmt.Println(B, ",", A, ":", happy, "=", res)
	}
	return res
}

func calc(way []int, show bool) int {
	var res int
	for i := 1; i < len(way); i++ {
		res += getHappy(way[i-1], way[i], show)
	}
	res += getHappy(way[len(way)-1], way[0], show)

	if show {
		fmt.Println("Total happy: ", res)
	}
	return res
}

type Result struct {
	Happy int
	Way   []int
}

func getMask(m uint) string {
	return fmt.Sprintf("%08b", m)
}

func mutate(way []int, mask uint, res *Result) {
	N := len(Names)

	//fmt.Println("way:", way, " N: ", N, "used:", getMask(mask))
	if N == len(way) {
		h := calc(way, false)
		if h > res.Happy {
			res.Happy = h
			res.Way = way

			//calc(way, true)
			//fmt.Println("BEST: ", h, way)
		}
		return
	}

	for i := 0; i < N; i++ {
		bit := mask & (1 << uint(i))
		if bit == 0 {
			w := append(way, i)
			m := mask | (1 << uint(i))
			//fmt.Println("way:", w, i, getMask(m))
			mutate(w, m, res) // fix level head element i
		}
	}
}

func main() {
	//f, err := os.OpenFile("easy_13", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_13.txt", os.O_RDONLY, 0666)
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
		line = strings.TrimRight(line, ".\n\r")
		from, to, happy := parse(line)
		fmt.Println(from, to, happy)
	}
	var res Result
	res.Happy = MinInt
	mutate([]int{}, uint(0), &res)
	fmt.Println("RESULT: ", res.Happy, res.Way)

	fmt.Println("With me:")
	Names = append(Names, "ME")
	res.Happy = MinInt
	mutate([]int{}, uint(0), &res)
	fmt.Println("RESULT: ", res.Happy, res.Way)
}
