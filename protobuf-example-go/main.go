package main

import (
	"fmt"
	simplepb "golang-practice/protobuf-example-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	// Always prefer to use helper funcs for such Id info instead of directly using sm.Id
	fmt.Println("The ID is: ", sm.GetId())
}
