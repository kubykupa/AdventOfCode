package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NOT    uint8 = 0
	AND          = 1
	OR           = 2
	LSHIFT       = 3
	RSHIFT       = 4
)

func T(b uint8) string {
	if b == NOT {
		return fmt.Sprint(" ! ")
	}
	if b == AND {
		return fmt.Sprint(" AND ")
	}
	if b == OR {
		return fmt.Sprint(" OR ")
	}
	if b == LSHIFT {
		return fmt.Sprint(" L ")
	}
	if b == RSHIFT {
		return fmt.Sprint(" R ")
	}
	return ""
}

type Wire struct {
	Label string
	Value *uint16
}

type Gate struct {
	Type uint8
	A, B *Wire
	Res  *Wire
}

func (t *Wire) String() string {
	var val string = "?"
	if t.Value != nil {
		val = fmt.Sprint(*t.Value)
	}
	if len(t.Label) == 0 {
		return val
	}
	return fmt.Sprint("{", t.Label, ":", val, "}")
}

func (t Gate) String() string {
	if t.B == nil {
		return fmt.Sprint("(", T(t.Type), t.A, " -> ", t.Res, ")")
	}
	return fmt.Sprint("(", t.A, T(t.Type), t.B, " -> ", t.Res, ")")
}

var Wires = make(map[string]*Wire)
var Gates []Gate

func isWireOk(wire *Wire) bool {
	if wire.Value == nil {
		return false
	}
	return true
}

func tryGate(gate *Gate) bool {
	if isWireOk(gate.Res) {
		return true
	}
	if gate.A.Value == nil {
		return false
	}

	a := *(gate.A.Value)
	res := new(uint16)
	if gate.Type == NOT {
		*res = ^a
		gate.Res.Value = res
		return true
	}

	if gate.B.Value == nil {
		return false
	}

	b := *(gate.B.Value)
	switch gate.Type {
	case AND:
		*res = a & b
		break
	case OR:
		*res = a | b
		break
	case RSHIFT:
		*res = a >> b
		break
	case LSHIFT:
		*res = a << b
		break
	}

	gate.Res.Value = res
	return true
}

func Get(val string) *Wire {
	v, e := strconv.Atoi(val)
	var w Wire
	if e == nil {
		u16 := uint16(v)
		w = Wire{"", &u16}
	} else {
		if w, ok := Wires[val]; ok {
			return w
		}
		w = Wire{val, nil} // bad value
		Wires[w.Label] = &w
	}
	return &w
}

func Parse(line string) {
	//fmt.Println(line)
	args := strings.Split(line, " ")
	if len(args) == 3 {
		from := Get(args[0])
		to := Get(args[2])
		if len(from.Label) > 0 {
			Wires[to.Label] = from
		} else {
			to.Value = from.Value
		}
		return
	}

	if args[0] == "NOT" {
		Gates = append(Gates, Gate{NOT, Get(args[1]), nil, Get(args[3])})
		return
	}

	var t uint8
	switch args[1] {
	case "AND":
		t = AND
		break
	case "OR":
		t = OR
		break
	case "RSHIFT":
		t = RSHIFT
		break
	case "LSHIFT":
		t = LSHIFT
		break
	}
	Gates = append(Gates, Gate{t, Get(args[0]), Get(args[2]), Get(args[4])})
}

func Show() {
	fmt.Println("Wires: ")
	for _, v := range Wires {
		fmt.Print(v, " ")
	}
	fmt.Println("\nGates: ", Gates)
}

func main() {
	f, err := os.OpenFile("input_7_star.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("input_07.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_07", os.O_RDONLY, 0666)
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
		Parse(line)
	}

	for len(Gates) > 0 {
		for g := 0; g < len(Gates); {
			if tryGate(&Gates[g]) {
				fmt.Println("COMPLETE: ", Gates[g])
				Gates = append(Gates[:g], Gates[g+1:]...)
			} else {
				g++
			}

			//Show()
		}
	}
	Show()

	fmt.Println("RESULT: ", Wires["a"])
}
