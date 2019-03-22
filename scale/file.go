package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	config := "./config.json"

	config_content, err := ioutil.ReadFile(config)
	checkError(err)

	result := string(config_content)

	fmt.Println("Read the file: ", result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
