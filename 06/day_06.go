package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1000

type LampMap [N][N]int

var lamps LampMap

type Rect struct {
	x0, y0 int
	x1, y1 int
}

const (
	TOGGLE int = 0
	ON         = 1
	OFF        = 2
)

func Show() {
	for i := 0; i < N; i++ {
		fmt.Println(lamps[i])
	}
}

func ParsePoint(str string) (int, int) {
	coords := strings.Split(str, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return x, y
}

func Parse(line string) (int, Rect) {
	args := strings.Split(line, " ")
	act := 0
	var x0, y0, x1, y1 int
	if args[0] == "toggle" {
		act = TOGGLE
		x0, y0 = ParsePoint(args[1])
		x1, y1 = ParsePoint(args[3])
	} else {
		if args[1] == "on" {
			act = ON
		} else {
			act = OFF
		}
		x0, y0 = ParsePoint(args[2])
		x1, y1 = ParsePoint(args[4])
	}
	return act, Rect{x0, y0, x1, y1}
}

func Apply(act int, rectPtr *Rect) {
	rect := *rectPtr
	for i := rect.x0; i <= rect.x1; i++ {
		for j := rect.y0; j <= rect.y1; j++ {
			var lamp *int = &lamps[i][j]
			switch act {
			case TOGGLE:
				//*lamp = !*lamp
				*lamp += 2
				break
			case ON:
				//*lamp = true
				*lamp += 1
				break
			case OFF:
				if *lamp > 0 {
					*lamp -= 1
				}
				//*lamp = false
				break
			}
		}
	}
}

func CountOn() int {
	var res int = 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res += lamps[i][j]
			//if lamps[i][j] {
			//res++
			//}
		}
	}
	return res
}

func main() {
	f, err := os.OpenFile("input_06.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_06", os.O_RDONLY, 0666)
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
		//fmt.Print(line, " ")
		action, Rect := Parse(line)
		//fmt.Println("Act: ", action, " Rect: ", Rect)
		Apply(action, &Rect)
		//fmt.Println("Result: ", CountOn())
		//Show()
	}
	fmt.Println("Result: ", CountOn())
}
