package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Seq []string
type State struct {
	Tokens     string
	Begin, End int
}

var Rules = map[string][]Seq{}
var Molecule string

var Cache = map[State]int{}

const (
	BAD = -8
)

func min(a, b int) int {
	if a == BAD {
		return b
	}
	if b == BAD {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func nonCached(tokens Seq, begin, end int) int {
	if strings.Join(tokens, "") == Molecule[begin:end] {
		//fmt.Println("Found:", Molecule[:begin], tokens, Molecule[end:], "in:", begin, ":", end)
		return 0
	}
	res := BAD
	if len(tokens) == 1 {
		fmt.Println("Tokens:", tokens)
		for _, v := range Rules[tokens[0]] {
			r := solve(v, begin, end)
			fmt.Println(tokens[0], "->", v, "in:", begin, end, "solve:", r)
			if r != BAD {
				res = min(res, r+1)
			}
		}
	} else {
		for l := 1; l < len(tokens); l++ {
			for r := begin + 1; r < end; r++ {
				ans1 := solve(tokens[:l], begin, r)
				if ans1 != BAD {
					ans2 := solve(tokens[l:], r, end)
					if ans2 != BAD {
						res = min(res, ans1+ans2)
					}
				}
			}
		}
	}
	return res
}

func solve(tokens Seq, begin, end int) int {
	state := State{strings.Join(tokens, "."), begin, end}
	if v, ok := Cache[state]; ok {
		fmt.Println("Cahed:", state)
		return v
	}
	fmt.Println("NonCached:", state)
	res := nonCached(tokens, begin, end)
	Cache[state] = res
	return res
}

func main() {
	f, err := os.OpenFile("input_19.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_19", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	re := regexp.MustCompile("([A-Z][a-z]*)")
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")

		// parse
		if len(line) == 0 {
			continue
		}
		sub := strings.Split(line, " ")
		if len(sub) == 3 {
			found := re.FindAllString(sub[2], -1)
			var s Seq
			for _, tok := range found {
				s = append(s, tok)
			}
			Rules[sub[0]] = append(Rules[sub[0]], s)
		} else {
			Molecule = line
		}
	}

	fmt.Println(Rules)
	//Molecule = "HOHOH"
	fmt.Println("Molecule:", Molecule)
	fmt.Println(solve(Seq{"e"}, 0, len(Molecule)))

	fmt.Println(Cache)
}
