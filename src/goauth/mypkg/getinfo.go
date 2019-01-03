package mypkg

import(
    //"fmt"
    //"encoding/json"
)

type MyString string

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

func GetinfoJson() string {
    getinfoJson := `{"result":{"version":1001550,"balance":10.16429765,"blocks":459,"name":"KMD"},"error":null,"id":"curltest"}`
    return getinfoJson
}

/*func (info MyString) MyMethod() GetInfo {
    var getinfo GetInfo
    json.Unmarshal([]byte(info), &getinfo)
    return getinfo
}*/


/*GetinfoJson := `{"result":{"version":1001550,"balance":10.16429765,"blocks":459,"name":"KMD"},"error":null,"id":"curltest"}`



func GetInfoAll() string {  
    return getinfoJson
}

func GetInfoVer() interface{
	var getinfo GetInfo   
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    //fmt.Println(getinfo.Result.Version)
    return getinfo.Result.Version
}*/