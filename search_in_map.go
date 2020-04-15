package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	data, err := ioutil.ReadFile(filepath.Join(dir, "/assets/subatomic.json"))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", data)

	// dec := json.NewDecoder(strings.NewReader(string(data[:])))
	// fmt.Println(dec)

	var parsed map[string][]map[string]string
	err = json.Unmarshal([]byte(data), &parsed)
	// fmt.Println(parsed["authorized"])
	// var auth map[string][string]interface{}

	// testpub := "03732f8ef851ff234c74d0df575c2c5b159e2bab3faca4ec52b3f217d5cda5361d"
	testpub := "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"

	for _, v := range parsed["authorized"] {
		// fmt.Println(v)
		for key, val := range v {
			if val == testpub {
				fmt.Println(key)
			}
		}
		// fmt.Println(findByKey(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"))
		// fmt.Println(findByValue(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"))
		// fmt.Println(findInMap(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"))
		// if testpub == findByKey(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61") {
		// 	fmt.Println("matched")
		// 	fmt.Println(findInMap(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"))
		// } else {
		// 	fmt.Println("Unmatched")
		// 	fmt.Println(findInMap(v, "02b27de3ee5335518b06f69f4fbabb029cfc737613b100996841d5532b324a5a61"))
		// }
	}
}

// Reference code from:
// https://forum.golangbridge.org/t/im-searching-for-something-like-a-both-way-map/2557/3
// https://play.golang.org/p/4c3UgqNHTZ

// findInMap uses findByValue and findByKey to check whether a
// string exists either as a key or a value inside of a map.
//
// findInMap returns either the key or value, if the string was found
// in the map, or an empty string if neither were found.
func findInMap(m map[string]string, str string) string {
	// Check if str exists as a value in map m.
	if val := findByKey(m, str); val != "" {
		return val
	}

	// Check if str exists as a key in map m.
	if val := findByValue(m, str); val != "" {
		return val
	}

	// str doesn't exist in map m.
	return ""
}

// findByValue searches for a string in a map based on map values.
//
// findByValue returns the key if the value was found, or an empty
// string if it wasn't.
func findByValue(m map[string]string, value string) string {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return ""
}

// findByKey searches for a string in a map based on map keys.
//
// findByKey returns the value if the key was found, or an empty
// string if it wasn't.
func findByKey(m map[string]string, key string) string {
	if val, found := m[key]; found {
		return val
	}
	return ""
}
