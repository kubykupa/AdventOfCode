package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func nice1(line string) bool {
	var vowels int
	var haspair bool
	bad := map[string]bool{"ab": true, "cd": true, "pq": true, "xy": true}
	for i := 0; i < len(line); i++ {
		c := line[i]

		if strings.Contains("aeiou", string(c)) {
			vowels++
		}
		if i > 0 {
			slice := line[i-1 : i+1]
			if bad[slice] {
				//fmt.Println("BAN for: ", slice)
				return false
			}
			if line[i-1] == line[i] {
				haspair = true
				//fmt.Print(" slice: ", slice, " ")
			}
		}
	}
	return haspair && vowels > 2
}

func nice2(line string) bool {
	var f1, f2 bool
	for i := 0; i < len(line)-2; i++ {
		slice := line[i : i+2]
		src := line[i+2:]
		//fmt.Println("src: ", src, "slice: ", slice)
		if !f1 && strings.Contains(src, slice) {
			f1 = true
		}
		if line[i] == line[i+2] {
			f2 = true
		}
		if f1 && f2 {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.OpenFile("input_05.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_05", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var n1, n2 int
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		//fmt.Println("Line: ", line)
		if nice1(line) {
			n1++
		}
		if nice2(line) {
			n2++
		}
		fmt.Println(line, " nice2: ", n2)
	}
	fmt.Println("NICE 1: ", n1)
	fmt.Println("NICE 2: ", n2)
}
