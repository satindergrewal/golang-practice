package main

import (
	"fmt"
)

func Sum[V int | int64 | float64](m ...V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4} ...))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(int64(4), int64(2), int64(9)))
	fmt.Println(Sum(1.3, 4.3, 9.2))
}