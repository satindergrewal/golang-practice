package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	// "net/http"
	"github.com/valyala/fasthttp"
)

type APIQuery struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	rpcuser, rpcpass, rpcport := "user2213678121", "passcd3dbbf76467e1b6c04adc51e2289af83c0c1921f997a65509785649aecd44c745", "27486"

	client := &fasthttp.Client{}
	url := `http://127.0.0.1:` + rpcport

	var method string
	var parameters interface{}

	method = `getinfo`
	parameters = []interface{}{}

	// paramsJSON, _ := json.Marshal(parameters)
	// fmt.Println(string(paramsJSON))

	// q := APIQuery{
	// 	Method: method,
	// 	Params: string(paramsJSON),
	// }

	// q := APIQuery{
	// 	Method: "getinfo",
	// 	Params: "[]",
	// }

	queryByte, err := json.Marshal(map[string]interface{}{
		"id":      "0",
		"jsonrpc": "1.0",
		"method":  method,
		"params":  parameters,
	})
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}
	fmt.Println("queryByte - ", queryByte)

	// var queryStr string
	// queryStr = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "` + q.Method + `", "params": ` + q.Params + ` }`
	// fmt.Println("queryStr\n", queryStr)

	// queryByte := []byte(queryStr)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.Add("Authorization", "Basic "+basicAuth(rpcuser, rpcpass))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetBody(queryByte)

	resp := fasthttp.AcquireResponse()
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))

	var res map[string]interface{}
	json.Unmarshal(bodyBytes, &res)

	fmt.Println(`res["result"] -- `, res["result"])
}
