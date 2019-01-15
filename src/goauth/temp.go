package main

import (
	"fmt"
	"encoding/json"
)

type meminfo struct {
	string
}

type memfalse struct {
	Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

//from https://mholt.github.io/json-to-go/
type memtrue struct {
	Result struct {
		Ff struct {
			Size             int           `json:"size"`
			Fee              float64       `json:"fee"`
			Time             int           `json:"time"`
			Height           int           `json:"height"`
			Startingpriority float64       `json:"startingpriority"`
			Currentpriority  float64       `json:"currentpriority"`
			Depends          []interface{} `json:"depends"`
		} `json:"ff"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

func (s meminfo) greeting() {
	fmt.Println("Hello",s.string)
}

func main() {
	a := meminfo{"World"}
	a.greeting()

	mfJson := `{"result":["8f549d7f360ebb90c6f6bd319490c71b1bccc28d70d53a82ed35d065ddce562f","49cfc2ca0b11f62ae39fedb484b76853575228993f8b7d283f5b6023e2c0a240","603d5d94b9f736f1ea88d551280389be846f2031fdd47c89b2aac07d5c0b8794","e3b88f232e9dc60d0065e4466b7754c33530f5bb0b1268d67dcc46c21c671ba9"],"error":null,"id":"curltest"}`
	
	var mf memfalse
	json.Unmarshal([]byte(mfJson), &mf)
	fmt.Println(mf.Result)
	fmt.Println(mf.Result[0])

	mtJson := `{"result":["8f549d7f360ebb90c6f6bd319490c71b1bccc28d70d53a82ed35d065ddce562f","49cfc2ca0b11f62ae39fedb484b76853575228993f8b7d283f5b6023e2c0a240","603d5d94b9f736f1ea88d551280389be846f2031fdd47c89b2aac07d5c0b8794","e3b88f232e9dc60d0065e4466b7754c33530f5bb0b1268d67dcc46c21c671ba9"],"error":null,"id":"curltest"}`
	var mt memtrue
	json.Unmarshal([]byte(mtJson), &mt)
	fmt.Println(mt.Result)
}