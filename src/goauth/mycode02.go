package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "errors"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

type appType string

type GetBestBlockhash struct {
    Result interface{} `json:"result"`
    Error  Error       `json:"error"`
    ID     string      `json:"id"`
}

type GetInfo struct {
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
}

type APIResult struct {
    Result interface{}  `json:"result"`
    Error interface{}   `json:"error"`
    ID string           `json:"id"`
}

type APIQuery struct {
    Method string `json:"method"`
    Params string `json:"params"`
}

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (appName appType) APICall(q APIQuery) string {
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(string(appName))

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport

    var query_str string
    query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "`+q.Method+`", "params": `+q.Params+` }`

    query_byte := []byte(query_str)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

    //req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(rpcuser, rpcpass)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)

    var query_result map[string]interface{}
    if err := json.Unmarshal(bodyText, &query_result); err != nil {
        panic(err)
    }

    s := string(bodyText)
    return s
}

func (appName appType) GetBestBlockhash() (GetBestBlockhash, error) {
    query := APIQuery {
        Method:     "getbestblockhash",
        Params:   "[v]",
    }

    var getbestblockhash GetBestBlockhash

    getbestblockhashJson := appName.APICall(query)
    fmt.Println(getbestblockhashJson)
    
    
    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    
    if getbestblockhash.Result == nil {
        fmt.Println("!!! Return Error !!!")
        answer_error, err := json.Marshal(getbestblockhash.Error)
        if err != nil {
            fmt.Println("error:", err)
        }
        return getbestblockhash, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    return getbestblockhash, nil
}

func (appName appType) GetInfo() (APIResult, error) {
    query := APIQuery {
        Method:     "getinfo",
        Params:   "[]",
    }

    var getinfo APIResult

    getinfoJson := appName.APICall(query)
    fmt.Println(getinfoJson)
    
    
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    
    if getinfo.Error != nil {
        fmt.Println("!!! Return Error !!!")
        answer_error, err := json.Marshal(getinfo.Error)
        if err != nil {
            fmt.Println("error:", err)
        }
        return getinfo, errors.New(string(answer_error))
    }

    json.Unmarshal([]byte(getinfoJson), &getinfo)
    return getinfo, nil
}

func main() {
    var appName appType
    appName = `komodo`
    fmt.Printf("\n")
    /*var bh GetBestBlockhash
    bh, err := appName.GetBestBlockhash()
    if err != nil {
        log.Println("err happened", err)
    }
    fmt.Println("bh value", bh)
    fmt.Println(bh.Result)
    fmt.Println(bh.Error.Code)
    fmt.Println(bh.Error.Message)
    fmt.Println(bh.ID)*/

    var res APIResult

    res, err := appName.GetInfo()
    if err != nil {
        log.Println("err happened", err)
    }
    fmt.Println(res.Result)

    fmt.Printf("\n\n\n")

    fmt.Printf("%T\n", res.Result)

    //var gi GetInfo
    //gi = res.Result
    //fmt.Println(gi)

}
