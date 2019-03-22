//authclient.go
package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
)



func basicAuth() string {
    var username string = "someuser"
    var passwd string = "somepassword"
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://0.0.0.0:7000", nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    return s
}

func main(){
    fmt.Println("requesting...")
    S := basicAuth()
    fmt.Println(S)
}