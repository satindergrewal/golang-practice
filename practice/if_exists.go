package main

import (
	"fmt"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	x := make([]interface{}, 1)
	x[0] = 42

	fmt.Println(x)

}
