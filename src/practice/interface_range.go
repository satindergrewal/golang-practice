package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	x := make([]interface{}, 10)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(cap(x))
	x[0] = 42
	x[1] = 0.0001
	x[2] = true
	x[3] = `hello world`

	fmt.Printf("%T\t %d\n", x[0], x[0])
	fmt.Printf("%T\t %f\n", x[1], x[1])
	fmt.Printf("%T\t %t\n", x[2], x[2])
	fmt.Printf("%T\t %s\n", x[3], x[3])
	fmt.Printf("%T\t %s\n", x[4], x[4])
	
	if x[4] == nil {
		fmt.Println("THIS IS NIL!!!")
	}

	for i, v := range x {
		fmt.Println(i)
		fmt.Println(v)
	}

	fmt.Println("#########################")

	bs, _ := json.Marshal(x)
	fmt.Println(x)
	fmt.Println(string(bs))
}
