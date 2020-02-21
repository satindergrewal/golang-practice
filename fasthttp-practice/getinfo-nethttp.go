package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type APIQuery struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

func main() {
	rpcuser, rpcpass, rpcport := "user2a66b68210b05bee", "84fd254fbc577286c0b139472c1cf72b", "7771"

	client := &http.Client{}
	url := `http://127.0.0.1:` + rpcport

	q := APIQuery{
		Method: "getinfo",
		Params: "[]",
	}

	var query_str string
	query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "` + q.Method + `", "params": ` + q.Params + ` }`
	fmt.Println("query_str\n", query_str)

	query_byte := []byte(query_str)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
	req.Header.Set("Content-Type", "application/json")

	//req, err := http.NewRequest("POST", , nil)
	req.SetBasicAuth(rpcuser, rpcpass)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	var query_result map[string]interface{}
	if err := json.Unmarshal(bodyText, &query_result); err != nil {
		panic(err)
	}

	fmt.Println(string(bodyText))

}
