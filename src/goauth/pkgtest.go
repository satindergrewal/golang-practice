package main

import (
	//"strings"
	"fmt"
	"goauth/mypkg"
	"encoding/json"
)

func main() {
	getinfoJson := `{"result":{"version":1001550,"balance":10.16429765,"blocks":459,"name":"KMD"},"error":null,"id":"curltest"}`
	fmt.Println(getinfoJson)

    var getinfo mypkg.GetInfo
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    fmt.Println(getinfo.Result.Version)
}