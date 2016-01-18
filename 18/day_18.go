package main

import (
	"bufio"
	"fmt"
	"os"
)

func getOn(x, y int, src [][]byte) (bool, int) {
	res := 0
	//var off = [][2]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	var off = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	for i := 0; i < len(off); i++ {
		xx := x + off[i][0]
		yy := y + off[i][1]
		if xx >= 0 && yy >= 0 && yy < len(src) && xx < len(src[yy]) {
			if src[yy][xx] == '#' {
				res++
			}
		}
	}
	return src[y][x] == '#', res
}

func show(t [][]byte) string {
	var val string = ""
	for y := 0; y < len(t); y++ {
		val += fmt.Sprintf(string(t[y]))
	}
	return fmt.Sprint(val)
}

func holdCornersOn(t [][]byte, X, Y int) {
	t[0][0], t[Y][0], t[0][X], t[Y][X] = '#', '#', '#', '#'
}

func main() {
	f, err := os.OpenFile("input_18.txt", os.O_RDONLY, 0666)
	//f, err := os.OpenFile("easy_18", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var src, dst [][]byte

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		src = append(src, []byte(line))
		dst = append(dst, []byte(line))
	}
	//fmt.Println("Initial state:")
	//fmt.Println(show(src))

	X := len(src) - 1
	Y := len(src[0]) - 2
	const N int = 100
	var pSrc, pDst *[][]byte = &src, &dst

	holdCornersOn(*pSrc, X, Y)
	for s := 0; s < N; s++ {
		//fmt.Println("Initial state src:")
		//fmt.Println(show(*pSrc))
		for y := 0; y <= Y; y++ {
			for x := 0; x <= X; x++ {
				state, env := getOn(x, y, *pSrc)
				//fmt.Print(x, y, "=", state, env)
				if state {
					if env == 2 || env == 3 {
						(*pDst)[y][x] = '#'
						//fmt.Println(" -> ON")
					} else {
						(*pDst)[y][x] = '.'
						//fmt.Println(" -> OFF")
					}
				} else {
					if env == 3 {
						(*pDst)[y][x] = '#'
						//fmt.Println(" -> ON")
					} else {
						(*pDst)[y][x] = (*pSrc)[y][x]
						//fmt.Println(" -> SKIP")
					}
				}
			}
		}

		pSrc, pDst = pDst, pSrc
		holdCornersOn(*pSrc, X, Y)
		//fmt.Println("After:", s)
		//fmt.Println(show(*pSrc))
	}

	on := 0
	for y := 0; y <= Y; y++ {
		for x := 0; x <= X; x++ {
			if (*pSrc)[y][x] == '#' {
				on++
			}
		}
	}
	fmt.Println("Lamps ON:", on)
}
