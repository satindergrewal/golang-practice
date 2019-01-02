//authclient.go
package main

import(
    "fmt"
    "encoding/json"
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


func main(){

    fmt.Println("requesting...\n")
    getinfoJson := `{"result":{"version":1001550,"protocolversion":170003,"KMDversion":"0.2.0","notarized":0,"prevMoMheight":0,"notarizedhash":"0000000000000000000000000000000000000000000000000000000000000000","notarizedtxid":"0000000000000000000000000000000000000000000000000000000000000000","notarizedtxid_height":"mempool","KMDnotarized_height":0,"notarized_confirms":0,"walletversion":60000,"balance":10.16429765,"blocks":459,"longestchain":0,"timeoffset":0,"tiptime":1536624090,"connections":0,"proxy":"","difficulty":1.000026345948652,"testnet":false,"keypoololdest":1536262464,"keypoolsize":101,"relayfee":0.000001,"paytxfee":0,"errors":"","name":"SIDD","p2pport":9800,"rpcport":9801,"magic":-759875707,"premine":10},"error":null,"id":"curltest"}`
    //fmt.Println(getinfoJson)

    //fmt.Printf("%T\n", getinfoJson)
    
    var getinfo GetInfo
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    fmt.Println(getinfo.Result.Version)
    //fmt.Printf("Version: %d", getinfo[0].Version)
}