package main

import (
	"fmt"
)

//const abc string = "abcdefghijklmnopqrstuvwxyz"

func valid(pass string) bool {
	L := len(pass)
	var lastPair int = -1
	var pairCount int
	var hasSeq bool = false
	for i := 0; i < L; i++ {
		if pass[i] == 'i' || pass[i] == 'o' || pass[i] == 'l' {
			return false
		}
		if i > 1 && pass[i-1]-pass[i-2] == 1 && pass[i]-pass[i-1] == 1 {
			hasSeq = true
			//fmt.Println("pass:", pass, "seq: ", pass[i-2:i+1])
		}
		if i > lastPair+1 && pass[i-1] == pass[i] {
			//fmt.Println("pass:", pass, "pair: ", i-1, i, pass[i-1:i+1])
			pairCount++
			lastPair = i
		}
	}
	return pairCount > 1 && hasSeq
}

func propagate(r *rune, base uint8) bool {
	if *r == 'z' {
		*r = 'a'
		return true
	} else {
		*r++
	}
	//if uint8(*r) == base {
	//return true
	//}
	if *r == 'i' || *r == 'o' || *r == 'l' {
		*r++
	}
	return false
}

func next(base string) string {
	pass := []rune(base)
	var i int = len(pass) - 1

	for {
		for j := i; j >= 0; j-- {
			//fmt.Println(string(pass))
			if !propagate(&pass[j], base[j]) {
				break
			}
		}
		res := string(pass)
		if valid(res) {
			return res
		}
	}
	return base
}

func main() {
	if valid("xyzdffaa") == false || valid("aaa") == true || valid("aabcaa") == false || valid("hijklmmn") == true || valid("abbceffg") == true || valid("abbcegjk") == true ||
		valid("abcdffaa") == false || valid("ghjaabcc") == false {
		fmt.Println("Bad valid func")
		return
	}

	if next("abcdefgh") != "abcdffaa" && next("ghijklmn") != "ghjaabcc" {
		fmt.Println("Bad gen func")
		return
	}

	var pass string = "hxbxwxba"
	res := next(pass)
	fmt.Println("Res: ", res)
	fmt.Println("Res 2:", next(res))
}
