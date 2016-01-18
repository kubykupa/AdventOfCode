package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	f, err := os.OpenFile("input_02.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_02", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var sum int = 0
	var rib int = 0
	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}

		lwh := strings.Split(strings.TrimRight(line, "\r\n"), "x")
		if len(lwh) != 3 {
			continue
		}

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])
		//fmt.Print("l: ", l, " w: ", w, " h:", h, "\n")

		s := []int{l, w, h}
		sort.Ints(s)

		lw := l * w
		wh := w * h
		hl := h * l
		min := Min(lw, Min(wh, hl))
		res := 2*lw + 2*wh + 2*hl + min
		sum = sum + res

		rib = rib + 2*s[0] + 2*s[1] + l*w*h
	}
	fmt.Println("Sum: ", sum)
	fmt.Println("Rib: ", rib)
}
