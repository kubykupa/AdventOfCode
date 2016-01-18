package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("(\\\\\\\\|\\\\\"|\\\\x[0123456789abcdef]{2})")
var er = regexp.MustCompile("(\\\\|\")")

func encode(line string) string {
	var res string
	for i := 0; i < len(line); i++ {
		if line[i] == '\\' || line[i] == '"' {
			res = res + "\\"
		}
		res = res + string(line[i])
	}

	//res := er.ReplaceAllStringFunc(line, func(s string) string {
	//if s == "\\" {
	//return "\\\\"
	//}
	//return "\\\""
	//})
	return "\"" + res + "\""
}

func decode(line string) string {

	line = line[1 : len(line)-1]
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		switch s {
		case "\\\\":
			return "\\"
		case "\\\"":
			return "\""
		}
		return "@"
	})
	return line
}

func count(line string) int {
	var res int

	for i := 1; i < len(line)-1; i++ {
		res++
		if line[i] == '\\' {
			if line[i+1] == '"' || line[i+1] == '\\' {
				i++
			} else {
				//fmt.Println(line[i : i+4])
				i += 3
			}
		}
	}
	return res
}

func main() {
	//f, err := os.OpenFile("easy_08", os.O_RDONLY, 0666)
	f, err := os.OpenFile("input_08.txt", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var res1 int
	var res2 int
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n\r")
		encoded := encode(line)
		//fmt.Println(line, "->", encoded)

		str := decode(line)
		if len(str) != count(line) {
			fmt.Println("count: ", count(line), "decode:", len(str))
		}

		res1 += len(line) - count(line)
		res2 += len(encoded) - len(line)
	}
	fmt.Println("Res1:", res1)
	fmt.Println("Res2:", res2)
}
