package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Mover struct {
	Pos     Point
	Visited map[Point]int
}

func savePos(mover *Mover) {
	mover.Visited[mover.Pos] += 1
}

func move(mover *Mover, r uint8) {
	switch r {
	case 94:
		mover.Pos.y += 1
		break
	case 118:
		mover.Pos.y -= 1
		break
	case 60:
		mover.Pos.x -= 1
		break
	case 62:
		mover.Pos.x += 1
		break
	}
}

func main() {
	f, err := os.OpenFile("input_03.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_03", os.O_RDONLY, 0666)
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
		fmt.Println(line)
		santa := Mover{Visited: make(map[Point]int)}
		robo := Mover{Visited: make(map[Point]int)}
		savePos(&santa)
		savePos(&robo)
		for i := 0; i < len(line); i++ {
			r := line[i]

			if i%2 == 0 {
				move(&santa, r)
				savePos(&santa)
			} else {
				move(&robo, r)
				savePos(&robo)
			}

			//fmt.Println(string(rune), " = ", rune, " x: ", x, " y: ", y)
			//fmt.Println(string(r), " santa: ", santa, " robo: ", robo)
		}
		fmt.Println("Result santa: ", len(santa.Visited))
		fmt.Println("Result robo: ", len(robo.Visited))

		for k, v := range robo.Visited {
			santa.Visited[k] += v
		}
		fmt.Println("merge visited: ", len(santa.Visited))
	}
}
