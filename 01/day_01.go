package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("input_01.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var sum int
	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.TrimRight(line, "\r\n")
		//fmt.Println("Line:\n\r", line)
		var flag bool
		for i := 0; i < len(line); i++ {
			if line[i] == '(' {
				sum++
			}
			if line[i] == ')' {
				sum--
			}
			//fmt.Printf("%d %c - sum: %d\n", i, line[i], sum)
			if sum < 0 && !flag {
				fmt.Println("HIT: ", i+1)
				flag = true
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
