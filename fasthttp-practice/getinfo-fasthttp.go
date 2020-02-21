package main

import (
	"encoding/base64"
	"fmt"
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
	rpcuser, rpcpass, rpcport := "user2a66b68210b05bee", "84fd254fbc577286c0b139472c1cf72b", "7771"

	client := &fasthttp.Client{}
	url := `http://127.0.0.1:` + rpcport

	q := APIQuery{
		Method: "getinfo",
		Params: "[]",
	}

	var query_str string
	query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "` + q.Method + `", "params": ` + q.Params + ` }`
	fmt.Println("query_str\n", query_str)

	query_byte := []byte(query_str)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.Add("Authorization", "Basic "+basicAuth(rpcuser, rpcpass))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetBody(query_byte)

	resp := fasthttp.AcquireResponse()
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))

}
