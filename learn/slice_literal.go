//composite literal; slice literal
//https://play.golang.org/p/nGekFavi0S

package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	fmt.Println(x)

	y := make([]int, 0, 100)
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y)

	// map
	// struct
}

//	Output
//	[7 9 42]
//	[8 12 43]
