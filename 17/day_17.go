package main

import (
	"fmt"
)

func main() {
	box := []int{33, 14, 18, 20, 45, 35, 16, 35, 1, 13, 18, 13, 50, 44, 48, 6, 24, 41, 30, 42}
	const Sum int = 150

	//box := []int{20, 15, 10, 5, 5}
	//const Sum int = 25

	var maxMask int
	for i := 0; i < len(box); i++ {
		maxMask |= 1 << uint(i)
	}
	fmt.Printf("%020b\n", maxMask)
	fmt.Println(box)
	masks := make(map[int]bool)
	counts := make(map[int]int)
	for m := 1; m <= maxMask; m++ {
		sum := 0
		used := 0
		count := 0
		for i := 0; i < len(box); i++ {
			if m&(1<<uint(i)) > 0 {
				sum += box[i]
				used |= 1 << uint(i)
				count++
				if sum == Sum {
					if !masks[used] {
						counts[count]++
						masks[used] = true
						fmt.Printf("Variant: %020b total: %d\n", used, len(masks))
					}
					break
				}
				if sum > Sum {
					break
				}
			}
		}
	}
	fmt.Println("Ways:", counts)
}
