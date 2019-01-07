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

type appType string

type GetBestBlockhash struct {
    Result string      `json:"result"`
    Error  interface{} `json:"error"`
    ID     string      `json:"id"`
}

type APIQuery struct {
    Method string `json:"method"`
    Params string `json:"params"`
}

type APIError struct {
    Result interface{} `json:"result"`
    Error  struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    } `json:"error"`
    ID string `json:"id"`
}

type APIAnswer struct {
    Result interface{} `json:"result"`
    Error  interface{} `json:"error"`
    ID     string      `json:"id"`
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

    var query_ans APIAnswer
    var query_err APIError

    json.Unmarshal([]byte(bodyText), &query_ans)
    fmt.Println(query_ans)
    
    if query_ans.Result != nil {
        fmt.Println("query got answer!")
    } else {
        fmt.Println("!!! Return Error !!!")
        json.Unmarshal([]byte(bodyText), &query_err)
        fmt.Println(query_err.Error.Code)
        fmt.Println(query_err.Error.Message)
    }


    s := string(bodyText)
    return s
}


func GetBestBlockhashJsonValue(appName appType) string {
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
}

func (appName appType) GetBestBlockhash() GetBestBlockhash {
    query := APIQuery {
        Method:     "getbestblockhash",
        Params:   "[]",
    }

    //getbestblockhashJson := GetBestBlockhashJsonValue(appName)
    getbestblockhashJson := appName.APICall(query)
    fmt.Println(getbestblockhashJson)
    var getbestblockhash GetBestBlockhash
    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    return getbestblockhash
}

func main() {
    var appName appType
    appName = `komodo`
    //fmt.Printf("%T\n", appName)
    var bh GetBestBlockhash
    bh = appName.GetBestBlockhash()
    fmt.Println(bh)
    fmt.Println(bh.Result)
    fmt.Println(bh.Error)
    fmt.Println(bh.ID)

    /*query := APIQuery {
        Method:     "getbestblockhash",
        Params:   "[]",
    }

    bh := appName.APICall(query)
    fmt.Println(bh)*/
}
