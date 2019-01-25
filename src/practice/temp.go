package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	type Txes []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	}

	x := Txes{{"d7ba45296c66e16eb61f27a4eef8848c7f5579fe801f277c1b0e074a4f47d6fd", 0}}

	fmt.Println(x)

	params_json, _ := json.Marshal(x)
	fmt.Println(string(params_json))
}
