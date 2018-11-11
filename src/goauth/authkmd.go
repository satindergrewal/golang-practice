//authclient.go
package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

//RPCUsername, RPCPassword string = "user60de7828fd8985d3", "ce3f74430f82aa34b58aeba4b37a3373"

func BytesToString(data []byte) string {
    return string(data[:])
}

func basicAuth() string {

    appName := "komodo"

    //appDir := kmdutil.AppDataDir(appName, false)
    //fmt.Println(appDir)

    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)

    client := &http.Client{}

    url := `http://127.0.0.1:`+rpcport
    fmt.Println("URL:>", url)

    query_byte := []byte(`{"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }`)
    fmt.Printf("Query: %s\n\n", query_byte)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

//    req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(rpcuser, rpcpass)

    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("%T\n\n", bodyText)

    var query_result map[string]interface{}

    if err := json.Unmarshal(bodyText, &query_result); err != nil {
        panic(err)
    }
    fmt.Println(query_result)
    fmt.Printf("\n\n")

    fmt.Println(query_result["result"].(map[string]interface{})["connections"])

    parsed_result := query_result["result"]
    fmt.Println("result: ", parsed_result)
    fmt.Printf("\n\n")

    s := string(bodyText)
    return s
}

func main(){
    
    //fmt.Printf("RPC Username: %s\nRPC Password: %s\n\n", RPCUsername, RPCPassword)
    

    fmt.Println("requesting...\n")
    S := basicAuth()
    fmt.Println(S)
}