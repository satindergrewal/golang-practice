package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

type GetInfo struct {
    Result struct {
        Version             int     `json:"version"`
        Protocolversion     int     `json:"protocolversion"`
        KMDversion          string  `json:"KMDversion"`
        Notarized           int     `json:"notarized"`
        PrevMoMheight       int     `json:"prevMoMheight"`
        Notarizedhash       string  `json:"notarizedhash"`
        Notarizedtxid       string  `json:"notarizedtxid"`
        NotarizedtxidHeight string  `json:"notarizedtxid_height"`
        KMDnotarizedHeight  int     `json:"KMDnotarized_height"`
        NotarizedConfirms   int     `json:"notarized_confirms"`
        Walletversion       int     `json:"walletversion"`
        Balance             float64 `json:"balance"`
        Blocks              int     `json:"blocks"`
        Longestchain        int     `json:"longestchain"`
        Timeoffset          int     `json:"timeoffset"`
        Tiptime             int     `json:"tiptime"`
        Connections         int     `json:"connections"`
        Proxy               string  `json:"proxy"`
        Difficulty          float64 `json:"difficulty"`
        Testnet             bool    `json:"testnet"`
        Keypoololdest       int     `json:"keypoololdest"`
        Keypoolsize         int     `json:"keypoolsize"`
        Paytxfee            float64 `json:"paytxfee"`
        Relayfee            float64 `json:"relayfee"`
        Errors              string  `json:"errors"`
        CCid                int     `json:"CCid"`
        Name                string  `json:"name"`
        P2Pport             int     `json:"p2pport"`
        Rpcport             int     `json:"rpcport"`
        Magic               int     `json:"magic"`
        Premine             int     `json:"premine"`
        Reward              int64   `json:"reward"`
        Halving             int     `json:"halving"`
        Commission          int     `json:"commission"`
    } `json:"result"`
    Error interface{} `json:"error"`
    ID    string      `json:"id"`
}

func GetinfoJsonValue() string {
    appName := "komodo"
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport
    //fmt.Println("URL:>", url)

    query_byte := []byte(`{"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }`)
    //fmt.Printf("Query: %s\n\n", query_byte)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

    //req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(rpcuser, rpcpass)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("%T\n\n", bodyText)

    var query_result map[string]interface{}
    if err := json.Unmarshal(bodyText, &query_result); err != nil {
        panic(err)
    }
    //fmt.Println(query_result)
    fmt.Printf("\n\n")

    //fmt.Println(query_result["result"].(map[string]interface{})["connections"])

/*  parsed_result := query_result["result"]
    fmt.Println("result: ", parsed_result)
    fmt.Printf("\n\n")
*/

    s := string(bodyText)
    return s
}

/*func GetinfoJsonValue() string {
    getinfoJson := `{"result":{"version":1001550,"protocolversion":170003,"KMDversion":"0.2.0","notarized":0,"prevMoMheight":0,"notarizedhash":"0000000000000000000000000000000000000000000000000000000000000000","notarizedtxid":"0000000000000000000000000000000000000000000000000000000000000000","notarizedtxid_height":"mempool","KMDnotarized_height":0,"notarized_confirms":0,"walletversion":60000,"balance":10.16429765,"blocks":459,"longestchain":0,"timeoffset":0,"tiptime":1536624090,"connections":0,"proxy":"","difficulty":1.000026345948652,"testnet":false,"keypoololdest":1536262464,"keypoolsize":101,"relayfee":0.000001,"paytxfee":0,"errors":"","name":"SIDD","p2pport":9800,"rpcport":9801,"magic":-759875707,"premine":10},"error":null,"id":"curltest"}`
    return getinfoJson
}*/

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
