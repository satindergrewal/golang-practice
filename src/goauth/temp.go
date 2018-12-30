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


    //empl := kmdutil.Employee()

    /*
    //creating structure using field names
    emp1 := empl{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := empl{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)*/
}