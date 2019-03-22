//for range loop & index position
//https://play.golang.org/p/1fzbabuaWD

package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 10)
	y = append(y, 777)

	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	// map
	// struct
}

//	Output
//	0 - 7
//	1 - 9
//	2 - 42
//	0 - 777
