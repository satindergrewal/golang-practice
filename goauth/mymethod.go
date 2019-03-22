package main

import (  
    "fmt"
    "encoding/json"
)

type GetInfo struct {
    Result struct {
        Version             int     `json:"version"`
        Balance             float64 `json:"balance"`
        Blocks              int     `json:"blocks"`
        Name                string  `json:"name"`
    } `json:"result"`
    Error interface{} `json:"error"`
    ID    string      `json:"id"`
}

func GetinfoJsonValue() string {
    getinfoJson := `{"result":{"version":1001550,"balance":10.16429765,"blocks":459,"name":"KMD"},"error":null,"id":"curltest"}`
    return getinfoJson
}

func (i GetInfo) DisplayGetinfo() GetInfo {
    //fmt.Println(i.Result.Version)
    return i
}

func ResultGetInfo() GetInfo {
    getinfoJson := GetinfoJsonValue()
    var getinfo GetInfo
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    return getinfo
}

func main() {  
    var rval GetInfo
    rval = ResultGetInfo()
    fmt.Println(rval)
    fmt.Println(rval.Result)
    fmt.Println(rval.Result.Version)
    fmt.Println(rval.Result.Balance)
    fmt.Println(rval.Result.Blocks)
    fmt.Println(rval.Result.Name)
    fmt.Println(rval.Error)
    fmt.Println(rval.ID)
}
