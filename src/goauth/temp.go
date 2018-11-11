package main

import "strings"
import "fmt"
import "github.com/satindergrewal/kmdgo/kmdutil"

func main() {
    initial := "<h1>Hello World!</h1>"

    out := strings.TrimLeft(strings.TrimRight(initial,"</h1>"),"<h1>")
    fmt.Println(out)

    appName := "komodo"

    dir := kmdutil.AppDataDir(appName, false)
    fmt.Println(dir)

    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)

    fmt.Printf("RPC User: %s\nRPC Password: %s\nRPC Port: %s\n", rpcuser, rpcpass, rpcport)
}