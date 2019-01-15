package main

import (
	"encoding/json"
	"fmt"
)

type meminfo struct {
	string
}

type memfalse struct {
	//Result interface{}    `json:"result"`
  Result map[string]memtrue    `json:"result"`
  //Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type memtrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

func (s meminfo) greeting() {
	fmt.Println("Hello", s.string)
}

func main() {
	a := meminfo{"World"}
	a.greeting()



  mfJson := `{"result":[
  "7921fbd3ad4d1593aa580833946cb123260f81480069a37f649d2c8cc10d770b",
  "7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448"
],"error":null,"id":"curltest"}`

	mtJson := `{"result":{
  "7921fbd3ad4d1593aa580833946cb123260f81480069a37f649d2c8cc10d770b": {
    "size": 4611,
    "fee": 0.00110000,
    "time": 1547431224,
    "height": 1183180,
    "startingpriority": 8135355.08111136,
    "currentpriority": 8135355.08111136,
    "depends": [
    ]
  },
  "7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448": {
    "size": 244,
    "fee": 0.00005000,
    "time": 1547431225,
    "height": 1183180,
    "startingpriority": 0,
    "currentpriority": 0,
    "depends": [
      "f9b4e1667bf01e9fc1507f35b77a90ef6b70508f6c5ce7a814238a2e0c893dc0"
    ]
  }
},"error":null,"id":"curltest"}`
  
  fmt.Println()
  fmt.Println(mfJson)
  fmt.Println()
  fmt.Println(mtJson)

	var result memfalse
  json.Unmarshal([]byte(mtJson), &result)


  fmt.Println("-----")
  fmt.Println(result.Result)
  fmt.Printf("%T\n", result.Result)

  fmt.Println(len(result.Result))
  fmt.Println(result.Result["7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448"])
  fmt.Println(result.Result["7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448"].Size)


  fmt.Println("-----")
  for i, v := range result.Result {
    fmt.Println(i)
    fmt.Println(v)
    fmt.Println(v.Size)
  }
}
