package main

import (
	//"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
)

type Config struct {
	Passphrase    string                 `json:"passphrase"`
	AC_Range      []int                  `json:"ac_range"`
	AC_Seed       []string               `json:"ac_seed"`
	Insight_Setup map[string]interface{} `json: "insight_setup"`
	Addresslist   string                 `json:"addresslist"`
}

/*type AddressListData struct {

}*/

type Bird struct {
	Species string
	//Description string
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	config := "./config.json"

	config_content, err := ioutil.ReadFile(config)
	checkError(err)

	result := string(config_content)

	//fmt.Println("Read the file: ", result)

	data := []byte(result)
	//fmt.Println(json.Get(data, "addresslist").ToString())
	//fmt.Println(data)

	var conf Config
	err = json.Unmarshal(data, &conf)
	//fmt.Println(conf)
	//fmt.Println("----------------")
	//fmt.Println(conf.Addresslist)

	/*birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)*/

	//birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	birdJson := `[{"RKV4MzePkH8H9N6K3RdC1JDDWBLbCE4H5p": 0.0001},{"RSzuS5dT5WYLxZekMwcoGXQjtcD9FufPaq": 0.0001}]`
	obj := []byte(json.Get([]byte(birdJson), 0).ToString())
	fmt.Println(string(obj))
	//fmt.Println(json.Get(obj).ToString())

	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(obj, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	// output result to STDOUT
	fmt.Printf("%v\n", k[0])

	var sati = `Hello World`
	fmt.Println(sati)

	//birdJson := `[{"species":"pigeon"},{"species":"eagle"}]`
	//var birds []Bird
	//json.Unmarshal([]byte(birdJson), &birds)
	//fmt.Printf("Birds : %+v\n", birds)

	/*for k, v := range conf.Addresslist {
		fmt.Printf("%v: %v\n", k, v)
		if k == 2 {
			break
		}
	}*/

	//keys := map[string]int{conf.Addresslist}
	//fmt.Println(keys)

	/*i := 0
	for k := range conf.Addresslist {
		keys[i] = k
		i++
	}
	//sort.Strings(keys)
	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(conf.Addresslist[keys[i]])
	}*/

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func trimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
