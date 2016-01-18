package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lines []string

func exec(a, b int) int {
	var pos int
	var off int

	for pos < len(lines) {
		cmd := strings.Split(lines[pos], " ")
		//fmt.Println(pos, " cmd:", cmd)

		off = 1
		switch cmd[0] {
		case "hlf":
			r := &b
			if cmd[1] == "a" {
				r = &a
			}
			(*r) /= 2
			break
		case "tpl":
			r := &b
			if cmd[1] == "a" {
				r = &a
			}
			(*r) *= 3
			break
		case "inc":
			r := &b
			if cmd[1] == "a" {
				r = &a
			}
			(*r)++
			break
		case "jmp":
			off, _ = strconv.Atoi(cmd[1])
			break
		case "jie":
			r := &b
			if cmd[1] == "a" {
				r = &a
			}
			if (*r)%2 == 0 {
				off, _ = strconv.Atoi(cmd[2])
			}
			break
		case "jio":
			r := &b
			if cmd[1] == "a" {
				r = &a
			}
			if (*r) == 1 {
				off, _ = strconv.Atoi(cmd[2])
			}
			break
		}

		pos += off
	}

	return b
}

func main() {
	f, err := os.OpenFile("input_23.txt", os.O_RDONLY, 0666)
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
		line = strings.Replace(line, ",", "", -1)
		lines = append(lines, line)
	}

	fmt.Println("Part 1:", exec(0, 0))
	fmt.Println("Part 2:", exec(1, 0))
}
