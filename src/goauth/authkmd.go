//authclient.go
package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "os/user"
    "regexp"
)

//RPCUsername, RPCPassword string = "user60de7828fd8985d3", "ce3f74430f82aa34b58aeba4b37a3373"

func BytesToString(data []byte) string {
    return string(data[:])
}

func basicAuth() string {

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    fmt.Println( usr.HomeDir )

    kmdconf := usr.HomeDir+`/.komodo/komodo.conf`
    fmt.Println(kmdconf)

    confdata, err := ioutil.ReadFile(kmdconf)
    if err != nil {
        log.Fatal( err )
    }
    //fmt.Printf("%s\n",confdata)

    var rpcu = regexp.MustCompile("(?m)^rpcuser=.+$")
    fmt.Println(rpcu)

    fmt.Println(rpcu.Match(confdata))

    bytestr := BytesToString(confdata)
    //fmt.Println("BytesToString: "+bytestr)

    rpcuser_line := rpcu.FindString(bytestr)
    fmt.Printf("%q\n", rpcuser_line)

    re := regexp.MustCompile("rpcuser=.?")
    fmt.Printf("%q\n", re.FindString(bytestr))

    fmt.Println(re.FindStringSubmatch(rpcuser_line))

    var username string = "user60de7828fd8985d3"
    var passwd string = "ce3f74430f82aa34b58aeba4b37a3373"
    
    client := &http.Client{}

    url := "http://127.0.0.1:7771"
    fmt.Println("URL:>", url)

    query_byte := []byte(`{"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }`)
    fmt.Printf("Query: %s\n\n", query_byte)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

//    req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(username, passwd)

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