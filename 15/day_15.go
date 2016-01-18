package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Resource struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func parse(line string) Resource {
	line = strings.Replace(line, ",", "", -1)
	sub := strings.Split(line, " ")

	name := sub[0]
	capacity, _ := strconv.Atoi(sub[2])
	durability, _ := strconv.Atoi(sub[4])
	flavor, _ := strconv.Atoi(sub[6])
	texture, _ := strconv.Atoi(sub[8])
	calories, _ := strconv.Atoi(sub[10])
	return Resource{name, capacity, durability, flavor, texture, calories}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//f, err := os.OpenFile("easy_15", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_15.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)

	var res []Resource
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		res = append(res, parse(line))
	}
	fmt.Println(res)

	const N int = 100
	weight := make([]int, len(res))

	best := 0
	cal_500 := 0
	for weight[len(weight)-1] != N {
		for i := 0; i < len(weight); i++ {
			weight[i]++
			if weight[i] > N {
				weight[i] = 0
				continue
			}
			break
		}

		sum := 0
		for i := 0; i < len(weight); i++ {
			sum += weight[i]
		}
		if sum == N {

			var capa, dura, flav, tex, cal int
			for i := 0; i < len(res); i++ {
				w := weight[i]
				capa += w * res[i].Capacity
				dura += w * res[i].Durability
				flav += w * res[i].Flavor
				tex += w * res[i].Texture
				cal += w * res[i].Calories
			}

			score := max(capa, 0) * max(dura, 0) * max(flav, 0) * max(tex, 0)
			fmt.Println(weight, "score:", score, "=", capa, dura, flav, tex)

			if score > best {
				best = score
			}
			if cal == 500 && score > cal_500 {
				cal_500 = score
			}
		}
	}
	fmt.Println("Result: ", best)
	fmt.Println("Calories: ", cal_500)
}
