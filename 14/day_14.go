package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Dear struct {
	Name   string
	Speed  int
	Time   int
	Rest   int
	Points int
}

func dist(dear *Dear, time int) int {
	run := dear.Time + dear.Rest
	hops := int(time / run)
	last := time % run

	sec := hops*dear.Time + min(dear.Time, last)
	return sec * dear.Speed
}

func parse(line string) (name string, speed int, time int, rest int) {
	sub := strings.Split(line, " ")

	name = sub[0]
	speed, _ = strconv.Atoi(sub[3])
	time, _ = strconv.Atoi(sub[6])
	rest, _ = strconv.Atoi(sub[13])
	return
}

func main() {
	//f, err := os.OpenFile("easy_14", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_14.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)

	const RaceTime int = 2503
	//const RaceTime int = 1000
	var dears []Dear
	var win int
	var winnerPoints int
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		name, speed, time, rest := parse(line)
		dears = append(dears, Dear{name, speed, time, rest, 0})

	}

	var dists = make([]int, len(dears))
	for t := 1; t <= RaceTime; t++ {
		for i := 0; i < len(dears); i++ {
			d := dist(&dears[i], t)
			dists[i] = d
			if d > win {
				win = d
			}
		}

		for i := 0; i < len(dears); i++ {
			dear := &dears[i]
			if dists[i] == win {
				p := &dear.Points
				*p++
				//fmt.Println("t =", t, " give point to:", dear.Name, "points:", *p)
				if *p > winnerPoints {
					winnerPoints = *p
				}
			}

			//fmt.Println(dear.Name, "(speed:", dear.Speed, "km/s", dear.Time, "s rest:", dear.Rest, "s) at: ", dists[i])
		}
	}
	fmt.Println("Winner: ", win)
	fmt.Println("Winner points: ", winnerPoints)
}
