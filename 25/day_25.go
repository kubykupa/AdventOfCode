package main

// To continue, please consult the code grid in the manual.
// Enter the code at row 2981, column 3075.

import (
	"fmt"
	//"strings"
)

func next(val uint) uint {
	const mul uint = 252533
	const mod uint = 33554393
	return (val * mul) % mod
}

func stop(x, y int) bool {
	x++
	y++
	//return x == 6 && y == 6
	//return x == 4 && y == 3
	return x == 3075 && y == 2981
}

func main() {
	var v uint = 20151125
	for diag := 0; ; diag++ {
		fmt.Println("diag: ", diag)
		for x := 0; x <= diag; x++ {
			y := diag - x
			//fmt.Printf("[%d, %d]%10d ", x, y, v)
			if stop(x, y) {
				fmt.Println("\nResult: ", v)
				return
			}
			v = next(v)
		}
		fmt.Print("\n\n")
	}
}
