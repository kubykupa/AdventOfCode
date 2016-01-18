package main

import (
	"fmt"
)

func lookAndSay(line string, T int) int {
	for t := 0; t < T; t++ {
		count := 1
		s := line[0]
		var l string
		for i := 1; i < len(line); i++ {
			if s == line[i] {
				count++
			} else {
				l += fmt.Sprintf("%d%c", count, s)
				s = line[i]
				count = 1
			}
		}
		l += fmt.Sprintf("%d%c", count, s)
		fmt.Print(t, " ")
		//fmt.Println(t, ":", line, "->", l)
		line = l
	}
	return len(line)
}

func main() {
	//const T int = 5
	//var line string = "1"
	//res := lookAndSay("11132", 50) + lookAndSay("22113", 50)
	res := lookAndSay("1", 5)
	fmt.Println("res: ", res)
	return

	const T int = 50
	var line string = "1113222113"

	for t := 0; t < T; t++ {
		count := 1
		s := line[0]
		var l string
		for i := 1; i < len(line); i++ {
			if s == line[i] {
				count++
			} else {
				l += fmt.Sprintf("%d%c", count, s)
				s = line[i]
				count = 1
			}
		}
		l += fmt.Sprintf("%d%c", count, s)
		fmt.Print(t, " ")
		//fmt.Println(t, ":", line, "->", l)
		line = l
	}
	fmt.Println("\nResult: ", len(line))
}
