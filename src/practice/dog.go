package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Breed    string
	WeightKg int
}

func main() {
	d := Dog{
		Breed: "dalmation",
		//WeightKg:	45,
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}
