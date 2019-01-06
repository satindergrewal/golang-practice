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

//{"result":null,"error":{"code":-28,"message":"Loading block index..."},"id":"kmdgo"}
//{"result":"0895451014bccc54040257849239bd1677142a3dfe5cd41a5b8f3c9b03a1ad21","error":null,"id":"kmdgo"}

type appType string

type GetBestBlockhash struct {
    Result string      `json:"result"`
    Error  interface{} `json:"error"`
    ID     string      `json:"id"`
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

func (appName appType) DisplayGetBestBlockhash() GetBestBlockhash {
    fmt.Println(appName)
    fmt.Printf("%T\n", appName)
    getbestblockhashJson := GetBestBlockhashJsonValue(appName)
    var getbestblockhash GetBestBlockhash
    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    return getbestblockhash
}

func ResultGetBestBlockhash(appName appType) GetBestBlockhash {
    getbestblockhashJson := GetBestBlockhashJsonValue(appName)
    var getbestblockhash GetBestBlockhash
    json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
    return getbestblockhash
}

func main() {
    //appName := `komodo`
    var appName appType
    appName = `komodo`
    fmt.Printf("%T\n", appName)
    var bh GetBestBlockhash
    //bh.DisplayGetBestBlockhash()
    //bh = ResultGetBestBlockhash(appName)
    bh = appName.DisplayGetBestBlockhash()
    fmt.Println(bh)
    fmt.Println(bh.Result)
    fmt.Println(bh.Error)
    fmt.Println(bh.ID)
}
