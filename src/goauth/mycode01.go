package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "errors"
    //"os"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

type appType string

type GetBestBlockhash struct {
    Result interface{}      `json:"result"`
    Error  Error `json:"error"`
    ID     string      `json:"id"`
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

    //fmt.Println(rpcuser)
    //fmt.Println(rpcpass)
    //fmt.Println(rpcport)
    //fmt.Printf("%s\n", url)
    //fmt.Println(q.Method)
    //fmt.Println(q.Params)

    var query_str string
    query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "`+q.Method+`", "params": `+q.Params+` }`
    //fmt.Println(query_str)

    query_byte := []byte(query_str)
    //fmt.Println(query_byte)

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


/*func GetBestBlockhashJsonValue(appName appType) string {
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(string(appName))

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport

    query_byte := []byte(`{"jsonrpc": "1.0", "id":"kmdgo", "method": "getbestblockhash", "params": [] }`)

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
}*/

func (appName appType) GetBestBlockhash() (GetBestBlockhash, error) {
    query := APIQuery {
        Method:     "getbestblockhash",
        Params:   "[]",
    }

    var getbestblockhash GetBestBlockhash

    //getbestblockhashJson := GetBestBlockhashJsonValue(appName)
    getbestblockhashJson := appName.APICall(query)
    fmt.Println(getbestblockhashJson)
    
    
    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    
    if getbestblockhash.Result != nil {
        fmt.Println("query got answer!")
    } else {
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

func main() {
    var appName appType
    appName = `komodo`
    fmt.Printf("\n")
    var bh GetBestBlockhash
    bh, err := appName.GetBestBlockhash()
    if err != nil {
        log.Println("err happened", err)
    }
    fmt.Println("bh value", bh)
    fmt.Println(bh.Result)
    fmt.Println(bh.Error.Code)
    fmt.Println(bh.Error.Message)
    fmt.Println(bh.ID)
}
