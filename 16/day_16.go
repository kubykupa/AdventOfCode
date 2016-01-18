package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Info struct {
	Id    int
	Score int
	Prop  map[string]int
}

var Base Info = Info{-1, 666, map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1}}

func score(i Info) int {
	res := 0
	for k, v := range Base.Prop {
		if p, ok := i.Prop[k]; ok {
			if (k == "cats" || k == "trees") && (p > v) {
				res += 100
				continue
			}
			if (k == "pomeranians" || k == "goldfish") && (p < v) {
				res += 100
				continue
			}

			if v == p {
				res += 100
			}
		}
	}
	return res
}

func parse(line string) Info {
	line = strings.Replace(line, ":", "", -1)
	line = strings.Replace(line, ",", "", -1)
	sub := strings.Split(line, " ")
	v, _ := strconv.Atoi(sub[1])
	res := Info{v, -666, make(map[string]int)}

	v, _ = strconv.Atoi(sub[3])
	res.Prop[sub[2]] = v

	v, _ = strconv.Atoi(sub[5])
	res.Prop[sub[4]] = v

	v, _ = strconv.Atoi(sub[7])
	res.Prop[sub[6]] = v

	res.Score = score(res)
	return res
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByScore []Info

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func main() {
	//f, err := os.OpenFile("easy_15", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_16.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)

	fmt.Println("base:", Base)
	var sues []Info
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		info := parse(line)
		//fmt.Println(line, info)

		sues = append(sues, info)
	}

	sort.Sort(ByScore(sues))
	fmt.Println("Sorted:")
	for i := 0; i < 2; i++ {
		fmt.Println(sues[i])
	}
}
