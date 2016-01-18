package main

//cat input_12.json | grep -oE "(-)?\d+" | awk 'BEGIN { sum=0 } { sum+=$1 } END {print "Sum: " sum}'

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	//"reflect"
)

func travese(any interface{}, level int) (int, bool) {
	//fmt.Println("\nStep in: ", any)
	var sum int
	var flag bool = true
	switch v := any.(type) {
	case map[string]interface{}:
		for _, v := range any.(map[string]interface{}) {
			s, f := travese(v, level+1)
			if f {
				sum += s
			} else {
				sum = 0
				break
			}
		}
		//fmt.Println("\nLVL:", level, "Step out:", any, "Sum:", sum, "flag:", flag)
		break
	case []interface{}:
		for _, v := range any.([]interface{}) {
			s, _ := travese(v, level+1)
			sum += s
		}
		break
		//fmt.Println("\nLVL:", level, "Step out:", any, "Sum:", sum, "flag:", flag)
	case string:
		if v == "red" {
			sum = 0
			flag = false
		}
		break
	case float64:
		sum = int(v)
	}

	//fmt.Println("\nLVL:", level, "Step out:", any, "Sum:", sum, "flag:", flag)
	return sum, flag
}

func main() {
	f, err := os.OpenFile("input_12.json", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_12", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("bad input file")
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, _ := r.ReadString('\n')
		var m interface{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			fmt.Println("ERORR:", err)
			return
		}

		//fmt.Print("Line:", line)
		s, _ := travese(m, 0)
		fmt.Println(" Sum: ", s)
	}
}
